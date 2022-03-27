// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AnarchyClient is the client API for Anarchy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnarchyClient interface {
	// Elements
	GetElem(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*AnarchyElement, error)
	GetCombination(ctx context.Context, in *AnarchyCombination, opts ...grpc.CallOption) (*AnarchyCombinationResult, error)
	GetAll(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (Anarchy_GetAllClient, error)
	CreateElement(ctx context.Context, in *AnarchyElementCreate, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Savefile
	GetInv(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*AnarchyInventory, error)
	AddFound(ctx context.Context, in *AnarchyUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Recents
	GetRecents(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AnarchyRecents, error)
	WaitForNextRecent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Anarchy_WaitForNextRecentClient, error)
}

type anarchyClient struct {
	cc grpc.ClientConnInterface
}

func NewAnarchyClient(cc grpc.ClientConnInterface) AnarchyClient {
	return &anarchyClient{cc}
}

func (c *anarchyClient) GetElem(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*AnarchyElement, error) {
	out := new(AnarchyElement)
	err := c.cc.Invoke(ctx, "/anarchy.Anarchy/GetElem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anarchyClient) GetCombination(ctx context.Context, in *AnarchyCombination, opts ...grpc.CallOption) (*AnarchyCombinationResult, error) {
	out := new(AnarchyCombinationResult)
	err := c.cc.Invoke(ctx, "/anarchy.Anarchy/GetCombination", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anarchyClient) GetAll(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (Anarchy_GetAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &Anarchy_ServiceDesc.Streams[0], "/anarchy.Anarchy/GetAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &anarchyGetAllClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Anarchy_GetAllClient interface {
	Recv() (*AnarchyGetAllChunk, error)
	grpc.ClientStream
}

type anarchyGetAllClient struct {
	grpc.ClientStream
}

func (x *anarchyGetAllClient) Recv() (*AnarchyGetAllChunk, error) {
	m := new(AnarchyGetAllChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *anarchyClient) CreateElement(ctx context.Context, in *AnarchyElementCreate, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/anarchy.Anarchy/CreateElement", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anarchyClient) GetInv(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*AnarchyInventory, error) {
	out := new(AnarchyInventory)
	err := c.cc.Invoke(ctx, "/anarchy.Anarchy/GetInv", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anarchyClient) AddFound(ctx context.Context, in *AnarchyUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/anarchy.Anarchy/AddFound", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anarchyClient) GetRecents(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AnarchyRecents, error) {
	out := new(AnarchyRecents)
	err := c.cc.Invoke(ctx, "/anarchy.Anarchy/GetRecents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anarchyClient) WaitForNextRecent(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (Anarchy_WaitForNextRecentClient, error) {
	stream, err := c.cc.NewStream(ctx, &Anarchy_ServiceDesc.Streams[1], "/anarchy.Anarchy/WaitForNextRecent", opts...)
	if err != nil {
		return nil, err
	}
	x := &anarchyWaitForNextRecentClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Anarchy_WaitForNextRecentClient interface {
	Recv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type anarchyWaitForNextRecentClient struct {
	grpc.ClientStream
}

func (x *anarchyWaitForNextRecentClient) Recv() (*emptypb.Empty, error) {
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AnarchyServer is the server API for Anarchy service.
// All implementations must embed UnimplementedAnarchyServer
// for forward compatibility
type AnarchyServer interface {
	// Elements
	GetElem(context.Context, *wrapperspb.StringValue) (*AnarchyElement, error)
	GetCombination(context.Context, *AnarchyCombination) (*AnarchyCombinationResult, error)
	GetAll(*wrapperspb.StringValue, Anarchy_GetAllServer) error
	CreateElement(context.Context, *AnarchyElementCreate) (*emptypb.Empty, error)
	// Savefile
	GetInv(context.Context, *wrapperspb.StringValue) (*AnarchyInventory, error)
	AddFound(context.Context, *AnarchyUserRequest) (*emptypb.Empty, error)
	// Recents
	GetRecents(context.Context, *emptypb.Empty) (*AnarchyRecents, error)
	WaitForNextRecent(*emptypb.Empty, Anarchy_WaitForNextRecentServer) error
	mustEmbedUnimplementedAnarchyServer()
}

// UnimplementedAnarchyServer must be embedded to have forward compatible implementations.
type UnimplementedAnarchyServer struct {
}

func (UnimplementedAnarchyServer) GetElem(context.Context, *wrapperspb.StringValue) (*AnarchyElement, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetElem not implemented")
}
func (UnimplementedAnarchyServer) GetCombination(context.Context, *AnarchyCombination) (*AnarchyCombinationResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCombination not implemented")
}
func (UnimplementedAnarchyServer) GetAll(*wrapperspb.StringValue, Anarchy_GetAllServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedAnarchyServer) CreateElement(context.Context, *AnarchyElementCreate) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateElement not implemented")
}
func (UnimplementedAnarchyServer) GetInv(context.Context, *wrapperspb.StringValue) (*AnarchyInventory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInv not implemented")
}
func (UnimplementedAnarchyServer) AddFound(context.Context, *AnarchyUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFound not implemented")
}
func (UnimplementedAnarchyServer) GetRecents(context.Context, *emptypb.Empty) (*AnarchyRecents, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecents not implemented")
}
func (UnimplementedAnarchyServer) WaitForNextRecent(*emptypb.Empty, Anarchy_WaitForNextRecentServer) error {
	return status.Errorf(codes.Unimplemented, "method WaitForNextRecent not implemented")
}
func (UnimplementedAnarchyServer) mustEmbedUnimplementedAnarchyServer() {}

// UnsafeAnarchyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnarchyServer will
// result in compilation errors.
type UnsafeAnarchyServer interface {
	mustEmbedUnimplementedAnarchyServer()
}

func RegisterAnarchyServer(s grpc.ServiceRegistrar, srv AnarchyServer) {
	s.RegisterService(&Anarchy_ServiceDesc, srv)
}

func _Anarchy_GetElem_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnarchyServer).GetElem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anarchy.Anarchy/GetElem",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(AnarchyServer).GetElem(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anarchy_GetCombination_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(AnarchyCombination)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnarchyServer).GetCombination(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anarchy.Anarchy/GetCombination",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(AnarchyServer).GetCombination(ctx, req.(*AnarchyCombination))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anarchy_GetAll_Handler(srv any, stream grpc.ServerStream) error {
	m := new(wrapperspb.StringValue)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AnarchyServer).GetAll(m, &anarchyGetAllServer{stream})
}

type Anarchy_GetAllServer interface {
	Send(*AnarchyGetAllChunk) error
	grpc.ServerStream
}

type anarchyGetAllServer struct {
	grpc.ServerStream
}

func (x *anarchyGetAllServer) Send(m *AnarchyGetAllChunk) error {
	return x.ServerStream.SendMsg(m)
}

func _Anarchy_CreateElement_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(AnarchyElementCreate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnarchyServer).CreateElement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anarchy.Anarchy/CreateElement",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(AnarchyServer).CreateElement(ctx, req.(*AnarchyElementCreate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anarchy_GetInv_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnarchyServer).GetInv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anarchy.Anarchy/GetInv",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(AnarchyServer).GetInv(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anarchy_AddFound_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(AnarchyUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnarchyServer).AddFound(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anarchy.Anarchy/AddFound",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(AnarchyServer).AddFound(ctx, req.(*AnarchyUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anarchy_GetRecents_Handler(srv any, ctx context.Context, dec func(any) error, interceptor grpc.UnaryServerInterceptor) (any, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnarchyServer).GetRecents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anarchy.Anarchy/GetRecents",
	}
	handler := func(ctx context.Context, req any) (any, error) {
		return srv.(AnarchyServer).GetRecents(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anarchy_WaitForNextRecent_Handler(srv any, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AnarchyServer).WaitForNextRecent(m, &anarchyWaitForNextRecentServer{stream})
}

type Anarchy_WaitForNextRecentServer interface {
	Send(*emptypb.Empty) error
	grpc.ServerStream
}

type anarchyWaitForNextRecentServer struct {
	grpc.ServerStream
}

func (x *anarchyWaitForNextRecentServer) Send(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

// Anarchy_ServiceDesc is the grpc.ServiceDesc for Anarchy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Anarchy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "anarchy.Anarchy",
	HandlerType: (*AnarchyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetElem",
			Handler:    _Anarchy_GetElem_Handler,
		},
		{
			MethodName: "GetCombination",
			Handler:    _Anarchy_GetCombination_Handler,
		},
		{
			MethodName: "CreateElement",
			Handler:    _Anarchy_CreateElement_Handler,
		},
		{
			MethodName: "GetInv",
			Handler:    _Anarchy_GetInv_Handler,
		},
		{
			MethodName: "AddFound",
			Handler:    _Anarchy_AddFound_Handler,
		},
		{
			MethodName: "GetRecents",
			Handler:    _Anarchy_GetRecents_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAll",
			Handler:       _Anarchy_GetAll_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "WaitForNextRecent",
			Handler:       _Anarchy_WaitForNextRecent_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "anarchy.proto",
}
