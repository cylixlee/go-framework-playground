.PHONY: help generate clean

help:
	@echo "Available Make recipes:"
	@echo "  generate - Generate boilerplate code using goctl"
	@echo "  clean    - Remove all generated code except internal/logic"
	@echo ""

generate: cmd/user/main.go

# Generate code from user.api
#
# We must use cp+rm instead of mv to make timestamp present, in order to make incremental make possible.
cmd/user/main.go: user.api
	@mkdir -p cmd/user
	@echo -n "Generating code... "
	@goctl api go -api user.api -dir .
	@cp user.go cmd/user/main.go
	@rm user.go

clean:
	@rm -rf cmd
	@rm -rf etc
	@for pkg in internal/*; do                                     \
		if [ -d "$$pkg" ] && [ "$$pkg" != "internal/logic" ]; then \
			rm -rf $$pkg;                                          \
		fi                                                         \
	done