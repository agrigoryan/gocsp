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
		builder.WriteString(fmt.Sprintf("%v", v.ValueIdx))
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
	}
}

func (v *Variable) Copy(other *Variable) {
	other.Index = v.Index
	other.ValueIdx = v.ValueIdx
	other.Assigned = v.Assigned
	other.Constraints = v.Constraints
}
