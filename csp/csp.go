package csp

type CSP interface {
	Domains() []ValueSet
	Constraints() []Constraint
}

type GenericCSP struct {
	domains     []ValueSet
	constraints []Constraint
}

func (c *GenericCSP) Domains() []ValueSet {
	return c.domains
}

func (c *GenericCSP) Constraints() []Constraint {
	return c.constraints
}

func NewGenericCSP(domains []ValueSet, constraints []Constraint) *GenericCSP {
	return &GenericCSP{
		domains:     domains,
		constraints: constraints,
	}
}
