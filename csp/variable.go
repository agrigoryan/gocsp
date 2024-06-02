package csp

import (
	"fmt"
	"strings"
)

type Variable struct {
	Id       any
	Domain   *Domain
	Value    Value
	Assigned bool
}

func (v *Variable) Assign(value Value) error {
	if !v.Domain.Contains(value) {
		return fmt.Errorf("can not assign, value %v is not in the variable's domain", value)
	}
	v.Domain.RemoveAllBut(value)
	v.Value = value
	v.Assigned = true
	return nil
}

func (v *Variable) String() string {
	builder := &strings.Builder{}
	builder.WriteString(fmt.Sprintf("variable %v: value=", v.Id))
	if v.Assigned {
		builder.WriteString(fmt.Sprintf("%d", v.Value))
	} else {
		builder.WriteString("<nil>")
	}
	builder.WriteString(fmt.Sprintf(", domain=%v", v.Domain))
	return builder.String()
}

func NewVariable(id any, domain *Domain) *Variable {
	return &Variable{
		Id:     id,
		Domain: domain,
	}
}
