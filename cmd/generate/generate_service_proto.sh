#!/bin/bash

if [ "$#" -ne 2 ]; then
    echo "用法: $0 <proto文件名> <输出目录>"
	echo "例如: $0 xxx/xxx/xxx.proto xxx/xx/xxx"
    exit 1
fi

PROTO_FILE=$1
OUTPUT_DIR=$2

echo $OUTPUT_DIR
mkdir -p "$GOPATH/src/yazl-tech/ai-bot/proto/$OUTPUT_DIR"

protoc \
	-I $GOPATH/src/github.com/yazl-tech/ai-bot/internal/proto/service \
	-I $GOPATH/src/github.com/yazl-tech/ai-bot/internal/proto \
	-I $GOPATH/src \
	-I $GOPATH/src/github.com/gogo/protobuf/protobuf \
	--go-grpc_out=../../proto/$OUTPUT_DIR \
	--go-grpc_opt=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor,paths=source_relative \
	$PROTO_FILE 

echo "$PROTO_FILE generate success"
