#!/bin/bash

# gowatch 文档：https://studygolang.com/articles/26039
# 支持的命令行参数：
#   -o : 非必须，指定build的目标文件路径
#   -p : 非必须，指定需要build的package（也可以是单个文件）
#   -args: 非必须，指定程序运行时参数，例如：-args='-host=:8080,-name=demo'
#   -v: 非必须，显示gowatch版本信息
gowatch -o bin/kit-scaffold.palm.grpc.api.dev -p main.go -args='--server.host=0.0.0.0,--server.port=8080,--server.gateway.host=0.0.0.0,--server.gateway.port=8081'