package i18n

import (
	"sort"
	"strings"
)

// Language represents a language. ISO 639-1.
type Language struct {
	Code       string // ISO code of the language
	NativeName string // native name of language
	Name       *MultiLanguageString
}

// Equal reports whether two languages are same.
// It compares the code.
func (x *Language) Equal(y *Language) bool {
	return strings.EqualFold(x.Code, y.Code)
}

// Languages represents a sortable collection of Language.
type Languages []*Language

// SortByCode sorts the list by code.
func (l Languages) SortByCode() {
	var byCode LanguageLessFunc = func(c1, c2 *Language) bool {
		return c1.Code < c2.Code
	}

	byCode.Sort(l)
}

// SortByName sorts the list by name with given language.
func (l Languages) SortByName(language *Language) {
	var byName LanguageLessFunc = func(c1, c2 *Language) bool {
		return c1.Name.Value(language) < c2.Name.Value(language)
	}

	byName.Sort(l)
}

// languageSorter joins a LanguageLessFunc function and Languages to be sorted.
type languageSorter struct {
	List     Languages
	LessFunc LanguageLessFunc // Closure used in the Less method.
}

// Len is part of sort.Interface.
func (ls languageSorter) Len() int {
	return len(ls.List)
}

// Swap is part of sort.Interface.
func (ls languageSorter) Swap(i, j int) {
	ls.List[i], ls.List[j] = ls.List[j], ls.List[i]
}

// Less is part of sort.Interface. It is implemented by calling the "less" closure in the languageSorter.
func (ls languageSorter) Less(i, j int) bool {
	return ls.LessFunc(ls.List[i], ls.List[j])
}

// LanguageLessFunc represents the less function for sorting languages.
type LanguageLessFunc func(l1, l2 *Language) bool

// Sort is a method on the function type that sorts the argument slic according to the function.
func (llf LanguageLessFunc) Sort(list Languages) {
	sorter := &languageSorter{
		List:     list,
		LessFunc: llf, // The sort method's receiver is the function (closure) that defines the sort order.
	}

	sort.Sort(sorter)
}

var (
	languageTable map[string]*Language
	languageList  Languages
)

func init() {
	languageTable = make(map[string]*Language)
	languageList = make(Languages, len(languageCodes))
	for i, v := range languageCodes {
		nativeName := languageNativeNames[v]
		language := &Language{
			Code:       v,
			NativeName: nativeName,
			Name:       NewMultiLanguageString(),
		}

		languageTable[strings.ToLower(v)] = language
		languageList[i] = language
	}
}

// AllLanguages returns the list of all languages.
func AllLanguages() Languages {
	return languageList
}

// LookupLanguage returns the language by given code.
func LookupLanguage(code string) (*Language, bool) {
	code = strings.TrimSpace(code)
	if len(code) > 0 {
		if lang, ok := languageTable[strings.ToLower(code)]; ok {
			return lang, true
		}
	}
	return nil, false
}
