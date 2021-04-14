// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stateservice.proto

package lnrpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type WalletState int32

const (
	WalletState_NON_EXISTING WalletState = 0
	WalletState_LOCKED       WalletState = 1
	WalletState_UNLOCKED     WalletState = 2
	WalletState_RPC_ACTIVE   WalletState = 3
)

var WalletState_name = map[int32]string{
	0: "NON_EXISTING",
	1: "LOCKED",
	2: "UNLOCKED",
	3: "RPC_ACTIVE",
}

var WalletState_value = map[string]int32{
	"NON_EXISTING": 0,
	"LOCKED":       1,
	"UNLOCKED":     2,
	"RPC_ACTIVE":   3,
}

func (x WalletState) String() string {
	return proto.EnumName(WalletState_name, int32(x))
}

func (WalletState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9cd64a11048f05dd, []int{0}
}

type SubscribeStateRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeStateRequest) Reset()         { *m = SubscribeStateRequest{} }
func (m *SubscribeStateRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeStateRequest) ProtoMessage()    {}
func (*SubscribeStateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cd64a11048f05dd, []int{0}
}

func (m *SubscribeStateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeStateRequest.Unmarshal(m, b)
}
func (m *SubscribeStateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeStateRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeStateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeStateRequest.Merge(m, src)
}
func (m *SubscribeStateRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeStateRequest.Size(m)
}
func (m *SubscribeStateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeStateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeStateRequest proto.InternalMessageInfo

type SubscribeStateResponse struct {
	State                WalletState `protobuf:"varint,1,opt,name=state,proto3,enum=lnrpc.WalletState" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SubscribeStateResponse) Reset()         { *m = SubscribeStateResponse{} }
func (m *SubscribeStateResponse) String() string { return proto.CompactTextString(m) }
func (*SubscribeStateResponse) ProtoMessage()    {}
func (*SubscribeStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cd64a11048f05dd, []int{1}
}

func (m *SubscribeStateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeStateResponse.Unmarshal(m, b)
}
func (m *SubscribeStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeStateResponse.Marshal(b, m, deterministic)
}
func (m *SubscribeStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeStateResponse.Merge(m, src)
}
func (m *SubscribeStateResponse) XXX_Size() int {
	return xxx_messageInfo_SubscribeStateResponse.Size(m)
}
func (m *SubscribeStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeStateResponse proto.InternalMessageInfo

func (m *SubscribeStateResponse) GetState() WalletState {
	if m != nil {
		return m.State
	}
	return WalletState_NON_EXISTING
}

func init() {
	proto.RegisterEnum("lnrpc.WalletState", WalletState_name, WalletState_value)
	proto.RegisterType((*SubscribeStateRequest)(nil), "lnrpc.SubscribeStateRequest")
	proto.RegisterType((*SubscribeStateResponse)(nil), "lnrpc.SubscribeStateResponse")
}

func init() { proto.RegisterFile("stateservice.proto", fileDescriptor_9cd64a11048f05dd) }

var fileDescriptor_9cd64a11048f05dd = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x2e, 0x49, 0x2c,
	0x49, 0x2d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0xcd, 0xc9, 0x2b, 0x2a, 0x48, 0x56, 0x12, 0xe7, 0x12, 0x0d, 0x2e, 0x4d, 0x2a, 0x4e, 0x2e, 0xca,
	0x4c, 0x4a, 0x0d, 0x06, 0xa9, 0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x51, 0x72, 0xe2, 0x12,
	0x43, 0x97, 0x28, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0xd2, 0xe0, 0x62, 0x05, 0x9b, 0x27, 0xc1,
	0xa8, 0xc0, 0xa8, 0xc1, 0x67, 0x24, 0xa4, 0x07, 0x36, 0x49, 0x2f, 0x3c, 0x31, 0x27, 0x27, 0xb5,
	0x04, 0xa2, 0x14, 0xa2, 0x40, 0xcb, 0x93, 0x8b, 0x1b, 0x49, 0x54, 0x48, 0x80, 0x8b, 0xc7, 0xcf,
	0xdf, 0x2f, 0xde, 0x35, 0xc2, 0x33, 0x38, 0xc4, 0xd3, 0xcf, 0x5d, 0x80, 0x41, 0x88, 0x8b, 0x8b,
	0xcd, 0xc7, 0xdf, 0xd9, 0xdb, 0xd5, 0x45, 0x80, 0x51, 0x88, 0x87, 0x8b, 0x23, 0xd4, 0x0f, 0xca,
	0x63, 0x12, 0xe2, 0xe3, 0xe2, 0x0a, 0x0a, 0x70, 0x8e, 0x77, 0x74, 0x0e, 0xf1, 0x0c, 0x73, 0x15,
	0x60, 0x36, 0x8a, 0xe0, 0x62, 0x85, 0x18, 0xe2, 0xcf, 0xc5, 0x87, 0xea, 0x2e, 0x21, 0x19, 0xa8,
	0x03, 0xb0, 0xfa, 0x43, 0x4a, 0x16, 0x87, 0x2c, 0xc4, 0x33, 0x06, 0x8c, 0x4e, 0xea, 0x51, 0xaa,
	0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x39, 0x99, 0xe9, 0x19, 0x25,
	0x79, 0x99, 0x79, 0xe9, 0x79, 0xa9, 0x25, 0xe5, 0xf9, 0x45, 0xd9, 0xfa, 0x39, 0x79, 0x29, 0xfa,
	0x60, 0x13, 0x92, 0xd8, 0xc0, 0x01, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x75, 0x6a, 0x2f,
	0xe2, 0x4e, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StateClient is the client API for State service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StateClient interface {
	// SubscribeState subscribes to the state of the wallet. The current wallet
	// state will always be delivered immediately.
	SubscribeState(ctx context.Context, in *SubscribeStateRequest, opts ...grpc.CallOption) (State_SubscribeStateClient, error)
}

type stateClient struct {
	cc *grpc.ClientConn
}

func NewStateClient(cc *grpc.ClientConn) StateClient {
	return &stateClient{cc}
}

func (c *stateClient) SubscribeState(ctx context.Context, in *SubscribeStateRequest, opts ...grpc.CallOption) (State_SubscribeStateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_State_serviceDesc.Streams[0], "/lnrpc.State/SubscribeState", opts...)
	if err != nil {
		return nil, err
	}
	x := &stateSubscribeStateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type State_SubscribeStateClient interface {
	Recv() (*SubscribeStateResponse, error)
	grpc.ClientStream
}

type stateSubscribeStateClient struct {
	grpc.ClientStream
}

func (x *stateSubscribeStateClient) Recv() (*SubscribeStateResponse, error) {
	m := new(SubscribeStateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StateServer is the server API for State service.
type StateServer interface {
	// SubscribeState subscribes to the state of the wallet. The current wallet
	// state will always be delivered immediately.
	SubscribeState(*SubscribeStateRequest, State_SubscribeStateServer) error
}

// UnimplementedStateServer can be embedded to have forward compatible implementations.
type UnimplementedStateServer struct {
}

func (*UnimplementedStateServer) SubscribeState(req *SubscribeStateRequest, srv State_SubscribeStateServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeState not implemented")
}

func RegisterStateServer(s *grpc.Server, srv StateServer) {
	s.RegisterService(&_State_serviceDesc, srv)
}

func _State_SubscribeState_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeStateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StateServer).SubscribeState(m, &stateSubscribeStateServer{stream})
}

type State_SubscribeStateServer interface {
	Send(*SubscribeStateResponse) error
	grpc.ServerStream
}

type stateSubscribeStateServer struct {
	grpc.ServerStream
}

func (x *stateSubscribeStateServer) Send(m *SubscribeStateResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _State_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lnrpc.State",
	HandlerType: (*StateServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeState",
			Handler:       _State_SubscribeState_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "stateservice.proto",
}