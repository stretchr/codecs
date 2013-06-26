package xml

import (
	"github.com/stretchr/codecs"
	"github.com/stretchr/codecs/constants"
	"github.com/stretchr/stew/objects"
	"github.com/stretchr/testify/assert"
	"testing"
)

var xmlCodec XmlCodec

func TestInterface(t *testing.T) {

	assert.Implements(t, (*codecs.Codec)(nil), new(XmlCodec), "XmlCodec")

}

func TestCanMarshalWithCallback(t *testing.T) {
	assert.False(t, xmlCodec.CanMarshalWithCallback(), "XmlCodec cannot marshal with callback")
}

func TestContentType(t *testing.T) {
	assert.Equal(t, constants.ContentTypeXML, xmlCodec.ContentType())
}

func TestExtension(t *testing.T) {
	assert.Equal(t, constants.FileExtensionXML, xmlCodec.FileExtension())
}

func TestMarshal(t *testing.T) {

	data := map[string]interface{}{"name": "Mat"}
	bytes, marshalErr := xmlCodec.Marshal(data, nil)

	if assert.NoError(t, marshalErr) {
		assert.Equal(t, "<?xml version=\"1.0\"?>", string(bytes), "Output")
	}

}

func TestMarshal_map(t *testing.T) {

	data := map[string]interface{}{"name": "Mat", "age": 30, "yesOrNo": true}
	bytes, marshalErr := marshal(data, false, 0, nil)

	if assert.NoError(t, marshalErr) {
		assert.Equal(t, "<object><name>Mat</name><age>30</age><yesOrNo>true</yesOrNo></object>", string(bytes), "Output")
	}

}

func TestMarshal_mapWithTypes(t *testing.T) {

	data := map[string]interface{}{"name": "Mat", "age": 30, "yesOrNo": true}
	options := objects.NewMap(OptionIncludeTypeAttributes, true)
	bytes, marshalErr := marshal(data, false, 0, options)

	if assert.NoError(t, marshalErr) {
		assert.Equal(t, "<object><name type=\"string\">Mat</name><age type=\"int\">30</age><yesOrNo type=\"bool\">true</yesOrNo></object>", string(bytes), "Output")
	}

}

func TestMarshal_arrayOfMaps(t *testing.T) {

	data1 := map[string]interface{}{"name": "Mat"}
	data2 := map[string]interface{}{"name": "Tyler"}
	data3 := map[string]interface{}{"name": "Ryan"}
	array := []map[string]interface{}{data1, data2, data3}
	bytes, marshalErr := marshal(array, false, 0, nil)

	if assert.NoError(t, marshalErr) {
		assert.Equal(t, "<objects><object><name>Mat</name></object><object><name>Tyler</name></object><object><name>Ryan</name></object></objects>", string(bytes), "Output")
	}

}
