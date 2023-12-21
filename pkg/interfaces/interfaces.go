package interfaces

import (
	"fmt"
)

type Uno struct {
	Name string `json:"name" yaml:"name"`
}

type Dos struct {
	Name string `json:"name" yaml:"name"`
}

type saver interface {
	Save()
}

func ExecInterfaces() {
	uno := NewUno()
	dos := NewDos()

	saveData(uno)
	saveData(dos)
}

func NewUno() *Uno {
	return &Uno{Name: "Uno"}
}

func NewDos() *Dos {
	return &Dos{Name: "Dos"}
}

// because both receivers have Save method
func (uno *Uno) Save() {
	fmt.Println("uno")
}

func (dos *Dos) Save() {
	fmt.Println("dos")
}

// saveData can receive both, Uno and Dos as an argument
func saveData(data saver) {
	data.Save()
}