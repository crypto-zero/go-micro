package proto

import (
	"github.com/crypto-zero/go-micro/v2/codec"
	"google.golang.org/protobuf/proto"
)

type Marshaler struct{}

func (Marshaler) Marshal(v interface{}) ([]byte, error) {
	pb, ok := v.(proto.Message)
	if !ok {
		return nil, codec.ErrInvalidMessage
	}

	data, err := proto.Marshal(pb)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (Marshaler) Unmarshal(data []byte, v interface{}) error {
	pb, ok := v.(proto.Message)
	if !ok {
		return codec.ErrInvalidMessage
	}

	return proto.Unmarshal(data, pb)
}

func (Marshaler) String() string {
	return "proto"
}
