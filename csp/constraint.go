package csp

type Constraint interface {
	IsSatisfied(assignment Assignment) bool

	// AppliesTo - The list of indices this constraint applies to
	AppliesTo() []int
}

type AllDifferent struct {
	Indices []int
}

func (c AllDifferent) IsSatisfied(assignment Assignment) bool {
	for i := 0; i < len(c.Indices); i++ {
		var1 := assignment.Variables[c.Indices[i]]
		if !var1.Assigned {
			continue
		}
		for j := i + 1; j < len(c.Indices); j++ {
			var2 := assignment.Variables[c.Indices[j]]
			if var2.Assigned && var1.Value == var2.Value {
				return false
			}
		}
	}
	return true
}

func (c AllDifferent) AppliesTo() []int {
	return c.Indices
}
