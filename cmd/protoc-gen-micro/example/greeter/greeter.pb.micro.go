// Code generated by protoc-gen-go-micro. DO NOT EDIT.
// versions:
// - protoc-gen-go-micro 69846db0b2cf6b0e0d49ea9576be84af5ea094fb
// - protoc              v3.19.4
// source: greeter/greeter.proto

package greeter

import (
	context "context"
	person "github.com/crypto-zero/go-micro/cmd/protoc-gen-micro/v2/example/person"
	api "github.com/crypto-zero/go-micro/v2/api"
	client "github.com/crypto-zero/go-micro/v2/client"
	server "github.com/crypto-zero/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// GreeterService is the client API for Greeter service.
type GreeterService interface {
	// function comments
	Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, opts ...client.CallOption) (Greeter_StreamService, error)
	StreamA(ctx context.Context, in *Request, opts ...client.CallOption) (Greeter_StreamAService, error)
	StreamB(ctx context.Context, opts ...client.CallOption) (Greeter_StreamBService, error)
	FindPerson(ctx context.Context, in *person.Person, opts ...client.CallOption) (*person.Person, error)
}

type greeterService struct {
	c    client.Client
	name string
}

func NewGreeterService(name string, c client.Client) GreeterService {
	return &greeterService{
		c:    c,
		name: name,
	}
}

func (c *greeterService) Hello(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Greeter.Hello", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterService) Stream(ctx context.Context, opts ...client.CallOption) (Greeter_StreamService, error) {
	req := c.c.NewRequest(c.name, "Greeter.Stream", &Request{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &greeterServiceStream{stream}, nil
}

type Greeter_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Request) error
	Recv() (*Response, error)
}

type greeterServiceStream struct {
	stream client.Stream
}

func (x *greeterServiceStream) Close() error {
	return x.stream.Close()
}

func (x *greeterServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *greeterServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *greeterServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *greeterServiceStream) Send(m *Request) error {
	return x.stream.Send(m)
}

func (x *greeterServiceStream) Recv() (*Response, error) {
	m := new(Response)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterService) StreamA(ctx context.Context, in *Request, opts ...client.CallOption) (Greeter_StreamAService, error) {
	req := c.c.NewRequest(c.name, "Greeter.StreamA", &Request{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &greeterServiceStreamA{stream}, nil
}

type Greeter_StreamAService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*Response, error)
}

type greeterServiceStreamA struct {
	stream client.Stream
}

func (x *greeterServiceStreamA) Close() error {
	return x.stream.Close()
}

func (x *greeterServiceStreamA) Context() context.Context {
	return x.stream.Context()
}

func (x *greeterServiceStreamA) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *greeterServiceStreamA) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *greeterServiceStreamA) Recv() (*Response, error) {
	m := new(Response)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterService) StreamB(ctx context.Context, opts ...client.CallOption) (Greeter_StreamBService, error) {
	req := c.c.NewRequest(c.name, "Greeter.StreamB", &Request{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &greeterServiceStreamB{stream}, nil
}

type Greeter_StreamBService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Request) error
}

type greeterServiceStreamB struct {
	stream client.Stream
}

func (x *greeterServiceStreamB) Close() error {
	return x.stream.Close()
}

func (x *greeterServiceStreamB) Context() context.Context {
	return x.stream.Context()
}

func (x *greeterServiceStreamB) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *greeterServiceStreamB) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *greeterServiceStreamB) Send(m *Request) error {
	return x.stream.Send(m)
}

func (c *greeterService) FindPerson(ctx context.Context, in *person.Person, opts ...client.CallOption) (*person.Person, error) {
	req := c.c.NewRequest(c.name, "Greeter.FindPerson", in)
	out := new(person.Person)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterHandler is the server API for Greeter service.
type GreeterHandler interface {
	// function comments
	Hello(context.Context, *Request, *Response) error
	Stream(context.Context, Greeter_StreamStream) error
	StreamA(context.Context, *Request, Greeter_StreamAStream) error
	StreamB(context.Context, Greeter_StreamBStream) error
	FindPerson(context.Context, *person.Person, *person.Person) error
}

func RegisterGreeterHandler(s server.Server, hdlr GreeterHandler, opts ...server.HandlerOption) error {
	type greeter interface {
		Hello(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		StreamA(ctx context.Context, stream server.Stream) error
		StreamB(ctx context.Context, stream server.Stream) error
		FindPerson(ctx context.Context, in *person.Person, out *person.Person) error
	}
	type Greeter struct {
		greeter
	}
	h := &greeterHandler{hdlr}
	return s.Handle(s.NewHandler(&Greeter{h}, opts...))
}

type greeterHandler struct {
	GreeterHandler
}

func (h *greeterHandler) Hello(ctx context.Context, in *Request, out *Response) error {
	return h.GreeterHandler.Hello(ctx, in, out)
}

func (h *greeterHandler) Stream(ctx context.Context, stream server.Stream) error {
	return h.GreeterHandler.Stream(ctx, &greeterStreamStream{stream})
}

type Greeter_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Response) error
	Recv() (*Request, error)
}

type greeterStreamStream struct {
	stream server.Stream
}

func (x *greeterStreamStream) Close() error {
	return x.stream.Close()
}

func (x *greeterStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *greeterStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *greeterStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *greeterStreamStream) Send(m *Response) error {
	return x.stream.Send(m)
}

func (x *greeterStreamStream) Recv() (*Request, error) {
	m := new(Request)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *greeterHandler) StreamA(ctx context.Context, stream server.Stream) error {
	m := new(Request)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.GreeterHandler.StreamA(ctx, m, &greeterStreamAStream{stream})
}

type Greeter_StreamAStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Response) error
}

type greeterStreamAStream struct {
	stream server.Stream
}

func (x *greeterStreamAStream) Close() error {
	return x.stream.Close()
}

func (x *greeterStreamAStream) Context() context.Context {
	return x.stream.Context()
}

func (x *greeterStreamAStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *greeterStreamAStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *greeterStreamAStream) Send(m *Response) error {
	return x.stream.Send(m)
}

func (h *greeterHandler) StreamB(ctx context.Context, stream server.Stream) error {
	return h.GreeterHandler.StreamB(ctx, &greeterStreamBStream{stream})
}

type Greeter_StreamBStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*Request, error)
}

type greeterStreamBStream struct {
	stream server.Stream
}

func (x *greeterStreamBStream) Close() error {
	return x.stream.Close()
}

func (x *greeterStreamBStream) Context() context.Context {
	return x.stream.Context()
}

func (x *greeterStreamBStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *greeterStreamBStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *greeterStreamBStream) Recv() (*Request, error) {
	m := new(Request)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *greeterHandler) FindPerson(ctx context.Context, in *person.Person, out *person.Person) error {
	return h.GreeterHandler.FindPerson(ctx, in, out)
}