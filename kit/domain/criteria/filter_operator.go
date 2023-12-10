package criteria

type filterOperator struct {
	value operator
}

func (f *filterOperator) string() string {
	return f.value.string()
}

func newFilterOperator(value string) *filterOperator {
	op := operator(value)
	if !isValidOperator(op) {
		return nil
	}

	return &filterOperator{value: op}
}

// operator enum
type operator string

const (
	equalOperator        operator = "=="
	notEqualOperator     operator = "!="
	lessThanOperator     operator = "<"
	lessEqualOperator    operator = "<="
	greaterThanOperator  operator = ">"
	greaterEqualOperator operator = ">="
	containingOperator   operator = "CONTAINS"
)

func (o operator) string() string {
	return string(o)
}

func isValidOperator(operator operator) bool {
	switch operator {
	case equalOperator, notEqualOperator, lessThanOperator, lessEqualOperator, greaterThanOperator, greaterEqualOperator, containingOperator:
		return true
	}
	return false
}
