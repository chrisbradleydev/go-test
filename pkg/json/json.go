package json

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	Name  string
	Class string
	Order string
}

var jsonBlob = []byte(`[
{"Name": "Eagle",   "Class": "Aves",     "Order": "Accipitriformes"},
{"Name": "Leopard", "Class": "Mammalia", "Order": "Carnivora"},
{"Name": "Orca",    "Class": "Mammalia", "Order": "Artiodactyla"}
]`)

func GetAnimals() {
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", animals)
}