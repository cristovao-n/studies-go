package basics

import "fmt"

// strings
func strings() {
	// Strings in Go are represented as an arrays of bytes and are encoding using UTF-8
	// If the UTF-8 character has more than one byte, it'll need more than one position in the array of bytes

	// noun[1] and noun[2] are used to store 'é' character
	var noun = "résumé"

	fmt.Println(noun)

}

// TODO: Pointers
