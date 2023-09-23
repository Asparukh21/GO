package economist

type Economist struct {
	position string
	salary   int
	address  string
}

func NewEconomist(position string, salary int, address string) Economist {
	return Economist{position: position, salary: salary, address: address}
}

func (e *Economist) GetPosition() string {
	return e.position
}

func (e *Economist) SetPosition(position string) {
	e.position = position
}

func (e *Economist) GetSalary() int {
	return e.salary
}

func (e *Economist) SetSalary(salary int) {
	e.salary = salary
}

func (e *Economist) GetAddress() string {
	return e.address
}

func (e *Economist) SetAddress(address string) {
	e.address = address
}
