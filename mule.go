package mule

import "regexp"

type Mule interface {
	Check(ok bool, key, message string)
	Valid() bool
	Errors() map[string]string

	Unique(values []string) bool
	Matches(value string, rx *regexp.Regexp) bool
	In(value string, list ...string) bool
}

type mule struct {
	errors map[string]string
}

func New() Mule {

	return &mule{errors: make(map[string]string)}
}

func (m *mule) Errors() map[string]string {

	return m.errors
}

func (v *mule) Valid() bool {
	return len(v.errors) == 0
}

func (m *mule) addError(key, message string) {
	if _, exists := m.errors[key]; !exists {
		m.errors[key] = message
	}
}

func (m *mule) Check(ok bool, key, message string) {
	if !ok {
		m.addError(key, message)
	}
}

func (m *mule) In(value string, list ...string) bool {
	for i := range list {
		if value == list[i] {
			return true
		}
	}

	return false
}

func (m *mule) Matches(value string, rx *regexp.Regexp) bool {

	return rx.MatchString(value)
}

func (m *mule) Unique(values []string) bool {
	uniqueValues := make(map[string]bool)
	for _, value := range values {
		uniqueValues[value] = true
	}

	return len(values) == len(uniqueValues)
}
