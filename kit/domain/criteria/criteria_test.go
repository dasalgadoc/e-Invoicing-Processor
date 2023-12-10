package criteria

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestNewCriteriaShouldBuildStruct(t *testing.T) {
	t.Parallel()

	offset := 0
	limit := 10

	criteria := NewCriteria(offset, limit, "", "", nil)

	assert.Equalf(t, offset, criteria.Offset(), "Offset should be %d", offset)
	assert.Equalf(t, limit, criteria.Limit(), "Limit should be %d", limit)
}

func TestNewCriteriaWithoutOrderShouldGenerateNilOrderAttribute(t *testing.T) {
	t.Parallel()

	offset := 0
	limit := 10

	criteria := NewCriteria(offset, limit, "", "", nil)

	assert.Nilf(t, criteria.order, "Order should be nil")
}

func TestNewCriteriaWithNoValidLimitShouldGenerateZeroLimit(t *testing.T) {
	t.Parallel()

	offset := 0
	limit := -1

	criteria := NewCriteria(offset, limit, "", "", nil)

	assert.Equalf(t, 0, criteria.Limit(), "Limit should be 0")
}

func TestNewCriteriaWithNoValidOffsetShouldGenerateZeroOffset(t *testing.T) {
	t.Parallel()

	offset := -1
	limit := 10

	criteria := NewCriteria(offset, limit, "", "", nil)

	assert.Equalf(t, 0, criteria.Offset(), "Offset should be 0")
}

func TestNewCriteriaWithoutFiltersGenerateNilFiltersAttribute(t *testing.T) {
	t.Parallel()

	offset := 0
	limit := 10

	criteria := NewCriteria(offset, limit, "", "", nil)

	assert.Nilf(t, criteria.filters, "Filters should be nil for nil entry")

	criteria = NewCriteria(offset, limit, "", "", [][]string{})
	assert.Nilf(t, criteria.filters, "Filters should be nil for empty entry")
}

func TestNewCriteriaWithOrderShouldGenerateNotNilOrderAttribute(t *testing.T) {
	t.Parallel()

	offset := 0
	limit := 10
	by := "name"
	orderType := "asc"
	expectedOrder := strings.ToUpper(orderType)

	criteria := NewCriteria(offset, limit, by, orderType, nil)

	assert.NotNilf(t, criteria.order, "Order should be not nil")
	assert.Equalf(t, by, criteria.OrderBy(), "By should be %s", by)
	assert.Equalf(t, expectedOrder, criteria.OrderType(), "OrderType should be %s", expectedOrder)
}

func TestCriteriaWithOrderAttributeButNoOrderTypeShouldBeASCByDefault(t *testing.T) {
	t.Parallel()

	offset := 0
	limit := 10
	by := "name"

	criteria := NewCriteria(offset, limit, by, "", nil)

	assert.Equalf(t, by, criteria.OrderBy(), "By should be %s", by)
	assert.Equalf(t, "ASC", criteria.OrderType(), "OrderType should be %s", "ASC")
}

func TestCriteriaWithFiltersShouldGenerateNotNilFiltersAttribute(t *testing.T) {
	t.Parallel()

	offset := 0
	limit := 10
	filters := [][]string{{"name", "==", "John", "OR"}}

	criteria := NewCriteria(offset, limit, "", "", filters)

	assert.NotNilf(t, criteria.Filters(), "Filters should be not nil")
	assert.Equalf(t, 1, len(criteria.Filters()), "Filters should be %d", 1)
}

func TestACriteriaWithInvalidFiltersShouldSkipBadFilters(t *testing.T) {
	t.Parallel()

	offset := 0
	limit := 10
	filters := [][]string{{"name", "==", "John", "AND"}, {"name", "=="}}

	criteria := NewCriteria(offset, limit, "", "", filters)

	assert.NotNilf(t, criteria.Filters(), "Filters should be not nil")
	assert.Equalf(t, 1, len(criteria.Filters()), "Filters should be %d", 1)
}

func TestACriteriaWithNoOperableFiltersShouldSkipBadFilters(t *testing.T) {
	t.Parallel()

	offset := 0
	limit := 10
	filters := [][]string{{"name", "==", "John", "AND"}, {"name", "equals", "Jane", "AND"}}

	criteria := NewCriteria(offset, limit, "", "", filters)

	assert.NotNilf(t, criteria.Filters(), "Filters should be not nil")
	assert.Equalf(t, 1, len(criteria.Filters()), "Filters should be %d", 1)
}

func TestCriteriaWithEmptyFilterFieldsShouldReturnNil(t *testing.T) {
	t.Parallel()

	offset := 0
	limit := 10
	filters := [][]string{{"", "==", "John"}}

	criteria := NewCriteria(offset, limit, "", "", filters)

	assert.Nilf(t, criteria.Filters(), "Filters should be nil")
}

func TestCriteriaFiltersShouldBuildAndReturnOk(t *testing.T) {
	t.Parallel()

	offset := 0
	limit := 10
	filters := [][]string{{"name", "==", "John", "AND"}, {"surname", "==", "Jones", "OR"}}

	criteria := NewCriteria(offset, limit, "", "", filters)

	assert.Equalf(t, filters, criteria.Filters(), "Filters should be %v", filters)

}

func TestCriteriaShouldProvideMethodToKnowIfHasLimit(t *testing.T) {
	t.Parallel()

	offset := 10

	limit := 0
	criteria := NewCriteria(offset, limit, "", "", nil)
	assert.Falsef(t, criteria.HasLimit(), "HasLimit should be false for limit %d", limit)

	limit = 5
	criteria = NewCriteria(offset, limit, "", "", nil)
	assert.Truef(t, criteria.HasLimit(), "HasLimit should be true for limit %d", limit)
}

func TestCriteriaShouldProvideMethodToKnowIfHasOffset(t *testing.T) {
	t.Parallel()

	limit := 10

	offset := -3
	criteria := NewCriteria(offset, limit, "", "", nil)
	assert.Falsef(t, criteria.HasOffset(), "HasLimit should be false for offset %d", offset)

	offset = 6
	criteria = NewCriteria(offset, limit, "", "", nil)
	assert.Truef(t, criteria.HasOffset(), "HasLimit should be true for offset %d", offset)
}

func TestCriteriaShouldProvideMethodToKnowIfHasOrder(t *testing.T) {
	t.Parallel()

	offset := 0
	limit := -1

	criteria := NewCriteria(offset, limit, "", "", nil)
	assert.Falsef(t, criteria.HasOrder(), "HasOrder should be false for nil order")

	criteria = NewCriteria(offset, limit, "name", "", nil)
	assert.Truef(t, criteria.HasOrder(), "HasOrder should be true for order by name")
}

func TestCriteriaShouldProvideMethodToKnowIfHasFilters(t *testing.T) {
	t.Parallel()

	offset := 0
	limit := -1

	criteria := NewCriteria(offset, limit, "", "", nil)
	assert.Falsef(t, criteria.HasFilters(), "HasFilters should be false for nil filters")

	criteria = NewCriteria(offset, limit, "", "", [][]string{{"name", "==", "John", "AND"}})
	assert.Truef(t, criteria.HasFilters(), "HasFilters should be true for filters")
}
