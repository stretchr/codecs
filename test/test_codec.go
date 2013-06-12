package test

import (
	"github.com/stretchr/testify/mock"
)

// TestCodec is a codec mock object useful for testing code
// that relies on codecs.
//
// The mocking capabilities are provided by stretchrcom/testify framework.
type TestCodec struct {
	mock.Mock
}

// Marshal is a mocked function that records the activity in the Mock object and
// returns the values setup in user code by the .On.Return pair.
func (c *TestCodec) Marshal(object interface{}, options map[string]interface{}) ([]byte, error) {
	// TODO: generalise this into the Mock framework - it's likely other
	// people will need to do similar things.
	allArgs := []interface{}{object, options}
	args := c.Mock.Called(allArgs...)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]byte), args.Error(1)
}

// Unmarshal is a mocked function that records the activity in the Mock object and
// returns the values setup in user code by the .On.Return pair.
func (c *TestCodec) Unmarshal(data []byte, obj interface{}) error {
	return c.Mock.Called(data, obj).Error(0)
}

// ContentType is a mocked function that records the activity in the Mock object and
// returns the values setup in user code by the .On.Return pair.
func (c *TestCodec) ContentType() string {
	return c.Mock.Called().String(0)
}

// FileExtension is a mocked function that records the activity in the Mock object and
// returns the values setup in user code by the .On.Return pair.
func (c *TestCodec) FileExtension() string {
	return c.Mock.Called().String(0)
}

// CanMarshalWithCallback is a mocked function that records the activity in the Mock object and
// returns the values setup in user code by the .On.Return pair.
func (c *TestCodec) CanMarshalWithCallback() bool {
	return c.Mock.Called().Bool(0)
}
