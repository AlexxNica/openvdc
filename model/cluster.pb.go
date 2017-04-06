// Code generated by protoc-gen-go.
// source: cluster.proto
// DO NOT EDIT!

package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Console_Transport int32

const (
	Console_SSH Console_Transport = 0
)

var Console_Transport_name = map[int32]string{
	0: "SSH",
}
var Console_Transport_value = map[string]int32{
	"SSH": 0,
}

func (x Console_Transport) String() string {
	return proto.EnumName(Console_Transport_name, int32(x))
}
func (Console_Transport) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0, 0} }

type NodeState_State int32

const (
	NodeState_REGISTERED NodeState_State = 0
)

var NodeState_State_name = map[int32]string{
	0: "REGISTERED",
}
var NodeState_State_value = map[string]int32{
	"REGISTERED": 0,
}

func (x NodeState_State) String() string {
	return proto.EnumName(NodeState_State_name, int32(x))
}
func (NodeState_State) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{3, 0} }

type Console struct {
	Type     Console_Transport `protobuf:"varint,1,opt,name=type,enum=model.Console_Transport" json:"type,omitempty"`
	BindAddr string            `protobuf:"bytes,2,opt,name=bind_addr" json:"bind_addr,omitempty"`
}

func (m *Console) Reset()                    { *m = Console{} }
func (m *Console) String() string            { return proto.CompactTextString(m) }
func (*Console) ProtoMessage()               {}
func (*Console) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Console) GetType() Console_Transport {
	if m != nil {
		return m.Type
	}
	return Console_SSH
}

func (m *Console) GetBindAddr() string {
	if m != nil {
		return m.BindAddr
	}
	return ""
}

type ExecutorNode struct {
	Id        string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created_at" json:"created_at,omitempty"`
	Console   *Console                   `protobuf:"bytes,3,opt,name=console" json:"console,omitempty"`
	GrpcAddr  string                     `protobuf:"bytes,4,opt,name=grpc_addr" json:"grpc_addr,omitempty"`
	LastState *NodeState                 `protobuf:"bytes,5,opt,name=last_state" json:"last_state,omitempty"`
	NodeId    string                     `protobuf:"bytes,6,opt,name=node_id" json:"node_id,omitempty"`
}

func (m *ExecutorNode) Reset()                    { *m = ExecutorNode{} }
func (m *ExecutorNode) String() string            { return proto.CompactTextString(m) }
func (*ExecutorNode) ProtoMessage()               {}
func (*ExecutorNode) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ExecutorNode) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ExecutorNode) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *ExecutorNode) GetConsole() *Console {
	if m != nil {
		return m.Console
	}
	return nil
}

func (m *ExecutorNode) GetGrpcAddr() string {
	if m != nil {
		return m.GrpcAddr
	}
	return ""
}

func (m *ExecutorNode) GetLastState() *NodeState {
	if m != nil {
		return m.LastState
	}
	return nil
}

func (m *ExecutorNode) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

type SchedulerNode struct {
	Id        string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created_at" json:"created_at,omitempty"`
}

func (m *SchedulerNode) Reset()                    { *m = SchedulerNode{} }
func (m *SchedulerNode) String() string            { return proto.CompactTextString(m) }
func (*SchedulerNode) ProtoMessage()               {}
func (*SchedulerNode) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *SchedulerNode) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SchedulerNode) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type NodeState struct {
	State     NodeState_State            `protobuf:"varint,1,opt,name=state,enum=model.NodeState_State" json:"state,omitempty"`
	CreatedAt *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created_at" json:"created_at,omitempty"`
}

func (m *NodeState) Reset()                    { *m = NodeState{} }
func (m *NodeState) String() string            { return proto.CompactTextString(m) }
func (*NodeState) ProtoMessage()               {}
func (*NodeState) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *NodeState) GetState() NodeState_State {
	if m != nil {
		return m.State
	}
	return NodeState_REGISTERED
}

func (m *NodeState) GetCreatedAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Console)(nil), "model.Console")
	proto.RegisterType((*ExecutorNode)(nil), "model.ExecutorNode")
	proto.RegisterType((*SchedulerNode)(nil), "model.SchedulerNode")
	proto.RegisterType((*NodeState)(nil), "model.NodeState")
	proto.RegisterEnum("model.Console_Transport", Console_Transport_name, Console_Transport_value)
	proto.RegisterEnum("model.NodeState_State", NodeState_State_name, NodeState_State_value)
}

func init() { proto.RegisterFile("cluster.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xb1, 0xce, 0xd3, 0x30,
	0x14, 0x85, 0xff, 0xf4, 0x6f, 0x1a, 0xe5, 0x42, 0xa3, 0xca, 0x42, 0x10, 0x55, 0x20, 0xaa, 0x4c,
	0x1d, 0x2a, 0x07, 0x95, 0x0d, 0x36, 0x20, 0x02, 0x16, 0x06, 0xa7, 0x13, 0x0c, 0x91, 0x63, 0x9b,
	0x34, 0x22, 0x89, 0x23, 0xdb, 0x41, 0xe5, 0x25, 0x78, 0x54, 0x9e, 0x01, 0xc5, 0x6e, 0xda, 0xc2,
	0x0a, 0x9b, 0x7d, 0xef, 0xb9, 0xe7, 0x7e, 0x47, 0x36, 0x2c, 0x59, 0x33, 0x68, 0x23, 0x14, 0xee,
	0x95, 0x34, 0x12, 0xf9, 0xad, 0xe4, 0xa2, 0x59, 0xbf, 0xae, 0x6a, 0x73, 0x1c, 0x4a, 0xcc, 0x64,
	0x9b, 0x56, 0xb2, 0xa1, 0x5d, 0x95, 0xda, 0x7e, 0x39, 0x7c, 0x4d, 0x7b, 0xf3, 0xa3, 0x17, 0x3a,
	0x35, 0x75, 0x2b, 0xb4, 0xa1, 0x6d, 0x7f, 0x3d, 0x39, 0x8f, 0xe4, 0x1b, 0x04, 0x6f, 0x65, 0xa7,
	0x65, 0x23, 0xd0, 0x0e, 0xe6, 0xa3, 0x3a, 0xf6, 0x36, 0xde, 0x36, 0xda, 0xc7, 0xd8, 0xba, 0xe3,
	0x73, 0x17, 0x1f, 0x14, 0xed, 0x74, 0x2f, 0x95, 0x21, 0x56, 0x85, 0x9e, 0x42, 0x58, 0xd6, 0x1d,
	0x2f, 0x28, 0xe7, 0x2a, 0x9e, 0x6d, 0xbc, 0x6d, 0x48, 0xae, 0x85, 0xe4, 0x11, 0x84, 0x97, 0x01,
	0x14, 0xc0, 0x7d, 0x9e, 0x7f, 0x58, 0xdd, 0x25, 0xbf, 0x3c, 0x78, 0x98, 0x9d, 0x04, 0x1b, 0x8c,
	0x54, 0x9f, 0x24, 0x17, 0x28, 0x82, 0x59, 0xcd, 0xed, 0xc2, 0x90, 0xcc, 0x6a, 0x8e, 0x5e, 0x01,
	0x30, 0x25, 0xa8, 0x11, 0xbc, 0xa0, 0xc6, 0xba, 0x3e, 0xd8, 0xaf, 0x71, 0x25, 0x65, 0xd5, 0x08,
	0x3c, 0x85, 0xc2, 0x87, 0x29, 0x03, 0xb9, 0x51, 0xa3, 0x2d, 0x04, 0xcc, 0xb1, 0xc6, 0xf7, 0x76,
	0x30, 0xfa, 0x33, 0x01, 0x99, 0xda, 0x23, 0x7a, 0xa5, 0x7a, 0xe6, 0xd0, 0xe7, 0x0e, 0xfd, 0x52,
	0x40, 0x2f, 0x00, 0x1a, 0xaa, 0x4d, 0xa1, 0x0d, 0x35, 0x22, 0xf6, 0xad, 0xd5, 0xea, 0x6c, 0x35,
	0x42, 0xe7, 0x63, 0x9d, 0xdc, 0x68, 0x50, 0x0c, 0x41, 0x27, 0xb9, 0x28, 0x6a, 0x1e, 0x2f, 0xac,
	0xdb, 0x74, 0x4d, 0xbe, 0xc0, 0x32, 0x67, 0x47, 0xc1, 0x87, 0x46, 0xfc, 0xf7, 0xc0, 0xc9, 0x4f,
	0x0f, 0xc2, 0x0b, 0x10, 0xda, 0x81, 0xef, 0x88, 0xdd, 0xf3, 0x3d, 0xfe, 0x9b, 0x18, 0x3b, 0x6e,
	0x27, 0xfa, 0xa7, 0xbd, 0x4f, 0xc0, 0x77, 0x2b, 0x23, 0x00, 0x92, 0xbd, 0xff, 0x98, 0x1f, 0x32,
	0x92, 0xbd, 0x5b, 0xdd, 0xbd, 0x79, 0xfe, 0xf9, 0xd9, 0xcd, 0x57, 0xa4, 0x27, 0x7d, 0x4c, 0x65,
	0x2f, 0xba, 0xef, 0x9c, 0xa5, 0x16, 0xa6, 0x5c, 0x58, 0xe7, 0x97, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x3f, 0x1e, 0x75, 0x4e, 0xc8, 0x02, 0x00, 0x00,
}
