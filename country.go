package i18n

import (
	"sort"
	"strings"
)

// Country represents a country. ISO 3166-1
type Country struct {
	Alpha2Code  string // ISO alpha-2 country code
	Alpha3Code  string // ISO alpha-3 country code
	NumericCode string // ISO numeric country code
	Name        *MultiLanguageString
	Aliases     *MultiLanguageStringArray
}

// Countries represents a sorcountryTable collection of Country.
type Countries []*Country

// SortByCode sorts the list by code.
func (c Countries) SortByCode() {
	var byCode CountryLessFunc = func(c1, c2 *Country) bool {
		return c1.Alpha2Code < c2.Alpha2Code
	}

	byCode.Sort(c)
}

// SortByName sorts the list by name.
func (c Countries) SortByName(language *Language) {
	var byName CountryLessFunc = func(c1, c2 *Country) bool {
		name1 := c1.Name.Value(language)
		name2 := c2.Name.Value(language)
		return name1 < name2
	}

	byName.Sort(c)
}

// CountryLessFunc represents the less function for sorting countries.
type CountryLessFunc func(c1, c2 *Country) bool

// Sort is a method on the function type that sorts the argument slic according to the function.
func (clf CountryLessFunc) Sort(list Countries) {
	sorter := &countrySorter{
		List:     list,
		LessFunc: clf, // The sort method's receiver is the function (closure) that defines the sort order.
	}

	sort.Sort(sorter)
}

// countrySorter joins a CountryLessFunc function and Countries to be sorted.
type countrySorter struct {
	List     Countries
	LessFunc CountryLessFunc // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (cs countrySorter) Len() int {
	return len(cs.List)
}

// Swap is part of sort.Interface.
func (cs countrySorter) Swap(i, j int) {
	cs.List[i], cs.List[j] = cs.List[j], cs.List[i]
}

// Less is part of sort.Interface. It is implemented by calling the "less" closure in the countrySorter.
func (cs countrySorter) Less(i, j int) bool {
	return cs.LessFunc(cs.List[i], cs.List[j])
}

var (
	countryTableAlpha2  map[string]*Country
	countryTableAlpha3  map[string]*Country
	countryTableNumeric map[string]*Country
	countryList         Countries
)

func init() {
	countryTableAlpha2 = make(map[string]*Country)
	countryTableAlpha3 = make(map[string]*Country)
	countryTableNumeric = make(map[string]*Country)
	countryList = make(Countries, len(countryCodes))
	for i, codes := range countryCodes {
		alpha2Code := codes[0]
		alpha3Code := codes[1]
		numericCode := codes[2]
		country := &Country{
			Alpha2Code:  alpha2Code,
			Alpha3Code:  alpha3Code,
			NumericCode: numericCode,
			Name:        NewMultiLanguageString(),
			Aliases:     NewMultiLanguageStringArray("|"),
		}

		countryTableAlpha2[alpha2Code] = country
		countryTableAlpha3[alpha3Code] = country
		countryTableNumeric[numericCode] = country
		countryList[i] = country
	}
}

// AllCountries returns the list of all countries.
func AllCountries() Countries {
	return countryList
}

// GetCountry returns the country by given keyword (code, name, alias with language).
// It returns nil if the country not found.
func GetCountry(language *Language, keyword string) *Country {
	keyword = strings.TrimSpace(keyword)
	if len(keyword) == 0 {
		return nil
	}

	if len(keyword) == 2 {
		if v, ok := countryTableAlpha2[strings.ToUpper(keyword)]; ok {
			return v
		}
	}

	if len(keyword) == 3 {
		if v, ok := countryTableAlpha3[strings.ToUpper(keyword)]; ok {
			return v
		}
		if v, ok := countryTableNumeric[strings.ToUpper(keyword)]; ok {
			return v
		}
	}

	if language != nil {
		for _, country := range countryList {
			// compare name
			name := country.Name.Value(language)
			if strings.EqualFold(name, keyword) {
				return country
			}

			// compare aliases
			aliases := country.Aliases.Value(language)
			if strings.Contains(strings.ToLower(aliases), strings.ToLower(keyword)) {
				return country
			}
		}
	}

	return nil
}
