package main

import (
	"io/ioutil"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/xeipuuv/gojsonreference"
	"github.com/xeipuuv/gojsonschema"
)

// YAMLFileLoader loads yaml files for json schema use
type YAMLFileLoader struct {
	source string
}

// NewYAMLFileLoader creates a new folder based on the passed path
func NewYAMLFileLoader(filename string) gojsonschema.JSONLoader {
	return YAMLFileLoader{source: filename}
}

// JsonSource returns the yaml file path
func (yfl YAMLFileLoader) JsonSource() interface{} {
	return yfl.source
}

// LoadJSON returns the yaml file data
func (yfl YAMLFileLoader) LoadJSON() (interface{}, error) {
	source := strings.Replace(yfl.source, "file://", "", 1)

	data, err := ioutil.ReadFile(source)
	if err != nil {
		return nil, err
	}

	var obj interface{}
	err = yaml.Unmarshal(data, &obj)
	return obj, err
}

// JsonReference returns the json schema refernce to the yaml
func (yfl YAMLFileLoader) JsonReference() (gojsonreference.JsonReference, error) {
	return gojsonreference.NewJsonReference(yfl.JsonSource().(string))
}

// LoaderFactory returns the YAMLFileLoader factory
func (yfl YAMLFileLoader) LoaderFactory() gojsonschema.JSONLoaderFactory {
	return yfl
}

// New retuns a new YAML file loader for the passed source
func (yfl YAMLFileLoader) New(source string) gojsonschema.JSONLoader {
	return NewYAMLFileLoader(source)
}
