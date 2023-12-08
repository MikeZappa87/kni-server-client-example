// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: pkg/apis/runtime/beta/api.proto

package beta

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
	KNI_AttachNetwork_FullMethodName    = "/kni.KNI/AttachNetwork"
	KNI_DetachNetwork_FullMethodName    = "/kni.KNI/DetachNetwork"
	KNI_QueryNetworks_FullMethodName    = "/kni.KNI/QueryNetworks"
	KNI_NetworkStatus_FullMethodName    = "/kni.KNI/NetworkStatus"
	KNI_SetupNodeNetwork_FullMethodName = "/kni.KNI/SetupNodeNetwork"
	KNI_QueryPod_FullMethodName         = "/kni.KNI/QueryPod"
)

// KNIClient is the client API for KNI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KNIClient interface {
	AttachNetwork(ctx context.Context, in *AttachNetworkRequest, opts ...grpc.CallOption) (*AttachNetworkResponse, error)
	DetachNetwork(ctx context.Context, in *DetachNetworkRequest, opts ...grpc.CallOption) (*DetachNetworkResponse, error)
	QueryNetworks(ctx context.Context, in *QueryNetworksRequest, opts ...grpc.CallOption) (*QueryNetworksResponse, error)
	NetworkStatus(ctx context.Context, in *NetworkStatusRequest, opts ...grpc.CallOption) (*NetworkStatusResponse, error)
	SetupNodeNetwork(ctx context.Context, in *SetupNodeNetworkRequest, opts ...grpc.CallOption) (*SetupNodeNetworkResponse, error)
	QueryPod(ctx context.Context, in *QueryPodRequest, opts ...grpc.CallOption) (*QueryPodResponse, error)
}

type kNIClient struct {
	cc grpc.ClientConnInterface
}

func NewKNIClient(cc grpc.ClientConnInterface) KNIClient {
	return &kNIClient{cc}
}

func (c *kNIClient) AttachNetwork(ctx context.Context, in *AttachNetworkRequest, opts ...grpc.CallOption) (*AttachNetworkResponse, error) {
	out := new(AttachNetworkResponse)
	err := c.cc.Invoke(ctx, KNI_AttachNetwork_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kNIClient) DetachNetwork(ctx context.Context, in *DetachNetworkRequest, opts ...grpc.CallOption) (*DetachNetworkResponse, error) {
	out := new(DetachNetworkResponse)
	err := c.cc.Invoke(ctx, KNI_DetachNetwork_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kNIClient) QueryNetworks(ctx context.Context, in *QueryNetworksRequest, opts ...grpc.CallOption) (*QueryNetworksResponse, error) {
	out := new(QueryNetworksResponse)
	err := c.cc.Invoke(ctx, KNI_QueryNetworks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kNIClient) NetworkStatus(ctx context.Context, in *NetworkStatusRequest, opts ...grpc.CallOption) (*NetworkStatusResponse, error) {
	out := new(NetworkStatusResponse)
	err := c.cc.Invoke(ctx, KNI_NetworkStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kNIClient) SetupNodeNetwork(ctx context.Context, in *SetupNodeNetworkRequest, opts ...grpc.CallOption) (*SetupNodeNetworkResponse, error) {
	out := new(SetupNodeNetworkResponse)
	err := c.cc.Invoke(ctx, KNI_SetupNodeNetwork_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kNIClient) QueryPod(ctx context.Context, in *QueryPodRequest, opts ...grpc.CallOption) (*QueryPodResponse, error) {
	out := new(QueryPodResponse)
	err := c.cc.Invoke(ctx, KNI_QueryPod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KNIServer is the server API for KNI service.
// All implementations should embed UnimplementedKNIServer
// for forward compatibility
type KNIServer interface {
	AttachNetwork(context.Context, *AttachNetworkRequest) (*AttachNetworkResponse, error)
	DetachNetwork(context.Context, *DetachNetworkRequest) (*DetachNetworkResponse, error)
	QueryNetworks(context.Context, *QueryNetworksRequest) (*QueryNetworksResponse, error)
	NetworkStatus(context.Context, *NetworkStatusRequest) (*NetworkStatusResponse, error)
	SetupNodeNetwork(context.Context, *SetupNodeNetworkRequest) (*SetupNodeNetworkResponse, error)
	QueryPod(context.Context, *QueryPodRequest) (*QueryPodResponse, error)
}

// UnimplementedKNIServer should be embedded to have forward compatible implementations.
type UnimplementedKNIServer struct {
}

func (UnimplementedKNIServer) AttachNetwork(context.Context, *AttachNetworkRequest) (*AttachNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachNetwork not implemented")
}
func (UnimplementedKNIServer) DetachNetwork(context.Context, *DetachNetworkRequest) (*DetachNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetachNetwork not implemented")
}
func (UnimplementedKNIServer) QueryNetworks(context.Context, *QueryNetworksRequest) (*QueryNetworksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryNetworks not implemented")
}
func (UnimplementedKNIServer) NetworkStatus(context.Context, *NetworkStatusRequest) (*NetworkStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NetworkStatus not implemented")
}
func (UnimplementedKNIServer) SetupNodeNetwork(context.Context, *SetupNodeNetworkRequest) (*SetupNodeNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetupNodeNetwork not implemented")
}
func (UnimplementedKNIServer) QueryPod(context.Context, *QueryPodRequest) (*QueryPodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPod not implemented")
}

// UnsafeKNIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KNIServer will
// result in compilation errors.
type UnsafeKNIServer interface {
	mustEmbedUnimplementedKNIServer()
}

func RegisterKNIServer(s grpc.ServiceRegistrar, srv KNIServer) {
	s.RegisterService(&KNI_ServiceDesc, srv)
}

func _KNI_AttachNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KNIServer).AttachNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KNI_AttachNetwork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KNIServer).AttachNetwork(ctx, req.(*AttachNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KNI_DetachNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetachNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KNIServer).DetachNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KNI_DetachNetwork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KNIServer).DetachNetwork(ctx, req.(*DetachNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KNI_QueryNetworks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNetworksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KNIServer).QueryNetworks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KNI_QueryNetworks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KNIServer).QueryNetworks(ctx, req.(*QueryNetworksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KNI_NetworkStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetworkStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KNIServer).NetworkStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KNI_NetworkStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KNIServer).NetworkStatus(ctx, req.(*NetworkStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KNI_SetupNodeNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetupNodeNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KNIServer).SetupNodeNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KNI_SetupNodeNetwork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KNIServer).SetupNodeNetwork(ctx, req.(*SetupNodeNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KNI_QueryPod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryPodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KNIServer).QueryPod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: KNI_QueryPod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KNIServer).QueryPod(ctx, req.(*QueryPodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// KNI_ServiceDesc is the grpc.ServiceDesc for KNI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KNI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kni.KNI",
	HandlerType: (*KNIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AttachNetwork",
			Handler:    _KNI_AttachNetwork_Handler,
		},
		{
			MethodName: "DetachNetwork",
			Handler:    _KNI_DetachNetwork_Handler,
		},
		{
			MethodName: "QueryNetworks",
			Handler:    _KNI_QueryNetworks_Handler,
		},
		{
			MethodName: "NetworkStatus",
			Handler:    _KNI_NetworkStatus_Handler,
		},
		{
			MethodName: "SetupNodeNetwork",
			Handler:    _KNI_SetupNodeNetwork_Handler,
		},
		{
			MethodName: "QueryPod",
			Handler:    _KNI_QueryPod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/apis/runtime/beta/api.proto",
}