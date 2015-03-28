API_JSON = $(wildcard */*/*-api.json)

all: generator
	$(GOPATH)/bin/google-api-go-generator -cache=false -install -api=*

cached: generator
	$(GOPATH)/bin/google-api-go-generator -cache=true -install -api=*

local: $(API_JSON:-api.json=-gen.go)

%-gen.go: %-api.json generator
	$(GOPATH)/bin/google-api-go-generator -api_json_file=$<

generator:
	go install google.golang.org/api/googleapi
	go install google.golang.org/api/google-api-go-generator

.PHONY: all cached local generator
