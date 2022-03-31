// Code generated by protoc-gen-go-micro. DO NOT EDIT.
// versions:
// - protoc-gen-go-micro d8794062c4bcd2ce7b76dcec4ac75a5dab39d600
// - protoc              v3.19.4
// source: service/grpc/proto/test.proto

package proto

import (
	context "context"
	api "github.com/crypto-zero/go-micro/v2/api"
	client "github.com/crypto-zero/go-micro/v2/client"
	server "github.com/crypto-zero/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// NewTestEndpoints API Endpoints for Test service
func NewTestEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// TestService is the client API for Test service.
type TestService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type testService struct {
	c    client.Client
	name string
}

func NewTestService(name string, c client.Client) TestService {
	return &testService{
		c:    c,
		name: name,
	}
}

func (c *testService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Test.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestHandler is the server API for Test service.
type TestHandler interface {
	Call(context.Context, *Request, *Response) error
}

func RegisterTestHandler(s server.Server, hdlr TestHandler, opts ...server.HandlerOption) error {
	type test interface {
		Call(ctx context.Context, in *Request, out *Response) error
	}
	type Test struct {
		test
	}
	h := &testHandler{hdlr}
	return s.Handle(s.NewHandler(&Test{h}, opts...))
}

type testHandler struct {
	TestHandler
}

func (h *testHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.TestHandler.Call(ctx, in, out)
}
