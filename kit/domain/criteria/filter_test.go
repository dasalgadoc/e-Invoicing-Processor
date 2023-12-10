package criteria

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterWithEmptyFieldShouldReturnNil(t *testing.T) {
	t.Parallel()

	filters := [][]string{{"", "==", "John", "OR"}}
	filter := newFilters(filters)

	assert.Nilf(t, filter, "Filter should be nil")
}

func TestFilterWithInvalidLogicalOperatorShouldBeANDLogical(t *testing.T) {
	t.Parallel()

	filters := [][]string{{"name", "==", "John", "NOR"}}
	filter := *newFilters(filters)

	assert.Equalf(t, "AND", filter[0].logical.string(), "Logical operator should be AND")
}

func TestFilterShouldProvideAMethodToGetOperatorAsString(t *testing.T) {
	t.Parallel()

	filters := [][]string{{"name", "==", "John", "OR"}}
	filter := *newFilters(filters)

	assert.Equalf(t, "==", filter[0].operator.string(), "Operator should be ==")
}
