// Code generated by protoc-gen-go. DO NOT EDIT.
// source: recommendpb/recommend.proto

package recommendpb

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

type Recommender struct {
	Movie                string   `protobuf:"bytes,1,opt,name=movie,proto3" json:"movie,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Recommender) Reset()         { *m = Recommender{} }
func (m *Recommender) String() string { return proto.CompactTextString(m) }
func (*Recommender) ProtoMessage()    {}
func (*Recommender) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1532813a15d5618, []int{0}
}

func (m *Recommender) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Recommender.Unmarshal(m, b)
}
func (m *Recommender) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Recommender.Marshal(b, m, deterministic)
}
func (m *Recommender) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Recommender.Merge(m, src)
}
func (m *Recommender) XXX_Size() int {
	return xxx_messageInfo_Recommender.Size(m)
}
func (m *Recommender) XXX_DiscardUnknown() {
	xxx_messageInfo_Recommender.DiscardUnknown(m)
}

var xxx_messageInfo_Recommender proto.InternalMessageInfo

func (m *Recommender) GetMovie() string {
	if m != nil {
		return m.Movie
	}
	return ""
}

type RecommendRequest struct {
	Recommender          *Recommender `protobuf:"bytes,1,opt,name=recommender,proto3" json:"recommender,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RecommendRequest) Reset()         { *m = RecommendRequest{} }
func (m *RecommendRequest) String() string { return proto.CompactTextString(m) }
func (*RecommendRequest) ProtoMessage()    {}
func (*RecommendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1532813a15d5618, []int{1}
}

func (m *RecommendRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecommendRequest.Unmarshal(m, b)
}
func (m *RecommendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecommendRequest.Marshal(b, m, deterministic)
}
func (m *RecommendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecommendRequest.Merge(m, src)
}
func (m *RecommendRequest) XXX_Size() int {
	return xxx_messageInfo_RecommendRequest.Size(m)
}
func (m *RecommendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecommendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecommendRequest proto.InternalMessageInfo

func (m *RecommendRequest) GetRecommender() *Recommender {
	if m != nil {
		return m.Recommender
	}
	return nil
}

type RecommendResponse struct {
	MovieTitle           string   `protobuf:"bytes,1,opt,name=movieTitle,proto3" json:"movieTitle,omitempty"`
	MovieId              int64    `protobuf:"varint,2,opt,name=movieId,proto3" json:"movieId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecommendResponse) Reset()         { *m = RecommendResponse{} }
func (m *RecommendResponse) String() string { return proto.CompactTextString(m) }
func (*RecommendResponse) ProtoMessage()    {}
func (*RecommendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1532813a15d5618, []int{2}
}

func (m *RecommendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecommendResponse.Unmarshal(m, b)
}
func (m *RecommendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecommendResponse.Marshal(b, m, deterministic)
}
func (m *RecommendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecommendResponse.Merge(m, src)
}
func (m *RecommendResponse) XXX_Size() int {
	return xxx_messageInfo_RecommendResponse.Size(m)
}
func (m *RecommendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RecommendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RecommendResponse proto.InternalMessageInfo

func (m *RecommendResponse) GetMovieTitle() string {
	if m != nil {
		return m.MovieTitle
	}
	return ""
}

func (m *RecommendResponse) GetMovieId() int64 {
	if m != nil {
		return m.MovieId
	}
	return 0
}

func init() {
	proto.RegisterType((*Recommender)(nil), "recommend.Recommender")
	proto.RegisterType((*RecommendRequest)(nil), "recommend.RecommendRequest")
	proto.RegisterType((*RecommendResponse)(nil), "recommend.RecommendResponse")
}

func init() { proto.RegisterFile("recommendpb/recommend.proto", fileDescriptor_f1532813a15d5618) }

var fileDescriptor_f1532813a15d5618 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x4a, 0x4d, 0xce,
	0xcf, 0xcd, 0x4d, 0xcd, 0x4b, 0x29, 0x48, 0xd2, 0x87, 0xb3, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0x38, 0xe1, 0x02, 0x4a, 0xca, 0x5c, 0xdc, 0x41, 0x30, 0x4e, 0x6a, 0x91, 0x90, 0x08, 0x17,
	0x6b, 0x6e, 0x7e, 0x59, 0x66, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x84, 0xa3, 0xe4,
	0xc3, 0x25, 0x00, 0x57, 0x14, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x64, 0xc1, 0xc5, 0x5d,
	0x84, 0xd0, 0x08, 0x56, 0xcf, 0x6d, 0x24, 0xa6, 0x87, 0xb0, 0x0a, 0xc9, 0xd8, 0x20, 0x64, 0xa5,
	0x4a, 0xbe, 0x5c, 0x82, 0x48, 0xa6, 0x15, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0xc9, 0x71, 0x71,
	0x81, 0xed, 0x0a, 0xc9, 0x2c, 0xc9, 0x81, 0xd9, 0x8e, 0x24, 0x22, 0x24, 0xc1, 0xc5, 0x0e, 0xe6,
	0x79, 0xa6, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x30, 0x07, 0xc1, 0xb8, 0x46, 0x71, 0x48, 0x8e, 0x0b,
	0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0xf2, 0xe2, 0xe2, 0x84, 0x8b, 0x09, 0x49, 0x63, 0x73,
	0x14, 0xd4, 0x1b, 0x52, 0x32, 0xd8, 0x25, 0x21, 0xae, 0x52, 0x62, 0x30, 0x60, 0x74, 0xe2, 0x8d,
	0xe2, 0x46, 0x0a, 0xcb, 0x24, 0x36, 0x70, 0x10, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb5,
	0x39, 0x76, 0x14, 0x61, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RecommendServiceClient is the client API for RecommendService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RecommendServiceClient interface {
	Recommend(ctx context.Context, in *RecommendRequest, opts ...grpc.CallOption) (RecommendService_RecommendClient, error)
}

type recommendServiceClient struct {
	cc *grpc.ClientConn
}

func NewRecommendServiceClient(cc *grpc.ClientConn) RecommendServiceClient {
	return &recommendServiceClient{cc}
}

func (c *recommendServiceClient) Recommend(ctx context.Context, in *RecommendRequest, opts ...grpc.CallOption) (RecommendService_RecommendClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RecommendService_serviceDesc.Streams[0], "/recommend.RecommendService/Recommend", opts...)
	if err != nil {
		return nil, err
	}
	x := &recommendServiceRecommendClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RecommendService_RecommendClient interface {
	Recv() (*RecommendResponse, error)
	grpc.ClientStream
}

type recommendServiceRecommendClient struct {
	grpc.ClientStream
}

func (x *recommendServiceRecommendClient) Recv() (*RecommendResponse, error) {
	m := new(RecommendResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RecommendServiceServer is the server API for RecommendService service.
type RecommendServiceServer interface {
	Recommend(*RecommendRequest, RecommendService_RecommendServer) error
}

// UnimplementedRecommendServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRecommendServiceServer struct {
}

func (*UnimplementedRecommendServiceServer) Recommend(req *RecommendRequest, srv RecommendService_RecommendServer) error {
	return status.Errorf(codes.Unimplemented, "method Recommend not implemented")
}

func RegisterRecommendServiceServer(s *grpc.Server, srv RecommendServiceServer) {
	s.RegisterService(&_RecommendService_serviceDesc, srv)
}

func _RecommendService_Recommend_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RecommendRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RecommendServiceServer).Recommend(m, &recommendServiceRecommendServer{stream})
}

type RecommendService_RecommendServer interface {
	Send(*RecommendResponse) error
	grpc.ServerStream
}

type recommendServiceRecommendServer struct {
	grpc.ServerStream
}

func (x *recommendServiceRecommendServer) Send(m *RecommendResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _RecommendService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "recommend.RecommendService",
	HandlerType: (*RecommendServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Recommend",
			Handler:       _RecommendService_Recommend_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "recommendpb/recommend.proto",
}
