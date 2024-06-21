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
	v.Assign(value)
	v.Domain.RemoveAllBut(value)
}

func (v *Variable) Unassign() {
	v.Value = 0
	v.Assigned = false
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

func (v *Variable) Copy() Variable {
	return Variable{
		Index:       v.Index,
		Value:       v.Value,
		Assigned:    v.Assigned,
		Constraints: v.Constraints,
		Domain:      v.Domain,
	}
}
