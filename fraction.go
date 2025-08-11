package trading

import "strconv"

// Fraction 是一个表示分数的结构, 由分子和分母两个整数构成
type Fraction struct {
	Numerator   int // 分子
	Denominator int // 分母
}

// NewFraction 新建并初始化一个 Fraction 结构.
//
// n: 一个整数,表示分子;
//
// d: 一个整数,表示分母;
func NewFraction(n, d int) Fraction {
	if d == 0 {
		d = 1 // 分母的默认值为 1
	}
	var res Fraction
	res.Numerator = n
	res.Denominator = d
	return res
}

// NewFraction
//
//	新建一个 Fraction 结构, 并使用传入的浮点数和精度来初始化这个 Fraction
//
// value: 一个双精度浮点数, 表示 Fraction 的值;
//
// precision: 一个整数, 表示精确到小数点后的位数;
func NewFractionWithFloat(value float64, precision int) Fraction {
	if precision <= 0 {
		precision = 5 // 默认的精度为5位小数
	}
	x := 1
	for i := 0; i < precision; i++ {
		if i > 8 {
			//  提供最大 10^8 的精度
			break
		}
		x *= 10
	}
	var res Fraction
	res.Numerator = int(float64(x) * value)
	res.Denominator = x
	return res
}

// 返回这个分数的字符串形式
func (inst *Fraction) String() string {
	n := inst.Numerator
	d := inst.Denominator
	if d == 0 {
		return "N/A"
	}
	return strconv.Itoa(n) + "/" + strconv.Itoa(d)
}

// Float 返回单精度的数值
func (inst *Fraction) Float() float32 {
	n := inst.Numerator
	d := inst.Denominator
	if d == 0 {
		return 0
	}
	f1 := float32(n)
	f2 := float32(d)
	return f1 / f2
}

// Double 返回双精度的数值
func (inst *Fraction) Double() float64 {
	n := inst.Numerator
	d := inst.Denominator
	if d == 0 {
		return 0
	}
	f1 := float64(n)
	f2 := float64(d)
	return f1 / f2
}

// Available 判断这个分数是否有效
func (inst *Fraction) Available() bool {
	return inst.Denominator != 0
}

// FormatAsDecimal 把分数格式化为小数形式.
// precision 表示小数的精度, 即小数点后的位数; 如果 precision<=0, 则只输出整数部分.
func (inst *Fraction) FormatAsDecimal(precision int) string {
	f := inst.Double()
	return strconv.FormatFloat(f, 'f', precision, 64)
}

// Add 把这个 Fraction 和另一个 Fraction 相加, 返回结果是一个新的 Fraction
func (inst *Fraction) Add(other *Fraction) Fraction {

	var res Fraction
	if inst == nil || other == nil {
		return res
	}

	n1 := inst.Numerator
	n2 := other.Numerator
	d1 := inst.Denominator
	d2 := other.Denominator
	if d1 == 0 || d2 == 0 {
		return res
	}

	if d1 == d2 {
		// NOP
	} else if d1 > d2 {
		// use: d1
		x := float32(d1) / float32(d2)
		n2 = int(x * float32(n2))
		d2 = d1
	} else {
		// use: d2
		x := float32(d2) / float32(d1)
		n1 = int(x * float32(n1))
		d1 = d2
	}

	res.Numerator = n1 + n2
	res.Denominator = d1
	return res
}

// MultiplyWithInt  把这个 Fraction  乘以一个整数, 返回结果是一个新的 Fraction
func (inst *Fraction) MultiplyWithInt(n int) Fraction {
	var res Fraction = *inst
	res.Numerator *= n
	return res
}

// MultiplyWithFloat  把这个 Fraction  乘以一个浮点数, 返回结果是一个新的 Fraction
func (inst *Fraction) MultiplyWithFloat(n float64) Fraction {
	var res Fraction = *inst
	n2 := float64(res.Numerator) * n
	res.Numerator = int(n2)
	return res
}
