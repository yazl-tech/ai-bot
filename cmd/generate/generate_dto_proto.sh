#!/bin/bash

######################################################################
# @author      : hoven (hoven@$HOSTNAME)
# @file        : generate_dto_proto
# @created     : Thursday Jan 09, 2025 22:48:34 CST
#
# @description : 
######################################################################

if [ "$#" -ne 2 ]; then
    echo "用法: $0 <proto文件名>"
    echo "例如: $0 xxx/xxx/xxx.proto"
    exit 1
fi

PROTO_FILE=$1
OUTPUT_DIR=$2

mkdir -p "$GOPATH/src/yazl-tech/ai-bot/pkg/dto/$OUTPUT_DIR"

protoc \
	-I $GOPATH/src/github.com/yazl-tech/ai-bot/internal/proto/dto \
	-I $GOPATH/src \
	-I $GOPATH/src/github.com/gogo/protobuf/protobuf \
	--gofast_out=paths=source_relative:../../pkg/dto/$OUTPUT_DIR \
	$PROTO_FILE

echo "$PROTO_FILE generate success"
