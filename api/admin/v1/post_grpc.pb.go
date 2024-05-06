// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: post.proto

package v1

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
	SysPost_ListPost_FullMethodName   = "/api.admin.v1.SysPost/ListPost"
	SysPost_CreatePost_FullMethodName = "/api.admin.v1.SysPost/CreatePost"
	SysPost_UpdatePost_FullMethodName = "/api.admin.v1.SysPost/UpdatePost"
	SysPost_DeletePost_FullMethodName = "/api.admin.v1.SysPost/DeletePost"
)

// SysPostClient is the client API for SysPost service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysPostClient interface {
	// 岗位列表
	ListPost(ctx context.Context, in *ListPostRequest, opts ...grpc.CallOption) (*ListPostReply, error)
	// 创建岗位
	CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostReply, error)
	// 更新岗位
	UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*UpdatePostReply, error)
	// 删除岗位
	DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostReply, error)
}

type sysPostClient struct {
	cc grpc.ClientConnInterface
}

func NewSysPostClient(cc grpc.ClientConnInterface) SysPostClient {
	return &sysPostClient{cc}
}

func (c *sysPostClient) ListPost(ctx context.Context, in *ListPostRequest, opts ...grpc.CallOption) (*ListPostReply, error) {
	out := new(ListPostReply)
	err := c.cc.Invoke(ctx, SysPost_ListPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysPostClient) CreatePost(ctx context.Context, in *CreatePostRequest, opts ...grpc.CallOption) (*CreatePostReply, error) {
	out := new(CreatePostReply)
	err := c.cc.Invoke(ctx, SysPost_CreatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysPostClient) UpdatePost(ctx context.Context, in *UpdatePostRequest, opts ...grpc.CallOption) (*UpdatePostReply, error) {
	out := new(UpdatePostReply)
	err := c.cc.Invoke(ctx, SysPost_UpdatePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysPostClient) DeletePost(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostReply, error) {
	out := new(DeletePostReply)
	err := c.cc.Invoke(ctx, SysPost_DeletePost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysPostServer is the server API for SysPost service.
// All implementations must embed UnimplementedSysPostServer
// for forward compatibility
type SysPostServer interface {
	// 岗位列表
	ListPost(context.Context, *ListPostRequest) (*ListPostReply, error)
	// 创建岗位
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostReply, error)
	// 更新岗位
	UpdatePost(context.Context, *UpdatePostRequest) (*UpdatePostReply, error)
	// 删除岗位
	DeletePost(context.Context, *DeletePostRequest) (*DeletePostReply, error)
	mustEmbedUnimplementedSysPostServer()
}

// UnimplementedSysPostServer must be embedded to have forward compatible implementations.
type UnimplementedSysPostServer struct {
}

func (UnimplementedSysPostServer) ListPost(context.Context, *ListPostRequest) (*ListPostReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPost not implemented")
}
func (UnimplementedSysPostServer) CreatePost(context.Context, *CreatePostRequest) (*CreatePostReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePost not implemented")
}
func (UnimplementedSysPostServer) UpdatePost(context.Context, *UpdatePostRequest) (*UpdatePostReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePost not implemented")
}
func (UnimplementedSysPostServer) DeletePost(context.Context, *DeletePostRequest) (*DeletePostReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePost not implemented")
}
func (UnimplementedSysPostServer) mustEmbedUnimplementedSysPostServer() {}

// UnsafeSysPostServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysPostServer will
// result in compilation errors.
type UnsafeSysPostServer interface {
	mustEmbedUnimplementedSysPostServer()
}

func RegisterSysPostServer(s grpc.ServiceRegistrar, srv SysPostServer) {
	s.RegisterService(&SysPost_ServiceDesc, srv)
}

func _SysPost_ListPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysPostServer).ListPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysPost_ListPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysPostServer).ListPost(ctx, req.(*ListPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysPost_CreatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysPostServer).CreatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysPost_CreatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysPostServer).CreatePost(ctx, req.(*CreatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysPost_UpdatePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysPostServer).UpdatePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysPost_UpdatePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysPostServer).UpdatePost(ctx, req.(*UpdatePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysPost_DeletePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysPostServer).DeletePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysPost_DeletePost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysPostServer).DeletePost(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SysPost_ServiceDesc is the grpc.ServiceDesc for SysPost service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SysPost_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.admin.v1.SysPost",
	HandlerType: (*SysPostServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPost",
			Handler:    _SysPost_ListPost_Handler,
		},
		{
			MethodName: "CreatePost",
			Handler:    _SysPost_CreatePost_Handler,
		},
		{
			MethodName: "UpdatePost",
			Handler:    _SysPost_UpdatePost_Handler,
		},
		{
			MethodName: "DeletePost",
			Handler:    _SysPost_DeletePost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "post.proto",
}
