// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: apps/stats/rpc/desc/proto/stats.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Stats_SortByHot_FullMethodName = "/analytics.stats/SortByHot"
)

// StatsClient is the client API for Stats service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatsClient interface {
	// 按热度排序, 返回排序完成的 anime 列表
	SortByHot(ctx context.Context, in *SortByHotReq, opts ...grpc.CallOption) (*SortByHotResp, error)
}

type statsClient struct {
	cc grpc.ClientConnInterface
}

func NewStatsClient(cc grpc.ClientConnInterface) StatsClient {
	return &statsClient{cc}
}

func (c *statsClient) SortByHot(ctx context.Context, in *SortByHotReq, opts ...grpc.CallOption) (*SortByHotResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SortByHotResp)
	err := c.cc.Invoke(ctx, Stats_SortByHot_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatsServer is the server API for Stats service.
// All implementations must embed UnimplementedStatsServer
// for forward compatibility.
type StatsServer interface {
	// 按热度排序, 返回排序完成的 anime 列表
	SortByHot(context.Context, *SortByHotReq) (*SortByHotResp, error)
	mustEmbedUnimplementedStatsServer()
}

// UnimplementedStatsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStatsServer struct{}

func (UnimplementedStatsServer) SortByHot(context.Context, *SortByHotReq) (*SortByHotResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SortByHot not implemented")
}
func (UnimplementedStatsServer) mustEmbedUnimplementedStatsServer() {}
func (UnimplementedStatsServer) testEmbeddedByValue()               {}

// UnsafeStatsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatsServer will
// result in compilation errors.
type UnsafeStatsServer interface {
	mustEmbedUnimplementedStatsServer()
}

func RegisterStatsServer(s grpc.ServiceRegistrar, srv StatsServer) {
	// If the following call pancis, it indicates UnimplementedStatsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Stats_ServiceDesc, srv)
}

func _Stats_SortByHot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SortByHotReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).SortByHot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Stats_SortByHot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).SortByHot(ctx, req.(*SortByHotReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Stats_ServiceDesc is the grpc.ServiceDesc for Stats service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Stats_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "analytics.stats",
	HandlerType: (*StatsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SortByHot",
			Handler:    _Stats_SortByHot_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps/stats/rpc/desc/proto/stats.proto",
}
