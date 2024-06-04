package csp

type CSP interface {
	Domains() []Domain
	Constraints() []Constraint
}

type GenericCSP struct {
	domains     []Domain
	constraints []Constraint
}

func (c *GenericCSP) Domains() []Domain {
	return c.domains
}

func (c *GenericCSP) Constraints() []Constraint {
	return c.constraints
}

func NewGenericCSP(domains []Domain, constraints []Constraint) *GenericCSP {
	return &GenericCSP{
		domains:     domains,
		constraints: constraints,
	}
}
