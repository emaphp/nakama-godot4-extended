NAKAMA_MODULE := nakama-godot4-extended

# Build the application
all: build

build:
	@echo "Building..."
	@go build --trimpath --mod=vendor --buildmode=plugin -o ./nakama/modules/$(NAKAMA_MODULE).so src/$(NAKAMA_MODULE)/main.go

# Clean the compiled binary
clean:
	@echo "Cleaning..."
	@rm -f nakama/modules/$(NAKAMA_MODULE).so

.PHONY: all build clean
