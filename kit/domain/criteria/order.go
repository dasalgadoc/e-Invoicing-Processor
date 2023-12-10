package criteria

type order struct {
	orderBy   *orderBy
	orderType *orderType
}

func newOrder(by string, orderType string) *order {
	if by == "" {
		return nil
	}

	return &order{
		orderBy:   newOrderBy(by),
		orderType: newOrderType(orderType),
	}
}
