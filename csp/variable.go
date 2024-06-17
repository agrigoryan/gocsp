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
}

func (v Variable) Assign(value Value) Variable {
	return Variable{
		Index:       v.Index,
		Value:       value,
		Assigned:    true,
		Constraints: v.Constraints,
	}
}

func (v Variable) Unassign() Variable {
	return Variable{
		Index:       v.Index,
		Assigned:    false,
		Constraints: v.Constraints,
	}
}

func (v Variable) String() string {
	builder := &strings.Builder{}
	builder.WriteString(fmt.Sprintf("variable %d: value=", v.Index))
	if v.Assigned {
		builder.WriteString(fmt.Sprintf("%d", v.Value))
	} else {
		builder.WriteString("<none>")
	}
	return builder.String()
}
