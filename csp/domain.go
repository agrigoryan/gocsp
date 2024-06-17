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
	for _, v := range d.values.Copy() {
		if v != value {
			d.Remove(v)
		}
	}
}

func (d *Domain) String() string {
	return d.values.String()
}

func (d *Domain) ShallowCopy() Domain {
	return Domain{values: d.values[:]}
}
