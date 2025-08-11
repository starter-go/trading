package trading

import (
	"bytes"
	"strconv"
)

////////////////////////////////////////////////////////////////////////////////

// Price 是表示价格的基本结构
type Price struct {
	Currency Currency // 计价货币
	Number   Fraction // 价值
}

func (inst *Price) String() string {
	return inst.Text().String()
}

func (inst *Price) Text() PriceText {

	if !inst.Available() {
		return "N/A"
	}

	// format: 'CC:xxxx.xx'
	// example: 'CNY:1234.56'
	s1 := inst.Number.FormatAsDecimal(2)
	s2 := inst.Currency.String()
	str := s2 + ":" + s1
	return PriceText(str)
}

func (inst *Price) Available() bool {
	b1 := inst.Currency.Available()
	b2 := inst.Number.Available()
	return b1 && b2
}

////////////////////////////////////////////////////////////////////////////////

// PriceText 是 Price 的文本形式, 它们之间可以互相转换
type PriceText string

func (txt PriceText) String() string {
	return string(txt)
}

func (txt PriceText) Price() Price {
	var dst Price

	str := txt.String()
	bin := []byte(str)

	idx := bytes.IndexByte(bin, ':')
	if idx < 0 {
		return dst
	}
	str1 := string(bin[0:idx])
	str2 := string(bin[idx+1:])

	num, err := strconv.ParseFloat(str2, 64)
	if err != nil {
		dst.Number.Denominator = 0
		return dst
	}

	dst.Currency = Currency(str1)
	dst.Number = NewFractionWithFloat(num, 0)
	return dst
}

////////////////////////////////////////////////////////////////////////////////

// UnitPrice 表示货物的单价
type UnitPrice struct {
	Price

	Unit PriceUnit // 计价单位

}

func (inst *UnitPrice) String() string {
	return inst.Text().String()
}

func (inst *UnitPrice) Text() UnitPriceText {

	if !inst.Available() {
		return "N/A"
	}

	s1 := inst.Price.String()
	s2 := inst.Unit.String()
	str := s1 + s2
	return UnitPriceText(str)
}

func (inst *UnitPrice) Available() bool {
	b1 := inst.Price.Available()
	b2 := inst.Unit.Available()
	return b1 && b2
}

////////////////////////////////////////////////////////////////////////////////

// UnitPriceText 是 UnitPrice 的文本形式, 它们之间可以互相转换
type UnitPriceText string

func (txt UnitPriceText) String() string {
	return string(txt)
}

func (txt UnitPriceText) UnitPrice() UnitPrice {
	var dst UnitPrice

	str := txt.String()
	bin := []byte(str)

	i1 := bytes.IndexByte(bin, ':')
	i2 := bytes.IndexByte(bin, '/')

	if 0 < i1 && i1 < i2 {
		str1 := string(bin[0:i1])
		str2 := string(bin[i1+1 : i2])
		str3 := string(bin[i2:])
		num, err := strconv.ParseFloat(str2, 64)
		if err == nil {
			dst.Currency = Currency(str1)
			dst.Number = NewFractionWithFloat(num, 0)
			dst.Unit = PriceUnit(str3)
		}
	} else {
		dst.Number.Denominator = 0
	}

	return dst
}

////////////////////////////////////////////////////////////////////////////////
