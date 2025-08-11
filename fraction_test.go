package trading

import "testing"

func TestNewFraction(t *testing.T) {

	f1 := NewFraction(147, 258)
	f2 := NewFractionWithFloat(147.369, 20)
	f3 := Fraction{108, 0}

	t.Logf("f1 = %v = %v", f1.String(), f1.Float())
	t.Logf("f2 = %v = %v", f2.String(), f2.Double())

	t.Logf("f2.Available() = %v", f2.Available())
	t.Logf("f2.FormatAsDecimal(5) = %v", f2.FormatAsDecimal(5))

	t.Logf("f3.Available() = %v", f3.Available())
	t.Logf("f3.FormatAsDecimal(5) = %v", f3.FormatAsDecimal(5))
	t.Logf("f3.String() = %v", f3.String())

}

func TestFractionAdd1(t *testing.T) {

	a := NewFraction(147, 258)
	b := NewFraction(369, 10000)
	c := a.Add(&b)

	t.Logf("a = %v", a.Float())
	t.Logf("b = %v", b.Float())
	t.Logf("c = a+b = %v", c.Float())

}

func TestFractionAdd2(t *testing.T) {

	a := NewFraction(147, 10000)
	b := NewFraction(369, 258)
	c := a.Add(&b)

	t.Logf("a = %v", a.Float())
	t.Logf("b = %v", b.Float())
	t.Logf("c = a+b = %v", c.Float())

}

func TestFractionMultiplyWithFloat(t *testing.T) {

	a := NewFractionWithFloat(1.4700, 10000)
	b := -2.5
	c := a.MultiplyWithFloat(b)

	t.Logf("a = %v", a.Float())
	t.Logf("b = %v", b)
	t.Logf("c = a*b = %v", c.Float())

}

func TestFractionMultiplyWithInt(t *testing.T) {

	a := NewFraction(147, 159)
	b := 10
	c := a.MultiplyWithInt(b)

	t.Logf("a = %v", a.Float())
	t.Logf("b = %v", b)
	t.Logf("c = a*b = %v", c.Float())

}
