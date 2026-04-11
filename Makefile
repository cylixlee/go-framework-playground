.PHONY: help clean proto

MODULE  := github.com/cylixlee/go-framework-playground
PROTOS  := $(wildcard proto/*.proto)
TARGETS := $(addsuffix /main.go,$(addprefix cmd/,$(basename $(notdir $(PROTOS)))))

help:
	@echo "Available Make recipes:"
	@echo "  clean - Cleanup generated code"
	@echo "  proto - Compile proto files"
	@echo ""

proto: $(TARGETS)

cmd/%/main.go: proto/%.proto
	@echo -n "Compiling $<... "
	@goctl rpc protoc $<                               \
		--go_out=.      --go_opt=module=$(MODULE)      \
		--go-grpc_out=. --go-grpc_opt=module=$(MODULE) \
		--zrpc_out=.                                   \
		-m -c=false --name-from-filename
	@mkdir -p "$(dir $@)"
	@mv "$(addsuffix .go,$(basename $(notdir $<)))" "$@"