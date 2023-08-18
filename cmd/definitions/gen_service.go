//go:build tools
// +build tools

package main

import (
	"fmt"

	"github.com/Xuanwo/gg"
	"github.com/Xuanwo/templateutils"
	log "github.com/sirupsen/logrus"
)

func generateSrv(data *Service, path string) {
	f := gg.NewGroup()
	f.AddLineComment("Code generated by go generate via cmd/definitions; DO NOT EDIT.")
	f.AddPackage(data.Name)
	f.NewImport().
		AddPath("context").
		AddPath("io").
		AddPath("net/http").
		AddPath("strings").
		AddPath("time").
		AddLine().
		AddDot("github.com/fastone-open/go-storage/pairs").
		AddPath("github.com/fastone-open/go-storage/pkg/httpclient").
		AddPath("github.com/fastone-open/go-storage/services").
		AddDot("github.com/fastone-open/go-storage/types")

	f.NewVar().
		AddDecl("_", "Storager").
		AddDecl("_", "services.ServiceError").
		AddDecl("_", "httpclient.Options").
		AddDecl("_", "time.Duration").
		AddDecl("_", "http.Request").
		AddDecl("_", "Error")

	f.AddLineComment("Type is the type for %s", data.Name)
	f.NewConst().AddField("Type", gg.Lit(data.Name).String())

	// Generate object system metadata.
	f.AddLineComment("ObjectSystemMetadata stores system metadata for object.")
	osm := f.NewStruct("ObjectSystemMetadata")
	for _, info := range data.SortedInfos() {
		if info.Scope != "object" {
			continue
		}
		if info.Global {
			continue
		}
		pname := templateutils.ToPascal(info.Name)
		if info.DisplayName() != "" {
			pname = info.DisplayName()
		}
		// FIXME: we will support comment on field later.
		osm.AddField(pname, info.Type)
	}

	f.AddLineComment(`
GetObjectSystemMetadata will get ObjectSystemMetadata from Object.

- This function should not be called by service implementer.
- The returning ObjectServiceMetadata is read only and should not be modified.
`)

	gosmfn := f.NewFunction("GetObjectSystemMetadata")
	gosmfn.AddParameter("o", "*Object").
		AddResult("", "ObjectSystemMetadata")
	gosmfn.AddBody(
		gg.S("sm, ok := o.GetSystemMetadata()"),
		gg.If(gg.S("ok")).
			AddBody(gg.Return("sm.(ObjectSystemMetadata)")),
		gg.Return(gg.Value("ObjectSystemMetadata")),
	)

	f.AddLineComment(`
setObjectSystemMetadata will set ObjectSystemMetadata into Object.

- This function should only be called once, please make sure all data has been written before set.
`)

	sosmfn := f.NewFunction("setObjectSystemMetadata")
	sosmfn.AddParameter("o", "*Object").
		AddParameter("sm", "ObjectSystemMetadata")
	sosmfn.AddBody(
		gg.S("o.SetSystemMetadata(sm)"),
	)

	// Generate storage system metadata.
	f.AddLineComment("StorageSystemMetadata stores system metadata for object.")
	ssm := f.NewStruct("StorageSystemMetadata")
	for _, info := range data.SortedInfos() {
		if info.Scope != "object" {
			continue
		}
		if info.Global {
			continue
		}
		pname := templateutils.ToPascal(info.Name)
		if info.DisplayName() != "" {
			pname = info.DisplayName()
		}
		// FIXME: we will support comment on field later.
		ssm.AddField(pname, info.Type)
	}

	f.AddLineComment(`
GetStorageSystemMetadata will get StorageSystemMetadata from Storage.

- This function should not be called by service implementer.
- The returning StorageServiceMetadata is read only and should not be modified.
`)

	gssmfn := f.NewFunction("GetStorageSystemMetadata")
	gssmfn.AddParameter("s", "*StorageMeta").
		AddResult("", "StorageSystemMetadata")
	gssmfn.AddBody(
		gg.S("sm, ok := s.GetSystemMetadata()"),
		gg.If(gg.S("ok")).
			AddBody(gg.Return("sm.(StorageSystemMetadata)")),
		gg.Return(gg.Value("StorageSystemMetadata")),
	)

	f.AddLineComment(`setStorageSystemMetadata will set StorageSystemMetadata into Storage.

- This function should only be called once, please make sure all data has been written before set.`)
	sssmfn := f.NewFunction("setStorageSystemMetadata")
	sssmfn.AddParameter("s", "*StorageMeta").
		AddParameter("sm", "StorageSystemMetadata")
	sssmfn.AddBody(
		gg.S("s.SetSystemMetadata(sm)"),
	)

	// Generate service pairs.
	for _, pair := range data.SortedPairs() {
		// We don't need to generate global pairs here.
		if pair.Global {
			continue
		}

		pname := templateutils.ToPascal(pair.Name)

		f.AddLineComment(`With%s will apply %s value to Options.

%s`, pname, pair.Name, pair.Description)
		fn := f.NewFunction("With" + pname)

		// Set to true as default.
		value := "true"
		// bool type pairs don't need input.
		if pair.Type != "bool" {
			fn.AddParameter("v", pair.Type)
			value = "v"
		}
		fn.AddResult("", "Pair")
		fn.AddBody(gg.Return(
			gg.Value("Pair").
				AddField("Key", gg.Lit(pair.Name)).
				AddField("Value", value)))
	}

	// Generate pair map
	f.NewVar().AddField("pairMap", gg.Embed(func() gg.Node {
		i := gg.Value("map[string]string")
		for _, pair := range data.SortedPairs() {
			i.AddField(gg.Lit(pair.Name), gg.Lit(pair.Type))
		}
		return i
	}))

	// Generate every namespace.
	for _, ns := range data.SortedNamespaces() {
		nsNameP := templateutils.ToPascal(ns.Name)

		// Generate interface assert.
		inters := f.NewVar()
		for _, inter := range ns.ParsedInterfaces() {
			interNameP := templateutils.ToPascal(inter.Name)
			inters.AddTypedField(
				"_", interNameP, gg.S("&%s{}", nsNameP))
		}

		// // Generate feature struct.
		// features := f.NewStruct(nsNameP + "Features")
		// for _, fs := range ns.ParsedFeatures() {
		// 	features.AddLineComment(fs.Description)
		// 	features.AddField(templateutils.ToPascal(fs.Name), "bool")
		// }

		// Generate pair new.
		fnNewNameP := templateutils.ToPascal(ns.New.Name)
		partStructName := fmt.Sprintf("pair%s%s", nsNameP, fnNewNameP)
		f.AddLineComment("%s is the parsed struct", partStructName)
		pairStruct := f.NewStruct(partStructName).
			AddField("pairs", "[]Pair").
			AddLine()

		// Generate required pairs.
		pairStruct.AddLineComment("Required pairs")
		for _, pair := range ns.New.ParsedRequired() {
			pairNameP := templateutils.ToPascal(pair.Name)
			pairStruct.AddField("Has"+pairNameP, "bool")
			pairStruct.AddField(pairNameP, pair.Type)
		}

		// Generate optional pairs.
		pairStruct.AddLineComment("Optional pairs")
		for _, pair := range ns.New.ParsedOptional() {
			pairNameP := templateutils.ToPascal(pair.Name)
			pairStruct.AddField("Has"+pairNameP, "bool")
			pairStruct.AddField(pairNameP, pair.Type)
		}

		// Generate feature handle logic.
		pairStruct.AddLineComment("Enable features")
		for _, feature := range ns.ParsedFeatures() {
			featureNameP := templateutils.ToPascal(feature.Name)
			pairStruct.AddField("hasEnable"+featureNameP, "bool")
			pairStruct.AddField("Enable"+featureNameP, "bool")
		}

		// Generate parse newPair.
		pairParseName := fmt.Sprintf("parsePair%s%s", nsNameP, fnNewNameP)
		f.AddLineComment("%s will parse Pair slice into *%s", pairParseName, partStructName)
		pairParse := f.NewFunction(pairParseName).
			AddParameter("opts", "[]Pair").
			AddResult("", partStructName).
			AddResult("", "error")
		pairParse.AddBody(
			gg.S("result :="),
			gg.Value(partStructName).AddField("pairs", "opts"),
			gg.Line(),
			gg.For(gg.S("_, v := range opts")).
				AddBody(gg.Embed(func() gg.Node {
					is := gg.Switch(gg.S("v.Key"))

					for _, pair := range ns.New.ParsedRequired() {
						pairNameP := templateutils.ToPascal(pair.Name)
						is.NewCase(gg.Lit(pair.Name)).AddBody(
							gg.If(gg.S("result.Has%s", pairNameP)).
								AddBody(gg.Continue()),
							gg.S("result.Has%s = true", pairNameP),
							gg.S("result.%s = v.Value.(%s)", pairNameP, pair.Type),
						)
					}
					for _, pair := range ns.New.ParsedOptional() {
						pairNameP := templateutils.ToPascal(pair.Name)
						is.NewCase(gg.Lit(pair.Name)).AddBody(
							gg.If(gg.S("result.Has%s", pairNameP)).
								AddBody(gg.Continue()),
							gg.S("result.Has%s = true", pairNameP),
							gg.S("result.%s = v.Value.(%s)", pairNameP, pair.Type),
						)
					}
					for _, feature := range ns.ParsedFeatures() {
						featureNameP := templateutils.ToPascal(feature.Name)
						is.NewCase(gg.Lit("enable_"+feature.Name)).AddBody(
							gg.If(gg.S("result.hasEnable%s", featureNameP)).
								AddBody(gg.Continue()),
							gg.S("result.hasEnable%s = true", featureNameP),
							gg.S("result.Enable%s = true", featureNameP),
						)
					}
					return is
				})),
			gg.LineComment("Enable features"),
			gg.Embed(func() gg.Node {
				// Generate features enable here.
				group := gg.NewGroup()
				for _, feature := range ns.ParsedFeatures() {
					featureNameP := templateutils.ToPascal(feature.Name)

					group.
						NewIf(gg.S("result.hasEnable%s", featureNameP)).
						AddBody(
							gg.S("result.Has%sFeatures = true", nsNameP),
							gg.S("result.%sFeatures.%s = true", nsNameP, featureNameP),
						)
				}
				return group
			}),
			gg.LineComment("Default pairs"),
			gg.Embed(func() gg.Node {
				// Generate default pair handle logic here.
				group := gg.NewGroup()

				for _, dp := range ns.ParsedDefaultable() {
					pairNameP := templateutils.ToPascal(dp.Pair.Name)

					xif := group.
						NewIf(gg.S("result.HasDefault%s", pairNameP)).
						AddBody(gg.S("result.HasDefault%sPairs = true", nsNameP))
					for _, op := range dp.Funcs {
						opN := templateutils.ToPascal(op.Name)

						xif.AddBody(gg.S(
							"result.Default%sPairs.%s = append(result.Default%sPairs.%s, With%s(result.Default%s))",
							nsNameP, opN, nsNameP, opN, pairNameP, pairNameP))
					}
				}
				return group
			}),
			gg.Embed(func() gg.Node {
				// Generate required validate logic here.
				group := gg.NewGroup()

				for _, pair := range ns.New.ParsedRequired() {
					pairNameP := templateutils.ToPascal(pair.Name)
					group.NewIf(gg.S("!result.Has%s", pairNameP)).
						AddBody(gg.S(
							`return pair%s%s{}, services.PairRequiredError{ Keys:[]string{ "%s" } }`,
							nsNameP, fnNewNameP, pair.Name))
				}
				return group
			}),
			gg.Return("result", "nil"),
		)

		// // Generate default pairs.
		// f.AddLineComment("Default%sPairs is default pairs for specific action", nsNameP)
		// dps := f.NewStruct(fmt.Sprintf("Default%sPairs", nsNameP))
		// for _, fn := range ns.ParsedFunctions() {
		// 	fnNameP := templateutils.ToPascal(fn.Name)
		// 	dps.AddField(fnNameP, "[]Pair")
		// }

		// Generate pair.
		for _, fn := range ns.ParsedFunctions() {
			fnNameP := templateutils.ToPascal(fn.Name)
			pairStructName := fmt.Sprintf("pair%s%s", nsNameP, fnNameP)

			// Generate pair struct
			pairStruct := f.NewStruct(pairStructName).
				AddField("pairs", "[]Pair")

			// Generate required pairs.
			pairStruct.AddLineComment("Required pairs")
			for _, pair := range fn.ParsedRequired() {
				pairNameP := templateutils.ToPascal(pair.Name)
				pairStruct.AddField("Has"+pairNameP, "bool")
				pairStruct.AddField(pairNameP, pair.Type)
			}

			// Generate optional pairs.
			pairStruct.AddLineComment("Optional pairs")
			for _, pair := range fn.ParsedOptional() {
				pairNameP := templateutils.ToPascal(pair.Name)
				pairStruct.AddField("Has"+pairNameP, "bool")
				pairStruct.AddField(pairNameP, pair.Type)
			}

			pairParseName := fmt.Sprintf("parsePair%s%s", nsNameP, fnNameP)
			pairParse := f.NewFunction(pairParseName).
				WithReceiver("s", "*"+nsNameP).
				AddParameter("opts", "[]Pair").
				AddResult("", pairStructName).
				AddResult("", "error")
			pairParse.AddBody(
				gg.S("result :="),
				gg.Value(pairStructName).AddField("pairs", "opts"),
				gg.Line(),
				gg.For(gg.S("_, v := range opts")).
					AddBody(gg.Embed(func() gg.Node {
						is := gg.Switch(gg.S("v.Key"))

						for _, pair := range fn.ParsedRequired() {
							pairNameP := templateutils.ToPascal(pair.Name)
							is.NewCase(gg.Lit(pair.Name)).AddBody(
								gg.If(gg.S("result.Has%s", pairNameP)).
									AddBody(gg.Continue()),
								gg.S("result.Has%s = true", pairNameP),
								gg.S("result.%s = v.Value.(%s)", pairNameP, pair.Type),
							)
						}
						for _, pair := range fn.ParsedOptional() {
							pairNameP := templateutils.ToPascal(pair.Name)
							is.NewCase(gg.Lit(pair.Name)).AddBody(
								gg.If(gg.S("result.Has%s", pairNameP)).
									AddBody(gg.Continue()),
								gg.S("result.Has%s = true", pairNameP),
								gg.S("result.%s = v.Value.(%s)", pairNameP, pair.Type),
							)
						}
						dcas := is.NewDefault()
						if ns.HasFeatureLoosePair {
							dcas.AddBody(
								gg.LineComment(`
loose_pair feature introduced in GSP-109.
If user enable this feature, service should ignore not support pair error.`),
								gg.If(gg.S("s.features.LoosePair")).
									AddBody(gg.Continue()),
							)
						}
						dcas.AddBody(gg.S("return pair%s%s{}, services.PairUnsupportedError{Pair:v}", nsNameP, fnNameP))
						return is
					})),

				gg.Embed(func() gg.Node {
					// Generate required validate logic here.
					group := gg.NewGroup()

					for _, pair := range fn.ParsedRequired() {
						pairNameP := templateutils.ToPascal(pair.Name)
						group.NewIf(gg.S("!result.Has%s", pairNameP)).
							AddBody(gg.S(
								`return pair%s%s{}, services.PairRequiredError{ Keys:[]string{ "%s" } }`,
								nsNameP, fnNameP, pair.Name))
					}
					return group
				}),
				gg.Return("result", "nil"),
			)
		}

		// Generate public functions.
		for _, fn := range ns.ParsedFunctions() {
			fnNameP := templateutils.ToPascal(fn.Name)
			op := fn.GetOperation()
			if op.Local {
				// Generate a local function.
				xfn := f.NewFunction(fnNameP).WithReceiver("s", "*"+nsNameP)
				for _, field := range op.ParsedParams() {
					xfn.AddParameter(field.Name, field.Type)
				}
				for _, field := range op.ParsedResults() {
					xfn.AddResult(field.Name, field.Type)
				}
				xfn.AddBody(
					gg.S("pairs = append(pairs, s.defaultPairs.%s...)", fnNameP),
					gg.S("var opt pair%s%s", nsNameP, fnNameP),
					gg.Line(),
					gg.LineComment("Ignore error while handling local functions."),
					gg.S("opt, _ = s.parsePair%s%s(pairs)", nsNameP, fnNameP),
					gg.Return(
						gg.Embed(func() gg.Node {
							ic := gg.Call(templateutils.ToCamel(fn.Name)).
								WithOwner("s")
							for _, v := range op.ParsedParams() {
								// We don't need to call pair again.
								if v.Type == "...Pair" {
									continue
								}
								ic.AddParameter(v.Name)
							}
							ic.AddParameter("opt")
							return ic
						})))
				continue
			}
			// Generate a non-local function.
			// TODO: generate comment here.
			xfn := f.NewFunction(fnNameP).
				WithReceiver("s", "*"+nsNameP)
			for _, field := range op.ParsedParams() {
				xfn.AddParameter(field.Name, field.Type)
			}
			for _, field := range op.ParsedResults() {
				xfn.AddResult(field.Name, field.Type)
			}
			xfn.AddBody(
				"ctx := context.Background()",
				gg.Return(
					gg.Embed(func() gg.Node {
						ic := gg.Call(fnNameP + "WithContext").
							WithOwner("s")
						ic.AddParameter("ctx")
						for _, v := range op.ParsedParams() {
							if v.Type == "...Pair" {
								ic.AddParameter("pairs...")
								continue
							}
							ic.AddParameter(v.Name)
						}
						return ic
					})))

			xfn = f.NewFunction(fnNameP+"WithContext").
				WithReceiver("s", "*"+nsNameP)
			xfn.AddParameter("ctx", "context.Context")
			for _, field := range op.ParsedParams() {
				xfn.AddParameter(field.Name, field.Type)
			}
			for _, field := range op.ParsedResults() {
				xfn.AddResult(field.Name, field.Type)
			}
			xfn.AddBody(
				gg.Defer(gg.Embed(func() gg.Node {
					caller := gg.Call("formatError").WithOwner("s")
					caller.AddParameter(gg.Lit(fn.Name)).AddParameter("err")
					isEmpty := true
					for _, v := range op.ParsedParams() {
						// formatError only accept string as input.
						if v.Type != "string" {
							continue
						}
						caller.AddParameter(v.Name)
						isEmpty = false
					}
					if isEmpty && nsNameP == "Service" {
						caller.AddParameter("\"\"")
					}

					fn := gg.Function("").
						AddBody(gg.S("err = "), caller).WithCall()
					return fn
				})),
				gg.Embed(func() gg.Node {
					if op.ObjectMode == "" {
						return gg.Line()
					}
					mode := templateutils.ToPascal(op.ObjectMode)
					return gg.If(gg.S("!o.Mode.Is%s()", mode)).AddBody(
						gg.S("err = services.ObjectModeInvalidError{Expected: Mode%s, Actual: o.Mode}", mode),
						gg.Return(),
					)
				}),
				gg.S("pairs = append(pairs, s.defaultPairs.%s...)", fnNameP),
				gg.S("var opt pair%s%s", nsNameP, fnNameP),
				gg.Line(),
				gg.S("opt, err = s.parsePair%s%s(pairs)", nsNameP, fnNameP),
				gg.If(gg.S("err != nil")).AddBody(gg.Return()),
				gg.Return(
					gg.Embed(func() gg.Node {
						ic := gg.Call(templateutils.ToCamel(fn.Name)).
							WithOwner("s")
						ic.AddParameter("ctx")
						for _, v := range op.ParsedParams() {
							// We don't need to call pair again.
							if v.Type == "...Pair" {
								continue
							}
							if v.Name == "path" || v.Name == "src" || v.Name == "dst" || v.Name == "target" {
								ic.AddParameter(gg.S(`strings.ReplaceAll(%s, "\\", "/")`, v.Name))
								continue
							}
							ic.AddParameter(v.Name)
						}
						ic.AddParameter("opt")
						return ic
					})))
		}
	}

	// Generate init function
	initFn := f.NewFunction("init")
	for _, ns := range data.SortedNamespaces() {
		nsNameP := templateutils.ToPascal(ns.Name)
		initFn.AddBody(gg.Call("Register" + nsNameP + "r").
			WithOwner("services").
			AddParameter("Type").
			AddParameter("New" + nsNameP + "r"))
	}
	initFn.AddBody("services.RegisterSchema(Type, pairMap)")

	err := f.WriteFile(path)
	if err != nil {
		log.Fatalf("generate to %s: %v", path, err)
	}
}
