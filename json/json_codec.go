package json

import (
	jsonEncoding "encoding/json"
	"github.com/stretchr/codecs/constants"
)

// JsonCodec converts objects to and from JSON.
type JsonCodec struct{}

// Converts an object to JSON.
func (c *JsonCodec) Marshal(object interface{}, options map[string]interface{}) ([]byte, error) {
	return jsonEncoding.Marshal(object)
}

// Unmarshal converts JSON into an object.
func (c *JsonCodec) Unmarshal(data []byte, obj interface{}) error {
	return jsonEncoding.Unmarshal(data, obj)
}

// ContentType returns the content type for this codec.
func (c *JsonCodec) ContentType() string {
	return constants.ContentTypeJSON
}

// FileExtension returns the file extension for this codec.
func (c *JsonCodec) FileExtension() string {
	return constants.FileExtensionJSON
}

// CanMarshalWithCallback returns whether this codec is capable of marshalling a response containing a callback.
func (c *JsonCodec) CanMarshalWithCallback() bool {
	return false
}
