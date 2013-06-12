package test

import (
	"github.com/stretchr/testify/mock"
)

// TestObjectWithFacade is a mock test object that implements the Facade interface.
//
// The mocking capabilities are provided by stretchrcom/testify framework.
type TestObjectWithFacade struct {
	mock.Mock
}

// PublicData is a mocked function, as defined by the Facade interface,
// that records the activity in the Mock object and
// returns the values setup in user code by the .On.Return pair.
func (o *TestObjectWithFacade) PublicData(options map[string]interface{}) (interface{}, error) {
	args := o.Mock.Called(options)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0), args.Error(1)
}
