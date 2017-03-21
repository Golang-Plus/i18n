package i18n

import (
	"sort"
	"strings"
)

// Currency represents a currency. ISO 4217
type Currency struct {
	Code string // ISO Alpha-3 Currency Code
	Name *MultiLanguageString
}

// Currencies represents a sorcurrencyTable collection of Currency.
type Currencies []*Currency

// SortByCode sorts the list by code.
func (c Currencies) SortByCode() {
	var byCode CurrencyLessFunc = func(c1, c2 *Currency) bool {
		return c1.Code < c2.Code
	}

	byCode.Sort(c)
}

// SortByName sorts the list by name.
func (c Currencies) SortByName(language *Language) {
	var byName CurrencyLessFunc = func(c1, c2 *Currency) bool {
		name1 := c1.Name.Value(language)
		name2 := c2.Name.Value(language)

		return name1 < name2
	}

	byName.Sort(c)
}

// currencySorter joins a CurrencyLessFunc function and Currencies to be sorted.
type currencySorter struct {
	List     Currencies
	LessFunc CurrencyLessFunc // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (cs currencySorter) Len() int {
	return len(cs.List)
}

// Swap is part of sort.Interface.
func (cs currencySorter) Swap(i, j int) {
	cs.List[i], cs.List[j] = cs.List[j], cs.List[i]
}

// Less is part of sort.Interface. It is implemented by calling the "less" closure in the currencySorter.
func (cs currencySorter) Less(i, j int) bool {
	return cs.LessFunc(cs.List[i], cs.List[j])
}

// CurrencyLessFunc represents the less function for sorting currencies.
type CurrencyLessFunc func(c1, c2 *Currency) bool

// Sort is a method on the function type that sorts the argument slic according to the function.
func (clf CurrencyLessFunc) Sort(list Currencies) {
	sorter := &currencySorter{
		List:     list,
		LessFunc: clf, // The sort method's receiver is the function (closure) that defines the sort order.
	}

	sort.Sort(sorter)
}

var (
	currencyTable map[string]*Currency
	currencyList  Currencies
)

func init() {
	currencyTable = make(map[string]*Currency)
	currencyList = make(Currencies, len(currencyCodes))
	for i, v := range currencyCodes {
		currency := &Currency{
			Code: v,
			Name: NewMultiLanguageString(),
		}

		currencyTable[v] = currency
		currencyList[i] = currency
	}
}

// AllCurrencies returns the list of all currencies.
func AllCurrencies() Currencies {
	return currencyList
}

// GetCurrency returns the currency by given code.
// It returns nil if the currency not found.
func GetCurrency(code string) *Currency {
	code = strings.TrimSpace(code)
	if len(code) > 0 {
		if curr, ok := currencyTable[strings.ToUpper(code)]; ok {
			return curr
		}
	}

	return nil
}
