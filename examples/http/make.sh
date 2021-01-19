#!/bin/bash

# CGO_ENABLED: 是否开启CGO; 1-开启(默认)、0-关闭(关闭CGO，编译纯静态Go程序，不依赖外部动态链接库，提高可移植性)
# CGO_ENABLED 参考资料: https://johng.cn/cgo-enabled-affect-go-static-compile
#                      https://wiki.jikexueyuan.com/project/go-command-tutorial/0.14.html
# GOOS: 程序构建环境的目标操作系统(Golang 支持交叉编译(注：交叉编译不支持 CGO 所以要禁用它)，在一个平台上生成另一个平台的可执行程序);
# GOOS=linux 编译Linux系统可执行程序 (linux: Linux系统、windows: Windows系统、darwin: Mac系统、freebsd: freebsd)
# GOARCH: 目标平台的体系架构（386、amd64、arm）
# 参数: -x 将运行期间所有实际执行的命令都打印到标准输出
#      -ldflags "-s -w" 减少编译后二进制文件大小，一般能减小20%的大小(注：不建议 -w 和-s 同时使用)
#      -s 去掉符号表信息(注: panic时, stace trace 没有任何文件名/行号信息)
#      -w 去掉DWARF调试信息(注: 程序不能使用gdb调试)
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -x -ldflags "-s -w" -o ./bin/kit-scaffold.palm.http.api.srv

# UPX 压缩
# -o：指定输出的文件名
# -k：保留备份原文件
# -1：最快压缩，共1-9九个级别
# -9：最优压缩，与上面对应
# -d：解压缩decompress，恢复原体积
# -l：显示压缩文件的详情，例如upx -l main.exe
# -t：测试压缩文件，例如upx -t main.exe
# -q：静默压缩be quiet
# -v：显示压缩细节be verbose
# -f：强制压缩
# -V：显示版本号
# -h：显示帮助信息
# --brute：尝试所有可用的压缩方法，slow
# --ultra-brute：比楼上更极端，very slow
upx -9 -v -o ./bin/kit-scaffold.palm.http.api ./bin/kit-scaffold.palm.http.api.srv && rm -rf ./bin/kit-scaffold.palm.http.api.srv