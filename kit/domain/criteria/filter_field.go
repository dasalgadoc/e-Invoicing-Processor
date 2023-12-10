package criteria

type filterField struct {
	value string
}

func newFilterField(value string) *filterField {
	if value == "" {
		return nil
	}

	return &filterField{value: value}
}
