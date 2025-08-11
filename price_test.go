package trading

import "testing"

func TestPrice(t *testing.T) {

	var p Price
	p.Currency = CurrencyRMB
	p.Number.Numerator = 1234567890
	p.Number.Denominator = 1000

	txt := p.Text()
	str := txt.String()
	t.Logf("price = %v", str)

}

func TestUnitPrice(t *testing.T) {

	var up UnitPrice
	up.Currency = CurrencyUSD
	up.Number.Numerator = 1234567890
	up.Number.Denominator = 100000
	up.Unit = PerKg

	txt := up.Text()
	str := txt.String()
	t.Logf("price = %v", str)

}

func TestParsePrice(t *testing.T) {

	p1 := &Price{}
	p1.Currency = CurrencyEUR
	p1.Number = NewFractionWithFloat(3.14, 0)

	text1 := p1.Text()
	p2 := text1.Price()
	text2 := p2.Text()

	t.Logf("p1 = %v", text1)
	t.Logf("p2 = %v", text2)

	if text1 != text2 {
		t.Error("text1 != text2, text1=" + text1 + ", text2=" + text2)
	}

}

func TestParseUnitPrice(t *testing.T) {

	p1 := &UnitPrice{}
	p1.Currency = CurrencyEUR
	p1.Number = NewFractionWithFloat(1024.567, 10000)
	p1.Unit = PerTon

	text1 := p1.Text()
	p2 := text1.UnitPrice()
	text2 := p2.Text()

	t.Logf("p1 = %v", text1)
	t.Logf("p2 = %v", text2)

	if text1 != text2 {
		t.Error("UnitPrice: text1 != text2, text1=" + text1 + ", text2=" + text2)
	}

}
