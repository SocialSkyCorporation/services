// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/xian/xian.proto

package go_micro_srv_xian

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Client API for Xian service

type XianService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Xian_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Xian_PingPongService, error)
}

type xianService struct {
	c    client.Client
	name string
}

func NewXianService(name string, c client.Client) XianService {
	return &xianService{
		c:    c,
		name: name,
	}
}

func (c *xianService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Xian.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *xianService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Xian_StreamService, error) {
	req := c.c.NewRequest(c.name, "Xian.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &xianServiceStream{stream}, nil
}

type Xian_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type xianServiceStream struct {
	stream client.Stream
}

func (x *xianServiceStream) Close() error {
	return x.stream.Close()
}

func (x *xianServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *xianServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *xianServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *xianServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *xianService) PingPong(ctx context.Context, opts ...client.CallOption) (Xian_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Xian.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &xianServicePingPong{stream}, nil
}

type Xian_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type xianServicePingPong struct {
	stream client.Stream
}

func (x *xianServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *xianServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *xianServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *xianServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *xianServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *xianServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Xian service

type XianHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Xian_StreamStream) error
	PingPong(context.Context, Xian_PingPongStream) error
}

func RegisterXianHandler(s server.Server, hdlr XianHandler, opts ...server.HandlerOption) error {
	type xian interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Xian struct {
		xian
	}
	h := &xianHandler{hdlr}
	return s.Handle(s.NewHandler(&Xian{h}, opts...))
}

type xianHandler struct {
	XianHandler
}

func (h *xianHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.XianHandler.Call(ctx, in, out)
}

func (h *xianHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.XianHandler.Stream(ctx, m, &xianStreamStream{stream})
}

type Xian_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type xianStreamStream struct {
	stream server.Stream
}

func (x *xianStreamStream) Close() error {
	return x.stream.Close()
}

func (x *xianStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *xianStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *xianStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *xianStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *xianHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.XianHandler.PingPong(ctx, &xianPingPongStream{stream})
}

type Xian_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type xianPingPongStream struct {
	stream server.Stream
}

func (x *xianPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *xianPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *xianPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *xianPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *xianPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *xianPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}