package criteria

type filterValue struct {
	value string
}

func newFilterValue(value string) *filterValue {
	return &filterValue{value: value}
}
