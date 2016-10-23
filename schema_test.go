package main

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xeipuuv/gojsonschema"
)

func TestValidObjects(t *testing.T) {
	files("test/valid/types/**/*.yaml", t, func(result *gojsonschema.Result) {
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
	files("test/invalid/types/**/*.yaml", t, func(result *gojsonschema.Result) {
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
		schemaPath := filepath.Join("types", parts[3]+".yaml")
		t.Logf("Testing %s -> %s\n", path, schemaPath)
		schemaLoader := NewYAMLFileLoader("file://" + schemaPath)
		documentLoader := NewYAMLFileLoader("file://" + path)

		result, err := gojsonschema.Validate(schemaLoader, documentLoader)
		if err != nil {
			assert.NoError(t, err)
		}
		r(result)
	}
}
