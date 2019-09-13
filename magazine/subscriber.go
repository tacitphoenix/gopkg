// Package magazine information about subscribers and employees
package magazine

type Subscriber struct {
	Name    string
	Rate    float64
	Active  bool
	Address Address
}

type Employee struct {
	Name    string
	Salary  float64
	Address Address
}

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode int
}
