package test

import (
	"bytes"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// FIXTURE OMIT
func TestSomething(t *testing.T) {
	f, err := os.Open("testdata/somefixture.json")
	//...
}

// END OMIT

// GOLDEN OMIT
var update = flag.Bool("update", false, "update .golden files") // HLUP

func init() {
	flag.Parse()
}

func TestSomething(t *testing.T) {
	actual := doSomething()
	golden := filepath.Join("testdata", tc.Name+".golden")
	if *update { // HLUP
		ioutil.WriteFile(golden, actual, 0644) // HLUP
	} // HLUP
	expected, _ := ioutil.ReadFile(golden)

	if !bytes.Equal(actual, expected) {
		// FAIL!
	}
}

// END OMIT

// GOLDENRES OMIT
$ go test 
…

$ go test -update
…
// END OMIT
