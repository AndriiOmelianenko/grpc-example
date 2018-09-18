SERVER_OUT := "bin/server"
CLIENT_OUT := "bin/client"
API_OUT := "api/api.pb.go"
PKG := "github.com/AndriiOmelianenko/grpc-example"
SERVER_PKG_BUILD := "${PKG}/server"
CLIENT_PKG_BUILD := "${PKG}/client"

.PHONY: build api build_server build_client

api/api.pb.go: api/api.proto
	@protoc -I api/ \
		--go_out=plugins=grpc:api \
		api/api.proto

api: api/api.pb.go ## Auto-generate grpc go sources

mod: ## Get the dependencies
	@go mod tidy -v

build_server: mod api ## Build the binary file for server
	@go build -i -v -o $(SERVER_OUT) $(SERVER_PKG_BUILD)

build_client: mod api ## Build the binary file for client
	@go build -i -v -o $(CLIENT_OUT) $(CLIENT_PKG_BUILD)

build: build_server build_client ## Build the binary file for server and client

clean: ## Remove previous builds
	@rm -v $(SERVER_OUT) $(CLIENT_OUT) $(API_OUT)

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
