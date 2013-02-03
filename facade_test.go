package codecs

import (
	"github.com/stretchrcom/codecs/constants"
	"github.com/stretchrcom/codecs/test"
	"github.com/stretchrcom/stew/objects"
	"github.com/stretchrcom/testify/assert"
	"github.com/stretchrcom/testify/mock"
	"testing"
)

/*
	Tests
*/

func TestPublicData(t *testing.T) {

	o := new(test.TestObjectWithFacade)
	o.Mock.On("PublicData", map[string]interface{}{}).Return(objects.Map{"theName": "Mat"}, nil)

	public, err := PublicData(o, map[string]interface{}{})

	if assert.Nil(t, err) {
		assert.Equal(t, public["theName"], "Mat")
	}

	mock.AssertExpectationsForObjects(t, o.Mock)

}

func TestPublicData_WithNil(t *testing.T) {

	o := new(test.TestObjectWithFacade)
	o.Mock.On("PublicData", map[string]interface{}{}).Return(nil, nil)

	public, err := PublicData(o, map[string]interface{}{})

	if assert.Nil(t, err) {
		assert.Nil(t, public, "Nil is OK")
	}

	mock.AssertExpectationsForObjects(t, o.Mock)

}

func TestPublicData_WithError(t *testing.T) {

	o := new(test.TestObjectWithFacade)
	o.Mock.On("PublicData", map[string]interface{}{}).Return(nil, assert.AnError)

	_, err := PublicData(o, map[string]interface{}{})

	assert.Equal(t, assert.AnError, err)
	mock.AssertExpectationsForObjects(t, o.Mock)

}

func TestPublicData_WithRecursion(t *testing.T) {

	o := new(test.TestObjectWithFacade)
	o1 := new(test.TestObjectWithFacade)
	o2 := new(test.TestObjectWithFacade)

	o.Mock.On("PublicData", map[string]interface{}{}).Return(o1, nil)
	o1.Mock.On("PublicData", map[string]interface{}{}).Return(o2, nil)
	o2.Mock.On("PublicData", map[string]interface{}{}).Return(objects.Map{"theName": "Mat"}, nil)

	public, err := PublicData(o, map[string]interface{}{})

	if assert.Nil(t, err) {
		assert.Equal(t, public["theName"], "Mat")
	}

	mock.AssertExpectationsForObjects(t, o.Mock, o1.Mock, o2.Mock)

}

func TestPublicData_WithRecursion_WithObjects(t *testing.T) {

	o := new(test.TestObjectWithFacade)
	o1 := new(test.TestObjectWithFacade)
	o2 := new(test.TestObjectWithFacade)

	args := map[string]interface{}{constants.OptionKeyClientCallback: "~d"}

	o.Mock.On("PublicData", args).Return(o1, nil)
	o1.Mock.On("PublicData", args).Return(o2, nil)
	o2.Mock.On("PublicData", args).Return(objects.Map{"theName": "Mat"}, nil)

	public, err := PublicData(o, args)

	if assert.Nil(t, err) {
		assert.Equal(t, public["theName"], "Mat")
	}

	mock.AssertExpectationsForObjects(t, o.Mock, o1.Mock, o2.Mock)

}
