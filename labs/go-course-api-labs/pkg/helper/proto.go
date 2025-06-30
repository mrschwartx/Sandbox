package helper

import (
	"database/sql"
	"errors"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var ErrInvalidPayload = errors.New("invalid payload")

func ProtoUnmarshal(b []byte, v proto.Message) error {
	data := &anypb.Any{}

	err := protojson.Unmarshal(b, data)
	if err != nil {
		return ErrInvalidPayload
	}

	err = anypb.UnmarshalTo(data, v, proto.UnmarshalOptions{})
	if err != nil {
		return ErrInvalidPayload
	}

	return nil
}

func ProtoMarshal(v proto.Message) ([]byte, error) {
	data, err := anypb.New(v)
	if err != nil {
		return nil, err
	}

	return protojson.Marshal(data)
}

func ProtoFromSQLNullTime(t sql.NullTime) *timestamppb.Timestamp {
	if !t.Valid {
		return nil
	}

	return timestamppb.New(t.Time)
}
