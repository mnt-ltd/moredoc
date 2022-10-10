package jsonpb

import (
	"bytes"
	"fmt"
	"io"
	"reflect"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	jsoniter "github.com/json-iterator/go"
)

var (
	typeProtoMessage = reflect.TypeOf((*proto.Message)(nil)).Elem()
	json             = jsoniter.ConfigCompatibleWithStandardLibrary
)

// JSONPb is a Marshaler which marshals/unmarshals into/from JSON
// with the "github.com/gogo/protobuf/jsonpb".
// It supports fully functionality of protobuf unlike JSONBuiltin.
type JSONPb jsonpb.Marshaler

// ContentType always returns "application/json".
func (*JSONPb) ContentType() string {
	return "application/json"
}

// Marshal marshals "v" into JSON
func (j *JSONPb) Marshal(v interface{}) ([]byte, error) {
	if pb, ok := v.(proto.Message); ok {
		buf, err := json.Marshal(&pb)
		if err != nil {
			return nil, err
		}
		return buf, nil
	}
	return nil, fmt.Errorf("unexpected type %T does not implement %s", v, typeProtoMessage)
}

// Unmarshal unmarshals JSON "data" into "v"
func (j *JSONPb) Unmarshal(data []byte, v interface{}) error {
	if pb, ok := v.(proto.Message); ok {
		return jsonpb.Unmarshal(bytes.NewReader(data), pb)
	}
	return fmt.Errorf("unexpected type %T does not implement %s", v, typeProtoMessage)
}

// NewDecoder returns a Decoder which reads JSON stream from "r".
func (j *JSONPb) NewDecoder(r io.Reader) gwruntime.Decoder {
	return gwruntime.DecoderFunc(func(v interface{}) error {
		if pb, ok := v.(proto.Message); ok {
			return jsonpb.Unmarshal(r, pb)
		}
		return fmt.Errorf("unexpected type %T does not implement %s", v, typeProtoMessage)
	})
}

// NewEncoder returns an Encoder which writes JSON stream into "w".
func (j *JSONPb) NewEncoder(w io.Writer) gwruntime.Encoder {
	return gwruntime.EncoderFunc(func(v interface{}) error {
		if pb, ok := v.(proto.Message); ok {
			marshalFn := (*jsonpb.Marshaler)(j).Marshal
			return marshalFn(w, pb)
		}
		return fmt.Errorf("unexpected type %T does not implement %s", v, typeProtoMessage)
	})
}
