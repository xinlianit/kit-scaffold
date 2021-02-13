#!/bin/bash

# 清除编译文件
if [ -d "../pb" ]; then
  rm -rf ../pb/* && echo "清除编译文件: ../pb/*"
fi

# 创建编译目录
if [ ! -d "../pb" ]; then
  mkdir ../pb && echo "创建编译目录: ../pb"
fi

# 编译谷歌API
protoc --proto_path=. ./google/*/*.proto \
  --go_opt paths=source_relative \
  --go_out=plugins=grpc:../pb

# 编译传输协议
protoc --proto_path=. ./transport/*/*.proto \
  --go_opt paths=source_relative \
  --go_out=plugins=grpc:../pb

# 编译服务协议
protoc --proto_path=. ./service/*.proto \
  --go_opt paths=source_relative \
  --go_out=plugins=grpc:../pb \
  --grpc-gateway_opt logtostderr=true \
  --grpc-gateway_opt paths=source_relative \
  --grpc-gateway_out ../pb
