YAML_PACKAGE=github.com/ghodss/yaml
SCHEMA_PACKAGE=github.com/xeipuuv/gojsonschema
TEST_PACKAGE=github.com/stretchr/testify/assert

test: $(GOPATH)/src/$(YAML_PACKAGE) $(GOPATH)/src/$(SCHEMA_PACKAGE) $(GOPATH)/src/$(TEST_PACKAGE)
	go test -v

$(GOPATH)/src/$(YAML_PACKAGE):
	go get $(YAML_PACKAGE)

$(GOPATH)/src/$(SCHEMA_PACKAGE):
	go get $(SCHEMA_PACKAGE)

$(GOPATH)/src/$(TEST_PACKAGE):
	go get $(TEST_PACKAGE)

.PHONY: test
