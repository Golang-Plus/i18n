package i18n

import (
	"testing"

	testing2 "github.com/golang-plus/testing"
)

func TestMoney(t *testing.T) {
	// test Exchange
	cny := MustNewExchangeableCurrency("CNY", 1)
	hkd := MustNewExchangeableCurrency("hkd", 1.13)
	usd := MustNewExchangeableCurrency("usd", 0.15)
	eur := MustNewExchangeableCurrency("eur", 0.14)
	data := map[*ExchangeableCurrency]map[float64]map[*ExchangeableCurrency]string{
		cny: {
			5.2: {
				hkd: "HKD 5.8760",
				usd: "USD 0.7800",
				eur: "EUR 0.7280",
			},
			132.25: {
				hkd: "HKD 149.4425",
				usd: "USD 19.8375",
				eur: "EUR 18.5150",
			},
		},
		hkd: {
			5.2: {
				hkd: "HKD 5.2000",
				cny: "CNY 4.6018",
				usd: "USD 0.6903",
				eur: "EUR 0.6442",
			},
		},
	}
	for k, v := range data {
		for k2, v2 := range v {
			for k3, v3 := range v2 {
				testing2.AssertEqual(t, MustNewMoney(k, k2).Exchange(k3).String(), v3)
			}
		}
	}

	// test Format
	cf := &CurrencyFormatter{
		Symbol:           "$",
		PositivePattern:  "$n",
		NegativePattern:  "-$n",
		DecimalDigits:    4,
		DecimalSeparator: ".",
		GroupSizes:       []int{3, 3, 3},
		GroupSeparator:   ",",
	}
	data2 := map[float64]string{
		0.8250:       "$0.8250",
		-0.8250:      "-$0.8250",
		22.22:        "$22.2200",
		-22.22:       "-$22.2200",
		2222.2222:    "$2,222.2222",
		-2222.2222:   "-$2,222.2222",
		2222.222251:  "$2,222.2223",
		-2222.222251: "-$2,222.2223",
	}
	for k, v := range data2 {
		m := MustNewMoney(usd, k)
		testing2.AssertEqual(t, m.Format(cf), v)
	}

	// test Add/Sub/Mul/Div
	testing2.AssertEqual(t, MustNewMoney(cny, 100).Add(MustNewMoney(usd, 100)).String(), "CNY 766.6667")
	testing2.AssertEqual(t, MustNewMoney(cny, 100).Sub(MustNewMoney(usd, 100)).String(), "CNY -566.6667")
	testing2.AssertEqual(t, MustNewMoney(cny, 102.54321).Mul(2.5).String(), "CNY 256.3580")
	testing2.AssertEqual(t, MustNewMoney(cny, 102.54321).Div(2.5).String(), "CNY 41.0173")
}
