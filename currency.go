package trading

// Currency 表示一种货币的类型
type Currency string

func (c Currency) String() string {
	return string(c)
}

func (c Currency) Available() bool {
	return c != ""
}

////////////////////////////////////////////////////////////////////////////////

const (
	CurrencyCNY Currency = "CNY" // 人民币 (中国)
	CurrencyEUR Currency = "EUR" // 欧元 (欧盟)
	CurrencyGBP Currency = "GBP" // 英镑 (英国)
	CurrencyJPY Currency = "JPY" // 日元 (日本)
	CurrencyUSD Currency = "USD" // 美元 (美国)
)

// aliases
const (
	CurrencyRMB = CurrencyCNY
)
