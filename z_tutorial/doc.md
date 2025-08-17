https://whimsical.com/graphql-grpc-go-microservice-LdA8wTyHe3pUaUnEdH99cj



account service
catalog service
order

// install protobuf and it's plugin
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest


and to account.proto folder and run protoc --go_out=./pb --go-grpc_out=./pb account.proto
