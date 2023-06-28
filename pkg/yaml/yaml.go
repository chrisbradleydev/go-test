package yaml

import (
	"fmt"
	"os"

	"github.com/chrisbradleydev/go-test/utils"
	"gopkg.in/yaml.v3"
)

type User struct {
	Name string `yaml:"name"`
	Age  int8    `yaml:"age"`
	Address struct {
		Street string `yaml:"street"`
		City string `yaml:"city"`
		State string `yaml:"state"`
		Zip int `yaml:"zip"`
	} `yaml:"address"`
	PhoneNumbers []struct {
		Type string `yaml:"type"`
		Number int `yaml:"number"`
	} `yaml:"phoneNumbers"`
}

func GetUser() {
	data, err := utils.GetData("data/user.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	var user User
	err = yaml.Unmarshal(data, &user)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user)
}

func GetFirstName() (string, error) {
	f := "First"
	data, err := os.ReadFile("data/user.yaml")
	if err != nil {
		return f, err
	}

	var user User
	err = yaml.Unmarshal(data, &user)
	if err != nil {
		return f, err
	}

	return user.Name, nil
}

func GetLastName() (string, error) {
	l := "Last"
	data, err := os.ReadFile("data/user.yaml")
	if err != nil {
		return l, err
	}

	var user User
	err = yaml.Unmarshal(data, &user)
	if err != nil {
		return l, err
	}

	return user.Name, nil
}

func GetAge() (int8, error) {
	var a int8 = 0
	data, err := os.ReadFile("data/user.yaml")
	if err != nil {
		return a, err
	}

	var user User
	err = yaml.Unmarshal(data, &user)
	if err != nil {
		return a, err
	}

	return user.Age, nil
}