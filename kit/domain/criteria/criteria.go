package criteria

type Criteria struct {
	offset  int
	limit   int
	order   *order
	filters *filters
}

func (c *Criteria) Offset() int {
	return c.offset
}

func (c *Criteria) HasOffset() bool {
	return c.offset > 0
}

func (c *Criteria) Limit() int {
	return c.limit
}

func (c *Criteria) HasLimit() bool {
	return c.limit > 0
}

func (c *Criteria) OrderBy() string {
	return c.order.orderBy.value
}

func (c *Criteria) OrderType() string {
	return c.order.orderType.string()
}

func (c *Criteria) HasOrder() bool {
	return c.order != nil
}

func (c *Criteria) Filters() [][]string {
	if c.filters == nil {
		return nil
	}
	var f [][]string
	for _, filter := range *c.filters {
		f = append(f, []string{filter.field.value, filter.operator.value.string(), filter.value.value, filter.logical.value.string()})
	}
	return f
}

func (c *Criteria) HasFilters() bool {
	return c.filters != nil && !c.filters.isEmpty()
}

func NewCriteria(offset int, limit int, by string, orderType string, filters [][]string) *Criteria {
	if offset < 0 {
		offset = 0
	}
	if limit < 0 {
		limit = 0
	}
	return &Criteria{
		offset:  offset,
		limit:   limit,
		order:   newOrder(by, orderType),
		filters: newFilters(filters),
	}
}
