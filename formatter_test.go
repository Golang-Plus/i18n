package i18n

import (
	"testing"

	testing2 "github.com/golang-plus/testing"
)

func TestFormatter(t *testing.T) {
	nf := &CurrencyFormatter{
		PositivePattern:  "$n",
		NegativePattern:  "-$n",
		DecimalDigits:    2,
		DecimalSeparator: ".",
		GroupSizes:       []int{3, 2, 0},
		GroupSeparator:   ",",
	}

	cf := &CurrencyFormatter{
		Symbol:           "$",
		PositivePattern:  "$n",
		NegativePattern:  "-$n",
		DecimalDigits:    2,
		DecimalSeparator: ".",
		GroupSizes:       []int{3, 2, 0},
		GroupSeparator:   ",",
	}

	str := cf.Format(1234567890.125)
	testing2.AssertEqual(t, str, "$12345,67,890.12")
	str = cf.Format(-1234567890.126)
	testing2.AssertEqual(t, str, "-$12345,67,890.13")

	str = nf.Format(1234567890.125)
	testing2.AssertEqual(t, str, "12345,67,890.12")
	str = nf.Format(-1234567890.126)
	testing2.AssertEqual(t, str, "-12345,67,890.13")
}
