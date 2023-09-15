// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.3
// - protoc             v4.23.3
// source: dictdata.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationDictDataCreateDictData = "/api.admin.v1.DictData/CreateDictData"
const OperationDictDataDeleteDictData = "/api.admin.v1.DictData/DeleteDictData"
const OperationDictDataGetDictData = "/api.admin.v1.DictData/GetDictData"
const OperationDictDataListDictData = "/api.admin.v1.DictData/ListDictData"
const OperationDictDataUpdateDictData = "/api.admin.v1.DictData/UpdateDictData"

type DictDataHTTPServer interface {
	// CreateDictData创建
	CreateDictData(context.Context, *CreateDictDataRequest) (*CreateDictDataReply, error)
	// DeleteDictData删除
	DeleteDictData(context.Context, *DeleteDictDataRequest) (*DeleteDictDataReply, error)
	// GetDictData获取信息
	GetDictData(context.Context, *GetDictDataRequest) (*GetDictDataReply, error)
	// ListDictData列表信息
	ListDictData(context.Context, *ListDictDataRequest) (*ListDictDataReply, error)
	// UpdateDictData更新
	UpdateDictData(context.Context, *UpdateDictDataRequest) (*UpdateDictDataReply, error)
}

func RegisterDictDataHTTPServer(s *http.Server, srv DictDataHTTPServer) {
	r := s.Route("/")
	r.GET("/system/dict/data/type", _DictData_ListDictData0_HTTP_Handler(srv))
	r.POST("/system/dict/data", _DictData_CreateDictData0_HTTP_Handler(srv))
	r.PUT("/system/dict/data", _DictData_UpdateDictData0_HTTP_Handler(srv))
	r.DELETE("/system/dict/data/{dictCode}", _DictData_DeleteDictData0_HTTP_Handler(srv))
	r.GET("/system/dict/data/{dictCode}", _DictData_GetDictData0_HTTP_Handler(srv))
}

func _DictData_ListDictData0_HTTP_Handler(srv DictDataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListDictDataRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictDataListDictData)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListDictData(ctx, req.(*ListDictDataRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListDictDataReply)
		return ctx.Result(200, reply)
	}
}

func _DictData_CreateDictData0_HTTP_Handler(srv DictDataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateDictDataRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictDataCreateDictData)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateDictData(ctx, req.(*CreateDictDataRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateDictDataReply)
		return ctx.Result(200, reply)
	}
}

func _DictData_UpdateDictData0_HTTP_Handler(srv DictDataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateDictDataRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictDataUpdateDictData)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateDictData(ctx, req.(*UpdateDictDataRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateDictDataReply)
		return ctx.Result(200, reply)
	}
}

func _DictData_DeleteDictData0_HTTP_Handler(srv DictDataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteDictDataRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictDataDeleteDictData)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteDictData(ctx, req.(*DeleteDictDataRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteDictDataReply)
		return ctx.Result(200, reply)
	}
}

func _DictData_GetDictData0_HTTP_Handler(srv DictDataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetDictDataRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDictDataGetDictData)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetDictData(ctx, req.(*GetDictDataRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetDictDataReply)
		return ctx.Result(200, reply)
	}
}

type DictDataHTTPClient interface {
	CreateDictData(ctx context.Context, req *CreateDictDataRequest, opts ...http.CallOption) (rsp *CreateDictDataReply, err error)
	DeleteDictData(ctx context.Context, req *DeleteDictDataRequest, opts ...http.CallOption) (rsp *DeleteDictDataReply, err error)
	GetDictData(ctx context.Context, req *GetDictDataRequest, opts ...http.CallOption) (rsp *GetDictDataReply, err error)
	ListDictData(ctx context.Context, req *ListDictDataRequest, opts ...http.CallOption) (rsp *ListDictDataReply, err error)
	UpdateDictData(ctx context.Context, req *UpdateDictDataRequest, opts ...http.CallOption) (rsp *UpdateDictDataReply, err error)
}

type DictDataHTTPClientImpl struct {
	cc *http.Client
}

func NewDictDataHTTPClient(client *http.Client) DictDataHTTPClient {
	return &DictDataHTTPClientImpl{client}
}

func (c *DictDataHTTPClientImpl) CreateDictData(ctx context.Context, in *CreateDictDataRequest, opts ...http.CallOption) (*CreateDictDataReply, error) {
	var out CreateDictDataReply
	pattern := "/system/dict/data"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDictDataCreateDictData))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictDataHTTPClientImpl) DeleteDictData(ctx context.Context, in *DeleteDictDataRequest, opts ...http.CallOption) (*DeleteDictDataReply, error) {
	var out DeleteDictDataReply
	pattern := "/system/dict/data/{dictCode}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDictDataDeleteDictData))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictDataHTTPClientImpl) GetDictData(ctx context.Context, in *GetDictDataRequest, opts ...http.CallOption) (*GetDictDataReply, error) {
	var out GetDictDataReply
	pattern := "/system/dict/data/{dictCode}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDictDataGetDictData))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictDataHTTPClientImpl) ListDictData(ctx context.Context, in *ListDictDataRequest, opts ...http.CallOption) (*ListDictDataReply, error) {
	var out ListDictDataReply
	pattern := "/system/dict/data/type"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDictDataListDictData))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *DictDataHTTPClientImpl) UpdateDictData(ctx context.Context, in *UpdateDictDataRequest, opts ...http.CallOption) (*UpdateDictDataReply, error) {
	var out UpdateDictDataReply
	pattern := "/system/dict/data"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDictDataUpdateDictData))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
