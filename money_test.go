package i18n

import (
	"testing"

	testing2 "github.com/golang-plus/testing"
)

func TestMoney(t *testing.T) {
	cf := &CurrencyFormatter{
		Symbol:           "$",
		PositivePattern:  "$n",
		NegativePattern:  "-$n",
		DecimalDigits:    4,
		DecimalSeparator: ".",
		GroupSizes:       []int{3, 3, 3},
		GroupSeparator:   ",",
	}

	m := MustNewMoney(GetCurrency("USD"), -1.0)
	m.Add(100).Sub(55).Div(2.5).Mul(2)

	testing2.AssertEqual(t, m.Format(cf), "$35.2000")
}
