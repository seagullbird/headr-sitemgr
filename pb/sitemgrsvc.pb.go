// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sitemgrsvc.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	sitemgrsvc.proto

It has these top-level messages:
	CreateNewSiteRequest
	CreateNewSiteReply
	ProxyDeleteSiteRequest
	ProxyDeleteSiteReply
	CheckSitenameExistsRequest
	CheckSitenameExistsReply
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type CreateNewSiteRequest struct {
	UserId   uint64 `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Sitename string `protobuf:"bytes,2,opt,name=sitename" json:"sitename,omitempty"`
}

func (m *CreateNewSiteRequest) Reset()                    { *m = CreateNewSiteRequest{} }
func (m *CreateNewSiteRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateNewSiteRequest) ProtoMessage()               {}
func (*CreateNewSiteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateNewSiteRequest) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CreateNewSiteRequest) GetSitename() string {
	if m != nil {
		return m.Sitename
	}
	return ""
}

type CreateNewSiteReply struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *CreateNewSiteReply) Reset()                    { *m = CreateNewSiteReply{} }
func (m *CreateNewSiteReply) String() string            { return proto.CompactTextString(m) }
func (*CreateNewSiteReply) ProtoMessage()               {}
func (*CreateNewSiteReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateNewSiteReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type ProxyDeleteSiteRequest struct {
	SiteId uint64 `protobuf:"varint,1,opt,name=site_id,json=siteId" json:"site_id,omitempty"`
}

func (m *ProxyDeleteSiteRequest) Reset()                    { *m = ProxyDeleteSiteRequest{} }
func (m *ProxyDeleteSiteRequest) String() string            { return proto.CompactTextString(m) }
func (*ProxyDeleteSiteRequest) ProtoMessage()               {}
func (*ProxyDeleteSiteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ProxyDeleteSiteRequest) GetSiteId() uint64 {
	if m != nil {
		return m.SiteId
	}
	return 0
}

type ProxyDeleteSiteReply struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *ProxyDeleteSiteReply) Reset()                    { *m = ProxyDeleteSiteReply{} }
func (m *ProxyDeleteSiteReply) String() string            { return proto.CompactTextString(m) }
func (*ProxyDeleteSiteReply) ProtoMessage()               {}
func (*ProxyDeleteSiteReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ProxyDeleteSiteReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type CheckSitenameExistsRequest struct {
	Sitename string `protobuf:"bytes,1,opt,name=sitename" json:"sitename,omitempty"`
}

func (m *CheckSitenameExistsRequest) Reset()                    { *m = CheckSitenameExistsRequest{} }
func (m *CheckSitenameExistsRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckSitenameExistsRequest) ProtoMessage()               {}
func (*CheckSitenameExistsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CheckSitenameExistsRequest) GetSitename() string {
	if m != nil {
		return m.Sitename
	}
	return ""
}

type CheckSitenameExistsReply struct {
	Exists bool   `protobuf:"varint,1,opt,name=exists" json:"exists,omitempty"`
	Err    string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *CheckSitenameExistsReply) Reset()                    { *m = CheckSitenameExistsReply{} }
func (m *CheckSitenameExistsReply) String() string            { return proto.CompactTextString(m) }
func (*CheckSitenameExistsReply) ProtoMessage()               {}
func (*CheckSitenameExistsReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CheckSitenameExistsReply) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

func (m *CheckSitenameExistsReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateNewSiteRequest)(nil), "pb.CreateNewSiteRequest")
	proto.RegisterType((*CreateNewSiteReply)(nil), "pb.CreateNewSiteReply")
	proto.RegisterType((*ProxyDeleteSiteRequest)(nil), "pb.ProxyDeleteSiteRequest")
	proto.RegisterType((*ProxyDeleteSiteReply)(nil), "pb.ProxyDeleteSiteReply")
	proto.RegisterType((*CheckSitenameExistsRequest)(nil), "pb.CheckSitenameExistsRequest")
	proto.RegisterType((*CheckSitenameExistsReply)(nil), "pb.CheckSitenameExistsReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Sitemgr service

type SitemgrClient interface {
	// Create a new site.
	NewSite(ctx context.Context, in *CreateNewSiteRequest, opts ...grpc.CallOption) (*CreateNewSiteReply, error)
	DeleteSite(ctx context.Context, in *ProxyDeleteSiteRequest, opts ...grpc.CallOption) (*ProxyDeleteSiteReply, error)
	CheckSitenameExists(ctx context.Context, in *CheckSitenameExistsRequest, opts ...grpc.CallOption) (*CheckSitenameExistsReply, error)
}

type sitemgrClient struct {
	cc *grpc.ClientConn
}

func NewSitemgrClient(cc *grpc.ClientConn) SitemgrClient {
	return &sitemgrClient{cc}
}

func (c *sitemgrClient) NewSite(ctx context.Context, in *CreateNewSiteRequest, opts ...grpc.CallOption) (*CreateNewSiteReply, error) {
	out := new(CreateNewSiteReply)
	err := grpc.Invoke(ctx, "/pb.Sitemgr/NewSite", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sitemgrClient) DeleteSite(ctx context.Context, in *ProxyDeleteSiteRequest, opts ...grpc.CallOption) (*ProxyDeleteSiteReply, error) {
	out := new(ProxyDeleteSiteReply)
	err := grpc.Invoke(ctx, "/pb.Sitemgr/DeleteSite", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sitemgrClient) CheckSitenameExists(ctx context.Context, in *CheckSitenameExistsRequest, opts ...grpc.CallOption) (*CheckSitenameExistsReply, error) {
	out := new(CheckSitenameExistsReply)
	err := grpc.Invoke(ctx, "/pb.Sitemgr/CheckSitenameExists", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Sitemgr service

type SitemgrServer interface {
	// Create a new site.
	NewSite(context.Context, *CreateNewSiteRequest) (*CreateNewSiteReply, error)
	DeleteSite(context.Context, *ProxyDeleteSiteRequest) (*ProxyDeleteSiteReply, error)
	CheckSitenameExists(context.Context, *CheckSitenameExistsRequest) (*CheckSitenameExistsReply, error)
}

func RegisterSitemgrServer(s *grpc.Server, srv SitemgrServer) {
	s.RegisterService(&_Sitemgr_serviceDesc, srv)
}

func _Sitemgr_NewSite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNewSiteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SitemgrServer).NewSite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sitemgr/NewSite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SitemgrServer).NewSite(ctx, req.(*CreateNewSiteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sitemgr_DeleteSite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProxyDeleteSiteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SitemgrServer).DeleteSite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sitemgr/DeleteSite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SitemgrServer).DeleteSite(ctx, req.(*ProxyDeleteSiteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sitemgr_CheckSitenameExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckSitenameExistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SitemgrServer).CheckSitenameExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Sitemgr/CheckSitenameExists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SitemgrServer).CheckSitenameExists(ctx, req.(*CheckSitenameExistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sitemgr_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Sitemgr",
	HandlerType: (*SitemgrServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewSite",
			Handler:    _Sitemgr_NewSite_Handler,
		},
		{
			MethodName: "DeleteSite",
			Handler:    _Sitemgr_DeleteSite_Handler,
		},
		{
			MethodName: "CheckSitenameExists",
			Handler:    _Sitemgr_CheckSitenameExists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sitemgrsvc.proto",
}

func init() { proto.RegisterFile("sitemgrsvc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x4d, 0x94, 0xa4, 0x9d, 0x53, 0x19, 0x4b, 0x0c, 0x8b, 0x48, 0xc9, 0x41, 0x72, 0x0a,
	0xa8, 0x17, 0x2f, 0x9e, 0x1a, 0x0f, 0x45, 0x10, 0x49, 0xf0, 0x2c, 0x4d, 0x33, 0x68, 0x30, 0x35,
	0xeb, 0xee, 0x56, 0x9b, 0x1f, 0xec, 0xff, 0x90, 0xdd, 0xa6, 0xb5, 0xc4, 0xed, 0x2d, 0x8f, 0xc9,
	0xbc, 0xef, 0xcd, 0x63, 0x61, 0x24, 0x2b, 0x45, 0xcb, 0x57, 0x21, 0xbf, 0x16, 0x09, 0x17, 0x8d,
	0x6a, 0xd0, 0xe5, 0x45, 0xf4, 0x00, 0xe3, 0xa9, 0xa0, 0xb9, 0xa2, 0x47, 0xfa, 0xce, 0x2b, 0x45,
	0x19, 0x7d, 0xae, 0x48, 0x2a, 0x3c, 0x03, 0x7f, 0x25, 0x49, 0xbc, 0x54, 0x65, 0xe8, 0x4c, 0x9c,
	0xf8, 0x24, 0xf3, 0xb4, 0x9c, 0x95, 0xc8, 0x60, 0xa0, 0x8d, 0x3e, 0xe6, 0x4b, 0x0a, 0xdd, 0x89,
	0x13, 0x0f, 0xb3, 0x9d, 0x8e, 0x2e, 0x01, 0x7b, 0x66, 0xbc, 0x6e, 0x71, 0x04, 0xc7, 0x24, 0x84,
	0xb1, 0x19, 0x66, 0xfa, 0x33, 0xba, 0x82, 0xe0, 0x49, 0x34, 0xeb, 0x36, 0xa5, 0x9a, 0x14, 0xf5,
	0xb0, 0xda, 0x6d, 0x0f, 0xab, 0xe5, 0xac, 0x8c, 0x62, 0x18, 0xff, 0x5b, 0xb1, 0x9b, 0xdf, 0x02,
	0x9b, 0xbe, 0xd1, 0xe2, 0x3d, 0xef, 0x52, 0xdd, 0xaf, 0x2b, 0xa9, 0xe4, 0x16, 0xb0, 0x1f, 0xdf,
	0xe9, 0xc5, 0x4f, 0x21, 0xb4, 0x6e, 0x6a, 0x4e, 0x00, 0x1e, 0x19, 0x69, 0xb6, 0x06, 0x59, 0xa7,
	0xb6, 0x7c, 0x77, 0xc7, 0xbf, 0xfe, 0x71, 0xc0, 0xcf, 0x37, 0x55, 0xe3, 0x1d, 0xf8, 0x5d, 0x15,
	0x18, 0x26, 0xbc, 0x48, 0x6c, 0x55, 0xb3, 0xc0, 0x32, 0xe1, 0x75, 0x1b, 0x1d, 0x61, 0x0a, 0xf0,
	0x77, 0x2f, 0x32, 0xfd, 0x9f, 0xbd, 0x37, 0x16, 0x5a, 0x67, 0x1b, 0x97, 0x67, 0x38, 0xb5, 0x9c,
	0x85, 0x17, 0x06, 0x7b, 0xb0, 0x29, 0x76, 0x7e, 0x70, 0x6e, 0x6c, 0x0b, 0xcf, 0x3c, 0xa2, 0x9b,
	0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xaa, 0x2f, 0xc8, 0x58, 0x02, 0x00, 0x00,
}
