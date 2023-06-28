package user

import (
	"fmt"

	"github.com/chrisbradleydev/go-test/pkg/yaml"
)

type UserNameAge struct {
	firstName string
	lastName string
	age int8
}

func (u *UserNameAge) incrementAge(i int8) {
	fmt.Printf("curr age: %d\n", u.age)

	if u.age < 127 {
		newAge := u.age + 1
		u.age = newAge
		fmt.Printf("next age: %d\n", newAge)
	} else {
		fmt.Printf("You're old AF!")
	}
}

func IncAge() {
	f, _ := yaml.GetFirstName()
	l, _ := yaml.GetLastName()
	a, _ := yaml.GetAge()

	p := &UserNameAge{
		firstName: f,
		lastName: l,
		age: a,
	}

	fmt.Printf("Hello, %s.\n", p.firstName)

	p.incrementAge(1)
}