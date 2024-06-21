package csp

type Domain struct {
	values ValueSet
}

func (d *Domain) Size() int {
	return d.values.Size()
}

func (d *Domain) Values() []Value {
	return d.values
}

func (d *Domain) Value(idx int) Value {
	return d.values[idx]
}

func (d *Domain) Contains(value Value) bool {
	return d.values.Contains(value)
}

func (d *Domain) Add(value Value) {
	d.values = append(d.values, value)
}

func (d *Domain) Remove(value Value) {
	d.values = d.values.Remove(value)
}

func (d *Domain) RemoveAllBut(value Value) {
	d.values = d.values.RemoveAllBut(value)
}

func (d *Domain) String() string {
	return d.values.String()
}

func (d *Domain) Copy() Domain {
	return Domain{
		values: d.values,
	}
}

func NewDomain(values []Value) Domain {
	return Domain{values: values}
}
