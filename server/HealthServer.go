package server

import (
	"context"
	service "google.golang.org/grpc/health/grpc_health_v1"
)

type HealthServer struct {
	
}

func (s HealthServer) Check(ctx context.Context, req *service.HealthCheckRequest) (*service.HealthCheckResponse, error) {
	return &service.HealthCheckResponse{
		Status: service.HealthCheckResponse_SERVING,
	}, nil
}

func (s HealthServer) Watch(req *service.HealthCheckRequest, server service.Health_WatchServer) error  {
	//log.Printf("service watch:%s", req.Service)
	r := &service.HealthCheckResponse{
		Status: service.HealthCheckResponse_SERVING,
	}
	for {
		server.Send(r)
	}
	return nil
}
