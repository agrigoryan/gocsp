package csp

import (
	"fmt"
	"strings"
)

type Variable struct {
	Id       any
	Value    Value
	Assigned bool
}

func (v *Variable) Assign(value Value) {
	v.Value = value
	v.Assigned = true
}

func (v *Variable) String() string {
	builder := &strings.Builder{}
	builder.WriteString(fmt.Sprintf("variable %v: value=", v.Id))
	if v.Assigned {
		builder.WriteString(fmt.Sprintf("%d", v.Value))
	} else {
		builder.WriteString("<none>")
	}
	return builder.String()
}
