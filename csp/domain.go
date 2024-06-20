package csp

type Domain struct {
	values ValueSet
}

func (d Domain) Size() int {
	return d.values.Size()
}

func (d Domain) Values() []Value {
	return d.values
}

func (d Domain) Contains(value Value) bool {
	return d.values.Contains(value)
}

func (d Domain) Add(value Value) Domain {
	d.values = append(d.values, value)
	return d
}

func (d Domain) Remove(value Value) Domain {
	d.values = d.values.Remove(value)
	return d
}

func (d Domain) RemoveAllBut(value Value) Domain {
	for i := 0; i < d.values.Size(); i++ {
		if d.values[i] == value {
			d.values[0], d.values[i] = d.values[i], d.values[0]
			d.values = d.values[:1]
			break
		}
	}
	return d
}

func (d Domain) String() string {
	return d.values.String()
}

func NewDomain(values []Value) Domain {
	return Domain{values: values}
}
