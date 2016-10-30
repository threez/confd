package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

type Dir string

func (d Dir) Open(name string) (http.File, error) {
	path := filepath.Join(string(d), name)
	fmt.Printf("Loading %s", path)
	return os.Open(path)
}

func TestValidObjects(t *testing.T) {
	files("test/valid/types/**/*.json", t, func(result *gojsonschema.Result) {
		if !result.Valid() {
			t.Error("The document is not valid. see errors")
			for _, desc := range result.Errors() {
				t.Error(desc)
			}
			t.FailNow()
		}
	})
}

func TestInvalidObjects(t *testing.T) {
	files("test/invalid/types/**/*.json", t, func(result *gojsonschema.Result) {
		if result.Valid() {
			t.Error("The document is valid but should not!")
			t.FailNow()
		}
	})
}

func files(path string, t *testing.T, r func(result *gojsonschema.Result)) {
	paths, err := filepath.Glob(path)
	assert.NoError(t, err)

	for _, path := range paths {
		parts := strings.Split(path, string(filepath.Separator))
		path := "file:///" + path
		schemaPath := "file:///types/" + parts[3] + ".json"
		t.Logf("Testing %s -> %s\n", path, schemaPath)

		schemaLoader := NewSimpleFileLoader(schemaPath)
		documentLoader := NewSimpleFileLoader(path)

		result, err := gojsonschema.Validate(schemaLoader, documentLoader)
		assert.NoError(t, err)

		r(result)
	}
}
