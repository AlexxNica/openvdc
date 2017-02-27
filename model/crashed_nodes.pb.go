// Code generated by protoc-gen-go.
// source: crashed_nodes.proto
// DO NOT EDIT!

/*
Package model is a generated protocol buffer package.

It is generated from these files:
	crashed_nodes.proto

It has these top-level messages:
	CrashedNode
*/
package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CrashedNode struct {
	Agentid      string                     `protobuf:"bytes,1,opt,name=agentid" json:"agentid,omitempty"`
	Agentmesosid string                     `protobuf:"bytes,2,opt,name=agentmesosid" json:"agentmesosid,omitempty"`
	Reconnected  bool                       `protobuf:"varint,3,opt,name=reconnected" json:"reconnected,omitempty"`
	CreatedAt    *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
}

func (m *CrashedNode) Reset()                    { *m = CrashedNode{} }
func (m *CrashedNode) String() string            { return proto.CompactTextString(m) }
func (*CrashedNode) ProtoMessage()               {}
func (*CrashedNode) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CrashedNode) GetAgentID() string {
	if m != nil {
		return m.Agentid
	}
	return ""
}

func (m *CrashedNode) GetAgentMesosID() string {
	if m != nil {
		return m.Agentmesosid
	}
	return ""
}

func (m *CrashedNode) GetReconnected() bool {
	if m != nil {
		return m.Reconnected
	}
	return false
}

func (m *CrashedNode) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*CrashedNode)(nil), "model.CrashedNode")
}

func init() { proto.RegisterFile("crashed_nodes.proto", fileDescriptor0) }

var fileDescriptor3 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x50, 0x3d, 0x4f, 0x04, 0x21,
	0x10, 0xcd, 0xfa, 0x7d, 0xac, 0x15, 0x36, 0xe4, 0x12, 0xe3, 0xe5, 0x2a, 0x2b, 0x48, 0xb4, 0x32,
	0x56, 0x6a, 0x6f, 0xb1, 0xb1, 0xb2, 0xb9, 0xb0, 0x30, 0xb2, 0x9b, 0x2c, 0x0c, 0x81, 0x39, 0xa3,
	0xff, 0xc8, 0x9f, 0x29, 0xe1, 0xdc, 0x78, 0xd7, 0xbd, 0x79, 0x5f, 0x79, 0xc0, 0xae, 0x4c, 0xd2,
	0x79, 0x00, 0xbb, 0x09, 0x68, 0x21, 0xcb, 0x98, 0x90, 0x90, 0x9f, 0xfa, 0x72, 0x4c, 0xcb, 0x47,
	0x37, 0xd2, 0xb0, 0xed, 0xa5, 0x41, 0xaf, 0x1c, 0x4e, 0x3a, 0x38, 0x55, 0xf5, 0x7e, 0xfb, 0xa1,
	0x22, 0x7d, 0x47, 0xc8, 0x8a, 0x46, 0x0f, 0x99, 0xb4, 0x8f, 0xff, 0x68, 0xd7, 0xb1, 0xfe, 0x69,
	0x58, 0xfb, 0xb2, 0xeb, 0x7e, 0x2d, 0x6d, 0x5c, 0xb0, 0x73, 0xed, 0x20, 0xd0, 0x68, 0x45, 0xb3,
	0x6a, 0x6e, 0x17, 0xdd, 0x7c, 0xf2, 0x35, 0xbb, 0xac, 0xb0, 0xe4, 0x31, 0x17, 0xf9, 0xa8, 0xca,
	0x07, 0x1c, 0x5f, 0xb1, 0x36, 0x81, 0xc1, 0x10, 0xc0, 0x10, 0x58, 0x71, 0x5c, 0x2c, 0x17, 0xdd,
	0x3e, 0xc5, 0x1f, 0x18, 0x33, 0x09, 0x74, 0x81, 0x1b, 0x4d, 0xe2, 0xa4, 0x18, 0xda, 0xbb, 0xa5,
	0x74, 0x88, 0x6e, 0x02, 0x39, 0xcf, 0x96, 0x6f, 0xf3, 0xca, 0x6e, 0xf1, 0xe7, 0x7e, 0xa2, 0xe7,
	0x9b, 0xf7, 0xeb, 0xbd, 0x97, 0xea, 0xaf, 0x3c, 0x28, 0x8c, 0x10, 0x3e, 0xad, 0x51, 0xf5, 0x23,
	0xfa, 0xb3, 0x9a, 0xbf, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xda, 0x16, 0xab, 0x49, 0x2d, 0x01,
	0x00, 0x00,
}