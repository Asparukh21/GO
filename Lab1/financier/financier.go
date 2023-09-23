package financier

type Financier struct {
	position string
	salary   int
	address  string
}

func NewFinancier(position string, salary int, address string) Financier {
	return Financier{position: position, salary: salary, address: address}
}

func (f *Financier) GetPosition() string {
	return f.position
}

func (f *Financier) SetPosition(position string) {
	f.position = position
}

func (f *Financier) GetSalary() int {
	return f.salary
}

func (f *Financier) SetSalary(salary int) {
	f.salary = salary
}

func (f *Financier) GetAddress() string {
	return f.address
}

func (f *Financier) SetAddress(address string) {
	f.address = address
}
