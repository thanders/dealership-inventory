// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: greet.proto

package proto

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

// HondaInventoryServiceClient is the client API for HondaInventoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HondaInventoryServiceClient interface {
	GetAllData(ctx context.Context, in *GetAllDataRequest, opts ...grpc.CallOption) (*DealerInventory, error)
	GetDealerInventory(ctx context.Context, in *GetDealerInventoryRequest, opts ...grpc.CallOption) (*TotalInventory, error)
	GetHondaInfo(ctx context.Context, in *GetHondaInfoRequest, opts ...grpc.CallOption) (*HondaVehicleInfo, error)
}

type hondaInventoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHondaInventoryServiceClient(cc grpc.ClientConnInterface) HondaInventoryServiceClient {
	return &hondaInventoryServiceClient{cc}
}

func (c *hondaInventoryServiceClient) GetAllData(ctx context.Context, in *GetAllDataRequest, opts ...grpc.CallOption) (*DealerInventory, error) {
	out := new(DealerInventory)
	err := c.cc.Invoke(ctx, "/greet.HondaInventoryService/GetAllData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hondaInventoryServiceClient) GetDealerInventory(ctx context.Context, in *GetDealerInventoryRequest, opts ...grpc.CallOption) (*TotalInventory, error) {
	out := new(TotalInventory)
	err := c.cc.Invoke(ctx, "/greet.HondaInventoryService/GetDealerInventory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hondaInventoryServiceClient) GetHondaInfo(ctx context.Context, in *GetHondaInfoRequest, opts ...grpc.CallOption) (*HondaVehicleInfo, error) {
	out := new(HondaVehicleInfo)
	err := c.cc.Invoke(ctx, "/greet.HondaInventoryService/GetHondaInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HondaInventoryServiceServer is the server API for HondaInventoryService service.
// All implementations must embed UnimplementedHondaInventoryServiceServer
// for forward compatibility
type HondaInventoryServiceServer interface {
	GetAllData(context.Context, *GetAllDataRequest) (*DealerInventory, error)
	GetDealerInventory(context.Context, *GetDealerInventoryRequest) (*TotalInventory, error)
	GetHondaInfo(context.Context, *GetHondaInfoRequest) (*HondaVehicleInfo, error)
	mustEmbedUnimplementedHondaInventoryServiceServer()
}

// UnimplementedHondaInventoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHondaInventoryServiceServer struct {
}

func (UnimplementedHondaInventoryServiceServer) GetAllData(context.Context, *GetAllDataRequest) (*DealerInventory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllData not implemented")
}
func (UnimplementedHondaInventoryServiceServer) GetDealerInventory(context.Context, *GetDealerInventoryRequest) (*TotalInventory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDealerInventory not implemented")
}
func (UnimplementedHondaInventoryServiceServer) GetHondaInfo(context.Context, *GetHondaInfoRequest) (*HondaVehicleInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHondaInfo not implemented")
}
func (UnimplementedHondaInventoryServiceServer) mustEmbedUnimplementedHondaInventoryServiceServer() {}

// UnsafeHondaInventoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HondaInventoryServiceServer will
// result in compilation errors.
type UnsafeHondaInventoryServiceServer interface {
	mustEmbedUnimplementedHondaInventoryServiceServer()
}

func RegisterHondaInventoryServiceServer(s grpc.ServiceRegistrar, srv HondaInventoryServiceServer) {
	s.RegisterService(&HondaInventoryService_ServiceDesc, srv)
}

func _HondaInventoryService_GetAllData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HondaInventoryServiceServer).GetAllData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.HondaInventoryService/GetAllData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HondaInventoryServiceServer).GetAllData(ctx, req.(*GetAllDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HondaInventoryService_GetDealerInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDealerInventoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HondaInventoryServiceServer).GetDealerInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.HondaInventoryService/GetDealerInventory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HondaInventoryServiceServer).GetDealerInventory(ctx, req.(*GetDealerInventoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HondaInventoryService_GetHondaInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHondaInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HondaInventoryServiceServer).GetHondaInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.HondaInventoryService/GetHondaInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HondaInventoryServiceServer).GetHondaInfo(ctx, req.(*GetHondaInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HondaInventoryService_ServiceDesc is the grpc.ServiceDesc for HondaInventoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HondaInventoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet.HondaInventoryService",
	HandlerType: (*HondaInventoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllData",
			Handler:    _HondaInventoryService_GetAllData_Handler,
		},
		{
			MethodName: "GetDealerInventory",
			Handler:    _HondaInventoryService_GetDealerInventory_Handler,
		},
		{
			MethodName: "GetHondaInfo",
			Handler:    _HondaInventoryService_GetHondaInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greet.proto",
}
