// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: RelationshipService.proto

package GrpcRelationshipService

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

// RelationshipServiceClient is the client API for RelationshipService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RelationshipServiceClient interface {
	SendFriendRequest(ctx context.Context, in *SendFriendRequestRequest, opts ...grpc.CallOption) (*SendFriendRequestResponse, error)
	AcceptFriendRequest(ctx context.Context, in *AcceptFriendRequestRequest, opts ...grpc.CallOption) (*AcceptFriendRequestResponse, error)
	RemoveFriendRequest(ctx context.Context, in *RemoveFriendRequestRequest, opts ...grpc.CallOption) (*RemoveFriendRequestResponse, error)
	BlockUser(ctx context.Context, in *BlockUserRequest, opts ...grpc.CallOption) (*BlockUserResponse, error)
	UnblockUser(ctx context.Context, in *UnblockUserRequest, opts ...grpc.CallOption) (*UnblockUserResponse, error)
	FollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*FollowUserResponse, error)
	UnfollowUser(ctx context.Context, in *UnfollowUserRequest, opts ...grpc.CallOption) (*UnfollowUserResponse, error)
	UnfriendUser(ctx context.Context, in *UnfriendUserRequest, opts ...grpc.CallOption) (*UnfriendUserResponse, error)
	GetRelationship(ctx context.Context, in *GetRelationshipRequest, opts ...grpc.CallOption) (*GetRelationshipResponse, error)
}

type relationshipServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRelationshipServiceClient(cc grpc.ClientConnInterface) RelationshipServiceClient {
	return &relationshipServiceClient{cc}
}

func (c *relationshipServiceClient) SendFriendRequest(ctx context.Context, in *SendFriendRequestRequest, opts ...grpc.CallOption) (*SendFriendRequestResponse, error) {
	out := new(SendFriendRequestResponse)
	err := c.cc.Invoke(ctx, "/RelationshipService/SendFriendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipServiceClient) AcceptFriendRequest(ctx context.Context, in *AcceptFriendRequestRequest, opts ...grpc.CallOption) (*AcceptFriendRequestResponse, error) {
	out := new(AcceptFriendRequestResponse)
	err := c.cc.Invoke(ctx, "/RelationshipService/AcceptFriendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipServiceClient) RemoveFriendRequest(ctx context.Context, in *RemoveFriendRequestRequest, opts ...grpc.CallOption) (*RemoveFriendRequestResponse, error) {
	out := new(RemoveFriendRequestResponse)
	err := c.cc.Invoke(ctx, "/RelationshipService/RemoveFriendRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipServiceClient) BlockUser(ctx context.Context, in *BlockUserRequest, opts ...grpc.CallOption) (*BlockUserResponse, error) {
	out := new(BlockUserResponse)
	err := c.cc.Invoke(ctx, "/RelationshipService/BlockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipServiceClient) UnblockUser(ctx context.Context, in *UnblockUserRequest, opts ...grpc.CallOption) (*UnblockUserResponse, error) {
	out := new(UnblockUserResponse)
	err := c.cc.Invoke(ctx, "/RelationshipService/UnblockUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipServiceClient) FollowUser(ctx context.Context, in *FollowUserRequest, opts ...grpc.CallOption) (*FollowUserResponse, error) {
	out := new(FollowUserResponse)
	err := c.cc.Invoke(ctx, "/RelationshipService/FollowUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipServiceClient) UnfollowUser(ctx context.Context, in *UnfollowUserRequest, opts ...grpc.CallOption) (*UnfollowUserResponse, error) {
	out := new(UnfollowUserResponse)
	err := c.cc.Invoke(ctx, "/RelationshipService/UnfollowUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipServiceClient) UnfriendUser(ctx context.Context, in *UnfriendUserRequest, opts ...grpc.CallOption) (*UnfriendUserResponse, error) {
	out := new(UnfriendUserResponse)
	err := c.cc.Invoke(ctx, "/RelationshipService/UnfriendUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *relationshipServiceClient) GetRelationship(ctx context.Context, in *GetRelationshipRequest, opts ...grpc.CallOption) (*GetRelationshipResponse, error) {
	out := new(GetRelationshipResponse)
	err := c.cc.Invoke(ctx, "/RelationshipService/GetRelationship", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RelationshipServiceServer is the server API for RelationshipService service.
// All implementations must embed UnimplementedRelationshipServiceServer
// for forward compatibility
type RelationshipServiceServer interface {
	SendFriendRequest(context.Context, *SendFriendRequestRequest) (*SendFriendRequestResponse, error)
	AcceptFriendRequest(context.Context, *AcceptFriendRequestRequest) (*AcceptFriendRequestResponse, error)
	RemoveFriendRequest(context.Context, *RemoveFriendRequestRequest) (*RemoveFriendRequestResponse, error)
	BlockUser(context.Context, *BlockUserRequest) (*BlockUserResponse, error)
	UnblockUser(context.Context, *UnblockUserRequest) (*UnblockUserResponse, error)
	FollowUser(context.Context, *FollowUserRequest) (*FollowUserResponse, error)
	UnfollowUser(context.Context, *UnfollowUserRequest) (*UnfollowUserResponse, error)
	UnfriendUser(context.Context, *UnfriendUserRequest) (*UnfriendUserResponse, error)
	GetRelationship(context.Context, *GetRelationshipRequest) (*GetRelationshipResponse, error)
	mustEmbedUnimplementedRelationshipServiceServer()
}

// UnimplementedRelationshipServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRelationshipServiceServer struct {
}

func (UnimplementedRelationshipServiceServer) SendFriendRequest(context.Context, *SendFriendRequestRequest) (*SendFriendRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendFriendRequest not implemented")
}
func (UnimplementedRelationshipServiceServer) AcceptFriendRequest(context.Context, *AcceptFriendRequestRequest) (*AcceptFriendRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptFriendRequest not implemented")
}
func (UnimplementedRelationshipServiceServer) RemoveFriendRequest(context.Context, *RemoveFriendRequestRequest) (*RemoveFriendRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFriendRequest not implemented")
}
func (UnimplementedRelationshipServiceServer) BlockUser(context.Context, *BlockUserRequest) (*BlockUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockUser not implemented")
}
func (UnimplementedRelationshipServiceServer) UnblockUser(context.Context, *UnblockUserRequest) (*UnblockUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnblockUser not implemented")
}
func (UnimplementedRelationshipServiceServer) FollowUser(context.Context, *FollowUserRequest) (*FollowUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowUser not implemented")
}
func (UnimplementedRelationshipServiceServer) UnfollowUser(context.Context, *UnfollowUserRequest) (*UnfollowUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnfollowUser not implemented")
}
func (UnimplementedRelationshipServiceServer) UnfriendUser(context.Context, *UnfriendUserRequest) (*UnfriendUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnfriendUser not implemented")
}
func (UnimplementedRelationshipServiceServer) GetRelationship(context.Context, *GetRelationshipRequest) (*GetRelationshipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelationship not implemented")
}
func (UnimplementedRelationshipServiceServer) mustEmbedUnimplementedRelationshipServiceServer() {}

// UnsafeRelationshipServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RelationshipServiceServer will
// result in compilation errors.
type UnsafeRelationshipServiceServer interface {
	mustEmbedUnimplementedRelationshipServiceServer()
}

func RegisterRelationshipServiceServer(s grpc.ServiceRegistrar, srv RelationshipServiceServer) {
	s.RegisterService(&RelationshipService_ServiceDesc, srv)
}

func _RelationshipService_SendFriendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendFriendRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServiceServer).SendFriendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RelationshipService/SendFriendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServiceServer).SendFriendRequest(ctx, req.(*SendFriendRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationshipService_AcceptFriendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptFriendRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServiceServer).AcceptFriendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RelationshipService/AcceptFriendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServiceServer).AcceptFriendRequest(ctx, req.(*AcceptFriendRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationshipService_RemoveFriendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFriendRequestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServiceServer).RemoveFriendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RelationshipService/RemoveFriendRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServiceServer).RemoveFriendRequest(ctx, req.(*RemoveFriendRequestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationshipService_BlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServiceServer).BlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RelationshipService/BlockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServiceServer).BlockUser(ctx, req.(*BlockUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationshipService_UnblockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnblockUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServiceServer).UnblockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RelationshipService/UnblockUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServiceServer).UnblockUser(ctx, req.(*UnblockUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationshipService_FollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServiceServer).FollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RelationshipService/FollowUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServiceServer).FollowUser(ctx, req.(*FollowUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationshipService_UnfollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnfollowUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServiceServer).UnfollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RelationshipService/UnfollowUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServiceServer).UnfollowUser(ctx, req.(*UnfollowUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationshipService_UnfriendUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnfriendUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServiceServer).UnfriendUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RelationshipService/UnfriendUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServiceServer).UnfriendUser(ctx, req.(*UnfriendUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RelationshipService_GetRelationship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRelationshipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RelationshipServiceServer).GetRelationship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RelationshipService/GetRelationship",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RelationshipServiceServer).GetRelationship(ctx, req.(*GetRelationshipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RelationshipService_ServiceDesc is the grpc.ServiceDesc for RelationshipService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RelationshipService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RelationshipService",
	HandlerType: (*RelationshipServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendFriendRequest",
			Handler:    _RelationshipService_SendFriendRequest_Handler,
		},
		{
			MethodName: "AcceptFriendRequest",
			Handler:    _RelationshipService_AcceptFriendRequest_Handler,
		},
		{
			MethodName: "RemoveFriendRequest",
			Handler:    _RelationshipService_RemoveFriendRequest_Handler,
		},
		{
			MethodName: "BlockUser",
			Handler:    _RelationshipService_BlockUser_Handler,
		},
		{
			MethodName: "UnblockUser",
			Handler:    _RelationshipService_UnblockUser_Handler,
		},
		{
			MethodName: "FollowUser",
			Handler:    _RelationshipService_FollowUser_Handler,
		},
		{
			MethodName: "UnfollowUser",
			Handler:    _RelationshipService_UnfollowUser_Handler,
		},
		{
			MethodName: "UnfriendUser",
			Handler:    _RelationshipService_UnfriendUser_Handler,
		},
		{
			MethodName: "GetRelationship",
			Handler:    _RelationshipService_GetRelationship_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "RelationshipService.proto",
}