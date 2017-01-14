package backend

import (
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

// ProtoClusterBackend is the ClusterBackend interface wrapper
// that accepts proto.Message instead of raw bytes.
type ProtoClusterBackend interface {
	Backend() ClusterBackend
	Connect(dest []string) error
	Close() error
	Register(nodeID string, value proto.Message) error
	Find(nodeID string, value proto.Message) error
}

type ProtoClusterWrapper struct {
	backend ClusterBackend
}

func NewProtoClusterWrapper(bk ClusterBackend) ProtoClusterBackend {
	return &ProtoClusterWrapper{bk}
}

func (p *ProtoClusterWrapper) Backend() ClusterBackend {
	return p.backend
}

func (p *ProtoClusterWrapper) Connect(dest []string) error {
	return p.backend.Connect(dest)
}

func (p *ProtoClusterWrapper) Close() error {
	return p.backend.Close()
}

func (p *ProtoClusterWrapper) Register(key string, value proto.Message) error {
	buf, err := proto.Marshal(value)
	if err != nil {
		return errors.Wrapf(err, "Failed to marshall %T", value)
	}

	return p.backend.Register(key, buf)
}

func (p *ProtoClusterWrapper) Find(nodeID string, v proto.Message) error {
	buf, err := p.backend.Find(nodeID)
	if err != nil {
		return errors.Wrapf(err, "Find to %s", nodeID)
	}
	return proto.Unmarshal(buf, v)
}
