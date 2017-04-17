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

	// calculation:
	data := map[string][4]float64{
		"-$0.8250":    {1.1, 2.2, 3.3, 4.4},
		"$2.2000":     {4.4, 3.3, 2.2, 1.1},
		"$22.2200":    {44.44, 33.33, 22.22, 11.11},
		"$2,222.2222": {4444.4444, 3333.3333, 2222.2222, 1111.1111},
		"-$833.3333":  {1111.1111, 2222.2222, 3333.3333, 4444.4444},
	}
	for k, v := range data {
		m := MustNewMoney(GetCurrency("usd"), 0)
		// + - * /
		m.Add(v[0]).Sub(v[1]).Mul(v[2]).Div(v[3])
		testing2.AssertEqual(t, m.Format(cf), k)
	}

	// rounding:

	// to nearest even
	data2 := map[float64]float64{
		2.34564:     2.3456,
		2.34566:     2.3457,
		2.34605001:  2.3461,
		2.34635:     2.3464,
		2.34645:     2.3464,
		-2.34564:    -2.3456,
		-2.34566:    -2.3457,
		-2.34605001: -2.3461,
		-2.34635:    -2.3464,
		-2.34645:    -2.3464,
		2.34555:     2.3456,
		2.34525:     2.3452,
	}
	for k, v := range data2 {
		m := MustNewMoney(GetCurrency("usd"), k)
		testing2.AssertEqual(t, m.RoundToNearestEven(4).Amount, v)
	}

}
