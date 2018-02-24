package test

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"testing"
)

// HELPER OMIT
func helperLoadFile(t *testing.T, name string) ([]byte, err) {
	path := filepath.Join("testdata", name)
	return ioutil.ReadFile(path)
}

func TestFile(t *testing.T) {
	bytes, err := helperLoadFile(t, "some_file")
	//...
}

// END OMIT

// SETUP OMIT
func setUp(t *testing.T) {
	// ...
}

func tearDown(t *testing.T) {
	// ...
}

func TestSomething(t *testing.T) {
	setUp(t)
	defer tearDown(t)
	//...
}

// END OMIT

// WRONG OMIT
func TestDeferFatal(t *testing.T) {
	defer helperFatal(t)
	t.Fatal("Test fatal") // HLF
}

func helperFatal(t *testing.T) {
	log.Fatal("Helper fatal") // HLF
}

// END OMIT

// WRONGRES OMIT
$ go test ./test --test.run=TestDeferFatal
2018/02/05 14:03:54 FATAL: Helper fatal // HLF
// END OMIT