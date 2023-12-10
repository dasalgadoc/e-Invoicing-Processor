package criteria

import "strings"

type orderType struct {
	value orderTypeEnum
}

func (o *orderType) string() string {
	return o.value.string()
}

func newOrderType(value string) *orderType {
	valueToLower := strings.ToUpper(value)
	oTypeEnum := orderTypeEnum(valueToLower)
	if !isValidOrderType(oTypeEnum) {
		oTypeEnum = orderTypeASC
	}
	return &orderType{value: oTypeEnum}
}

// orderTypeEnum ASC, DESC
type orderTypeEnum string

const (
	orderTypeASC  orderTypeEnum = "ASC"
	orderTypeDESC orderTypeEnum = "DESC"
)

func (o orderTypeEnum) string() string {
	return string(o)
}

func isValidOrderType(value orderTypeEnum) bool {
	switch value {
	case orderTypeASC, orderTypeDESC:
		return true
	}
	return false
}
