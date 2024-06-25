package csp

import (
	"fmt"
	"strings"
)

type Value int

type ValueSet []Value

func (s ValueSet) indexOf(value Value) int {
	for i := 0; i < len(s); i++ {
		if s[i] == value {
			return i
		}
	}
	return -1
}

func (s ValueSet) Contains(value Value) bool {
	return s.indexOf(value) >= 0
}

func (s ValueSet) Size() int {
	return len(s)
}

func (s ValueSet) String() string {
	var b strings.Builder
	b.WriteRune('{')
	for i, v := range s {
		b.WriteString(fmt.Sprintf("%v", v))
		if i != len(s)-1 {
			b.WriteString(", ")
		}
	}
	b.WriteRune('}')
	return b.String()
}

func NewValueSet(values []Value) ValueSet {
	set := make(ValueSet, len(values))
	for i, v := range values {
		set[i] = v
	}
	return set
}
