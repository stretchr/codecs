package csv

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/stretchr/codecs/constants"
	"reflect"
	"strings"
)

type CsvCodec struct{}

// Converts an object to JSON.
func (c *CsvCodec) Marshal(object interface{}, options map[string]interface{}) ([]byte, error) {

	// collect the data rows in a consistent type

	dataRows := make([]map[string]interface{}, 0)
	switch object.(type) {
	case map[string]interface{}:
		dataRows = append(dataRows, object.(map[string]interface{}))
	case []map[string]interface{}:
		dataRows = object.([]map[string]interface{})
	}

	// collect the fields
	var fields []string
	for _, m := range dataRows {

		// for each field
		for k, _ := range m {

			shouldAdd := true
			for _, field := range fields {
				if strings.ToLower(field) == strings.ToLower(k) {
					shouldAdd = false
					break
				}
			}

			if shouldAdd {
				// add this new field
				fields = append(fields, k)
			}

		}

	}

	// make a new CSV writer
	byteBuffer := new(bytes.Buffer)
	writer := csv.NewWriter(byteBuffer)

	// write the fields
	writer.Write(fields)

	// now write the data
	for _, row := range dataRows {

		rowData := make([]string, len(fields))

		// do it each field at a time
		for k, v := range row {

			// find the field index
			var fieldIndex int
			for index, f := range fields {
				if strings.ToLower(f) == strings.ToLower(k) {
					fieldIndex = index
					break
				}
			}

			// set the field
			rowData[fieldIndex] = fmt.Sprintf("%v", v)

		}

		// write the row
		writer.Write(rowData)

	}

	// finish writing
	writer.Flush()
	if writer.Error() != nil {
		return nil, writer.Error()
	}

	return byteBuffer.Bytes(), nil
}

// Unmarshal converts JSON into an object.
func (c *CsvCodec) Unmarshal(data []byte, obj interface{}) error {

	// check the value
	rv := reflect.ValueOf(obj)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(obj)}
	}

	reader := csv.NewReader(bytes.NewReader(data))
	records, readErr := reader.ReadAll()

	if readErr != nil {
		return readErr
	}

	lenRecords := len(records)

	if lenRecords == 0 {

		// no records
		return nil

	} else if lenRecords == 1 {

		// no records (first line should be header)
		return nil

	} else if lenRecords == 2 {

		// one record

		// get the object
		object := mapFromFieldsAndRow(records[0], records[1])

		// set the obj value
		rv.Elem().Set(reflect.ValueOf(object))

	} else {

		// multiple records

		// make a new array to hold the data
		rows := make([]map[string]interface{}, lenRecords-1)

		// collect the fields
		fields := records[0]

		// add each row
		for i := 1; i < lenRecords; i++ {
			rows[i-1] = mapFromFieldsAndRow(fields, records[i])
		}

		// set the obj value
		rv.Elem().Set(reflect.ValueOf(rows))

	}

	return nil
}

// ContentType returns the content type for this codec.
func (c *CsvCodec) ContentType() string {
	return constants.ContentTypeCSV
}

// FileExtension returns the file extension for this codec.
func (c *CsvCodec) FileExtension() string {
	return constants.FileExtensionCSV
}

// CanMarshalWithCallback returns whether this codec is capable of marshalling a response containing a callback.
func (c *CsvCodec) CanMarshalWithCallback() bool {
	return false
}

// mapFromFieldsAndRow makes a map[string]interface{} from the given fields and
// row data.
func mapFromFieldsAndRow(fields, row []string) map[string]interface{} {
	m := make(map[string]interface{})

	for index, item := range row {
		m[fields[index]] = item
	}

	return m
}
