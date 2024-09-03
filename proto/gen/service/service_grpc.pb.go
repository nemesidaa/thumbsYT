// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: service.proto

package nemesidaa_thumbsYT

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
	Mainstream_Load_FullMethodName = "/service.Mainstream/Load"
)

// MainstreamClient is the client API for Mainstream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MainstreamClient interface {
	// Loads & saves thumbs in DB
	Load(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (*LoadResponse, error)
}

type mainstreamClient struct {
	cc grpc.ClientConnInterface
}

func NewMainstreamClient(cc grpc.ClientConnInterface) MainstreamClient {
	return &mainstreamClient{cc}
}

func (c *mainstreamClient) Load(ctx context.Context, in *LoadRequest, opts ...grpc.CallOption) (*LoadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoadResponse)
	err := c.cc.Invoke(ctx, Mainstream_Load_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MainstreamServer is the server API for Mainstream service.
// All implementations must embed UnimplementedMainstreamServer
// for forward compatibility.
type MainstreamServer interface {
	// Loads & saves thumbs in DB
	Load(context.Context, *LoadRequest) (*LoadResponse, error)
	mustEmbedUnimplementedMainstreamServer()
}

// UnimplementedMainstreamServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMainstreamServer struct{}

func (UnimplementedMainstreamServer) Load(context.Context, *LoadRequest) (*LoadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Load not implemented")
}
func (UnimplementedMainstreamServer) mustEmbedUnimplementedMainstreamServer() {}
func (UnimplementedMainstreamServer) testEmbeddedByValue()                    {}

// UnsafeMainstreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MainstreamServer will
// result in compilation errors.
type UnsafeMainstreamServer interface {
	mustEmbedUnimplementedMainstreamServer()
}

func RegisterMainstreamServer(s grpc.ServiceRegistrar, srv MainstreamServer) {
	// If the following call pancis, it indicates UnimplementedMainstreamServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Mainstream_ServiceDesc, srv)
}

func _Mainstream_Load_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MainstreamServer).Load(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Mainstream_Load_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MainstreamServer).Load(ctx, req.(*LoadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mainstream_ServiceDesc is the grpc.ServiceDesc for Mainstream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mainstream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Mainstream",
	HandlerType: (*MainstreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Load",
			Handler:    _Mainstream_Load_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
