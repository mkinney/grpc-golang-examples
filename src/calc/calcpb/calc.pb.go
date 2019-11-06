// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calc/calcpb/calc.proto

package calcpb

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

type Calc struct {
	FirstValue           int64    `protobuf:"varint,1,opt,name=first_value,json=firstValue,proto3" json:"first_value,omitempty"`
	SecondValue          int64    `protobuf:"varint,2,opt,name=second_value,json=secondValue,proto3" json:"second_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Calc) Reset()         { *m = Calc{} }
func (m *Calc) String() string { return proto.CompactTextString(m) }
func (*Calc) ProtoMessage()    {}
func (*Calc) Descriptor() ([]byte, []int) {
	return fileDescriptor_53e5b7e388c7a849, []int{0}
}

func (m *Calc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Calc.Unmarshal(m, b)
}
func (m *Calc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Calc.Marshal(b, m, deterministic)
}
func (m *Calc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Calc.Merge(m, src)
}
func (m *Calc) XXX_Size() int {
	return xxx_messageInfo_Calc.Size(m)
}
func (m *Calc) XXX_DiscardUnknown() {
	xxx_messageInfo_Calc.DiscardUnknown(m)
}

var xxx_messageInfo_Calc proto.InternalMessageInfo

func (m *Calc) GetFirstValue() int64 {
	if m != nil {
		return m.FirstValue
	}
	return 0
}

func (m *Calc) GetSecondValue() int64 {
	if m != nil {
		return m.SecondValue
	}
	return 0
}

type CalcRequest struct {
	Calc                 *Calc    `protobuf:"bytes,1,opt,name=calc,proto3" json:"calc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalcRequest) Reset()         { *m = CalcRequest{} }
func (m *CalcRequest) String() string { return proto.CompactTextString(m) }
func (*CalcRequest) ProtoMessage()    {}
func (*CalcRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_53e5b7e388c7a849, []int{1}
}

func (m *CalcRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalcRequest.Unmarshal(m, b)
}
func (m *CalcRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalcRequest.Marshal(b, m, deterministic)
}
func (m *CalcRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcRequest.Merge(m, src)
}
func (m *CalcRequest) XXX_Size() int {
	return xxx_messageInfo_CalcRequest.Size(m)
}
func (m *CalcRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CalcRequest proto.InternalMessageInfo

func (m *CalcRequest) GetCalc() *Calc {
	if m != nil {
		return m.Calc
	}
	return nil
}

type CalcResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CalcResponse) Reset()         { *m = CalcResponse{} }
func (m *CalcResponse) String() string { return proto.CompactTextString(m) }
func (*CalcResponse) ProtoMessage()    {}
func (*CalcResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_53e5b7e388c7a849, []int{2}
}

func (m *CalcResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CalcResponse.Unmarshal(m, b)
}
func (m *CalcResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CalcResponse.Marshal(b, m, deterministic)
}
func (m *CalcResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CalcResponse.Merge(m, src)
}
func (m *CalcResponse) XXX_Size() int {
	return xxx_messageInfo_CalcResponse.Size(m)
}
func (m *CalcResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CalcResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CalcResponse proto.InternalMessageInfo

func (m *CalcResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Calc)(nil), "example.calc.Calc")
	proto.RegisterType((*CalcRequest)(nil), "example.calc.CalcRequest")
	proto.RegisterType((*CalcResponse)(nil), "example.calc.CalcResponse")
}

func init() { proto.RegisterFile("calc/calcpb/calc.proto", fileDescriptor_53e5b7e388c7a849) }

var fileDescriptor_53e5b7e388c7a849 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x4e, 0xcc, 0x49,
	0xd6, 0x07, 0x11, 0x05, 0x49, 0x60, 0x4a, 0xaf, 0xa0, 0x28, 0xbf, 0x24, 0x5f, 0x88, 0x27, 0xb5,
	0x22, 0x31, 0xb7, 0x20, 0x27, 0x55, 0x0f, 0x24, 0xa6, 0xe4, 0xc5, 0xc5, 0xe2, 0x9c, 0x98, 0x93,
	0x2c, 0x24, 0xcf, 0xc5, 0x9d, 0x96, 0x59, 0x54, 0x5c, 0x12, 0x5f, 0x96, 0x98, 0x53, 0x9a, 0x2a,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1c, 0xc4, 0x05, 0x16, 0x0a, 0x03, 0x89, 0x08, 0x29, 0x72, 0xf1,
	0x14, 0xa7, 0x26, 0xe7, 0xe7, 0xa5, 0x40, 0x55, 0x30, 0x81, 0x55, 0x70, 0x43, 0xc4, 0xc0, 0x4a,
	0x94, 0x4c, 0xb9, 0xb8, 0x41, 0x66, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xa9, 0x71,
	0xb1, 0x80, 0xac, 0x00, 0x9b, 0xc5, 0x6d, 0x24, 0xa4, 0x87, 0x6c, 0xaf, 0x1e, 0x58, 0x21, 0x58,
	0x5e, 0x49, 0x8d, 0x8b, 0x07, 0xa2, 0xad, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x8c, 0x8b,
	0xad, 0x28, 0xb5, 0xb8, 0x34, 0xa7, 0x04, 0xac, 0x93, 0x33, 0x08, 0xca, 0x33, 0xf2, 0x83, 0x18,
	0x1f, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0x64, 0x0f, 0x75, 0xb9, 0x24, 0x16, 0x83, 0x21,
	0x2e, 0x90, 0x92, 0xc2, 0x26, 0x05, 0xb1, 0x45, 0x89, 0xc1, 0x89, 0x23, 0x8a, 0x0d, 0x12, 0x3a,
	0x49, 0x6c, 0xe0, 0x90, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x88, 0x19, 0x0a, 0xd8, 0x33,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CalcServiceClient is the client API for CalcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalcServiceClient interface {
	// unary
	Calc(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error)
}

type calcServiceClient struct {
	cc *grpc.ClientConn
}

func NewCalcServiceClient(cc *grpc.ClientConn) CalcServiceClient {
	return &calcServiceClient{cc}
}

func (c *calcServiceClient) Calc(ctx context.Context, in *CalcRequest, opts ...grpc.CallOption) (*CalcResponse, error) {
	out := new(CalcResponse)
	err := c.cc.Invoke(ctx, "/example.calc.CalcService/Calc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalcServiceServer is the server API for CalcService service.
type CalcServiceServer interface {
	// unary
	Calc(context.Context, *CalcRequest) (*CalcResponse, error)
}

// UnimplementedCalcServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalcServiceServer struct {
}

func (*UnimplementedCalcServiceServer) Calc(ctx context.Context, req *CalcRequest) (*CalcResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calc not implemented")
}

func RegisterCalcServiceServer(s *grpc.Server, srv CalcServiceServer) {
	s.RegisterService(&_CalcService_serviceDesc, srv)
}

func _CalcService_Calc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Calc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.calc.CalcService/Calc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Calc(ctx, req.(*CalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CalcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.calc.CalcService",
	HandlerType: (*CalcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calc",
			Handler:    _CalcService_Calc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "calc/calcpb/calc.proto",
}
