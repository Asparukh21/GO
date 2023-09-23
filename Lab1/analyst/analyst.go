package analyst

type Analyst struct {
	position string
	salary   int
	address  string
}

func NewAnalyst(position string, salary int, address string) Analyst {
	return Analyst{position: position, salary: salary, address: address}
}

func (a *Analyst) GetPosition() string {
	return a.position
}

func (a *Analyst) SetPosition(position string) {
	a.position = position
}

func (a *Analyst) GetSalary() int {
	return a.salary
}

func (a *Analyst) SetSalary(salary int) {
	a.salary = salary
}

func (a *Analyst) GetAddress() string {
	return a.address
}

func (a *Analyst) SetAddress(address string) {
	a.address = address
}
