BUF_VERSION:=v1.9.0

all: gen

.PHONY: gen
gen: gen/proto

gen/proto:
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) mod update
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) generate
