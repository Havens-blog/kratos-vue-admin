// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: dictdata.proto

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
	DictData_ListDictData_FullMethodName   = "/api.admin.v1.DictData/ListDictData"
	DictData_CreateDictData_FullMethodName = "/api.admin.v1.DictData/CreateDictData"
	DictData_UpdateDictData_FullMethodName = "/api.admin.v1.DictData/UpdateDictData"
	DictData_DeleteDictData_FullMethodName = "/api.admin.v1.DictData/DeleteDictData"
	DictData_GetDictData_FullMethodName    = "/api.admin.v1.DictData/GetDictData"
)

// DictDataClient is the client API for DictData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DictDataClient interface {
	// 列表信息
	ListDictData(ctx context.Context, in *ListDictDataRequest, opts ...grpc.CallOption) (*ListDictDataReply, error)
	// 创建
	CreateDictData(ctx context.Context, in *CreateDictDataRequest, opts ...grpc.CallOption) (*CreateDictDataReply, error)
	// 更新
	UpdateDictData(ctx context.Context, in *UpdateDictDataRequest, opts ...grpc.CallOption) (*UpdateDictDataReply, error)
	// 删除
	DeleteDictData(ctx context.Context, in *DeleteDictDataRequest, opts ...grpc.CallOption) (*DeleteDictDataReply, error)
	// 获取信息
	GetDictData(ctx context.Context, in *GetDictDataRequest, opts ...grpc.CallOption) (*GetDictDataReply, error)
}

type dictDataClient struct {
	cc grpc.ClientConnInterface
}

func NewDictDataClient(cc grpc.ClientConnInterface) DictDataClient {
	return &dictDataClient{cc}
}

func (c *dictDataClient) ListDictData(ctx context.Context, in *ListDictDataRequest, opts ...grpc.CallOption) (*ListDictDataReply, error) {
	out := new(ListDictDataReply)
	err := c.cc.Invoke(ctx, DictData_ListDictData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictDataClient) CreateDictData(ctx context.Context, in *CreateDictDataRequest, opts ...grpc.CallOption) (*CreateDictDataReply, error) {
	out := new(CreateDictDataReply)
	err := c.cc.Invoke(ctx, DictData_CreateDictData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictDataClient) UpdateDictData(ctx context.Context, in *UpdateDictDataRequest, opts ...grpc.CallOption) (*UpdateDictDataReply, error) {
	out := new(UpdateDictDataReply)
	err := c.cc.Invoke(ctx, DictData_UpdateDictData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictDataClient) DeleteDictData(ctx context.Context, in *DeleteDictDataRequest, opts ...grpc.CallOption) (*DeleteDictDataReply, error) {
	out := new(DeleteDictDataReply)
	err := c.cc.Invoke(ctx, DictData_DeleteDictData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dictDataClient) GetDictData(ctx context.Context, in *GetDictDataRequest, opts ...grpc.CallOption) (*GetDictDataReply, error) {
	out := new(GetDictDataReply)
	err := c.cc.Invoke(ctx, DictData_GetDictData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DictDataServer is the server API for DictData service.
// All implementations must embed UnimplementedDictDataServer
// for forward compatibility
type DictDataServer interface {
	// 列表信息
	ListDictData(context.Context, *ListDictDataRequest) (*ListDictDataReply, error)
	// 创建
	CreateDictData(context.Context, *CreateDictDataRequest) (*CreateDictDataReply, error)
	// 更新
	UpdateDictData(context.Context, *UpdateDictDataRequest) (*UpdateDictDataReply, error)
	// 删除
	DeleteDictData(context.Context, *DeleteDictDataRequest) (*DeleteDictDataReply, error)
	// 获取信息
	GetDictData(context.Context, *GetDictDataRequest) (*GetDictDataReply, error)
	mustEmbedUnimplementedDictDataServer()
}

// UnimplementedDictDataServer must be embedded to have forward compatible implementations.
type UnimplementedDictDataServer struct {
}

func (UnimplementedDictDataServer) ListDictData(context.Context, *ListDictDataRequest) (*ListDictDataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDictData not implemented")
}
func (UnimplementedDictDataServer) CreateDictData(context.Context, *CreateDictDataRequest) (*CreateDictDataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDictData not implemented")
}
func (UnimplementedDictDataServer) UpdateDictData(context.Context, *UpdateDictDataRequest) (*UpdateDictDataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDictData not implemented")
}
func (UnimplementedDictDataServer) DeleteDictData(context.Context, *DeleteDictDataRequest) (*DeleteDictDataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDictData not implemented")
}
func (UnimplementedDictDataServer) GetDictData(context.Context, *GetDictDataRequest) (*GetDictDataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDictData not implemented")
}
func (UnimplementedDictDataServer) mustEmbedUnimplementedDictDataServer() {}

// UnsafeDictDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DictDataServer will
// result in compilation errors.
type UnsafeDictDataServer interface {
	mustEmbedUnimplementedDictDataServer()
}

func RegisterDictDataServer(s grpc.ServiceRegistrar, srv DictDataServer) {
	s.RegisterService(&DictData_ServiceDesc, srv)
}

func _DictData_ListDictData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDictDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictDataServer).ListDictData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictData_ListDictData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictDataServer).ListDictData(ctx, req.(*ListDictDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictData_CreateDictData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDictDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictDataServer).CreateDictData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictData_CreateDictData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictDataServer).CreateDictData(ctx, req.(*CreateDictDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictData_UpdateDictData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDictDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictDataServer).UpdateDictData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictData_UpdateDictData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictDataServer).UpdateDictData(ctx, req.(*UpdateDictDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictData_DeleteDictData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDictDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictDataServer).DeleteDictData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictData_DeleteDictData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictDataServer).DeleteDictData(ctx, req.(*DeleteDictDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DictData_GetDictData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDictDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DictDataServer).GetDictData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DictData_GetDictData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DictDataServer).GetDictData(ctx, req.(*GetDictDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DictData_ServiceDesc is the grpc.ServiceDesc for DictData service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DictData_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.admin.v1.DictData",
	HandlerType: (*DictDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDictData",
			Handler:    _DictData_ListDictData_Handler,
		},
		{
			MethodName: "CreateDictData",
			Handler:    _DictData_CreateDictData_Handler,
		},
		{
			MethodName: "UpdateDictData",
			Handler:    _DictData_UpdateDictData_Handler,
		},
		{
			MethodName: "DeleteDictData",
			Handler:    _DictData_DeleteDictData_Handler,
		},
		{
			MethodName: "GetDictData",
			Handler:    _DictData_GetDictData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dictdata.proto",
}
