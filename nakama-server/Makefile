NAKAMA_APP := nakama-godot4-extended

M_WORLD_CONTROL = world_control
M_WORLD_RPC = world_rpc

# Build the application
all: build

build:
	@echo "Building..."
	@go build --trimpath --mod=vendor --buildmode=plugin -o ./nakama/modules/$(M_WORLD_CONTROL).so src/$(M_WORLD_CONTROL)/main.go
	@go build --trimpath --mod=vendor --buildmode=plugin -o ./nakama/modules/$(M_WORLD_RPC).so src/$(M_WORLD_RPC)/main.go

# Clean the compiled binary
clean:
	@echo "Cleaning..."
	@rm -f nakama/modules/$(M_WORLD_CONTROL).so
	@rm -f nakama/modules/$(M_WORLD_RPC).so

.PHONY: all build clean
