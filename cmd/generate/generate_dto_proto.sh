#!/bin/bash

######################################################################
# @author      : hoven (hoven@$HOSTNAME)
# @file        : generate_dto_proto
# @created     : Thursday Jan 09, 2025 22:48:34 CST
#
# @description : 
######################################################################

if [ "$#" -ne 1 ]; then
    echo "用法: $0 <proto文件名>"
    echo "例如: $0 xxx/xxx/xxx.proto"
    exit 1
fi

PROTO_FILE=$1

protoc \
	-I $GOPATH/src/github.com/yazl-tech/ai-bot/internal/proto \
	-I $GOPATH/src \
	-I $GOPATH/src/github.com/gogo/protobuf/protobuf \
	--gofast_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor,paths=source_relative:../../pkg \
	$PROTO_FILE

echo "$PROTO_FILE generate success"
