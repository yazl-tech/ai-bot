#!/usr/bin/env sh

######################################################################
# @author      : hoven (hoven@$HOSTNAME)
# @file        : generate_proto
# @created     : Sunday Feb 16, 2025 23:35:32 CST
#
# @description : 
######################################################################

if [ "$#" -ne 2 ]; then
    echo "用法: $0 <proto文件名> <输出目录>"
	echo "例如: $0 xxx/xxx/xxx.proto xxx/xx/xxx"
    exit 1
fi


PROTO_FILE=$1
OUTPUT_DIR=$2

mkdir -p "$GOPATH/src/github.com/yazl-tech/ai-bot/pkg/proto/$OUTPUT_DIR"


protoc \
	-I ../../internal/proto \
	--go_out=paths=source_relative:../../pkg/proto/$OUTPUT_DIR \
	--go-grpc_out=paths=source_relative:../../pkg/proto/$OUTPUT_DIR \
	$PROTO_FILE	

echo "generate success"

