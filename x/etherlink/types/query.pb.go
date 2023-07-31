// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: etherlink/etherlink/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd78c672efb73f25, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd78c672efb73f25, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetEthStateRequest struct {
}

func (m *QueryGetEthStateRequest) Reset()         { *m = QueryGetEthStateRequest{} }
func (m *QueryGetEthStateRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetEthStateRequest) ProtoMessage()    {}
func (*QueryGetEthStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd78c672efb73f25, []int{2}
}
func (m *QueryGetEthStateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetEthStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetEthStateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetEthStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetEthStateRequest.Merge(m, src)
}
func (m *QueryGetEthStateRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetEthStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetEthStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetEthStateRequest proto.InternalMessageInfo

type QueryGetEthStateResponse struct {
	EthState EthState `protobuf:"bytes,1,opt,name=EthState,proto3" json:"EthState"`
}

func (m *QueryGetEthStateResponse) Reset()         { *m = QueryGetEthStateResponse{} }
func (m *QueryGetEthStateResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetEthStateResponse) ProtoMessage()    {}
func (*QueryGetEthStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd78c672efb73f25, []int{3}
}
func (m *QueryGetEthStateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetEthStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetEthStateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetEthStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetEthStateResponse.Merge(m, src)
}
func (m *QueryGetEthStateResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetEthStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetEthStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetEthStateResponse proto.InternalMessageInfo

func (m *QueryGetEthStateResponse) GetEthState() EthState {
	if m != nil {
		return m.EthState
	}
	return EthState{}
}

type QueryGetEthInputRequest struct {
}

func (m *QueryGetEthInputRequest) Reset()         { *m = QueryGetEthInputRequest{} }
func (m *QueryGetEthInputRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetEthInputRequest) ProtoMessage()    {}
func (*QueryGetEthInputRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd78c672efb73f25, []int{4}
}
func (m *QueryGetEthInputRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetEthInputRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetEthInputRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetEthInputRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetEthInputRequest.Merge(m, src)
}
func (m *QueryGetEthInputRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetEthInputRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetEthInputRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetEthInputRequest proto.InternalMessageInfo

type QueryGetEthInputResponse struct {
	EthInput EthInput `protobuf:"bytes,1,opt,name=EthInput,proto3" json:"EthInput"`
}

func (m *QueryGetEthInputResponse) Reset()         { *m = QueryGetEthInputResponse{} }
func (m *QueryGetEthInputResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetEthInputResponse) ProtoMessage()    {}
func (*QueryGetEthInputResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd78c672efb73f25, []int{5}
}
func (m *QueryGetEthInputResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetEthInputResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetEthInputResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetEthInputResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetEthInputResponse.Merge(m, src)
}
func (m *QueryGetEthInputResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetEthInputResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetEthInputResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetEthInputResponse proto.InternalMessageInfo

func (m *QueryGetEthInputResponse) GetEthInput() EthInput {
	if m != nil {
		return m.EthInput
	}
	return EthInput{}
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "etherlink.etherlink.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "etherlink.etherlink.QueryParamsResponse")
	proto.RegisterType((*QueryGetEthStateRequest)(nil), "etherlink.etherlink.QueryGetEthStateRequest")
	proto.RegisterType((*QueryGetEthStateResponse)(nil), "etherlink.etherlink.QueryGetEthStateResponse")
	proto.RegisterType((*QueryGetEthInputRequest)(nil), "etherlink.etherlink.QueryGetEthInputRequest")
	proto.RegisterType((*QueryGetEthInputResponse)(nil), "etherlink.etherlink.QueryGetEthInputResponse")
}

func init() { proto.RegisterFile("etherlink/etherlink/query.proto", fileDescriptor_bd78c672efb73f25) }

var fileDescriptor_bd78c672efb73f25 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x4e, 0xfa, 0x40,
	0x10, 0xc7, 0xdb, 0xdf, 0x1f, 0x62, 0xd6, 0xdb, 0x42, 0x22, 0x16, 0x29, 0xa4, 0x24, 0x4a, 0x8c,
	0x76, 0x03, 0xc6, 0x83, 0x27, 0x13, 0x12, 0x63, 0xbc, 0x21, 0xde, 0xf4, 0x60, 0x16, 0xb3, 0x29,
	0x8d, 0xd0, 0x2d, 0xdd, 0xc5, 0xc8, 0xcd, 0xf8, 0x02, 0x9a, 0xf8, 0x1e, 0x3e, 0x07, 0x47, 0x12,
	0x2f, 0x9e, 0x8c, 0x01, 0x1f, 0xc4, 0x74, 0x77, 0xcb, 0x1f, 0x29, 0x96, 0xdb, 0x76, 0xe6, 0x3b,
	0xf3, 0xfd, 0x74, 0x66, 0x17, 0x14, 0x08, 0x6f, 0x91, 0xa0, 0xed, 0x7a, 0xb7, 0x68, 0x7a, 0xea,
	0xf6, 0x48, 0xd0, 0xb7, 0xfd, 0x80, 0x72, 0x0a, 0xd3, 0x93, 0xb0, 0x3d, 0x39, 0x19, 0x19, 0x87,
	0x3a, 0x54, 0xe4, 0x51, 0x78, 0x92, 0x52, 0x63, 0xcb, 0xa1, 0xd4, 0x69, 0x13, 0x84, 0x7d, 0x17,
	0x61, 0xcf, 0xa3, 0x1c, 0x73, 0x97, 0x7a, 0x4c, 0x65, 0x77, 0x6f, 0x28, 0xeb, 0x50, 0x86, 0x9a,
	0x98, 0x11, 0xe9, 0x80, 0xee, 0x2a, 0x4d, 0xc2, 0x71, 0x05, 0xf9, 0xd8, 0x71, 0x3d, 0x21, 0x56,
	0xda, 0x62, 0x1c, 0x95, 0x8f, 0x03, 0xdc, 0x89, 0xba, 0x95, 0xe2, 0x14, 0x84, 0xb7, 0xae, 0x19,
	0xc7, 0x9c, 0x24, 0x89, 0x5c, 0xcf, 0xef, 0x71, 0x29, 0xb2, 0x32, 0x00, 0x9e, 0x87, 0x34, 0x75,
	0xd1, 0xbe, 0x41, 0xba, 0x3d, 0xc2, 0xb8, 0x55, 0x07, 0xe9, 0xb9, 0x28, 0xf3, 0xa9, 0xc7, 0x08,
	0x3c, 0x02, 0x29, 0x89, 0x91, 0xd5, 0x8b, 0x7a, 0x79, 0xbd, 0x9a, 0xb3, 0x63, 0xc6, 0x63, 0xcb,
	0xa2, 0xda, 0xbf, 0xc1, 0x47, 0x41, 0x6b, 0xa8, 0x02, 0x6b, 0x13, 0x6c, 0x88, 0x8e, 0xa7, 0x84,
	0x9f, 0xf0, 0xd6, 0x45, 0x88, 0x19, 0x99, 0x5d, 0x81, 0xec, 0x62, 0x4a, 0x39, 0x1e, 0x83, 0xb5,
	0x28, 0xa6, 0x3c, 0xf3, 0xb1, 0x9e, 0x91, 0x48, 0xb9, 0x4e, 0x8a, 0x7e, 0xf8, 0x9e, 0x85, 0x7f,
	0x1e, 0xef, 0xab, 0x52, 0x73, 0xbe, 0x22, 0x96, 0xe4, 0x2b, 0x44, 0x33, 0xbe, 0xe2, 0xbb, 0xfa,
	0xfa, 0x17, 0xfc, 0x17, 0xdd, 0xe1, 0x83, 0x0e, 0x52, 0x72, 0x24, 0x70, 0x27, 0xb6, 0xc7, 0xe2,
	0xfc, 0x8d, 0x72, 0xb2, 0x50, 0x82, 0x5a, 0xa5, 0xc7, 0xb7, 0xaf, 0x97, 0x3f, 0x79, 0x98, 0x43,
	0xcb, 0x2f, 0x0d, 0x7c, 0xd2, 0xa7, 0x63, 0x84, 0x7b, 0xcb, 0x7b, 0x2f, 0x2e, 0xc7, 0xd8, 0x5f,
	0x51, 0xad, 0x70, 0xb6, 0x05, 0x4e, 0x11, 0x9a, 0xe8, 0xd7, 0x1b, 0x1a, 0x11, 0x89, 0x59, 0x25,
	0x13, 0xcd, 0xae, 0x2d, 0x99, 0x68, 0x6e, 0x93, 0x2b, 0x10, 0x89, 0xe7, 0x50, 0x3b, 0x1c, 0x8c,
	0x4c, 0x7d, 0x38, 0x32, 0xf5, 0xcf, 0x91, 0xa9, 0x3f, 0x8f, 0x4d, 0x6d, 0x38, 0x36, 0xb5, 0xf7,
	0xb1, 0xa9, 0x5d, 0xe6, 0xa6, 0xf2, 0xfb, 0x99, 0x52, 0xde, 0xf7, 0x09, 0x6b, 0xa6, 0xc4, 0x33,
	0x3a, 0xf8, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x22, 0x60, 0x81, 0x4a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a EthState by index.
	EthState(ctx context.Context, in *QueryGetEthStateRequest, opts ...grpc.CallOption) (*QueryGetEthStateResponse, error)
	// Queries a EthInput by index.
	EthInput(ctx context.Context, in *QueryGetEthInputRequest, opts ...grpc.CallOption) (*QueryGetEthInputResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/etherlink.etherlink.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EthState(ctx context.Context, in *QueryGetEthStateRequest, opts ...grpc.CallOption) (*QueryGetEthStateResponse, error) {
	out := new(QueryGetEthStateResponse)
	err := c.cc.Invoke(ctx, "/etherlink.etherlink.Query/EthState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EthInput(ctx context.Context, in *QueryGetEthInputRequest, opts ...grpc.CallOption) (*QueryGetEthInputResponse, error) {
	out := new(QueryGetEthInputResponse)
	err := c.cc.Invoke(ctx, "/etherlink.etherlink.Query/EthInput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a EthState by index.
	EthState(context.Context, *QueryGetEthStateRequest) (*QueryGetEthStateResponse, error)
	// Queries a EthInput by index.
	EthInput(context.Context, *QueryGetEthInputRequest) (*QueryGetEthInputResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) EthState(ctx context.Context, req *QueryGetEthStateRequest) (*QueryGetEthStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EthState not implemented")
}
func (*UnimplementedQueryServer) EthInput(ctx context.Context, req *QueryGetEthInputRequest) (*QueryGetEthInputResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EthInput not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etherlink.etherlink.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EthState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetEthStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EthState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etherlink.etherlink.Query/EthState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EthState(ctx, req.(*QueryGetEthStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EthInput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetEthInputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EthInput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/etherlink.etherlink.Query/EthInput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EthInput(ctx, req.(*QueryGetEthInputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "etherlink.etherlink.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "EthState",
			Handler:    _Query_EthState_Handler,
		},
		{
			MethodName: "EthInput",
			Handler:    _Query_EthInput_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "etherlink/etherlink/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetEthStateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetEthStateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetEthStateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryGetEthStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetEthStateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetEthStateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.EthState.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetEthInputRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetEthInputRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetEthInputRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryGetEthInputResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetEthInputResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetEthInputResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.EthInput.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetEthStateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryGetEthStateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EthState.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetEthInputRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryGetEthInputResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EthInput.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetEthStateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetEthStateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetEthStateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetEthStateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetEthStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetEthStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EthState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetEthInputRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetEthInputRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetEthInputRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetEthInputResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetEthInputResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetEthInputResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthInput", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EthInput.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
