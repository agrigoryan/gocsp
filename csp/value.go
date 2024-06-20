package csp

import (
	"fmt"
	"strings"
)

type Value int32

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

func (s ValueSet) Remove(value Value) ValueSet {
	idx := s.indexOf(value)
	if idx >= 0 {
		s[idx], s[len(s)-1] = s[len(s)-1], s[idx]
		return s[:len(s)-1]
	}
	return s
}

func (s ValueSet) RemoveAllBut(valueToKeep Value) ValueSet {
	for i := 0; i < s.Size(); i++ {
		if s[i] == valueToKeep {
			s[0], s[i] = s[i], s[0]
			return s[:1]
		}
	}
	return s
}

func (s ValueSet) Add(value Value) ValueSet {
	if s.Contains(value) {
		return s
	}
	return append(s, value)
}

func (s ValueSet) Copy() ValueSet {
	res := make([]Value, len(s))
	copy(res, s)
	return res
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
	set := ValueSet{}
	for _, v := range values {
		set = set.Add(v)
	}
	return set
}
