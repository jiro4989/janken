BIN := bin/generate_janken_html
ifeq ($(ENV),prd)
	TEMPLATE_PATH := "/opt/janken/template"
	DEST_DIR := "/var/www/html/hand"
	LDFLAGS := -ldflags=" \
		-X \"main.TEMPLATE_PATH=$(TEMPLATE_PATH)\" \
		-X \"main.DEST_DIR=$(DEST_DIR)\" \
		"
endif

build:
	mkdir -p bin
	go build -o $(BIN) $(LDFLAGS) .
	cp -p $(BIN) ../bin/

run: bin/generate_janken_html
	$< rock paper

test: *.go
	go test $(LDFLAGS) -cover

.PHONY: test run build
