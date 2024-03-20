# gRPC-Gateway REST API Example
This project demonstrates the use of gRPC-Gateway to serve both gRPC and REST clients from a single application, adhering to the Google APIs standards for interface definitions. 
It leverages the gRPC-Gateway plugin to generate a reverse-proxy server that translates a RESTful HTTP API into gRPC. This allows the API to be accessible via both gRPC methods and traditional HTTP REST calls.

## Features

- gRPC server implementation in Go.
- RESTful API gateway to the gRPC server following [Google APIs standards](https://github.com/googleapis/googleapis).
- Example service definition with gRPC-Gateway annotations.
- Makefile for easy project setup and operations.


## Prerequisites

- Go 1.15 or later.
- Protocol Buffer Compiler (protoc) installed on your system.
- GOPATH environment variable set.

## Installation

### Clone the Repository

First, clone this repository to your local machine:

```bash
git clone https://github.com/squaick/grpc-gateway-rest-api-example
cd grpc-gateway-rest-api-example
```

### Install Dependencies
Run the following command to install the necessary dependencies, including the Protocol Buffer compiler plugins for Go, gRPC, and gRPC-Gateway:

```bash
make install-deps
```

### Clone Additional Protobuf Sources
Some protobuf definitions require Google's API common protos. Clone these into your GOPATH as follows:

```bash
make clone-googleapis
```

### Running go mod tidy
Ensure your project's dependencies are clean and up to date by running:

```bash
go mod tidy
```
This command will download the necessary Go modules and remove any unused dependencies.

### Generating protobuffs
To generate the Go bindings for your .proto files and the reverse-proxy server, run:

```bash
make generate
```
This command compiles the protobuf definitions and generates the necessary Go files in the generated_pb directory.


## Running the Project
### Start the gRPC and REST Servers
To start both the gRPC server and the REST gateway, simply run:
```bash
go run server.go
```
This command starts the gRPC server listening on port 3000 and the REST gateway on port 3001.

## Testing the API
You can test the gRPC service using a gRPC client and the REST API using any HTTP client (e.g., curl, Postman).

### gRPC Example
Use a gRPC client to call the service:

```go
// Example gRPC client call (pseudo-code)
client := NewSayHelloServiceClient(conn)
response, err := client.SayHello(context.Background(), &HelloRequest{Name: "World"})
```

or with grpcurl: 

```bash
grpcurl -plaintext -d '{"name": "World"}' localhost:3000 grpcGateway.sayHelloService/sayHello
```

### REST Example
To test the REST API, you can use curl:

```bash
curl http://localhost:3001/v1/hello/World
```
This should return a JSON response from the REST gateway, which in turn communicates with the gRPC server:

```json
{
  "message": "Hello World"
}
```


