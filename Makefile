
# Install protoc plugins agnostic to OS
install-deps:
	@echo "Installing protoc & deps..."
	@echo "If protoc is not installed on your system, checkout https://grpc.io/docs/protoc-installation/ for installation instructions"
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

# Clone googleapis and protobuf required for generating google api standard restful services
clone-googleapis:
	cd ${GOPATH}/src && git clone https://github.com/protocolbuffers/protobuf/ && git clone https://github.com/googleapis/api-common-protos

# Generate protobuffs in Go and gRPC
generate:
	@echo "Generating protobuffs..."
	protoc \
	--proto_path=${GOPATH}/src:. \
	--proto_path=${GOPATH}/src/api-common-protos:. \
	--proto_path=.  *.proto --go_out=. --go-grpc_out=. --grpc-gateway_out=. --openapiv2_out=.
