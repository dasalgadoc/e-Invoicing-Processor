package criteria

type filterLogical struct {
	value filterLogicalEnum
}

func (f *filterLogical) string() string {
	return f.value.string()
}

func newFilterLogical(value string) *filterLogical {
	logicalEnum := filterLogicalEnum(value)
	if !isValidFilterLogical(logicalEnum) {
		logicalEnum = filterLogicalAND
	}
	return &filterLogical{value: logicalEnum}
}

// filterLogicalEnum AND, OR
type filterLogicalEnum string

const (
	filterLogicalAND filterLogicalEnum = "AND"
	filterLogicalOR  filterLogicalEnum = "OR"
)

func (o filterLogicalEnum) string() string {
	return string(o)
}

func isValidFilterLogical(value filterLogicalEnum) bool {
	switch value {
	case filterLogicalAND, filterLogicalOR:
		return true
	}
	return false
}
