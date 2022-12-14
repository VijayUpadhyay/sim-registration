// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: simRegistration.proto

package registration

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

// SimRegistrationClient is the client API for SimRegistration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimRegistrationClient interface {
	RegisterCustomer(ctx context.Context, in *CustomerDetails, opts ...grpc.CallOption) (*CreateResponseMessage, error)
}

type simRegistrationClient struct {
	cc grpc.ClientConnInterface
}

func NewSimRegistrationClient(cc grpc.ClientConnInterface) SimRegistrationClient {
	return &simRegistrationClient{cc}
}

func (c *simRegistrationClient) RegisterCustomer(ctx context.Context, in *CustomerDetails, opts ...grpc.CallOption) (*CreateResponseMessage, error) {
	out := new(CreateResponseMessage)
	err := c.cc.Invoke(ctx, "/simRegistration.SimRegistration/RegisterCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimRegistrationServer is the server API for SimRegistration service.
// All implementations must embed UnimplementedSimRegistrationServer
// for forward compatibility
type SimRegistrationServer interface {
	RegisterCustomer(context.Context, *CustomerDetails) (*CreateResponseMessage, error)
	mustEmbedUnimplementedSimRegistrationServer()
}

// UnimplementedSimRegistrationServer must be embedded to have forward compatible implementations.
type UnimplementedSimRegistrationServer struct {
}

func (UnimplementedSimRegistrationServer) RegisterCustomer(context.Context, *CustomerDetails) (*CreateResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterCustomer not implemented")
}
func (UnimplementedSimRegistrationServer) mustEmbedUnimplementedSimRegistrationServer() {}

// UnsafeSimRegistrationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimRegistrationServer will
// result in compilation errors.
type UnsafeSimRegistrationServer interface {
	mustEmbedUnimplementedSimRegistrationServer()
}

func RegisterSimRegistrationServer(s grpc.ServiceRegistrar, srv SimRegistrationServer) {
	s.RegisterService(&SimRegistration_ServiceDesc, srv)
}

func _SimRegistration_RegisterCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimRegistrationServer).RegisterCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simRegistration.SimRegistration/RegisterCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimRegistrationServer).RegisterCustomer(ctx, req.(*CustomerDetails))
	}
	return interceptor(ctx, in, info, handler)
}

// SimRegistration_ServiceDesc is the grpc.ServiceDesc for SimRegistration service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimRegistration_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "simRegistration.SimRegistration",
	HandlerType: (*SimRegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterCustomer",
			Handler:    _SimRegistration_RegisterCustomer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "simRegistration.proto",
}
