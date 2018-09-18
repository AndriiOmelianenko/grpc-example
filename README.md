# grpc-example

gRPC client-server example, with golang modules.

## Build

### install dependencies

```shell
export GO111MODULE=on
brew install protobuf
go get github.com/golang/protobuf/proto
go get github.com/golang/protobuf/protoc-gen-go
go get google.golang.org/grpc
```

### generate grpc go sources

```
protoc \
    -I api/ \
    --go_out=plugins=grpc:api \
    api/api.proto
```

OR

```
make api
```

### build server

```
go build -i -v -o bin/server github.com/AndriiOmelianenko/grpc-example/server
```

OR

```
make build_server
```

### build client

```
go build -i -v -o bin/client github.com/AndriiOmelianenko/grpc-example/client
```

OR

```
make build_client
```

### build client and server

```
make build
```

## Run

### Run server
```
$ ./bin/server
2018/09/18 14:51:25 Listening on :7777
2018/09/18 14:51:28 Receive message hi John
```

### run client
```
$ bin/client
2018/09/18 14:51:28 Response from server: Hello John
```

## makefile help

```shell
make help
```
