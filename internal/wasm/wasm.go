package wasm

import _ "embed"

//go:embed protoc-gen-grpc-java.wasm
var ProtocGenGRPCJava []byte

//go:embed memory.wasm
var Memory []byte
