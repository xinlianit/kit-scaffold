package common

import (
	"context"
	"github.com/google/uuid"
	"github.com/xinlianit/go-util"
	"google.golang.org/grpc/metadata"
)

// Request 请求实例
func Request(ctx context.Context) *request {
	var (
		incomingContext = ctx
		//outgoingContext context.Context
		requestId string
		requestReferer string
	)

	metadataUtil := util.NewMetadataUtil()
	// 解析 metadata 元数据
	md, ok := metadataUtil.ParseMetadata(incomingContext)
	if ok {
		// 从 metadata 获取请求元数据
		requestId = metadataUtil.GetStringValue("X-Request-Id")
		requestReferer = metadataUtil.GetStringValue("X-Request-Referer")
	}

	// metadata 无请求ID，生成请求ID
	if requestId == "" {
		id, _ := uuid.NewRandom()
		requestId = id.String()

		// 请求ID设置到传入请求元数据
		md.Set("X-Request-Id", requestId)
		incomingContext = metadata.NewIncomingContext(ctx, md)
	}

	// 设置 metadata 元数据到传出上下文
	//outgoingContext = metadata.AppendToOutgoingContext(outgoingContext, "X-Request-Id", requestId)

	return &request{
		incomingContext: incomingContext,
		//outgoingContext: outgoingContext,
		md:md,
		requestId: requestId,
		requestReferer: requestReferer,
	}
}

// request 请求结构体
type request struct {
	// incomingContext 请求传入上下文
	incomingContext context.Context
	// outgoingContext 请求传出上下文
	//outgoingContext context.Context
	// 请求元数据
	md metadata.MD
	// 请求ID
	requestId string
	// 请求来源
	requestReferer string
}

// GetContext 获取请求上下文
// @return context.Context 请求上下文
func (r request) GetContext() context.Context {
	return r.incomingContext
}

// GetRequestId 获取请求ID
// @return string 请求ID
func (r request) GetRequestId() string {
	return r.requestId
}

// GetMetadata 获取请求元数据
// @return metadata.MD 请求元数据
func (r request) GetMetadata() metadata.MD {
	return r.md
}

// GetRequestReferer 请求来源
// @return string 请求来源
func (r request) GetRequestReferer() string {
	return r.requestReferer
}
