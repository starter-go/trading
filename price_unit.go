package trading

import "strings"

// PriceUnit 是一个字符串, 表示计价单位
type PriceUnit string

func (pu PriceUnit) String() string {
	return string(pu)
}

func (pu PriceUnit) Available() bool {
	str := string(pu)
	return (strings.HasPrefix(str, "/")) && (len(str) > 1)
}

////////////////////////////////////////////////////////////////////////////////

const (
	PerSecond PriceUnit = "/second" // 每秒价格
	PerMinute PriceUnit = "/minute" // 每分价格
	PerHour   PriceUnit = "/hour"   // 每时价格
	PerDay    PriceUnit = "/day"    // 每天价格
	PerWeek   PriceUnit = "/week"   // 每周价格
	PerMonth  PriceUnit = "/month"  // 每月价格
	PerYear   PriceUnit = "/year"   // 每年价格

	PerG    PriceUnit = "/g"    // 每克
	PerKg   PriceUnit = "/kg"   // 每千克
	PerTon  PriceUnit = "/ton"  // 每吨 (1000kg)
	Per500G PriceUnit = "/500g" // 每500克 (aka.每市斤)
)
