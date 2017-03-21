package i18n

import (
	"fmt"
	"strconv"
	"strings"
)

// NumberFormatter represents a number formatter.
type NumberFormatter struct {
	PositivePattern  string
	NegativePattern  string
	DecimalDigits    int
	DecimalSeparator string
	GroupSizes       []int
	GroupSeparator   string
}

// Format formats the float value to string.
func (nf *NumberFormatter) Format(v float64) string {
	value := fmt.Sprintf("%."+strconv.Itoa(nf.DecimalDigits)+"f", v)
	parts := strings.Split(value, ".")
	integer := parts[0]
	negative := strings.HasPrefix(integer, "-")
	if negative {
		integer = integer[1:]
	}
	index := len(integer)
	for i := 0; i < len(nf.GroupSizes); i++ {
		size := nf.GroupSizes[i]
		if size > 0 && index > size {
			index -= size
			integer = integer[0:index] + nf.GroupSeparator + integer[index:]
		}

		for i == len(nf.GroupSizes)-1 && size > 0 && index > size {
			index -= size
			integer = integer[0:index] + nf.GroupSeparator + integer[index:]
		}
	}
	if len(parts) == 2 { // has decimal
		value = integer + nf.DecimalSeparator + parts[1]
	} else {
		value = integer
	}
	if negative { // has pattern
		value = strings.Replace(nf.NegativePattern, "n", value, -1)
	} else {
		value = strings.Replace(nf.PositivePattern, "n", value, -1)
	}

	return value
}

// CurrencyFormatter represents a currency formatter.
type CurrencyFormatter struct {
	Symbol           string
	PositivePattern  string
	NegativePattern  string
	DecimalDigits    int
	DecimalSeparator string
	GroupSizes       []int
	GroupSeparator   string
}

// Format formats the float value to string.
func (cf *CurrencyFormatter) Format(v float64) string {
	value := fmt.Sprintf("%."+strconv.Itoa(cf.DecimalDigits)+"f", v)
	parts := strings.Split(value, ".")
	integer := parts[0]
	negative := strings.HasPrefix(integer, "-")
	if negative {
		integer = integer[1:]
	}
	index := len(integer)
	for i := 0; i < len(cf.GroupSizes); i++ {
		size := cf.GroupSizes[i]
		if size > 0 && index > size {
			index -= size
			integer = integer[0:index] + cf.GroupSeparator + integer[index:]
		}
		for i == len(cf.GroupSizes)-1 && size > 0 && index > size {
			index -= size
			integer = integer[0:index] + cf.GroupSeparator + integer[index:]
		}
	}
	if len(parts) == 2 { // has decimal
		value = integer + cf.DecimalSeparator + parts[1]
	} else {
		value = integer
	}
	if negative { // has pattern
		value = strings.Replace(strings.Replace(cf.NegativePattern, "n", value, -1), "$", cf.Symbol, -1)
	} else {
		value = strings.Replace(strings.Replace(cf.PositivePattern, "n", value, -1), "$", cf.Symbol, -1)
	}

	return value
}

// Formatter represents a formatter for number & currency.
type Formatter struct {
	Number   *NumberFormatter
	Currency *CurrencyFormatter
}
