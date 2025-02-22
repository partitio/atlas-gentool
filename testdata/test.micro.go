// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/partitio/atlas-gentool/testdata/test.proto

package test

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/infobloxopen/protoc-gen-atlas-query-validate/options"
	_ "github.com/infobloxopen/protoc-gen-atlas-validate/options"
	_ "github.com/infobloxopen/protoc-gen-preprocess/options"
	_ "github.com/lyft/protoc-gen-validate/validate"
	_ "github.com/partitio/protoc-gen-gorm/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ExampleService service

type ExampleService interface {
	Something(ctx context.Context, in *Example, opts ...client.CallOption) (*Example, error)
	Create(ctx context.Context, in *Example, opts ...client.CallOption) (*Example, error)
	Test(ctx context.Context, in *Example, opts ...client.CallOption) (*Example, error)
}

type exampleService struct {
	c    client.Client
	name string
}

func NewExampleService(name string, c client.Client) ExampleService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "test"
	}
	return &exampleService{
		c:    c,
		name: name,
	}
}

func (c *exampleService) Something(ctx context.Context, in *Example, opts ...client.CallOption) (*Example, error) {
	req := c.c.NewRequest(c.name, "ExampleService.Something", in)
	out := new(Example)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleService) Create(ctx context.Context, in *Example, opts ...client.CallOption) (*Example, error) {
	req := c.c.NewRequest(c.name, "ExampleService.Create", in)
	out := new(Example)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleService) Test(ctx context.Context, in *Example, opts ...client.CallOption) (*Example, error) {
	req := c.c.NewRequest(c.name, "ExampleService.Test", in)
	out := new(Example)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ExampleService service

type ExampleServiceHandler interface {
	Something(context.Context, *Example, *Example) error
	Create(context.Context, *Example, *Example) error
	Test(context.Context, *Example, *Example) error
}

func RegisterExampleServiceHandler(s server.Server, hdlr ExampleServiceHandler, opts ...server.HandlerOption) error {
	type exampleService interface {
		Something(ctx context.Context, in *Example, out *Example) error
		Create(ctx context.Context, in *Example, out *Example) error
		Test(ctx context.Context, in *Example, out *Example) error
	}
	type ExampleService struct {
		exampleService
	}
	h := &exampleServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ExampleService{h}, opts...))
}

type exampleServiceHandler struct {
	ExampleServiceHandler
}

func (h *exampleServiceHandler) Something(ctx context.Context, in *Example, out *Example) error {
	return h.ExampleServiceHandler.Something(ctx, in, out)
}

func (h *exampleServiceHandler) Create(ctx context.Context, in *Example, out *Example) error {
	return h.ExampleServiceHandler.Create(ctx, in, out)
}

func (h *exampleServiceHandler) Test(ctx context.Context, in *Example, out *Example) error {
	return h.ExampleServiceHandler.Test(ctx, in, out)
}
