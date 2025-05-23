// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: band/base/node/v1/query.proto

package node

import (
	context "context"
	fmt "fmt"
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

// ChainIDRequest is request type for the Service/ChainID RPC method.
type ChainIDRequest struct {
}

func (m *ChainIDRequest) Reset()         { *m = ChainIDRequest{} }
func (m *ChainIDRequest) String() string { return proto.CompactTextString(m) }
func (*ChainIDRequest) ProtoMessage()    {}
func (*ChainIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b9808b0dc6d368a, []int{0}
}
func (m *ChainIDRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainIDRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainIDRequest.Merge(m, src)
}
func (m *ChainIDRequest) XXX_Size() int {
	return m.Size()
}
func (m *ChainIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChainIDRequest proto.InternalMessageInfo

// ChainIDResponse is response type for the Service/ChainID RPC method.
type ChainIDResponse struct {
	ChainID string `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (m *ChainIDResponse) Reset()         { *m = ChainIDResponse{} }
func (m *ChainIDResponse) String() string { return proto.CompactTextString(m) }
func (*ChainIDResponse) ProtoMessage()    {}
func (*ChainIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b9808b0dc6d368a, []int{1}
}
func (m *ChainIDResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChainIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChainIDResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChainIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainIDResponse.Merge(m, src)
}
func (m *ChainIDResponse) XXX_Size() int {
	return m.Size()
}
func (m *ChainIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChainIDResponse proto.InternalMessageInfo

func (m *ChainIDResponse) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

// EVMValidatorsRequest is request type for the Service/EVMValidators RPC method.
type EVMValidatorsRequest struct {
}

func (m *EVMValidatorsRequest) Reset()         { *m = EVMValidatorsRequest{} }
func (m *EVMValidatorsRequest) String() string { return proto.CompactTextString(m) }
func (*EVMValidatorsRequest) ProtoMessage()    {}
func (*EVMValidatorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b9808b0dc6d368a, []int{2}
}
func (m *EVMValidatorsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EVMValidatorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EVMValidatorsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EVMValidatorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EVMValidatorsRequest.Merge(m, src)
}
func (m *EVMValidatorsRequest) XXX_Size() int {
	return m.Size()
}
func (m *EVMValidatorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EVMValidatorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EVMValidatorsRequest proto.InternalMessageInfo

// EVMValidatorsResponse is response type for the Service/EVMValidators RPC method.
type EVMValidatorsResponse struct {
	// BlockHeight is the latest block height
	BlockHeight int64 `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
	// Validators is list of validator's address and voting power
	Validators []ValidatorMinimal `protobuf:"bytes,2,rep,name=validators,proto3" json:"validators"`
}

func (m *EVMValidatorsResponse) Reset()         { *m = EVMValidatorsResponse{} }
func (m *EVMValidatorsResponse) String() string { return proto.CompactTextString(m) }
func (*EVMValidatorsResponse) ProtoMessage()    {}
func (*EVMValidatorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b9808b0dc6d368a, []int{3}
}
func (m *EVMValidatorsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EVMValidatorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EVMValidatorsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EVMValidatorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EVMValidatorsResponse.Merge(m, src)
}
func (m *EVMValidatorsResponse) XXX_Size() int {
	return m.Size()
}
func (m *EVMValidatorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EVMValidatorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EVMValidatorsResponse proto.InternalMessageInfo

func (m *EVMValidatorsResponse) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *EVMValidatorsResponse) GetValidators() []ValidatorMinimal {
	if m != nil {
		return m.Validators
	}
	return nil
}

// ValidatorMinimal is the data structure for storing validator's address and voting power
type ValidatorMinimal struct {
	Address     string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	VotingPower int64  `protobuf:"varint,2,opt,name=voting_power,json=votingPower,proto3" json:"voting_power,omitempty"`
}

func (m *ValidatorMinimal) Reset()         { *m = ValidatorMinimal{} }
func (m *ValidatorMinimal) String() string { return proto.CompactTextString(m) }
func (*ValidatorMinimal) ProtoMessage()    {}
func (*ValidatorMinimal) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b9808b0dc6d368a, []int{4}
}
func (m *ValidatorMinimal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorMinimal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorMinimal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorMinimal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorMinimal.Merge(m, src)
}
func (m *ValidatorMinimal) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorMinimal) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorMinimal.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorMinimal proto.InternalMessageInfo

func (m *ValidatorMinimal) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ValidatorMinimal) GetVotingPower() int64 {
	if m != nil {
		return m.VotingPower
	}
	return 0
}

func init() {
	proto.RegisterType((*ChainIDRequest)(nil), "band.base.node.v1.ChainIDRequest")
	proto.RegisterType((*ChainIDResponse)(nil), "band.base.node.v1.ChainIDResponse")
	proto.RegisterType((*EVMValidatorsRequest)(nil), "band.base.node.v1.EVMValidatorsRequest")
	proto.RegisterType((*EVMValidatorsResponse)(nil), "band.base.node.v1.EVMValidatorsResponse")
	proto.RegisterType((*ValidatorMinimal)(nil), "band.base.node.v1.ValidatorMinimal")
}

func init() { proto.RegisterFile("band/base/node/v1/query.proto", fileDescriptor_0b9808b0dc6d368a) }

var fileDescriptor_0b9808b0dc6d368a = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcf, 0x6f, 0xd3, 0x30,
	0x18, 0x6d, 0x3a, 0x44, 0xc1, 0xe5, 0xc7, 0xb0, 0xc6, 0x14, 0x55, 0xc3, 0xdb, 0x02, 0x82, 0x5e,
	0x88, 0xd5, 0xed, 0xc4, 0xb5, 0x80, 0x44, 0x85, 0x26, 0x50, 0x91, 0x76, 0xe0, 0x52, 0x39, 0x89,
	0xe5, 0x58, 0xa4, 0xfe, 0xb2, 0xd8, 0x0d, 0xe2, 0x8a, 0x04, 0xe2, 0x88, 0xc4, 0x3f, 0xb5, 0xe3,
	0x24, 0x2e, 0x9c, 0x26, 0x94, 0xf2, 0x87, 0x20, 0x3b, 0x4d, 0xa1, 0xa3, 0x62, 0x37, 0xfb, 0xbd,
	0xf7, 0x3d, 0x3d, 0x3f, 0x7f, 0xe8, 0x5e, 0xc4, 0x54, 0x42, 0x23, 0xa6, 0x39, 0x55, 0x90, 0x70,
	0x5a, 0x0e, 0xe8, 0xc9, 0x8c, 0x17, 0x1f, 0xc2, 0xbc, 0x00, 0x03, 0xf8, 0x8e, 0xa5, 0x43, 0x4b,
	0x87, 0x96, 0x0e, 0xcb, 0x41, 0x6f, 0x47, 0x00, 0x88, 0x8c, 0x53, 0x96, 0x4b, 0xca, 0x94, 0x02,
	0xc3, 0x8c, 0x04, 0xa5, 0xeb, 0x81, 0xde, 0x96, 0x00, 0x01, 0xee, 0x48, 0xed, 0xa9, 0x46, 0x83,
	0x4d, 0x74, 0xeb, 0x69, 0xca, 0xa4, 0x1a, 0x3d, 0x1b, 0xf3, 0x93, 0x19, 0xd7, 0x26, 0x78, 0x82,
	0x6e, 0x2f, 0x11, 0x9d, 0x83, 0xd2, 0x1c, 0x3f, 0x44, 0xd7, 0x62, 0x0b, 0x4d, 0x64, 0xe2, 0x7b,
	0x7b, 0x5e, 0xff, 0xfa, 0xb0, 0x5b, 0x9d, 0xef, 0x76, 0x1a, 0x59, 0xc7, 0x91, 0xa3, 0x24, 0xd8,
	0x46, 0x5b, 0xcf, 0x8f, 0x8f, 0x8e, 0x59, 0x26, 0x13, 0x66, 0xa0, 0xd0, 0x8d, 0xe5, 0x27, 0x0f,
	0xdd, 0xbd, 0x40, 0x2c, 0x9c, 0xf7, 0xd1, 0x8d, 0x28, 0x83, 0xf8, 0xdd, 0x24, 0xe5, 0x52, 0xa4,
	0xc6, 0xb9, 0x6f, 0x8c, 0xbb, 0x0e, 0x7b, 0xe1, 0x20, 0x3c, 0x42, 0xa8, 0x5c, 0x0e, 0xfa, 0xed,
	0xbd, 0x8d, 0x7e, 0xf7, 0xe0, 0x7e, 0xf8, 0xcf, 0xeb, 0xc3, 0xa5, 0xfb, 0x91, 0x54, 0x72, 0xca,
	0xb2, 0xe1, 0x95, 0xd3, 0xf3, 0xdd, 0xd6, 0xf8, 0xaf, 0xe1, 0xe0, 0x15, 0xda, 0xbc, 0xa8, 0xc2,
	0x3e, 0xea, 0xb0, 0x24, 0x29, 0xb8, 0xd6, 0xf5, 0xd3, 0xc6, 0xcd, 0xd5, 0x66, 0x2b, 0xc1, 0x48,
	0x25, 0x26, 0x39, 0xbc, 0xe7, 0x85, 0xdf, 0xae, 0xb3, 0xd5, 0xd8, 0x6b, 0x0b, 0x1d, 0x7c, 0x6e,
	0xa3, 0xce, 0x1b, 0x5e, 0x94, 0x32, 0xe6, 0x38, 0x47, 0x4d, 0x21, 0x78, 0x7f, 0x4d, 0xbc, 0xd5,
	0x96, 0x7b, 0xc1, 0xff, 0x24, 0x75, 0x39, 0x01, 0xf9, 0xf8, 0xfd, 0xd7, 0xb7, 0xb6, 0x8f, 0xb7,
	0xa9, 0xd5, 0xba, 0x92, 0xed, 0x16, 0x34, 0x5f, 0x81, 0xbf, 0x78, 0xe8, 0xe6, 0x4a, 0xad, 0xf8,
	0xd1, 0x1a, 0xd7, 0x75, 0x3f, 0xd2, 0xeb, 0x5f, 0x2e, 0x5c, 0x84, 0x78, 0xe0, 0x42, 0x10, 0xbc,
	0xb3, 0x1a, 0x82, 0x97, 0xd3, 0xc7, 0x7f, 0x9a, 0x1d, 0xbe, 0x3c, 0xad, 0x88, 0x77, 0x56, 0x11,
	0xef, 0x67, 0x45, 0xbc, 0xaf, 0x73, 0xd2, 0x3a, 0x9b, 0x93, 0xd6, 0x8f, 0x39, 0x69, 0xbd, 0x1d,
	0x08, 0x69, 0xd2, 0x59, 0x14, 0xc6, 0x30, 0x75, 0x0e, 0x6e, 0xed, 0x62, 0xc8, 0xe8, 0xc2, 0xea,
	0x90, 0xc6, 0x99, 0xe4, 0xca, 0x50, 0x51, 0xe4, 0xb1, 0xdb, 0xf4, 0xe8, 0xaa, 0xd3, 0x1c, 0xfe,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0x92, 0x7c, 0xf9, 0x69, 0x02, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	// ChainID queries the chain ID of this node
	ChainID(ctx context.Context, in *ChainIDRequest, opts ...grpc.CallOption) (*ChainIDResponse, error)
	// EVMValidators queries current list of validator's address and power
	EVMValidators(ctx context.Context, in *EVMValidatorsRequest, opts ...grpc.CallOption) (*EVMValidatorsResponse, error)
}

type serviceClient struct {
	cc grpc1.ClientConn
}

func NewServiceClient(cc grpc1.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) ChainID(ctx context.Context, in *ChainIDRequest, opts ...grpc.CallOption) (*ChainIDResponse, error) {
	out := new(ChainIDResponse)
	err := c.cc.Invoke(ctx, "/band.base.node.v1.Service/ChainID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) EVMValidators(ctx context.Context, in *EVMValidatorsRequest, opts ...grpc.CallOption) (*EVMValidatorsResponse, error) {
	out := new(EVMValidatorsResponse)
	err := c.cc.Invoke(ctx, "/band.base.node.v1.Service/EVMValidators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	// ChainID queries the chain ID of this node
	ChainID(context.Context, *ChainIDRequest) (*ChainIDResponse, error)
	// EVMValidators queries current list of validator's address and power
	EVMValidators(context.Context, *EVMValidatorsRequest) (*EVMValidatorsResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) ChainID(ctx context.Context, req *ChainIDRequest) (*ChainIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChainID not implemented")
}
func (*UnimplementedServiceServer) EVMValidators(ctx context.Context, req *EVMValidatorsRequest) (*EVMValidatorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EVMValidators not implemented")
}

func RegisterServiceServer(s grpc1.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_ChainID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChainIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).ChainID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/band.base.node.v1.Service/ChainID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).ChainID(ctx, req.(*ChainIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_EVMValidators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EVMValidatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).EVMValidators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/band.base.node.v1.Service/EVMValidators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).EVMValidators(ctx, req.(*EVMValidatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "band.base.node.v1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ChainID",
			Handler:    _Service_ChainID_Handler,
		},
		{
			MethodName: "EVMValidators",
			Handler:    _Service_EVMValidators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "band/base/node/v1/query.proto",
}

func (m *ChainIDRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainIDRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainIDRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ChainIDResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChainIDResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChainIDResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EVMValidatorsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EVMValidatorsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EVMValidatorsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *EVMValidatorsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EVMValidatorsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EVMValidatorsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Validators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.BlockHeight != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorMinimal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorMinimal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorMinimal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.VotingPower != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.VotingPower))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
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
func (m *ChainIDRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ChainIDResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *EVMValidatorsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *EVMValidatorsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeight != 0 {
		n += 1 + sovQuery(uint64(m.BlockHeight))
	}
	if len(m.Validators) > 0 {
		for _, e := range m.Validators {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *ValidatorMinimal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.VotingPower != 0 {
		n += 1 + sovQuery(uint64(m.VotingPower))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChainIDRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ChainIDRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainIDRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *ChainIDResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ChainIDResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChainIDResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
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
func (m *EVMValidatorsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EVMValidatorsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EVMValidatorsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *EVMValidatorsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EVMValidatorsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EVMValidatorsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
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
			m.Validators = append(m.Validators, ValidatorMinimal{})
			if err := m.Validators[len(m.Validators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *ValidatorMinimal) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ValidatorMinimal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorMinimal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingPower", wireType)
			}
			m.VotingPower = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotingPower |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
