package server

import (
	"context"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type HealthServer struct {
	
}

func (s HealthServer) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (s HealthServer) Watch(req *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error  {
	//log.Printf("service watch:%s", req.Service)
	r := &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}
	for {
		server.Send(r)
	}
	return nil
}
