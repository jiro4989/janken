TEMPLATE_PATH := "../template/janken.html"

ifeq ($(ENV),prd)
  TEMPLATE_PATH := "/opt/janken/template"
endif

LDFLAGS := -ldflags="-X \"main.TEMPLATE_PATH=$(TEMPLATE_PATH)\""
bin/generate_janken_html: *.go
	mkdir -p bin
	go build -o $@ $(LDFLAGS) .
	cp -p $@ ../bin/

run: bin/generate_janken_html
	$< rock paper

test: *.go
	go test $(LDFLAGS) -cover

.PHONY: test run
