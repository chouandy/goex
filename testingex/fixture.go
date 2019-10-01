package testingex

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

// LoadFixture load fixture
func LoadFixture(filename string) []byte {
	// Get file absolute path
	path, err := filepath.Abs("./fixtures/" + filename)
	if err != nil {
		panic(err)
	}

	// Read file
	f, err := ioutil.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Cannot load fixture %s", filename))
	}

	return f
}
