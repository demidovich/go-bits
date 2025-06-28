package promise

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PromiseSuccess(t *testing.T) {
	p := NewPromise(successJob())

	var res string
	var err error
	select {
	case res = <-p.Result:
	case err = <-p.Error:
	}

	assert.Equal(t, "success", res)
	assert.Nil(t, err)
}

func Test_PromiseError(t *testing.T) {
	p := NewPromise(failedJob())

	var res string
	var err error
	select {
	case res = <-p.Result:
	case err = <-p.Error:
	}

	assert.Equal(t, "", res)
	assert.NotNil(t, err)
}

func successJob() func() (string, error) {
	return func() (string, error) {
		return "success", nil
	}
}

func failedJob() func() (string, error) {
	return func() (string, error) {
		return "success", errors.New("failed")
	}
}
