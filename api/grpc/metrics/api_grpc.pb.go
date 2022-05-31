// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: api/grpc/metrics/api.proto

package api_metrics

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MetricsClient is the client API for Metrics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsClient interface {
	AuditInc(ctx context.Context, opts ...grpc.CallOption) (Metrics_AuditIncClient, error)
	BpfInc(ctx context.Context, opts ...grpc.CallOption) (Metrics_BpfIncClient, error)
}

type metricsClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsClient(cc grpc.ClientConnInterface) MetricsClient {
	return &metricsClient{cc}
}

func (c *metricsClient) AuditInc(ctx context.Context, opts ...grpc.CallOption) (Metrics_AuditIncClient, error) {
	stream, err := c.cc.NewStream(ctx, &Metrics_ServiceDesc.Streams[0], "/api_metrics.Metrics/AuditInc", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsAuditIncClient{stream}
	return x, nil
}

type Metrics_AuditIncClient interface {
	Send(*AuditRequest) error
	CloseAndRecv() (*EmptyResponse, error)
	grpc.ClientStream
}

type metricsAuditIncClient struct {
	grpc.ClientStream
}

func (x *metricsAuditIncClient) Send(m *AuditRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metricsAuditIncClient) CloseAndRecv() (*EmptyResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EmptyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *metricsClient) BpfInc(ctx context.Context, opts ...grpc.CallOption) (Metrics_BpfIncClient, error) {
	stream, err := c.cc.NewStream(ctx, &Metrics_ServiceDesc.Streams[1], "/api_metrics.Metrics/BpfInc", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsBpfIncClient{stream}
	return x, nil
}

type Metrics_BpfIncClient interface {
	Send(*BpfRequest) error
	CloseAndRecv() (*EmptyResponse, error)
	grpc.ClientStream
}

type metricsBpfIncClient struct {
	grpc.ClientStream
}

func (x *metricsBpfIncClient) Send(m *BpfRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metricsBpfIncClient) CloseAndRecv() (*EmptyResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(EmptyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MetricsServer is the server API for Metrics service.
// All implementations must embed UnimplementedMetricsServer
// for forward compatibility
type MetricsServer interface {
	AuditInc(Metrics_AuditIncServer) error
	BpfInc(Metrics_BpfIncServer) error
	mustEmbedUnimplementedMetricsServer()
}

// UnimplementedMetricsServer must be embedded to have forward compatible implementations.
type UnimplementedMetricsServer struct {
}

func (UnimplementedMetricsServer) AuditInc(Metrics_AuditIncServer) error {
	return status.Errorf(codes.Unimplemented, "method AuditInc not implemented")
}
func (UnimplementedMetricsServer) BpfInc(Metrics_BpfIncServer) error {
	return status.Errorf(codes.Unimplemented, "method BpfInc not implemented")
}
func (UnimplementedMetricsServer) mustEmbedUnimplementedMetricsServer() {}

// UnsafeMetricsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsServer will
// result in compilation errors.
type UnsafeMetricsServer interface {
	mustEmbedUnimplementedMetricsServer()
}

func RegisterMetricsServer(s grpc.ServiceRegistrar, srv MetricsServer) {
	s.RegisterService(&Metrics_ServiceDesc, srv)
}

func _Metrics_AuditInc_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetricsServer).AuditInc(&metricsAuditIncServer{stream})
}

type Metrics_AuditIncServer interface {
	SendAndClose(*EmptyResponse) error
	Recv() (*AuditRequest, error)
	grpc.ServerStream
}

type metricsAuditIncServer struct {
	grpc.ServerStream
}

func (x *metricsAuditIncServer) SendAndClose(m *EmptyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metricsAuditIncServer) Recv() (*AuditRequest, error) {
	m := new(AuditRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Metrics_BpfInc_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetricsServer).BpfInc(&metricsBpfIncServer{stream})
}

type Metrics_BpfIncServer interface {
	SendAndClose(*EmptyResponse) error
	Recv() (*BpfRequest, error)
	grpc.ServerStream
}

type metricsBpfIncServer struct {
	grpc.ServerStream
}

func (x *metricsBpfIncServer) SendAndClose(m *EmptyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metricsBpfIncServer) Recv() (*BpfRequest, error) {
	m := new(BpfRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Metrics_ServiceDesc is the grpc.ServiceDesc for Metrics service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Metrics_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api_metrics.Metrics",
	HandlerType: (*MetricsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AuditInc",
			Handler:       _Metrics_AuditInc_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BpfInc",
			Handler:       _Metrics_BpfInc_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api/grpc/metrics/api.proto",
}
