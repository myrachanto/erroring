package httperors

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Code())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "Bad Request", err.Errors())
}
func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Code())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "Not Found", err.Errors())

}
func TestNewAnuthorizedError(t *testing.T) {
	err := NewAnuthorizedError("You are unathorized")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusUnauthorized, err.Code())
	assert.EqualValues(t, "You are unathorized", err.Message())
	assert.EqualValues(t, "Unauthorized", err.Errors())

}

func TestValidStringNotEmpty(t *testing.T) {
	err := ValidStringNotEmpty("string not empty")
	assert.EqualValues(t, false, err.Noerror())
	err1 := ValidStringNotEmpty("")
	assert.EqualValues(t, "The string must not be empty!", err1.Message())
	assert.EqualValues(t, true, err1.Noerror())

}
