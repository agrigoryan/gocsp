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

func (v *Variable) Clone() Variable {
	return Variable{
		Index:       v.Index,
		Value:       v.Value,
		Assigned:    v.Assigned,
		Constraints: v.Constraints,
		Domain:      v.Domain.Clone(),
	}
}

func (v *Variable) Copy(other *Variable) {
	other.Index = v.Index
	other.Value = v.Value
	other.Assigned = v.Assigned
	other.Constraints = v.Constraints
	other.Domain = v.Domain.Clone()
}
