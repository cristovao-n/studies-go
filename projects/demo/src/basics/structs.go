package basics

import "fmt"

// types and structs

type address struct {
	city   string
	street street
}

type flatAddress struct {
	city string
	street
}

type street struct {
	name   string
	number uint
}

func typesStructs() {
	var myAddress address = address{city: "Esperança", street: street{name: "Rua 1", number: 71}}
	var myAddress2 = struct {
		city   string
		name   string
		number uint
	}{city: "Esperança", name: "Rua 1", number: 71}
	fmt.Println(myAddress.city, myAddress.street.name)
	fmt.Println(myAddress2.city, myAddress2.name, myAddress2.number)
}

// TODO: methods

// TODO: interfaces
