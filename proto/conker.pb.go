// Code generated by protoc-gen-go. DO NOT EDIT.
// source: conker.proto

package conker_v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ConnectRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectRequest) Reset()         { *m = ConnectRequest{} }
func (m *ConnectRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectRequest) ProtoMessage()    {}
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_751d5aef113abfed, []int{0}
}

func (m *ConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectRequest.Unmarshal(m, b)
}
func (m *ConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectRequest.Marshal(b, m, deterministic)
}
func (m *ConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectRequest.Merge(m, src)
}
func (m *ConnectRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectRequest.Size(m)
}
func (m *ConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectRequest proto.InternalMessageInfo

type ControlMessage struct {
	Destination          string            `protobuf:"bytes,1,opt,name=destination,proto3" json:"destination,omitempty"`
	Payload              []byte            `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Txid                 string            `protobuf:"bytes,3,opt,name=txid,proto3" json:"txid,omitempty"`
	Meta                 map[string]string `protobuf:"bytes,4,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ControlMessage) Reset()         { *m = ControlMessage{} }
func (m *ControlMessage) String() string { return proto.CompactTextString(m) }
func (*ControlMessage) ProtoMessage()    {}
func (*ControlMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_751d5aef113abfed, []int{1}
}

func (m *ControlMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlMessage.Unmarshal(m, b)
}
func (m *ControlMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlMessage.Marshal(b, m, deterministic)
}
func (m *ControlMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlMessage.Merge(m, src)
}
func (m *ControlMessage) XXX_Size() int {
	return xxx_messageInfo_ControlMessage.Size(m)
}
func (m *ControlMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ControlMessage proto.InternalMessageInfo

func (m *ControlMessage) GetDestination() string {
	if m != nil {
		return m.Destination
	}
	return ""
}

func (m *ControlMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ControlMessage) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

func (m *ControlMessage) GetMeta() map[string]string {
	if m != nil {
		return m.Meta
	}
	return nil
}

type ControlEnd struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ControlEnd) Reset()         { *m = ControlEnd{} }
func (m *ControlEnd) String() string { return proto.CompactTextString(m) }
func (*ControlEnd) ProtoMessage()    {}
func (*ControlEnd) Descriptor() ([]byte, []int) {
	return fileDescriptor_751d5aef113abfed, []int{2}
}

func (m *ControlEnd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlEnd.Unmarshal(m, b)
}
func (m *ControlEnd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlEnd.Marshal(b, m, deterministic)
}
func (m *ControlEnd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlEnd.Merge(m, src)
}
func (m *ControlEnd) XXX_Size() int {
	return xxx_messageInfo_ControlEnd.Size(m)
}
func (m *ControlEnd) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlEnd.DiscardUnknown(m)
}

var xxx_messageInfo_ControlEnd proto.InternalMessageInfo

func (m *ControlEnd) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *ControlEnd) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type ControlAck struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Stage                string   `protobuf:"bytes,2,opt,name=stage,proto3" json:"stage,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Txid                 string   `protobuf:"bytes,4,opt,name=txid,proto3" json:"txid,omitempty"`
	Seq                  int32    `protobuf:"varint,5,opt,name=seq,proto3" json:"seq,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ControlAck) Reset()         { *m = ControlAck{} }
func (m *ControlAck) String() string { return proto.CompactTextString(m) }
func (*ControlAck) ProtoMessage()    {}
func (*ControlAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_751d5aef113abfed, []int{3}
}

func (m *ControlAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ControlAck.Unmarshal(m, b)
}
func (m *ControlAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ControlAck.Marshal(b, m, deterministic)
}
func (m *ControlAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ControlAck.Merge(m, src)
}
func (m *ControlAck) XXX_Size() int {
	return xxx_messageInfo_ControlAck.Size(m)
}
func (m *ControlAck) XXX_DiscardUnknown() {
	xxx_messageInfo_ControlAck.DiscardUnknown(m)
}

var xxx_messageInfo_ControlAck proto.InternalMessageInfo

func (m *ControlAck) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *ControlAck) GetStage() string {
	if m != nil {
		return m.Stage
	}
	return ""
}

func (m *ControlAck) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ControlAck) GetTxid() string {
	if m != nil {
		return m.Txid
	}
	return ""
}

func (m *ControlAck) GetSeq() int32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

type ConnectAck struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectAck) Reset()         { *m = ConnectAck{} }
func (m *ConnectAck) String() string { return proto.CompactTextString(m) }
func (*ConnectAck) ProtoMessage()    {}
func (*ConnectAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_751d5aef113abfed, []int{4}
}

func (m *ConnectAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectAck.Unmarshal(m, b)
}
func (m *ConnectAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectAck.Marshal(b, m, deterministic)
}
func (m *ConnectAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectAck.Merge(m, src)
}
func (m *ConnectAck) XXX_Size() int {
	return xxx_messageInfo_ConnectAck.Size(m)
}
func (m *ConnectAck) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectAck.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectAck proto.InternalMessageInfo

type StatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusRequest) Reset()         { *m = StatusRequest{} }
func (m *StatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()    {}
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_751d5aef113abfed, []int{5}
}

func (m *StatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusRequest.Unmarshal(m, b)
}
func (m *StatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusRequest.Marshal(b, m, deterministic)
}
func (m *StatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusRequest.Merge(m, src)
}
func (m *StatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatusRequest.Size(m)
}
func (m *StatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatusRequest proto.InternalMessageInfo

type StatusMessage struct {
	ConnectedRobots      []string `protobuf:"bytes,1,rep,name=ConnectedRobots,proto3" json:"ConnectedRobots,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusMessage) Reset()         { *m = StatusMessage{} }
func (m *StatusMessage) String() string { return proto.CompactTextString(m) }
func (*StatusMessage) ProtoMessage()    {}
func (*StatusMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_751d5aef113abfed, []int{6}
}

func (m *StatusMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusMessage.Unmarshal(m, b)
}
func (m *StatusMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusMessage.Marshal(b, m, deterministic)
}
func (m *StatusMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusMessage.Merge(m, src)
}
func (m *StatusMessage) XXX_Size() int {
	return xxx_messageInfo_StatusMessage.Size(m)
}
func (m *StatusMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StatusMessage proto.InternalMessageInfo

func (m *StatusMessage) GetConnectedRobots() []string {
	if m != nil {
		return m.ConnectedRobots
	}
	return nil
}

func init() {
	proto.RegisterType((*ConnectRequest)(nil), "conker.v1.ConnectRequest")
	proto.RegisterType((*ControlMessage)(nil), "conker.v1.ControlMessage")
	proto.RegisterMapType((map[string]string)(nil), "conker.v1.ControlMessage.MetaEntry")
	proto.RegisterType((*ControlEnd)(nil), "conker.v1.ControlEnd")
	proto.RegisterType((*ControlAck)(nil), "conker.v1.ControlAck")
	proto.RegisterType((*ConnectAck)(nil), "conker.v1.ConnectAck")
	proto.RegisterType((*StatusRequest)(nil), "conker.v1.StatusRequest")
	proto.RegisterType((*StatusMessage)(nil), "conker.v1.StatusMessage")
}

func init() { proto.RegisterFile("conker.proto", fileDescriptor_751d5aef113abfed) }

var fileDescriptor_751d5aef113abfed = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xed, 0xd6, 0x49, 0xc1, 0x53, 0x43, 0xa3, 0x15, 0x20, 0x27, 0x12, 0x52, 0x64, 0x2e, 0x3e,
	0x45, 0xa1, 0x1c, 0x5a, 0xb8, 0x45, 0x25, 0x88, 0x4b, 0x2f, 0x1b, 0x21, 0xce, 0xdb, 0xf5, 0xa8,
	0x44, 0x76, 0x76, 0x5b, 0xef, 0xb8, 0x22, 0x12, 0x5f, 0xc1, 0xa7, 0xf1, 0x45, 0x68, 0xb7, 0x6b,
	0xcb, 0x29, 0x01, 0x89, 0xde, 0xe6, 0x3d, 0xf9, 0xbd, 0x79, 0x33, 0xb3, 0x86, 0x44, 0x19, 0x5d,
	0x62, 0x3d, 0xbb, 0xa9, 0x0d, 0x19, 0x1e, 0x07, 0x74, 0xf7, 0x36, 0x1b, 0xc1, 0xf3, 0x0b, 0xa3,
	0x35, 0x2a, 0x12, 0x78, 0xdb, 0xa0, 0xa5, 0xec, 0x17, 0xf3, 0x14, 0xd5, 0xa6, 0xba, 0x44, 0x6b,
	0xe5, 0x35, 0xf2, 0x29, 0x1c, 0x17, 0x68, 0x69, 0xad, 0x25, 0xad, 0x8d, 0x4e, 0xd9, 0x94, 0xe5,
	0xb1, 0xe8, 0x53, 0x3c, 0x85, 0x27, 0x37, 0x72, 0x5b, 0x19, 0x59, 0xa4, 0x87, 0x53, 0x96, 0x27,
	0xa2, 0x85, 0x9c, 0xc3, 0x80, 0xbe, 0xaf, 0x8b, 0x34, 0xf2, 0x22, 0x5f, 0xf3, 0x33, 0x18, 0x6c,
	0x90, 0x64, 0x3a, 0x98, 0x46, 0xf9, 0xf1, 0xe9, 0x9b, 0x59, 0x17, 0x67, 0xb6, 0xdb, 0x78, 0x76,
	0x89, 0x24, 0x97, 0x9a, 0xea, 0xad, 0xf0, 0x82, 0xc9, 0x19, 0xc4, 0x1d, 0xc5, 0x47, 0x10, 0x95,
	0xb8, 0x0d, 0x69, 0x5c, 0xc9, 0x5f, 0xc0, 0xf0, 0x4e, 0x56, 0x0d, 0xfa, 0x0c, 0xb1, 0xb8, 0x07,
	0x1f, 0x0e, 0xcf, 0x59, 0x76, 0x0e, 0x10, 0xac, 0x97, 0xda, 0x67, 0x52, 0xa6, 0x40, 0x2f, 0x1d,
	0x0a, 0x5f, 0xf3, 0x57, 0x70, 0x54, 0xa3, 0xb4, 0x46, 0x07, 0x71, 0x40, 0xd9, 0x8f, 0x4e, 0xb9,
	0x50, 0xa5, 0x9b, 0xd3, 0x36, 0x4a, 0xa1, 0xb5, 0x5e, 0xfc, 0x54, 0xb4, 0xd0, 0xf5, 0xb6, 0x24,
	0xaf, 0xbb, 0xde, 0x1e, 0xb8, 0xef, 0x37, 0xf7, 0xb3, 0x84, 0x05, 0xb4, 0xb0, 0xdb, 0xcb, 0xa0,
	0xb7, 0x97, 0x11, 0x44, 0x16, 0x6f, 0xd3, 0xa1, 0x8f, 0xe5, 0xca, 0x2c, 0xf1, 0xdd, 0xdd, 0x79,
	0x16, 0xaa, 0xcc, 0x4e, 0xe0, 0xd9, 0x8a, 0x24, 0x35, 0xb6, 0xbd, 0xd5, 0xfb, 0x96, 0x68, 0x2f,
	0x95, 0xc3, 0x49, 0xf8, 0x1e, 0x0b, 0x61, 0xae, 0x0c, 0xb9, 0x9c, 0x51, 0x1e, 0x8b, 0x87, 0xf4,
	0xe9, 0xcf, 0x08, 0x12, 0x5f, 0x86, 0xe9, 0xf8, 0x67, 0x48, 0xda, 0x97, 0xe0, 0x68, 0x3e, 0xde,
	0x3d, 0x4b, 0xef, 0x89, 0x4c, 0xc6, 0x7f, 0xbd, 0x58, 0x76, 0x30, 0x67, 0x7c, 0x05, 0xe3, 0xbe,
	0xd3, 0xd7, 0x35, 0x7d, 0x5b, 0xa8, 0x72, 0x45, 0x35, 0xca, 0x0d, 0x7f, 0xf9, 0xa7, 0x76, 0xa1,
	0xca, 0x7f, 0x5a, 0xe6, 0x6c, 0xce, 0xf8, 0x47, 0x1f, 0xcf, 0xf1, 0x7b, 0xe3, 0xf5, 0x04, 0x93,
	0x3d, 0x2d, 0x96, 0xba, 0x70, 0x3e, 0xfc, 0x13, 0x8c, 0xfa, 0x2e, 0x17, 0xb2, 0xaa, 0xfe, 0xd3,
	0xc9, 0xdd, 0xe1, 0x80, 0x7f, 0x81, 0xd7, 0x0f, 0x7d, 0x76, 0xc7, 0x7c, 0x84, 0xe9, 0x9c, 0x5d,
	0x1d, 0xf9, 0xff, 0xf3, 0xdd, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x11, 0xb1, 0x82, 0xee, 0xaf,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RobotControlClient is the client API for RobotControl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RobotControlClient interface {
	ConnectRobot(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (RobotControl_ConnectRobotClient, error)
	ConnectRobotWithAckStream(ctx context.Context, opts ...grpc.CallOption) (RobotControl_ConnectRobotWithAckStreamClient, error)
	ControlRobot(ctx context.Context, opts ...grpc.CallOption) (RobotControl_ControlRobotClient, error)
	ControlRobotCall(ctx context.Context, in *ControlMessage, opts ...grpc.CallOption) (*ControlAck, error)
	ControlRobotCallWithAckStream(ctx context.Context, in *ControlMessage, opts ...grpc.CallOption) (RobotControl_ControlRobotCallWithAckStreamClient, error)
}

type robotControlClient struct {
	cc *grpc.ClientConn
}

func NewRobotControlClient(cc *grpc.ClientConn) RobotControlClient {
	return &robotControlClient{cc}
}

func (c *robotControlClient) ConnectRobot(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (RobotControl_ConnectRobotClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RobotControl_serviceDesc.Streams[0], "/conker.v1.RobotControl/ConnectRobot", opts...)
	if err != nil {
		return nil, err
	}
	x := &robotControlConnectRobotClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RobotControl_ConnectRobotClient interface {
	Recv() (*ControlMessage, error)
	grpc.ClientStream
}

type robotControlConnectRobotClient struct {
	grpc.ClientStream
}

func (x *robotControlConnectRobotClient) Recv() (*ControlMessage, error) {
	m := new(ControlMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *robotControlClient) ConnectRobotWithAckStream(ctx context.Context, opts ...grpc.CallOption) (RobotControl_ConnectRobotWithAckStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RobotControl_serviceDesc.Streams[1], "/conker.v1.RobotControl/ConnectRobotWithAckStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &robotControlConnectRobotWithAckStreamClient{stream}
	return x, nil
}

type RobotControl_ConnectRobotWithAckStreamClient interface {
	Send(*ControlAck) error
	Recv() (*ControlMessage, error)
	grpc.ClientStream
}

type robotControlConnectRobotWithAckStreamClient struct {
	grpc.ClientStream
}

func (x *robotControlConnectRobotWithAckStreamClient) Send(m *ControlAck) error {
	return x.ClientStream.SendMsg(m)
}

func (x *robotControlConnectRobotWithAckStreamClient) Recv() (*ControlMessage, error) {
	m := new(ControlMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *robotControlClient) ControlRobot(ctx context.Context, opts ...grpc.CallOption) (RobotControl_ControlRobotClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RobotControl_serviceDesc.Streams[2], "/conker.v1.RobotControl/ControlRobot", opts...)
	if err != nil {
		return nil, err
	}
	x := &robotControlControlRobotClient{stream}
	return x, nil
}

type RobotControl_ControlRobotClient interface {
	Send(*ControlMessage) error
	CloseAndRecv() (*ControlEnd, error)
	grpc.ClientStream
}

type robotControlControlRobotClient struct {
	grpc.ClientStream
}

func (x *robotControlControlRobotClient) Send(m *ControlMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *robotControlControlRobotClient) CloseAndRecv() (*ControlEnd, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ControlEnd)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *robotControlClient) ControlRobotCall(ctx context.Context, in *ControlMessage, opts ...grpc.CallOption) (*ControlAck, error) {
	out := new(ControlAck)
	err := c.cc.Invoke(ctx, "/conker.v1.RobotControl/ControlRobotCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *robotControlClient) ControlRobotCallWithAckStream(ctx context.Context, in *ControlMessage, opts ...grpc.CallOption) (RobotControl_ControlRobotCallWithAckStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RobotControl_serviceDesc.Streams[3], "/conker.v1.RobotControl/ControlRobotCallWithAckStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &robotControlControlRobotCallWithAckStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RobotControl_ControlRobotCallWithAckStreamClient interface {
	Recv() (*ControlAck, error)
	grpc.ClientStream
}

type robotControlControlRobotCallWithAckStreamClient struct {
	grpc.ClientStream
}

func (x *robotControlControlRobotCallWithAckStreamClient) Recv() (*ControlAck, error) {
	m := new(ControlAck)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RobotControlServer is the server API for RobotControl service.
type RobotControlServer interface {
	ConnectRobot(*ConnectRequest, RobotControl_ConnectRobotServer) error
	ConnectRobotWithAckStream(RobotControl_ConnectRobotWithAckStreamServer) error
	ControlRobot(RobotControl_ControlRobotServer) error
	ControlRobotCall(context.Context, *ControlMessage) (*ControlAck, error)
	ControlRobotCallWithAckStream(*ControlMessage, RobotControl_ControlRobotCallWithAckStreamServer) error
}

func RegisterRobotControlServer(s *grpc.Server, srv RobotControlServer) {
	s.RegisterService(&_RobotControl_serviceDesc, srv)
}

func _RobotControl_ConnectRobot_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConnectRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RobotControlServer).ConnectRobot(m, &robotControlConnectRobotServer{stream})
}

type RobotControl_ConnectRobotServer interface {
	Send(*ControlMessage) error
	grpc.ServerStream
}

type robotControlConnectRobotServer struct {
	grpc.ServerStream
}

func (x *robotControlConnectRobotServer) Send(m *ControlMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _RobotControl_ConnectRobotWithAckStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RobotControlServer).ConnectRobotWithAckStream(&robotControlConnectRobotWithAckStreamServer{stream})
}

type RobotControl_ConnectRobotWithAckStreamServer interface {
	Send(*ControlMessage) error
	Recv() (*ControlAck, error)
	grpc.ServerStream
}

type robotControlConnectRobotWithAckStreamServer struct {
	grpc.ServerStream
}

func (x *robotControlConnectRobotWithAckStreamServer) Send(m *ControlMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *robotControlConnectRobotWithAckStreamServer) Recv() (*ControlAck, error) {
	m := new(ControlAck)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RobotControl_ControlRobot_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RobotControlServer).ControlRobot(&robotControlControlRobotServer{stream})
}

type RobotControl_ControlRobotServer interface {
	SendAndClose(*ControlEnd) error
	Recv() (*ControlMessage, error)
	grpc.ServerStream
}

type robotControlControlRobotServer struct {
	grpc.ServerStream
}

func (x *robotControlControlRobotServer) SendAndClose(m *ControlEnd) error {
	return x.ServerStream.SendMsg(m)
}

func (x *robotControlControlRobotServer) Recv() (*ControlMessage, error) {
	m := new(ControlMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RobotControl_ControlRobotCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControlMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RobotControlServer).ControlRobotCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/conker.v1.RobotControl/ControlRobotCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RobotControlServer).ControlRobotCall(ctx, req.(*ControlMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _RobotControl_ControlRobotCallWithAckStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ControlMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RobotControlServer).ControlRobotCallWithAckStream(m, &robotControlControlRobotCallWithAckStreamServer{stream})
}

type RobotControl_ControlRobotCallWithAckStreamServer interface {
	Send(*ControlAck) error
	grpc.ServerStream
}

type robotControlControlRobotCallWithAckStreamServer struct {
	grpc.ServerStream
}

func (x *robotControlControlRobotCallWithAckStreamServer) Send(m *ControlAck) error {
	return x.ServerStream.SendMsg(m)
}

var _RobotControl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "conker.v1.RobotControl",
	HandlerType: (*RobotControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ControlRobotCall",
			Handler:    _RobotControl_ControlRobotCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ConnectRobot",
			Handler:       _RobotControl_ConnectRobot_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ConnectRobotWithAckStream",
			Handler:       _RobotControl_ConnectRobotWithAckStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ControlRobot",
			Handler:       _RobotControl_ControlRobot_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ControlRobotCallWithAckStream",
			Handler:       _RobotControl_ControlRobotCallWithAckStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "conker.proto",
}
