.PHONY: help generate containers migrate clean

MACHINE_STATE := $(shell podman machine info | grep -oP '^\s*machinestate:\s*\K\S+$$')

help:
	@echo ""
	@echo "Available Make recipes:"
	@echo "  generate   - Generate boilerplate code using goctl"
	@echo "  containers - Pull up containers using Podman"
	@echo "  migrate    - Generate query boilerplate code and migrate database from SQL file."
	@echo "  clean      - Shut down containers using Podman"
	@echo ""

# Generate code from user.api
generate: cmd/user/main.go
# Must use cp + rm instead of mv to make timestamp present, in order to make incremental make possible.
cmd/user/main.go: user.api
	@mkdir -p cmd/user
	@echo "Generating from user.api"
	@goctl api go -api user.api -dir . 2> /dev/null
	@cp user.go cmd/user/main.go
	@rm user.go

containers:
	@if [ "$(MACHINE_STATE)" != "Running" ]; then \
		podman machine start;                     \
	fi
	@podman compose up -d 2> /dev/null

migrate: containers
	@goctl model mysql ddl -s "user.sql" -dir "internal/model"
	@podman exec -i mysql mysql -uroot -p123456 < user.sql

clean:
	@podman compose down 2> /dev/null