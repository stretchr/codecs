package xml

import (
	"errors"
	"fmt"
	"github.com/stretchr/codecs/constants"
	"github.com/stretchr/stew/objects"
	"reflect"
	"strings"
)

const (
	OptionIncludeTypeAttributes string = "types"
)

var (
	Indentation                               string = "  "
	XMLDeclaration                            string = "<?xml version=\"1.0\"?>"
	XMLElementFormat                          string = "<%s>%s</%s>"
	XMLElementFormatIndented                  string = "<%s>\n%s%s\n</%s>"
	XMLElementWithTypeAttributeFormat         string = "<%s type=\"%s\">%s</%s>"
	XMLElementWithTypeAttributeFormatIndented string = "<%s type=\"%s\">\n%s%s\n</%s>"
	XMLObjectElementName                      string = "object"
	XMLObjectsElementName                     string = "objects"
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
	return errors.New("codecs: xml: Unmarshalling XML is not supported.")
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

func marshal(object interface{}, doIndent bool, indentLevel int, options objects.Map) ([]byte, error) {

	var nextIndent int = indentLevel + 1
	var output []string

	switch object.(type) {
	case map[string]interface{}:

		var objects []string
		for k, v := range object.(map[string]interface{}) {

			valueBytes, valueMarshalErr := marshal(v, doIndent, nextIndent, options)

			// handle errors
			if valueMarshalErr != nil {
				return nil, valueMarshalErr
			}

			// add the key and value
			el := element(k, v, string(valueBytes), doIndent, nextIndent, options)
			objects = append(objects, el)

		}

		output = append(output, element(XMLObjectElementName, nil, strings.Join(objects, ""), doIndent, nextIndent, nil))

	case []map[string]interface{}:

		var objects []string
		for _, v := range object.([]map[string]interface{}) {

			valueBytes, err := marshal(v, doIndent, nextIndent, options)

			if err != nil {
				return nil, err
			}

			objects = append(objects, string(valueBytes))

		}

		el := strings.Join(objects, "")
		output = append(output, element(XMLObjectsElementName, nil, el, doIndent, nextIndent, nil))

	default:
		// return the value
		output = append(output, fmt.Sprintf("%v", object))
	}

	return []byte(strings.Join(output, "")), nil

}

func element(k string, v interface{}, vString string, doIndent bool, indentLevel int, options objects.Map) string {

	var typeString string
	if v != nil && options.Has(OptionIncludeTypeAttributes) {
		typeString = reflect.TypeOf(v).Name()
	}

	if doIndent {
		indent := strings.Repeat(Indentation, indentLevel)

		if options.Has(OptionIncludeTypeAttributes) {
			return fmt.Sprintf(XMLElementWithTypeAttributeFormatIndented, k, typeString, indent, vString, k)
		} else {
			return fmt.Sprintf(XMLElementFormatIndented, k, indent, vString, k)
		}

	}

	if options.Has(OptionIncludeTypeAttributes) {
		return fmt.Sprintf(XMLElementWithTypeAttributeFormat, k, typeString, vString, k)
	} else {
		return fmt.Sprintf(XMLElementFormat, k, vString, k)
	}

}
