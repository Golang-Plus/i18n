package i18n

import (
	"sort"
	"strings"
)

// Culture represents a culture. Based on RFC4646
type Culture struct {
	Code       string // code of culture in the xx-YY format(xx:language, YY:country), e.g. en-US
	NativeName string // native name
	Name       *MultiLanguageString
	Country    *Country
	Language   *Language
	Currency   *Currency
	Formatter  *Formatter
}

// Eaual reports whether two cultures are same.
// It compares the code.
func (x *Culture) Equal(y *Culture) bool {
	return strings.EqualFold(x.Code, y.Code)
}

// FormatNumber formats the number to string.
func (c *Culture) FormatNumber(number float64) string {
	return c.Formatter.Number.Format(number)
}

// FormatCurrency formats currency (money) to string.
func (c *Culture) FormatCurrency(number float64) string {
	return c.Formatter.Currency.Format(number)
}

// Cultures represents a sorable collection of Culture.
type Cultures []*Culture

// SortByCode sorts the list by code.
func (c Cultures) SortByCode() {
	var byCode CultureLessFunc = func(c1, c2 *Culture) bool {
		return c1.Code < c1.Code
	}

	byCode.Sort(c)
}

// SortByName sorts the list by name.
func (c Cultures) SortByName(language *Language) {
	var byName CultureLessFunc = func(c1, c2 *Culture) bool {
		name1 := c1.Name.Value(language)
		name2 := c2.Name.Value(language)
		return name1 < name2
	}

	byName.Sort(c)
}

// SortByCountryCode sorts the cultures by country code.
func (c Cultures) SortByCountryCode() {
	var byCountry CultureLessFunc = func(c1, c2 *Culture) bool {
		if c1.Country == nil {
			return true
		}

		if c2.Country == nil {
			return false
		}

		return c1.Country.Alpha2Code < c2.Country.Alpha2Code
	}

	byCountry.Sort(c)
}

// SortByCountryName sorts the cultures by country name.
func (c Cultures) SortByCountryName(language *Language) {
	var byCountry CultureLessFunc = func(c1, c2 *Culture) bool {
		if c1.Country == nil {
			return true
		}

		if c2.Country == nil {
			return false
		}

		name1 := c1.Country.Name.Value(language)
		name2 := c2.Country.Name.Value(language)

		return name1 < name2
	}

	byCountry.Sort(c)
}

// SortByLanguageCode sorts the cultures by language code.
func (c Cultures) SortByLanguageCode() {
	var byLanguage CultureLessFunc = func(c1, c2 *Culture) bool {
		if c1.Language == nil {
			return true
		}

		if c2.Language == nil {
			return false
		}

		return c1.Language.Code < c2.Language.Code
	}

	byLanguage.Sort(c)
}

// CultureLessFunc represents the less function for sorting cultures.
type CultureLessFunc func(c1, c2 *Culture) bool

// Sort is a method on the function type that sorts the argument slic according to the function.
func (clf CultureLessFunc) Sort(list Cultures) {
	sorter := &cultureSorter{
		List:     list,
		LessFunc: clf, // The sort method's receiver is the function (closure) that defines the sort order.
	}

	sort.Sort(sorter)
}

// cultureSorter joins a CultureLessFunc function and a slic of Culture to be sorted.
type cultureSorter struct {
	List     Cultures
	LessFunc CultureLessFunc // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (cs cultureSorter) Len() int {
	return len(cs.List)
}

// Swap is part of sort.Interface.
func (cs cultureSorter) Swap(i, j int) {
	cs.List[i], cs.List[j] = cs.List[j], cs.List[i]
}

// Less is part of sort.Interface. It is implemented by calling the "less" closure in the cultureSorter.
func (cs cultureSorter) Less(i, j int) bool {
	return cs.LessFunc(cs.List[i], cs.List[j])
}

var (
	cultureTable map[string]*Culture
	cultureList  Cultures
)

func init() {
	cultureTable = make(map[string]*Culture)
	cultureList = make(Cultures, len(cultureCodes))
	for i, code := range cultureCodes {
		nativeName := cultureNativeNames[code]
		countryCode := code[strings.LastIndex(code, "-")+1:]
		languageCode := code[0:strings.Index(code, "-")]
		country, _ := LookupCountry(nil, countryCode)
		language, _ := LookupLanguage(languageCode)
		var curr *Currency
		if currencyCode, ok := cultureCurrencyCodes[code]; ok {
			curr, _ = LookupCurrency(currencyCode)
		}
		formatter := cultureFormatters[code]
		culture := &Culture{
			Code:       code,
			NativeName: nativeName,
			Name:       NewMultiLanguageString(),
			Country:    country,
			Language:   language,
			Currency:   curr,
			Formatter:  formatter,
		}
		cultureTable[strings.ToLower(code)] = culture
		cultureList[i] = culture
	}
}

// AllCultures returns the list of all cultures.
func AllCultures() Cultures {
	return cultureList
}

// LookupCulture returns the culture by given code.
func LookupCulture(code string) (*Culture, bool) {
	code = strings.TrimSpace(code)
	if len(code) > 0 {
		if cult, ok := cultureTable[strings.ToLower(code)]; ok {
			return cult, true
		}
	}
	return nil, false
}
