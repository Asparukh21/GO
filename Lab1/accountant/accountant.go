package accountant

type Accountant struct {
	position string
	salary   int
	address  string
}

func NewAccountant(position string, salary int, address string) Accountant {
	return Accountant{position: position, salary: salary, address: address}
}

func (a *Accountant) GetPosition() string {
	return a.position
}

func (a *Accountant) SetPosition(position string) {
	a.position = position
}

func (a *Accountant) GetSalary() int {
	return a.salary
}

func (a *Accountant) SetSalary(salary int) {
	a.salary = salary
}

func (a *Accountant) GetAddress() string {
	return a.address
}

func (a *Accountant) SetAddress(address string) {
	a.address = address
}
