package csv

import (
	"github.com/stretchr/codecs"
	"github.com/stretchr/codecs/constants"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestInterface(t *testing.T) {

	assert.Implements(t, (*codecs.Codec)(nil), new(CsvCodec), "CsvCodec")

}

func TestMapFromFieldsAndRow(t *testing.T) {

	fields := []string{"field1", "field2", "field3"}
	row := []string{"one", "two", "three"}

	m := mapFromFieldsAndRow(fields, row)

	if assert.NotNil(t, m) {

		assert.Equal(t, "one", m["field1"])
		assert.Equal(t, "two", m["field2"])
		assert.Equal(t, "three", m["field3"])

	}

}

func TestMarshal_SingleObject(t *testing.T) {

	obj := map[string]interface{}{"field1": "one", "field2": "two", "field3": "three"}

	csvCodec := new(CsvCodec)
	bytes, marshalErr := csvCodec.Marshal(obj, nil)

	if assert.NoError(t, marshalErr) {

		assert.Equal(t, "field1,field2,field3\none,two,three\n", string(bytes))

	}

}

func TestMarshal_MultipleObjects(t *testing.T) {

	arr := make([]map[string]interface{}, 3)
	arr[0] = map[string]interface{}{"field1": "oneA", "field2": "twoA", "field3": "threeA"}
	arr[1] = map[string]interface{}{"field1": "oneB", "field2": "twoB", "field3": "threeB"}
	arr[2] = map[string]interface{}{"field1": "oneC", "field2": "twoC", "field3": "threeC"}

	csvCodec := new(CsvCodec)
	bytes, marshalErr := csvCodec.Marshal(arr, nil)

	if assert.NoError(t, marshalErr) {

		assert.Equal(t, "field1,field2,field3\noneA,twoA,threeA\noneB,twoB,threeB\noneC,twoC,threeC\n", string(bytes))

	}

}

func TestMarshal_MultipleObjects_WithDisimilarSchema(t *testing.T) {

	arr := make([]map[string]interface{}, 3)
	arr[0] = map[string]interface{}{"name": "Mat", "age": 30, "language": "en"}
	arr[1] = map[string]interface{}{"first_name": "Tyler", "age": 28, "last_name": "Bunnell"}
	arr[2] = map[string]interface{}{"name": "Ryan", "age": 26, "speaks": "english"}

	csvCodec := new(CsvCodec)
	bytes, marshalErr := csvCodec.Marshal(arr, nil)

	if assert.NoError(t, marshalErr) {

		assert.Equal(t, "name,age,language,first_name,last_name,speaks\nMat,30,en,\"\",\"\",\"\"\n\"\",28,\"\",Tyler,Bunnell,\"\"\nRyan,26,\"\",\"\",\"\",english\n", string(bytes))

	}

}

func TestUnmarshal_SingleObject(t *testing.T) {

	raw := "field_a,field_b,field_c\nrow1a,row1b,row1c\n"

	csvCodec := new(CsvCodec)

	var obj interface{}
	csvCodec.Unmarshal([]byte(raw), &obj)

	if assert.NotNil(t, obj, "Unmarshal should make an object") {
		if object, ok := obj.(map[string]interface{}); ok {

			assert.Equal(t, "row1a", object["field_a"])
			assert.Equal(t, "row1b", object["field_b"])
			assert.Equal(t, "row1c", object["field_c"])

		} else {
			t.Errorf("Expected to be array type, not %s.", reflect.TypeOf(obj).Elem().Name())
		}
	}

}

func TestUnmarshal_MultipleObjects(t *testing.T) {

	raw := "field_a,field_b,field_c\nrow1a,row1b,row1c\nrow2a,row2b,row2c\nrow3a,row3b,row3c"

	csvCodec := new(CsvCodec)

	var obj interface{}
	csvCodec.Unmarshal([]byte(raw), &obj)

	if assert.NotNil(t, obj, "Unmarshal should make an object") {
		if array, ok := obj.([]map[string]interface{}); ok {

			if assert.Equal(t, 3, len(array), "Should be 3 items") {

				assert.Equal(t, "row1a", array[0]["field_a"])
				assert.Equal(t, "row1b", array[0]["field_b"])
				assert.Equal(t, "row1c", array[0]["field_c"])

				assert.Equal(t, "row2a", array[1]["field_a"])
				assert.Equal(t, "row2b", array[1]["field_b"])
				assert.Equal(t, "row2c", array[1]["field_c"])

				assert.Equal(t, "row3a", array[2]["field_a"])
				assert.Equal(t, "row3b", array[2]["field_b"])
				assert.Equal(t, "row3c", array[2]["field_c"])

			}

		} else {
			t.Errorf("Expected to be array type, not %s.", reflect.TypeOf(obj).Elem().Name())
		}
	}

}

func TestResponseContentType(t *testing.T) {

	codec := new(CsvCodec)
	assert.Equal(t, codec.ContentType(), constants.ContentTypeCSV)

}

func TestFileExtension(t *testing.T) {

	codec := new(CsvCodec)
	assert.Equal(t, constants.FileExtensionCSV, codec.FileExtension())

}

func TestCanMarshalWithCallback(t *testing.T) {

	codec := new(CsvCodec)
	assert.False(t, codec.CanMarshalWithCallback())

}
