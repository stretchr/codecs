package xml

import (
	"fmt"
	"github.com/stretchr/codecs/constants"
	"strings"
)

var (
	Indentation              string = "  "
	XMLDeclaration           string = "<?xml version=\"1.0\"?>"
	XMLElementFormat         string = "<%s>%s</%s>"
	XMLElementFormatIndented string = "<%s>\n%s%s\n</%s>"
	XMLObjectElementName     string = "object"
	XMLObjectsElementName    string = "objects"
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

	var nextIndent int = indentLevel + 1
	var output []string

	switch object.(type) {
	case map[string]interface{}:

		for k, v := range object.(map[string]interface{}) {

			valueBytes, valueMarshalErr := marshal(v, doIndent, nextIndent)

			// handle errors
			if valueMarshalErr != nil {
				return nil, valueMarshalErr
			}

			// add the key and value
			el := element(k, string(valueBytes), doIndent, nextIndent)
			output = appends(output, element(XMLObjectElementName, el, doIndent, nextIndent), doIndent, nextIndent)

		}

	case []map[string]interface{}:

		var objects []string
		for _, v := range object.([]map[string]interface{}) {

			valueBytes, err := marshal(v, doIndent, nextIndent)

			if err != nil {
				return nil, err
			}

			objects = appends(objects, string(valueBytes), doIndent, nextIndent)

		}

		el := strings.Join(objects, "")
		output = appends(output, element(XMLObjectsElementName, el, doIndent, nextIndent), doIndent, nextIndent)

	default:
		// return the value
		output = appends(output, fmt.Sprintf("%v", object), doIndent, nextIndent)
	}

	return []byte(strings.Join(output, "")), nil

}

func appends(a []string, s string, doIndent bool, indentLevel int) []string {
	return append(a, s)
}

func element(k, v string, doIndent bool, indentLevel int) string {
	if doIndent {
		indent := strings.Repeat(Indentation, indentLevel)
		return fmt.Sprintf(XMLElementFormatIndented, k, indent, v, k)
	}
	return fmt.Sprintf(XMLElementFormat, k, v, k)
}
