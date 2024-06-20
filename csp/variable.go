package csp

import (
	"fmt"
	"strings"
)

type Variable struct {
	Index    int
	Value    Value
	Assigned bool

	Constraints []Constraint
	Domain      Domain
}

func (v *Variable) Assign(value Value) {
	v.Value = value
	v.Assigned = true
}

func (v *Variable) AssignAndReduceDomain(value Value) {
	v.Value = value
	v.Assigned = true
	v.RemoveRestFromDomain(value)
}

func (v *Variable) Unassign() {
	v.Value = 0
	v.Assigned = false
}

func (v *Variable) RemoveFromDomain(value Value) {
	v.Domain = v.Domain.Remove(value)
}

func (v *Variable) RemoveRestFromDomain(valueToKeep Value) {
	v.Domain = v.Domain.RemoveAllBut(valueToKeep)
}

func (v *Variable) String() string {
	builder := &strings.Builder{}
	builder.WriteString(fmt.Sprintf("variable %d: value=", v.Index))
	if v.Assigned {
		builder.WriteString(fmt.Sprintf("%d", v.Value))
	} else {
		builder.WriteString("<none>")
	}
	return builder.String()
}
