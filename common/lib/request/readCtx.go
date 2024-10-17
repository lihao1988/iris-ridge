package request

import (
	"fmt"

	"ridge/common/global"
	"ridge/common/lib/response"
	"ridge/common/pkg/validator"

	"github.com/kataras/iris/v12"
)

// ValidReadJson read json data and validate
func ValidReadJson(c iris.Context, obj interface{}) bool {
	if err := ReadJSON(c, obj); err != nil {
		response.SendMsg(c, nil, err)
		return false
	}

	return validStruct(c, obj)
}

// ValidReadForm read form data and validate
func ValidReadForm(c iris.Context, obj interface{}) bool {
	if err := ReadForm(c, obj); err != nil {
		response.SendMsg(c, nil, err)
		return false
	}

	return validStruct(c, obj)
}

// ValidReadQuery read query data and validate
func ValidReadQuery(c iris.Context, obj interface{}) bool {
	if err := ReadQuery(c, obj); err != nil {
		response.SendMsg(c, nil, err)
		return false
	}

	return validStruct(c, obj)
}

// ReadJSON read json data
func ReadJSON(c iris.Context, obj interface{}) error {
	if err := c.ReadJSON(obj); err != nil {
		return WrapError(err, "Reads JSON fail!")
	}

	return nil
}

// ReadForm read form data
func ReadForm(c iris.Context, obj interface{}) error {
	if err := c.ReadForm(obj); err != nil {
		return WrapError(err, "Reads Form fail!")
	}

	return nil
}

// ReadQuery read query data
func ReadQuery(c iris.Context, obj interface{}) error {
	if err := c.ReadQuery(obj); err != nil {
		return WrapError(err, "Reads Query fail!")
	}

	return nil
}

// validStruct validate data
func validStruct(c iris.Context, obj interface{}) bool {
	err := validator.ValidStruct(global.GValidator, global.GVTranslator, obj)
	if err != nil {
		response.SendMsg(c, nil, err)
		return false
	}

	return true
}

// WrapError warp error
func WrapError(err error, msg string) error {
	if err != nil {
		return fmt.Errorf(msg)
	}

	return nil
}
