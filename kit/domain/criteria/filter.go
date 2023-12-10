package criteria

type filters []filter

func (f filters) isEmpty() bool {
	return len(f) == 0
}

func newFilters(f [][]string) *filters {
	var newFilters filters
	for _, filter := range f {
		if len(filter) < 4 {
			continue
		}

		newFilter := newFilter(filter[0], filter[1], filter[2], filter[3])
		if newFilter == nil {
			continue
		}

		newFilters = append(newFilters, *newFilter)
	}

	if newFilters.isEmpty() {
		return nil
	}

	return &newFilters
}

type filter struct {
	field    *filterField
	value    *filterValue
	operator *filterOperator
	logical  *filterLogical
}

func newFilter(field, operator, value, logical string) *filter {
	f := newFilterField(field)
	v := newFilterValue(value)
	o := newFilterOperator(operator)
	l := newFilterLogical(logical)
	if f == nil || v == nil || o == nil || l == nil {
		return nil
	}

	return &filter{
		field:    f,
		value:    v,
		operator: o,
		logical:  l,
	}
}
