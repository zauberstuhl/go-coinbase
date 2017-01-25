ALL_SRCS := $(wildcard *.go)

all: compile
test: gotest
gotest:
	go test -v ./test
compile:
	go build $(SRCS)
