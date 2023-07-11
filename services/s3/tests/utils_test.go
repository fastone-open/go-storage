package tests

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"sync"
	"testing"

	ps "github.com/fastone-open/go-storage/pairs"
	s3 "github.com/fastone-open/go-storage/services/s3"
	"github.com/fastone-open/go-storage/types"
	typ "github.com/fastone-open/go-storage/types"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/assert"
)

func TestMy(t *testing.T) {

	type Conf struct {
		STORAGE_S3_CREDENTIAL string
		STORAGE_S3_NAME       string
		STORAGE_S3_LOCATION   string
		STORAGE_S3_ENDPOINT   string
	}

	buf, err := os.ReadFile("../volc.toml")
	assert.Nil(t, err)

	var conf Conf
	err = toml.Unmarshal(buf, &conf)
	assert.Nil(t, err)

	c := conf.STORAGE_S3_CREDENTIAL
	n := conf.STORAGE_S3_NAME
	l := conf.STORAGE_S3_LOCATION
	e := conf.STORAGE_S3_ENDPOINT

	store, err := s3.NewStorager(
		ps.WithCredential(c),
		ps.WithName(n),
		ps.WithLocation(l),
		ps.WithEndpoint(e),
	)
	assert.Nil(t, err)

	to := new(bytes.Buffer)
	_, err = store.Read("__test", to)
	assert.Nil(t, err)
	fmt.Println(n)
	fmt.Println(to.String())
}

func TestList(t *testing.T) {
	type Conf struct {
		STORAGE_S3_CREDENTIAL string
		STORAGE_S3_NAME       string
		STORAGE_S3_LOCATION   string
		STORAGE_S3_ENDPOINT   string
	}

	buf, err := os.ReadFile("../debug_tencent2.toml")
	assert.Nil(t, err)

	var conf Conf
	err = toml.Unmarshal(buf, &conf)
	assert.Nil(t, err)

	c := conf.STORAGE_S3_CREDENTIAL
	n := conf.STORAGE_S3_NAME
	l := conf.STORAGE_S3_LOCATION
	e := conf.STORAGE_S3_ENDPOINT

	store, err := s3.NewStorager(
		ps.WithCredential(c),
		ps.WithName(n),
		ps.WithLocation(l),
		ps.WithEndpoint(e),
	)
	assert.Nil(t, err)

	var wg sync.WaitGroup
	list := func(prefix, dest string) {
		defer wg.Done()
		oi, err := store.List(prefix)
		assert.Nil(t, err)
		var first = true
		var object *types.Object

		recordFile, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 0644)
		assert.Nil(t, err)
		defer recordFile.Close()
		writer := bufio.NewWriter(recordFile)

		for first || object != nil {
			object, err = oi.Next()
			if err != nil {
				if errors.Is(err, types.IterateDone) {
					return
				} else {
					fmt.Println(err)
				}
			}

			_, err := writer.WriteString(fmt.Sprintf("%s--%s\n", object.Path,
				object.MustGetLastModified().String()))
			assert.Nil(t, err)

			if first {
				first = false
			}
		}
	}

	for _, e := range []struct {
		prefix, dest string
	}{
		{
			prefix: "nott/__read/",
			dest:   "read.txt",
		},
	} {
		wg.Add(1)
		go list(e.prefix, e.dest)
	}
	wg.Wait()
}

func TestGetRoot(t *testing.T) {
	type Conf struct {
		STORAGE_S3_CREDENTIAL string
		STORAGE_S3_NAME       string
		STORAGE_S3_LOCATION   string
		STORAGE_S3_ENDPOINT   string
	}

	buf, err := os.ReadFile("../debug_tencent2.toml")
	assert.Nil(t, err)

	var conf Conf
	err = toml.Unmarshal(buf, &conf)
	assert.Nil(t, err)

	c := conf.STORAGE_S3_CREDENTIAL
	n := conf.STORAGE_S3_NAME
	l := conf.STORAGE_S3_LOCATION
	e := conf.STORAGE_S3_ENDPOINT

	store, err := s3.NewStorager(
		ps.WithCredential(c),
		ps.WithName(n),
		ps.WithLocation(l),
		ps.WithEndpoint(e),
	)
	assert.Nil(t, err)

	readRootState, err := store.Stat("nott/__read/root")
	fmt.Println(readRootState.Path, readRootState.MustGetLastModified())
	assert.Nil(t, err)
	writeRootState, err := store.Stat("nott/__write/root")
	assert.Nil(t, err)
	fmt.Println(writeRootState.Path, writeRootState.MustGetLastModified())
}

func TestGetObj(t *testing.T) {
	type Conf struct {
		STORAGE_S3_CREDENTIAL string
		STORAGE_S3_NAME       string
		STORAGE_S3_LOCATION   string
		STORAGE_S3_ENDPOINT   string
	}

	buf, err := os.ReadFile("../debug_tencent2.toml")
	assert.Nil(t, err)

	var conf Conf
	err = toml.Unmarshal(buf, &conf)
	assert.Nil(t, err)

	c := conf.STORAGE_S3_CREDENTIAL
	n := conf.STORAGE_S3_NAME
	l := conf.STORAGE_S3_LOCATION
	e := conf.STORAGE_S3_ENDPOINT

	store, err := s3.NewStorager(
		ps.WithCredential(c),
		ps.WithName(n),
		ps.WithLocation(l),
		ps.WithEndpoint(e),
	)
	assert.Nil(t, err)

	getObject(t, store, "nott/__read/1678585122024062976")
}

func getObject(t *testing.T, store typ.Storager, path string) {
	o, err := store.Stat(path)
	fmt.Println(o.Path, o.MustGetLastModified())
	assert.Nil(t, err)

	_, err = store.Read(path, os.Stdout)
	assert.Nil(t, err)
}

func setupTest(t *testing.T) types.Storager {
	t.Log("Setup test for s3")

	store, err := s3.NewStorager(
		ps.WithCredential(os.Getenv("STORAGE_S3_CREDENTIAL")),
		ps.WithName(os.Getenv("STORAGE_S3_NAME")),
		ps.WithLocation(os.Getenv("STORAGE_S3_LOCATION")),
		ps.WithEndpoint(os.Getenv("STORAGE_S3_ENDPOINT")),
		ps.WithEnableVirtualDir(),
		ps.WithEnableVirtualLink(),
		s3.WithForcePathStyle(),
	)
	if err != nil {
		t.Errorf("new storager: %v", err)
	}
	return store
}
