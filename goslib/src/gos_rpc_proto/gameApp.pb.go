// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gameApp.proto

package gos_rpc_proto

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

type DispatchGameRequest struct {
	AccountId string `protobuf:"bytes,1,opt,name=accountId" json:"accountId,omitempty"`
	ServerId  string `protobuf:"bytes,2,opt,name=serverId" json:"serverId,omitempty"`
	SceneId   string `protobuf:"bytes,3,opt,name=sceneId" json:"sceneId,omitempty"`
}

func (m *DispatchGameRequest) Reset()                    { *m = DispatchGameRequest{} }
func (m *DispatchGameRequest) String() string            { return proto.CompactTextString(m) }
func (*DispatchGameRequest) ProtoMessage()               {}
func (*DispatchGameRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *DispatchGameRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *DispatchGameRequest) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *DispatchGameRequest) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

type DispatchGameReply struct {
	GameAppId   string `protobuf:"bytes,1,opt,name=gameAppId" json:"gameAppId,omitempty"`
	GameAppHost string `protobuf:"bytes,2,opt,name=gameAppHost" json:"gameAppHost,omitempty"`
	GameAppPort string `protobuf:"bytes,3,opt,name=gameAppPort" json:"gameAppPort,omitempty"`
	SceneId     string `protobuf:"bytes,4,opt,name=sceneId" json:"sceneId,omitempty"`
}

func (m *DispatchGameReply) Reset()                    { *m = DispatchGameReply{} }
func (m *DispatchGameReply) String() string            { return proto.CompactTextString(m) }
func (*DispatchGameReply) ProtoMessage()               {}
func (*DispatchGameReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *DispatchGameReply) GetGameAppId() string {
	if m != nil {
		return m.GameAppId
	}
	return ""
}

func (m *DispatchGameReply) GetGameAppHost() string {
	if m != nil {
		return m.GameAppHost
	}
	return ""
}

func (m *DispatchGameReply) GetGameAppPort() string {
	if m != nil {
		return m.GameAppPort
	}
	return ""
}

func (m *DispatchGameReply) GetSceneId() string {
	if m != nil {
		return m.SceneId
	}
	return ""
}

type ReportGameRequest struct {
	Uuid string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Host string `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
	Port string `protobuf:"bytes,3,opt,name=port" json:"port,omitempty"`
	Ccu  int32  `protobuf:"varint,4,opt,name=ccu" json:"ccu,omitempty"`
}

func (m *ReportGameRequest) Reset()                    { *m = ReportGameRequest{} }
func (m *ReportGameRequest) String() string            { return proto.CompactTextString(m) }
func (*ReportGameRequest) ProtoMessage()               {}
func (*ReportGameRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *ReportGameRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ReportGameRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *ReportGameRequest) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func (m *ReportGameRequest) GetCcu() int32 {
	if m != nil {
		return m.Ccu
	}
	return 0
}

type ReportGameReply struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *ReportGameReply) Reset()                    { *m = ReportGameReply{} }
func (m *ReportGameReply) String() string            { return proto.CompactTextString(m) }
func (*ReportGameReply) ProtoMessage()               {}
func (*ReportGameReply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *ReportGameReply) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*DispatchGameRequest)(nil), "connectApp.DispatchGameRequest")
	proto.RegisterType((*DispatchGameReply)(nil), "connectApp.DispatchGameReply")
	proto.RegisterType((*ReportGameRequest)(nil), "connectApp.ReportGameRequest")
	proto.RegisterType((*ReportGameReply)(nil), "connectApp.ReportGameReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GameDispatcher service

type GameDispatcherClient interface {
	DispatchGame(ctx context.Context, in *DispatchGameRequest, opts ...grpc.CallOption) (*DispatchGameReply, error)
}

type gameDispatcherClient struct {
	cc *grpc.ClientConn
}

func NewGameDispatcherClient(cc *grpc.ClientConn) GameDispatcherClient {
	return &gameDispatcherClient{cc}
}

func (c *gameDispatcherClient) DispatchGame(ctx context.Context, in *DispatchGameRequest, opts ...grpc.CallOption) (*DispatchGameReply, error) {
	out := new(DispatchGameReply)
	err := grpc.Invoke(ctx, "/connectApp.GameDispatcher/DispatchGame", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GameDispatcher service

type GameDispatcherServer interface {
	DispatchGame(context.Context, *DispatchGameRequest) (*DispatchGameReply, error)
}

func RegisterGameDispatcherServer(s *grpc.Server, srv GameDispatcherServer) {
	s.RegisterService(&_GameDispatcher_serviceDesc, srv)
}

func _GameDispatcher_DispatchGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DispatchGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameDispatcherServer).DispatchGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectApp.GameDispatcher/DispatchGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameDispatcherServer).DispatchGame(ctx, req.(*DispatchGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GameDispatcher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "connectApp.GameDispatcher",
	HandlerType: (*GameDispatcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DispatchGame",
			Handler:    _GameDispatcher_DispatchGame_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gameApp.proto",
}

func init() { proto.RegisterFile("gameApp.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x8d, 0xad, 0xda, 0x8e, 0xd6, 0xda, 0xf5, 0x12, 0x8a, 0x62, 0xc9, 0x49, 0x10, 0x72,
	0xd0, 0x4f, 0x50, 0x11, 0x34, 0xb7, 0x92, 0xa3, 0x97, 0x92, 0x4e, 0x86, 0x36, 0xd0, 0x66, 0xc7,
	0xfd, 0x23, 0xf4, 0x53, 0xf8, 0x95, 0x65, 0x97, 0xa4, 0xd9, 0x82, 0x9e, 0xf2, 0xde, 0xdb, 0x81,
	0xdf, 0xdb, 0x9d, 0xc0, 0x68, 0x5d, 0xec, 0x68, 0xce, 0x9c, 0xb2, 0x92, 0x46, 0x0a, 0x40, 0x59,
	0xd7, 0x84, 0x66, 0xce, 0x9c, 0x54, 0x70, 0xfb, 0x56, 0x69, 0x2e, 0x0c, 0x6e, 0xde, 0x8b, 0x1d,
	0xe5, 0xf4, 0x65, 0x49, 0x1b, 0x71, 0x07, 0xc3, 0x02, 0x51, 0xda, 0xda, 0x64, 0x65, 0x1c, 0xcd,
	0xa2, 0xc7, 0x61, 0xde, 0x05, 0x62, 0x0a, 0x03, 0x4d, 0xea, 0x9b, 0x54, 0x56, 0xc6, 0xa7, 0xfe,
	0xf0, 0xe0, 0x45, 0x0c, 0x17, 0x1a, 0xa9, 0xa6, 0xac, 0x8c, 0x7b, 0xfe, 0xa8, 0xb5, 0xc9, 0x4f,
	0x04, 0x93, 0x63, 0x16, 0x6f, 0xf7, 0x8e, 0xd4, 0xb4, 0xeb, 0x48, 0x87, 0x40, 0xcc, 0xe0, 0xb2,
	0x31, 0x1f, 0x52, 0x9b, 0x06, 0x16, 0x46, 0xc1, 0xc4, 0x42, 0x2a, 0xd3, 0x30, 0xc3, 0x28, 0x6c,
	0xd4, 0x3f, 0x6e, 0x54, 0xc0, 0x24, 0x27, 0x96, 0xca, 0x84, 0x57, 0x17, 0xd0, 0xb7, 0xb6, 0x6a,
	0xbb, 0x78, 0xed, 0xb2, 0x4d, 0xc7, 0xf7, 0xda, 0x65, 0xdc, 0x11, 0xbd, 0x16, 0x37, 0xd0, 0x43,
	0xb4, 0x1e, 0x73, 0x96, 0x3b, 0x99, 0x3c, 0xc1, 0x38, 0x44, 0xb8, 0x1b, 0xbb, 0x3e, 0x16, 0x91,
	0xb4, 0xf6, 0x8c, 0x41, 0xde, 0xda, 0xe7, 0x15, 0x5c, 0xbb, 0xb1, 0xf6, 0x91, 0x48, 0x89, 0x05,
	0x5c, 0x85, 0x4f, 0x26, 0x1e, 0xd2, 0x6e, 0x77, 0xe9, 0x1f, 0x8b, 0x9b, 0xde, 0xff, 0x3f, 0xc0,
	0xdb, 0x7d, 0x72, 0xf2, 0x3a, 0xfe, 0x1c, 0xad, 0xa5, 0x5e, 0x2a, 0xc6, 0xa5, 0xff, 0x1b, 0x56,
	0xe7, 0xfe, 0xf3, 0xf2, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xf1, 0xf4, 0x8a, 0x78, 0x25, 0x02, 0x00,
	0x00,
}
