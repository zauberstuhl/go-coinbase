ALL_SRCS := $(wildcard *.go)
PKG_DIR := $(GOPATH)/src/zauberstuhl/coinbase

all: compile
test: gotest
gotest:
	rm -rv $(PKG_DIR) || true
	mkdir -vp $(PKG_DIR) && cp -rv *.go $(PKG_DIR)
	go test -v ./test
compile:
	go build -v $(SRCS)
