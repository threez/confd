package main

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/xeipuuv/gojsonreference"
	"github.com/xeipuuv/gojsonschema"
)

// SimpleFileLoader loads Simple files for json schema use
type SimpleFileLoader struct {
	source string
}

// NewSimpleFileLoader creates a new folder based on the passed path
func NewSimpleFileLoader(filename string) gojsonschema.JSONLoader {
	return SimpleFileLoader{source: filename}
}

// JsonSource returns the Simple file path
func (simple SimpleFileLoader) JsonSource() interface{} {
	return simple.source
}

// LoadJSON returns the Simple file data
func (simple SimpleFileLoader) LoadJSON() (interface{}, error) {
	source := strings.Replace(simple.source, "file:///", "", 1)

	data, err := ioutil.ReadFile(source)
	if err != nil {
		return nil, err
	}

	var obj interface{}
	err = json.Unmarshal(data, &obj)
	return obj, err
}

// JsonReference returns the json schema refernce to the Simple
func (simple SimpleFileLoader) JsonReference() (gojsonreference.JsonReference, error) {
	return gojsonreference.NewJsonReference(simple.JsonSource().(string))
}

// LoaderFactory returns the SimpleFileLoader factory
func (simple SimpleFileLoader) LoaderFactory() gojsonschema.JSONLoaderFactory {
	return simple
}

// New retuns a new Simple file loader for the passed source
func (simple SimpleFileLoader) New(source string) gojsonschema.JSONLoader {
	return NewSimpleFileLoader(source)
}
