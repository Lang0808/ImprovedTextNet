// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: NewsFeedService.proto

package GrpcNewsFeedService

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

const (
	NewsFeedService_GetNewsFeed_FullMethodName = "/NewsFeedService/GetNewsFeed"
)

// NewsFeedServiceClient is the client API for NewsFeedService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewsFeedServiceClient interface {
	GetNewsFeed(ctx context.Context, in *GetNewsFeedRequest, opts ...grpc.CallOption) (*GetNewsFeedResponse, error)
}

type newsFeedServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNewsFeedServiceClient(cc grpc.ClientConnInterface) NewsFeedServiceClient {
	return &newsFeedServiceClient{cc}
}

func (c *newsFeedServiceClient) GetNewsFeed(ctx context.Context, in *GetNewsFeedRequest, opts ...grpc.CallOption) (*GetNewsFeedResponse, error) {
	out := new(GetNewsFeedResponse)
	err := c.cc.Invoke(ctx, NewsFeedService_GetNewsFeed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewsFeedServiceServer is the server API for NewsFeedService service.
// All implementations must embed UnimplementedNewsFeedServiceServer
// for forward compatibility
type NewsFeedServiceServer interface {
	GetNewsFeed(context.Context, *GetNewsFeedRequest) (*GetNewsFeedResponse, error)
	mustEmbedUnimplementedNewsFeedServiceServer()
}

// UnimplementedNewsFeedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNewsFeedServiceServer struct {
}

func (UnimplementedNewsFeedServiceServer) GetNewsFeed(context.Context, *GetNewsFeedRequest) (*GetNewsFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewsFeed not implemented")
}
func (UnimplementedNewsFeedServiceServer) mustEmbedUnimplementedNewsFeedServiceServer() {}

// UnsafeNewsFeedServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewsFeedServiceServer will
// result in compilation errors.
type UnsafeNewsFeedServiceServer interface {
	mustEmbedUnimplementedNewsFeedServiceServer()
}

func RegisterNewsFeedServiceServer(s grpc.ServiceRegistrar, srv NewsFeedServiceServer) {
	s.RegisterService(&NewsFeedService_ServiceDesc, srv)
}

func _NewsFeedService_GetNewsFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewsFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsFeedServiceServer).GetNewsFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsFeedService_GetNewsFeed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsFeedServiceServer).GetNewsFeed(ctx, req.(*GetNewsFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NewsFeedService_ServiceDesc is the grpc.ServiceDesc for NewsFeedService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NewsFeedService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "NewsFeedService",
	HandlerType: (*NewsFeedServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNewsFeed",
			Handler:    _NewsFeedService_GetNewsFeed_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "NewsFeedService.proto",
}
