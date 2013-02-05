package msgpack

import (
	"github.com/stretchrcom/codecs/constants"
	"github.com/ugorji/go-msgpack"
)

// MsgpackCodec converts objects to and from Msgpack.
type MsgpackCodec struct{}

// Converts an object to Msgpack.
func (c *MsgpackCodec) Marshal(object interface{}, options map[string]interface{}) ([]byte, error) {
	return msgpack.Marshal(object)
}

// Unmarshal converts Msgpack into an object.
func (c *MsgpackCodec) Unmarshal(data []byte, obj interface{}) error {
	return msgpack.Unmarshal(data, obj, nil)
}

// ContentType returns the content type for this codec.
func (c *MsgpackCodec) ContentType() string {
	return constants.ContentTypeMsgpack
}

// FileExtension returns the file extension for this codec.
func (c *MsgpackCodec) FileExtension() string {
	return constants.FileExtensionMsgpack
}

// CanMarshalWithCallback returns whether this codec is capable of marshalling a response containing a callback.
func (c *MsgpackCodec) CanMarshalWithCallback() bool {
	return false
}
