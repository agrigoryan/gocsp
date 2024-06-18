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
	for i := 0; i < d.values.Size(); i++ {
		if d.values[i] != value {
			d.Remove(d.values[i])
			i--
		}
	}
}

func (d *Domain) String() string {
	return d.values.String()
}

func (d *Domain) ShallowCopy() Domain {
	return Domain{values: d.values[:]}
}
