<!-- buat readme bagus untuk projek ini -->
# Go gRPC Simple Project

This project demonstrates a simple implementation of gRPC in Go. It includes both the server and client implementations to showcase basic gRPC functionalities.

## Features

- Simple gRPC server and client
- Unary RPC example
- Protocol Buffers for message definition
- Easy to understand and extend

## Prerequisites

- Go 1.16 or higher
- `protoc` Protocol Buffers compiler
- `protoc-gen-go` plugin

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/go-grpc-simple.git
   cd go-grpc-simple
   ```

2. Install Protocol Buffers compiler:
   ```sh
   # For Ubuntu
   sudo apt-get install -y protobuf-compiler

   # For MacOS
   brew install protobuf
   ```

3. Install `protoc-gen-go` plugin:
   ```sh
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

## Usage

1. Generate the gRPC code from the `.proto` file:
   ```sh
   protoc --go_out=. --go-grpc_out=. proto/helloworld.proto
   ```

2. Run the server:
   ```sh
   go run server/main.go
   ```

3. In a new terminal, run the client:
   ```sh
   go run client/main.go
   ```

## Project Structure

- `proto/`: Contains the Protocol Buffers definitions.
- `server/`: Contains the gRPC server implementation.
- `client/`: Contains the gRPC client implementation.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any changes.

## Acknowledgements

- [gRPC](https://grpc.io/)
- [Protocol Buffers](https://developers.google.com/protocol-buffers)
