package csp

import (
	"fmt"
	"strings"
)

type Variable struct {
	Index    int
	ValueIdx int
	Assigned bool

	Constraints []Constraint
	Domain      Domain
}

func (v *Variable) Value() (value Value) {
	if v.Assigned {
		value = v.Domain.Value(v.ValueIdx)
	}
	return value
}

func (v *Variable) Assign(idx int) {
	v.ValueIdx = idx
	v.Assigned = true
}

func (v *Variable) Unassign() {
	v.ValueIdx = 0
	v.Assigned = false
}

func (v *Variable) String() string {
	builder := &strings.Builder{}
	builder.WriteString(fmt.Sprintf("variable %d: value=", v.Index))
	if v.Assigned {
		builder.WriteString(fmt.Sprintf("%v", v.Domain.Value(v.ValueIdx)))
	} else {
		builder.WriteString("<none>")
	}
	return builder.String()
}

func (v *Variable) Clone() Variable {
	return Variable{
		Index:       v.Index,
		ValueIdx:    v.ValueIdx,
		Assigned:    v.Assigned,
		Constraints: v.Constraints,
		Domain:      v.Domain.Clone(),
	}
}

func (v *Variable) Copy(other *Variable) {
	other.Index = v.Index
	other.ValueIdx = v.ValueIdx
	other.Assigned = v.Assigned
	other.Constraints = v.Constraints
	other.Domain = v.Domain.Clone()
}
