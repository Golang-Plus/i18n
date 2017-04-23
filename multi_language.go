package i18n

import (
	"strings"
)

// MultiLanguageString represents a string supports multi-language.
type MultiLanguageString struct {
	Values map[string]string
}

// SupportedLanguages returns the supported languages.
func (mlf *MultiLanguageString) SupportedLanguages() Languages {
	if len(mlf.Values) == 0 {
		return nil
	}

	langs := make(Languages, len(mlf.Values))
	index := 0
	for langCode, _ := range mlf.Values {
		if lang, ok := LookupLanguage(langCode); ok {
			langs[index] = lang
			index += 1
		}
	}

	return langs
}

// IsEmpty reports whether the values is empty.
func (mlf *MultiLanguageString) IsEmpty() bool {
	return len(mlf.Values) == 0
}

// Value returns the value of string with given language.
func (mlf *MultiLanguageString) Value(language *Language) string {
	for lang, val := range mlf.Values {
		if strings.EqualFold(lang, language.Code) {
			return val
		}
	}

	return ""
}

// SetValue sets the value with language.
func (mlf *MultiLanguageString) SetValue(language *Language, value string) {
	if mlf.Values == nil {
		mlf.Values = make(map[string]string)
	}

	if len(value) == 0 {
		delete(mlf.Values, language.Code)
		return
	}

	mlf.Values[language.Code] = value
}

// NewMultiLanguageString returns a new multi-language string.
func NewMultiLanguageString() *MultiLanguageString {
	return &MultiLanguageString{}
}

// MultiLanguageStringArray represents a string array supports multi-language.
type MultiLanguageStringArray struct {
	*MultiLanguageString
	Separator string
}

// Values returns values by given language.
func (mlsa *MultiLanguageStringArray) Values(language *Language) []string {
	val := mlsa.MultiLanguageString.Value(language)
	if len(val) == 0 {
		return nil
	}

	vals := strings.Split(val, mlsa.Separator)
	return vals
}

// SetValues sets the values.
func (mlsa *MultiLanguageStringArray) SetValues(language *Language, values []string) {
	vals := strings.Join(values, mlsa.Separator)
	mlsa.MultiLanguageString.SetValue(language, vals)
}

// NewMultiLanguageStringArray returns a new multi-language string array.
// Default separator sets to comma (,) if empty separator passed.
func NewMultiLanguageStringArray(separator string) *MultiLanguageStringArray {
	sep := separator
	if len(sep) == 0 {
		sep = ","
	}

	return &MultiLanguageStringArray{
		MultiLanguageString: NewMultiLanguageString(),
		Separator:           sep,
	}
}
