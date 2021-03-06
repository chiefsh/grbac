// Code generated by protoc-gen-go.
// source: common.proto
// DO NOT EDIT!

package igrpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Reply struct {
	ErrCode int64  `protobuf:"varint,1,opt,name=ErrCode" json:"ErrCode,omitempty"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=ErrMsg" json:"ErrMsg,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Reply) GetErrCode() int64 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *Reply) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type Integer32 struct {
	Value int32 `protobuf:"zigzag32,1,opt,name=Value" json:"Value,omitempty"`
}

func (m *Integer32) Reset()                    { *m = Integer32{} }
func (m *Integer32) String() string            { return proto.CompactTextString(m) }
func (*Integer32) ProtoMessage()               {}
func (*Integer32) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *Integer32) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Integer64 struct {
	Value int64 `protobuf:"zigzag64,1,opt,name=Value" json:"Value,omitempty"`
}

func (m *Integer64) Reset()                    { *m = Integer64{} }
func (m *Integer64) String() string            { return proto.CompactTextString(m) }
func (*Integer64) ProtoMessage()               {}
func (*Integer64) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *Integer64) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type String struct {
	Value string `protobuf:"bytes,1,opt,name=Value" json:"Value,omitempty"`
}

func (m *String) Reset()                    { *m = String{} }
func (m *String) String() string            { return proto.CompactTextString(m) }
func (*String) ProtoMessage()               {}
func (*String) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *String) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Bool struct {
	Value bool `protobuf:"varint,1,opt,name=Value" json:"Value,omitempty"`
}

func (m *Bool) Reset()                    { *m = Bool{} }
func (m *Bool) String() string            { return proto.CompactTextString(m) }
func (*Bool) ProtoMessage()               {}
func (*Bool) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *Bool) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type KeyValue struct {
	Key   string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=Value" json:"Value,omitempty"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *KeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Reply)(nil), "igrpc.Reply")
	proto.RegisterType((*Integer32)(nil), "igrpc.Integer32")
	proto.RegisterType((*Integer64)(nil), "igrpc.Integer64")
	proto.RegisterType((*String)(nil), "igrpc.String")
	proto.RegisterType((*Bool)(nil), "igrpc.Bool")
	proto.RegisterType((*KeyValue)(nil), "igrpc.KeyValue")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0x4c, 0x2f, 0x2a, 0x48, 0x56,
	0xb2, 0xe4, 0x62, 0x0d, 0x4a, 0x2d, 0xc8, 0xa9, 0x14, 0x92, 0xe0, 0x62, 0x77, 0x2d, 0x2a, 0x72,
	0xce, 0x4f, 0x49, 0x95, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x82, 0x71, 0x85, 0xc4, 0xb8, 0xd8,
	0x5c, 0x8b, 0x8a, 0x7c, 0x8b, 0xd3, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xa0, 0x3c, 0x25,
	0x45, 0x2e, 0x4e, 0xcf, 0xbc, 0x92, 0xd4, 0xf4, 0xd4, 0x22, 0x63, 0x23, 0x21, 0x11, 0x2e, 0xd6,
	0xb0, 0xc4, 0x9c, 0x52, 0x88, 0x66, 0xc1, 0x20, 0x08, 0x07, 0x49, 0x89, 0x99, 0x09, 0xaa, 0x12,
	0x21, 0x98, 0x12, 0x39, 0x2e, 0xb6, 0xe0, 0x92, 0xa2, 0xcc, 0xbc, 0x74, 0x54, 0x79, 0x4e, 0x98,
	0xbc, 0x0c, 0x17, 0x8b, 0x53, 0x7e, 0x7e, 0x0e, 0xaa, 0x2c, 0x07, 0x4c, 0xd6, 0x88, 0x8b, 0xc3,
	0x3b, 0xb5, 0x12, 0xcc, 0x16, 0x12, 0xe0, 0x62, 0xf6, 0x4e, 0xad, 0x84, 0xea, 0x06, 0x31, 0x11,
	0x7a, 0x98, 0x90, 0x4c, 0x4c, 0x62, 0x03, 0x07, 0x80, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x99,
	0xe2, 0xdf, 0xf5, 0x10, 0x01, 0x00, 0x00,
}
