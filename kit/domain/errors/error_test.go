package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExperimentErrorShouldBeASentinelError(t *testing.T) {
	t.Parallel()

	err := NewProjectError("error_test", TestingError, "test")
	assert.Error(t, err, "error should be an error")
}

func TestExperimentErrorShouldReturnDetail(t *testing.T) {
	t.Parallel()

	err := NewProjectError("error_test", TestingError, "test")
	assert.NotEmptyf(t, err.Error(), "error should not be empty")
}

func TestExperimentErrorShouldProvideAMethodToGetType(t *testing.T) {
	t.Parallel()

	err := NewProjectError("error_test", TestingError, "test")
	assert.Equal(t, TestingError, err.ErrorType(), "error type should be the same")
}

func TestExperimentErrorShouldBeAnError(t *testing.T) {
	t.Parallel()

	err := NewProjectError("error_test", TestingError, "test")
	assert.Implements(t, (*error)(nil), err, "error should implement error interface")
}
