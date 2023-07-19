## grpc docs
![grpcDocs](https://grpc.io/docs/languages/go/quickstart/)

1. go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
2. go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
3. export PATH="$PATH:$(go env GOPATH)/bin"
4. protoc --go_out=. --go-grpc_out=. proto/greet.proto
    - proto要對應在greet.proto中定義的位置
    ```proto3
    option go_package = "./proto";
    ```
