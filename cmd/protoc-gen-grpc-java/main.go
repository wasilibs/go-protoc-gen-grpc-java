package main

import (
	"os"

	"github.com/wasilibs/go-protoc-gen-grpc-java/internal/runner"
	"github.com/wasilibs/go-protoc-gen-grpc-java/internal/wasm"
)

func main() {
	os.Exit(runner.Run("protoc-gen-swift", os.Args[1:], wasm.ProtocGenGRPCJava, os.Stdin, os.Stdout, os.Stderr))
}
