NAKAMA_APP := nakama-godot4-extended

# Docker volume
PG_VOLUME := nakama-server_pg_data

# Server modules
M_WORLD_CONTROL = world_control
M_WORLD_RPC = world_rpc

all: build

# Build the modules
build:
	@echo "Building..."
	@go build --trimpath --mod=vendor --buildmode=plugin -o ./nakama/modules/$(M_WORLD_CONTROL).so src/$(M_WORLD_CONTROL)/main.go
	@go build --trimpath --mod=vendor --buildmode=plugin -o ./nakama/modules/$(M_WORLD_RPC).so src/$(M_WORLD_RPC)/main.go

# Clean the compiled binary
clean:
	@echo "Cleaning..."
	@rm -f nakama/modules/$(M_WORLD_CONTROL).so
	@rm -f nakama/modules/$(M_WORLD_RPC).so

# Build modules and start the containers
up:
	@make build
	@docker compose up -d postgres nakama

# Stop the containers
down:
	@docker compose down

# Stop the containers and delete database
purge:
	@docker compose down
	@make clean
	@docker volume rm $(PG_VOLUME)

.PHONY: all build clean up down purge
