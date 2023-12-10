package criteria

type orderBy struct {
	value string
}

func newOrderBy(value string) *orderBy {
	return &orderBy{
		value: value,
	}
}
