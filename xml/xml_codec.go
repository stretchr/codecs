package xml

import (
	"fmt"
	"github.com/stretchr/codecs/constants"
	"strings"
)

var (
	XMLDeclaration             string = "<?xml version=\"1.0\"?>"
	XMLElementFormat           string = "<%s>%s</%s>"
	XMLCollectionWrapperFormat string = "<objects>%s</objects>"
	XMLObjectWrapperFormat     string = "<object>%s</object>"
)

type XmlCodec struct{}

// Marshal converts an object to a []byte representation.
// You can optionally pass additional arguments to further customize this call.
func (c *XmlCodec) Marshal(object interface{}, options map[string]interface{}) ([]byte, error) {

	var output []string

	// add the declaration
	output = append(output, XMLDeclaration)

	// return the output
	return []byte(strings.Join(output, "")), nil
}

// Unmarshal converts a []byte representation into an object.
func (c *XmlCodec) Unmarshal(data []byte, obj interface{}) error {
	return nil
}

// ContentType gets the content type that this codec handles.
func (c *XmlCodec) ContentType() string {
	return constants.ContentTypeXML
}

// FileExtension returns the file extension by which this codec is represented.
func (c *XmlCodec) FileExtension() string {
	return constants.FileExtensionXML
}

// CanMarshalWithCallback indicates whether this codec is capable of marshalling a response with
// a callback parameter.
func (c *XmlCodec) CanMarshalWithCallback() bool {
	return false
}

/*
  Custom XML marshalling
*/

func marshal(object interface{}, doIndent bool, indentLevel int) ([]byte, error) {

	var output []string

	switch object.(type) {
	case map[string]interface{}:

		for k, v := range object.(map[string]interface{}) {

			valueBytes, valueMarshalErr := marshal(v, doIndent, indentLevel+1)

			// handle errors
			if valueMarshalErr != nil {
				return nil, valueMarshalErr
			}

			// add the key and value
			element := fmt.Sprintf(XMLElementFormat, k, string(valueBytes), k)
			output = appends(output, fmt.Sprintf(XMLObjectWrapperFormat, element), doIndent, indentLevel+1)

		}

	case []map[string]interface{}:

		var objects []string
		for _, v := range object.([]map[string]interface{}) {

			valueBytes, err := marshal(v, doIndent, indentLevel+1)

			if err != nil {
				return nil, err
			}

			objects = appends(objects, string(valueBytes), doIndent, indentLevel+1)

		}

		output = appends(output, fmt.Sprintf(XMLCollectionWrapperFormat, strings.Join(objects, "")), doIndent, indentLevel+1)

	default:
		// return the value
		output = appends(output, fmt.Sprintf("%v", object), doIndent, indentLevel+1)
	}

	return []byte(strings.Join(output, "")), nil

}

func appends(a []string, s string, doIndent bool, indentLevel int) []string {
	return append(a, s)
}
