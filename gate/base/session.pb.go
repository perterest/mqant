// Code generated by protoc-gen-go. DO NOT EDIT.
// source: session.proto

/*
Package basegate is a generated protocol buffer package.

It is generated from these files:
	session.proto

It has these top-level messages:
	SessionImp
*/
package basegate

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SessionImp struct {
	IP        string            `protobuf:"bytes,1,opt,name=IP" json:"IP,omitempty"`
	Network   string            `protobuf:"bytes,2,opt,name=Network" json:"Network,omitempty"`
	UserId    string            `protobuf:"bytes,3,opt,name=UserId" json:"UserId,omitempty"`
	SessionId string            `protobuf:"bytes,4,opt,name=SessionId" json:"SessionId,omitempty"`
	ServerId  string            `protobuf:"bytes,5,opt,name=ServerId" json:"ServerId,omitempty"`
	TraceId   string            `protobuf:"bytes,6,opt,name=TraceId" json:"TraceId,omitempty"`
	SpanId    string            `protobuf:"bytes,7,opt,name=SpanId" json:"SpanId,omitempty"`
	Settings  map[string]string `protobuf:"bytes,8,rep,name=Settings" json:"Settings,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Carrier   map[string]string `protobuf:"bytes,9,rep,name=Carrier" json:"Carrier,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *SessionImp) Reset()                    { *m = SessionImp{} }
func (m *SessionImp) String() string            { return proto.CompactTextString(m) }
func (*SessionImp) ProtoMessage()               {}
func (*SessionImp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SessionImp) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *SessionImp) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *SessionImp) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SessionImp) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *SessionImp) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *SessionImp) GetTraceId() string {
	if m != nil {
		return m.TraceId
	}
	return ""
}

func (m *SessionImp) GetSpanId() string {
	if m != nil {
		return m.SpanId
	}
	return ""
}

func (m *SessionImp) GetSettings() map[string]string {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *SessionImp) GetCarrier() map[string]string {
	if m != nil {
		return m.Carrier
	}
	return nil
}

func init() {
	proto.RegisterType((*SessionImp)(nil), "basegate.sessionImp")
}

func init() { proto.RegisterFile("session.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x49, 0x62, 0xf3, 0x67, 0xb4, 0x22, 0x83, 0xc8, 0x12, 0x3c, 0xd4, 0x9e, 0x7a, 0xca,
	0x41, 0x2f, 0xd2, 0x82, 0x17, 0xf1, 0xb0, 0x17, 0x29, 0x8d, 0x7e, 0x80, 0xad, 0x19, 0x4a, 0xa8,
	0x26, 0x61, 0x77, 0xad, 0xf4, 0x5b, 0xf9, 0x11, 0x65, 0xff, 0xa4, 0x51, 0xf0, 0xd2, 0x5b, 0xde,
	0xbe, 0xf9, 0xbd, 0xc7, 0x4c, 0x60, 0xac, 0x48, 0xa9, 0xba, 0x6d, 0x8a, 0x4e, 0xb6, 0xba, 0xc5,
	0x74, 0x2d, 0x14, 0x6d, 0x84, 0xa6, 0xe9, 0x77, 0x04, 0xe0, 0x3d, 0xfe, 0xd1, 0xe1, 0x39, 0x84,
	0x7c, 0xc9, 0x82, 0x49, 0x30, 0xcb, 0x56, 0x21, 0x5f, 0x22, 0x83, 0xe4, 0x99, 0xf4, 0x57, 0x2b,
	0xb7, 0x2c, 0xb4, 0x8f, 0xbd, 0xc4, 0x2b, 0x88, 0x5f, 0x15, 0x49, 0x5e, 0xb1, 0xc8, 0x1a, 0x5e,
	0xe1, 0x35, 0x64, 0xa5, 0xcf, 0xab, 0xd8, 0x89, 0xb5, 0x86, 0x07, 0xcc, 0x21, 0x2d, 0x49, 0xee,
	0x2c, 0x37, 0xb2, 0xe6, 0x41, 0x9b, 0xae, 0x17, 0x29, 0xde, 0x88, 0x57, 0x2c, 0x76, 0x5d, 0x5e,
	0x9a, 0xae, 0xb2, 0x13, 0x26, 0x30, 0x71, 0x5d, 0x4e, 0xe1, 0x83, 0x49, 0xd3, 0xba, 0x6e, 0x36,
	0x8a, 0xa5, 0x93, 0x68, 0x76, 0x7a, 0x3b, 0x2d, 0xfa, 0xcd, 0x8a, 0x61, 0xab, 0xa2, 0x1f, 0x7a,
	0x6a, 0xb4, 0xdc, 0xaf, 0x0e, 0x0c, 0x2e, 0x20, 0x79, 0x14, 0x52, 0xd6, 0x24, 0x59, 0x66, 0xf1,
	0x9b, 0x7f, 0x71, 0x3f, 0xe3, 0xe8, 0x9e, 0xc8, 0x17, 0x30, 0xfe, 0x93, 0x8b, 0x17, 0x10, 0x6d,
	0x69, 0xef, 0x8f, 0x67, 0x3e, 0xf1, 0x12, 0x46, 0x3b, 0xf1, 0xfe, 0x49, 0xfe, 0x76, 0x4e, 0xcc,
	0xc3, 0xfb, 0x20, 0x9f, 0xc3, 0xd9, 0xef, 0xd4, 0x63, 0xd8, 0x75, 0x6c, 0xff, 0xe1, 0xdd, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x9e, 0x5d, 0x17, 0xd4, 0x01, 0x00, 0x00,
}
