#--------------------------------- 构建阶段 ---------------------------------
# 基础镜像
FROM golang:alpine AS builder

###### 设置环境变量
# 应用名称
ARG APP_NAME=go-test-for-k8s
# 应用目录
ARG APP_DIR=/home/data/webroot/${APP_NAME}
# 应用日志目录
ARG APP_LOG_DIR=/home/data/logs/${APP_NAME}
# GOPROXY
ENV GOPROXY https://goproxy.cn,https://goproxy.io,https://mirrors.aliyun.com/goproxy/,direct

# 运行命令; 创建应用及应用日志目录
# 匹配 "dl-cdn.alpinelinux.org" 替换 "mirrors.aliyun.com" 阿里仓库源
# 使用 apk 包管理，禁用缓存，添加依赖: ca-certificates、tzdata、upx
RUN mkdir -p ${APP_DIR} ${APP_LOG_DIR} && \
    sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk --no-cache add ca-certificates tzdata upx

# 拷贝当前目录文件到项目目录
COPY . ${APP_DIR}

# 工作目录
WORKDIR ${APP_DIR}

# 编译项目
# 使用 upx 压缩二进制包
# 重命名压缩后的二进制包
# 编译项目
#RUN go build -o ${APP_NAME} .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -x -ldflags "-s -w" -o ${APP_NAME} . && \
    upx -9 -v -o ${APP_DIR}/_upx_${APP_NAME} ${APP_DIR}/${APP_NAME} && \
    mv -f ${APP_DIR}/_upx_${APP_NAME} ${APP_DIR}/${APP_NAME}

#--------------------------------- 运行阶段 ---------------------------------
# 基础镜像
FROM alpine AS runner

###### 设置环境变量
# 应用名称
ARG APP_NAME=go-test-for-k8s
# 应用目录
ARG APP_DIR=/home/data/webroot/${APP_NAME}
# 应用日志目录
ARG APP_LOG_DIR=/home/data/logs/${APP_NAME}

# 运行命令; 创建应用及应用日志目录
RUN mkdir -p ${APP_DIR} ${APP_LOG_DIR}

# 工作目录
WORKDIR ${APP_DIR}

# 从 builder 镜像中同步较准时间
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
# 从 builder 镜像中拷贝证书文件
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
# 从 builder 镜像中拷贝编译的二进制文件
COPY --from=builder ${APP_DIR}/${APP_NAME} ${APP_DIR}

# 运行应用
CMD ["/home/data/webroot/go-test-for-k8s/go-test-for-k8s"]
#CMD $APP_DIR/$APP_NAME