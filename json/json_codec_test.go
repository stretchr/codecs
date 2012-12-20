package json

import (
	"github.com/stretchrcom/codecs"
	"github.com/stretchrcom/codecs/constants"
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestInterface(t *testing.T) {

	assert.Implements(t, (*codecs.Codec)(nil), new(JsonCodec), "JsonCodec")

}

func TestMarshal(t *testing.T) {

	codec := new(JsonCodec)

	obj := make(map[string]string)
	obj["name"] = "Mat"

	jsonString, jsonError := codec.Marshal(obj, nil)

	if jsonError != nil {
		t.Errorf("Shouldn't return error: %s", jsonError)
	}

	assert.Equal(t, string(jsonString), `{"name":"Mat"}`)

}

func TestUnmarshal(t *testing.T) {

	codec := new(JsonCodec)
	jsonString := `{"name":"Mat"}`
	var object map[string]interface{}

	err := codec.Unmarshal([]byte(jsonString), &object)

	if err != nil {
		t.Errorf("Shouldn't return error: %s", err)
	}

	assert.Equal(t, "Mat", object["name"])

}

func TestResponseContentType(t *testing.T) {

	codec := new(JsonCodec)
	assert.Equal(t, codec.ContentType(), constants.ContentTypeJSON)

}

func TestFileExtension(t *testing.T) {

	codec := new(JsonCodec)
	assert.Equal(t, constants.FileExtensionJSON, codec.FileExtension())

}

func TestCanMarshalWithCallback(t *testing.T) {

	codec := new(JsonCodec)
	assert.False(t, codec.CanMarshalWithCallback())

}
