// Code generated by protoc-gen-go. DO NOT EDIT.
// source: repoctlsvc.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	repoctlsvc.proto

It has these top-level messages:
	NewSiteRequest
	NewSiteReply
	DeleteSiteRequest
	DeleteSiteReply
	WritePostRequest
	WritePostReply
	RemovePostRequest
	RemovePostReply
	ReadPostRequest
	ReadPostReply
	WriteConfigRequest
	WriteConfigReply
	ReadConfigRequest
	ReadConfigReply
	UpdateAboutRequest
	UpdateAboutReply
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

type NewSiteRequest struct {
	SiteId uint64 `protobuf:"varint,1,opt,name=site_id,json=siteId" json:"site_id,omitempty"`
	Theme  string `protobuf:"bytes,2,opt,name=theme" json:"theme,omitempty"`
}

func (m *NewSiteRequest) Reset()                    { *m = NewSiteRequest{} }
func (m *NewSiteRequest) String() string            { return proto.CompactTextString(m) }
func (*NewSiteRequest) ProtoMessage()               {}
func (*NewSiteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *NewSiteRequest) GetSiteId() uint64 {
	if m != nil {
		return m.SiteId
	}
	return 0
}

func (m *NewSiteRequest) GetTheme() string {
	if m != nil {
		return m.Theme
	}
	return ""
}

type NewSiteReply struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *NewSiteReply) Reset()                    { *m = NewSiteReply{} }
func (m *NewSiteReply) String() string            { return proto.CompactTextString(m) }
func (*NewSiteReply) ProtoMessage()               {}
func (*NewSiteReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *NewSiteReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type DeleteSiteRequest struct {
	SiteId uint64 `protobuf:"varint,1,opt,name=site_id,json=siteId" json:"site_id,omitempty"`
}

func (m *DeleteSiteRequest) Reset()                    { *m = DeleteSiteRequest{} }
func (m *DeleteSiteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteSiteRequest) ProtoMessage()               {}
func (*DeleteSiteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DeleteSiteRequest) GetSiteId() uint64 {
	if m != nil {
		return m.SiteId
	}
	return 0
}

type DeleteSiteReply struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *DeleteSiteReply) Reset()                    { *m = DeleteSiteReply{} }
func (m *DeleteSiteReply) String() string            { return proto.CompactTextString(m) }
func (*DeleteSiteReply) ProtoMessage()               {}
func (*DeleteSiteReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DeleteSiteReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type WritePostRequest struct {
	SiteId   uint64 `protobuf:"varint,1,opt,name=site_id,json=siteId" json:"site_id,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=filename" json:"filename,omitempty"`
	Content  string `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
}

func (m *WritePostRequest) Reset()                    { *m = WritePostRequest{} }
func (m *WritePostRequest) String() string            { return proto.CompactTextString(m) }
func (*WritePostRequest) ProtoMessage()               {}
func (*WritePostRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *WritePostRequest) GetSiteId() uint64 {
	if m != nil {
		return m.SiteId
	}
	return 0
}

func (m *WritePostRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *WritePostRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type WritePostReply struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *WritePostReply) Reset()                    { *m = WritePostReply{} }
func (m *WritePostReply) String() string            { return proto.CompactTextString(m) }
func (*WritePostReply) ProtoMessage()               {}
func (*WritePostReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *WritePostReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type RemovePostRequest struct {
	SiteId   uint64 `protobuf:"varint,1,opt,name=site_id,json=siteId" json:"site_id,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=filename" json:"filename,omitempty"`
}

func (m *RemovePostRequest) Reset()                    { *m = RemovePostRequest{} }
func (m *RemovePostRequest) String() string            { return proto.CompactTextString(m) }
func (*RemovePostRequest) ProtoMessage()               {}
func (*RemovePostRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RemovePostRequest) GetSiteId() uint64 {
	if m != nil {
		return m.SiteId
	}
	return 0
}

func (m *RemovePostRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

type RemovePostReply struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *RemovePostReply) Reset()                    { *m = RemovePostReply{} }
func (m *RemovePostReply) String() string            { return proto.CompactTextString(m) }
func (*RemovePostReply) ProtoMessage()               {}
func (*RemovePostReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RemovePostReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type ReadPostRequest struct {
	SiteId   uint64 `protobuf:"varint,1,opt,name=site_id,json=siteId" json:"site_id,omitempty"`
	Filename string `protobuf:"bytes,2,opt,name=filename" json:"filename,omitempty"`
}

func (m *ReadPostRequest) Reset()                    { *m = ReadPostRequest{} }
func (m *ReadPostRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadPostRequest) ProtoMessage()               {}
func (*ReadPostRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ReadPostRequest) GetSiteId() uint64 {
	if m != nil {
		return m.SiteId
	}
	return 0
}

func (m *ReadPostRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

type ReadPostReply struct {
	Content string `protobuf:"bytes,1,opt,name=content" json:"content,omitempty"`
	Err     string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *ReadPostReply) Reset()                    { *m = ReadPostReply{} }
func (m *ReadPostReply) String() string            { return proto.CompactTextString(m) }
func (*ReadPostReply) ProtoMessage()               {}
func (*ReadPostReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ReadPostReply) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ReadPostReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type WriteConfigRequest struct {
	SiteId uint64 `protobuf:"varint,1,opt,name=site_id,json=siteId" json:"site_id,omitempty"`
	Config string `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
}

func (m *WriteConfigRequest) Reset()                    { *m = WriteConfigRequest{} }
func (m *WriteConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*WriteConfigRequest) ProtoMessage()               {}
func (*WriteConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *WriteConfigRequest) GetSiteId() uint64 {
	if m != nil {
		return m.SiteId
	}
	return 0
}

func (m *WriteConfigRequest) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

type WriteConfigReply struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *WriteConfigReply) Reset()                    { *m = WriteConfigReply{} }
func (m *WriteConfigReply) String() string            { return proto.CompactTextString(m) }
func (*WriteConfigReply) ProtoMessage()               {}
func (*WriteConfigReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *WriteConfigReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type ReadConfigRequest struct {
	SiteId uint64 `protobuf:"varint,1,opt,name=site_id,json=siteId" json:"site_id,omitempty"`
}

func (m *ReadConfigRequest) Reset()                    { *m = ReadConfigRequest{} }
func (m *ReadConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadConfigRequest) ProtoMessage()               {}
func (*ReadConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *ReadConfigRequest) GetSiteId() uint64 {
	if m != nil {
		return m.SiteId
	}
	return 0
}

type ReadConfigReply struct {
	Config string `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	Err    string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *ReadConfigReply) Reset()                    { *m = ReadConfigReply{} }
func (m *ReadConfigReply) String() string            { return proto.CompactTextString(m) }
func (*ReadConfigReply) ProtoMessage()               {}
func (*ReadConfigReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *ReadConfigReply) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

func (m *ReadConfigReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

type UpdateAboutRequest struct {
	SiteId  uint64 `protobuf:"varint,1,opt,name=site_id,json=siteId" json:"site_id,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
}

func (m *UpdateAboutRequest) Reset()                    { *m = UpdateAboutRequest{} }
func (m *UpdateAboutRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateAboutRequest) ProtoMessage()               {}
func (*UpdateAboutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *UpdateAboutRequest) GetSiteId() uint64 {
	if m != nil {
		return m.SiteId
	}
	return 0
}

func (m *UpdateAboutRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type UpdateAboutReply struct {
	Err string `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
}

func (m *UpdateAboutReply) Reset()                    { *m = UpdateAboutReply{} }
func (m *UpdateAboutReply) String() string            { return proto.CompactTextString(m) }
func (*UpdateAboutReply) ProtoMessage()               {}
func (*UpdateAboutReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *UpdateAboutReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*NewSiteRequest)(nil), "pb.NewSiteRequest")
	proto.RegisterType((*NewSiteReply)(nil), "pb.NewSiteReply")
	proto.RegisterType((*DeleteSiteRequest)(nil), "pb.DeleteSiteRequest")
	proto.RegisterType((*DeleteSiteReply)(nil), "pb.DeleteSiteReply")
	proto.RegisterType((*WritePostRequest)(nil), "pb.WritePostRequest")
	proto.RegisterType((*WritePostReply)(nil), "pb.WritePostReply")
	proto.RegisterType((*RemovePostRequest)(nil), "pb.RemovePostRequest")
	proto.RegisterType((*RemovePostReply)(nil), "pb.RemovePostReply")
	proto.RegisterType((*ReadPostRequest)(nil), "pb.ReadPostRequest")
	proto.RegisterType((*ReadPostReply)(nil), "pb.ReadPostReply")
	proto.RegisterType((*WriteConfigRequest)(nil), "pb.WriteConfigRequest")
	proto.RegisterType((*WriteConfigReply)(nil), "pb.WriteConfigReply")
	proto.RegisterType((*ReadConfigRequest)(nil), "pb.ReadConfigRequest")
	proto.RegisterType((*ReadConfigReply)(nil), "pb.ReadConfigReply")
	proto.RegisterType((*UpdateAboutRequest)(nil), "pb.UpdateAboutRequest")
	proto.RegisterType((*UpdateAboutReply)(nil), "pb.UpdateAboutReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Repoctl service

type RepoctlClient interface {
	NewSite(ctx context.Context, in *NewSiteRequest, opts ...grpc.CallOption) (*NewSiteReply, error)
	DeleteSite(ctx context.Context, in *DeleteSiteRequest, opts ...grpc.CallOption) (*DeleteSiteReply, error)
	WritePost(ctx context.Context, in *WritePostRequest, opts ...grpc.CallOption) (*WritePostReply, error)
	RemovePost(ctx context.Context, in *RemovePostRequest, opts ...grpc.CallOption) (*RemovePostReply, error)
	ReadPost(ctx context.Context, in *ReadPostRequest, opts ...grpc.CallOption) (*ReadPostReply, error)
	WriteConfig(ctx context.Context, in *WriteConfigRequest, opts ...grpc.CallOption) (*WriteConfigReply, error)
	ReadConfig(ctx context.Context, in *ReadConfigRequest, opts ...grpc.CallOption) (*ReadConfigReply, error)
	UpdateAbout(ctx context.Context, in *UpdateAboutRequest, opts ...grpc.CallOption) (*UpdateAboutReply, error)
}

type repoctlClient struct {
	cc *grpc.ClientConn
}

func NewRepoctlClient(cc *grpc.ClientConn) RepoctlClient {
	return &repoctlClient{cc}
}

func (c *repoctlClient) NewSite(ctx context.Context, in *NewSiteRequest, opts ...grpc.CallOption) (*NewSiteReply, error) {
	out := new(NewSiteReply)
	err := grpc.Invoke(ctx, "/pb.Repoctl/NewSite", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoctlClient) DeleteSite(ctx context.Context, in *DeleteSiteRequest, opts ...grpc.CallOption) (*DeleteSiteReply, error) {
	out := new(DeleteSiteReply)
	err := grpc.Invoke(ctx, "/pb.Repoctl/DeleteSite", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoctlClient) WritePost(ctx context.Context, in *WritePostRequest, opts ...grpc.CallOption) (*WritePostReply, error) {
	out := new(WritePostReply)
	err := grpc.Invoke(ctx, "/pb.Repoctl/WritePost", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoctlClient) RemovePost(ctx context.Context, in *RemovePostRequest, opts ...grpc.CallOption) (*RemovePostReply, error) {
	out := new(RemovePostReply)
	err := grpc.Invoke(ctx, "/pb.Repoctl/RemovePost", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoctlClient) ReadPost(ctx context.Context, in *ReadPostRequest, opts ...grpc.CallOption) (*ReadPostReply, error) {
	out := new(ReadPostReply)
	err := grpc.Invoke(ctx, "/pb.Repoctl/ReadPost", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoctlClient) WriteConfig(ctx context.Context, in *WriteConfigRequest, opts ...grpc.CallOption) (*WriteConfigReply, error) {
	out := new(WriteConfigReply)
	err := grpc.Invoke(ctx, "/pb.Repoctl/WriteConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoctlClient) ReadConfig(ctx context.Context, in *ReadConfigRequest, opts ...grpc.CallOption) (*ReadConfigReply, error) {
	out := new(ReadConfigReply)
	err := grpc.Invoke(ctx, "/pb.Repoctl/ReadConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *repoctlClient) UpdateAbout(ctx context.Context, in *UpdateAboutRequest, opts ...grpc.CallOption) (*UpdateAboutReply, error) {
	out := new(UpdateAboutReply)
	err := grpc.Invoke(ctx, "/pb.Repoctl/UpdateAbout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Repoctl service

type RepoctlServer interface {
	NewSite(context.Context, *NewSiteRequest) (*NewSiteReply, error)
	DeleteSite(context.Context, *DeleteSiteRequest) (*DeleteSiteReply, error)
	WritePost(context.Context, *WritePostRequest) (*WritePostReply, error)
	RemovePost(context.Context, *RemovePostRequest) (*RemovePostReply, error)
	ReadPost(context.Context, *ReadPostRequest) (*ReadPostReply, error)
	WriteConfig(context.Context, *WriteConfigRequest) (*WriteConfigReply, error)
	ReadConfig(context.Context, *ReadConfigRequest) (*ReadConfigReply, error)
	UpdateAbout(context.Context, *UpdateAboutRequest) (*UpdateAboutReply, error)
}

func RegisterRepoctlServer(s *grpc.Server, srv RepoctlServer) {
	s.RegisterService(&_Repoctl_serviceDesc, srv)
}

func _Repoctl_NewSite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewSiteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoctlServer).NewSite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Repoctl/NewSite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoctlServer).NewSite(ctx, req.(*NewSiteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repoctl_DeleteSite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSiteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoctlServer).DeleteSite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Repoctl/DeleteSite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoctlServer).DeleteSite(ctx, req.(*DeleteSiteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repoctl_WritePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WritePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoctlServer).WritePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Repoctl/WritePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoctlServer).WritePost(ctx, req.(*WritePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repoctl_RemovePost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoctlServer).RemovePost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Repoctl/RemovePost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoctlServer).RemovePost(ctx, req.(*RemovePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repoctl_ReadPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoctlServer).ReadPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Repoctl/ReadPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoctlServer).ReadPost(ctx, req.(*ReadPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repoctl_WriteConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoctlServer).WriteConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Repoctl/WriteConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoctlServer).WriteConfig(ctx, req.(*WriteConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repoctl_ReadConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoctlServer).ReadConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Repoctl/ReadConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoctlServer).ReadConfig(ctx, req.(*ReadConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Repoctl_UpdateAbout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAboutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepoctlServer).UpdateAbout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Repoctl/UpdateAbout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepoctlServer).UpdateAbout(ctx, req.(*UpdateAboutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Repoctl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Repoctl",
	HandlerType: (*RepoctlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewSite",
			Handler:    _Repoctl_NewSite_Handler,
		},
		{
			MethodName: "DeleteSite",
			Handler:    _Repoctl_DeleteSite_Handler,
		},
		{
			MethodName: "WritePost",
			Handler:    _Repoctl_WritePost_Handler,
		},
		{
			MethodName: "RemovePost",
			Handler:    _Repoctl_RemovePost_Handler,
		},
		{
			MethodName: "ReadPost",
			Handler:    _Repoctl_ReadPost_Handler,
		},
		{
			MethodName: "WriteConfig",
			Handler:    _Repoctl_WriteConfig_Handler,
		},
		{
			MethodName: "ReadConfig",
			Handler:    _Repoctl_ReadConfig_Handler,
		},
		{
			MethodName: "UpdateAbout",
			Handler:    _Repoctl_UpdateAbout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "repoctlsvc.proto",
}

func init() { proto.RegisterFile("repoctlsvc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x5b, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x9b, 0x56, 0x9b, 0xed, 0x51, 0xd7, 0xf4, 0x6c, 0x5d, 0x43, 0x9e, 0xca, 0xb8, 0x0f,
	0xfb, 0x20, 0x05, 0x2f, 0x20, 0xb8, 0x88, 0x88, 0xf7, 0x17, 0x91, 0x88, 0xf8, 0x28, 0xb9, 0x9c,
	0xd5, 0x40, 0x36, 0x13, 0x93, 0x69, 0xa5, 0x1f, 0xd6, 0xef, 0x22, 0x93, 0xc9, 0x65, 0x72, 0x29,
	0x2d, 0xf4, 0xad, 0xe7, 0xf4, 0xcc, 0xf9, 0xfd, 0xe7, 0xcc, 0xff, 0x04, 0xac, 0x8c, 0x52, 0x1e,
	0x88, 0x38, 0xdf, 0x04, 0xab, 0x34, 0xe3, 0x82, 0xe3, 0x38, 0xf5, 0xd9, 0x6b, 0x38, 0xfd, 0x42,
	0x7f, 0xbf, 0x45, 0x82, 0x5c, 0xfa, 0xb3, 0xa6, 0x5c, 0xe0, 0x43, 0x30, 0xf3, 0x48, 0xd0, 0xcf,
	0x28, 0xb4, 0x8d, 0xa5, 0x71, 0x79, 0xcb, 0x9d, 0xca, 0xf0, 0x73, 0x88, 0x0b, 0xb8, 0x2d, 0x7e,
	0xd3, 0x0d, 0xd9, 0xe3, 0xa5, 0x71, 0x39, 0x73, 0x55, 0xc0, 0x96, 0x70, 0xb7, 0x6e, 0x90, 0xc6,
	0x5b, 0xb4, 0x60, 0x42, 0x59, 0x56, 0x1c, 0x9d, 0xb9, 0xf2, 0x27, 0x7b, 0x0c, 0xf3, 0x77, 0x14,
	0x93, 0xa0, 0x43, 0x28, 0xec, 0x11, 0xdc, 0xd7, 0xab, 0x87, 0x5b, 0x7a, 0x60, 0xfd, 0xc8, 0x22,
	0x41, 0x5f, 0x79, 0x2e, 0xf6, 0xea, 0x76, 0xe0, 0xe4, 0x3a, 0x8a, 0x29, 0xf1, 0x6a, 0xe9, 0x75,
	0x8c, 0x36, 0x98, 0x01, 0x4f, 0x04, 0x25, 0xc2, 0x9e, 0x14, 0x7f, 0x55, 0x21, 0x63, 0x70, 0xaa,
	0x21, 0x86, 0x65, 0x7c, 0x82, 0xb9, 0x4b, 0x37, 0x7c, 0x73, 0xb4, 0x0e, 0x79, 0x6b, 0xbd, 0xd3,
	0x30, 0xee, 0x83, 0x2c, 0xf2, 0xc2, 0xa3, 0x61, 0x57, 0x70, 0xaf, 0xe9, 0x23, 0x51, 0xda, 0x14,
	0x8c, 0xd6, 0x14, 0x2a, 0x11, 0xe3, 0x46, 0xc4, 0x7b, 0xc0, 0x62, 0x2e, 0x6f, 0x79, 0x72, 0x1d,
	0xfd, 0xda, 0xab, 0xe3, 0x1c, 0xa6, 0x41, 0x51, 0x59, 0xf6, 0x28, 0x23, 0x76, 0x51, 0xbe, 0x60,
	0xd5, 0x66, 0xa7, 0x75, 0xa4, 0xd2, 0xc3, 0x58, 0xec, 0x4a, 0xcd, 0x47, 0x6f, 0xd9, 0xe0, 0x0d,
	0x1d, 0x3f, 0x70, 0xaf, 0x8f, 0x80, 0xdf, 0xd3, 0xd0, 0x13, 0xf4, 0xc6, 0xe7, 0xeb, 0xfd, 0xf3,
	0xdd, 0x6d, 0x9c, 0x0b, 0xb0, 0x5a, 0x8d, 0x06, 0x6f, 0xf6, 0xf4, 0xdf, 0x04, 0x4c, 0x57, 0x2d,
	0x24, 0x3e, 0x01, 0xb3, 0x5c, 0x21, 0xc4, 0x55, 0xea, 0xaf, 0xda, 0x0b, 0xe9, 0x58, 0xad, 0x5c,
	0x1a, 0x6f, 0xd9, 0x08, 0x5f, 0x02, 0x34, 0x5b, 0x82, 0x0f, 0x64, 0x45, 0x6f, 0xc7, 0x9c, 0xb3,
	0x6e, 0x5a, 0x9d, 0x7d, 0x01, 0xb3, 0xda, 0xd9, 0xb8, 0x90, 0x35, 0xdd, 0x5d, 0x72, 0xb0, 0x93,
	0xad, 0xa1, 0x8d, 0x49, 0x15, 0xb4, 0x67, 0x7f, 0x05, 0xed, 0x78, 0x99, 0x8d, 0xf0, 0x39, 0x9c,
	0x54, 0x9e, 0xc3, 0xb2, 0xa4, 0xe5, 0x64, 0x67, 0xde, 0x4e, 0xaa, 0x53, 0xaf, 0xe0, 0x8e, 0xe6,
	0x12, 0x3c, 0xaf, 0x65, 0xb5, 0x1c, 0xe1, 0x2c, 0x7a, 0x79, 0x4d, 0x70, 0x65, 0x88, 0x4a, 0x70,
	0xc7, 0x4e, 0xce, 0x59, 0x37, 0x5d, 0xa3, 0xb5, 0x67, 0x54, 0xe8, 0xbe, 0x41, 0x14, 0xba, 0xfb,
	0xde, 0x6c, 0xe4, 0x4f, 0x8b, 0x4f, 0xec, 0xb3, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x03, 0x73,
	0x20, 0x5c, 0x76, 0x05, 0x00, 0x00,
}
