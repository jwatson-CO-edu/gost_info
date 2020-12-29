package main

/*** Imports ***/
import (
	"fmt" //- Printing
	//"os" // Terminal args
)

/*** Aliases ***/
var fprint = fmt.Println // Alias for PrintLine
var fprintf = fmt.Printf //- Alias for PrintFormat

// 27. Struct
type Vertex struct {
	X int
	y int
}

/********** MAIN *********************************************************************************/
func main(){ /*Terminal args*/ //progArgs := os.Args[1:]

	// 26. Initialize Pointers
	i , j := 42 , 2701 // Initialize 2 ints
	p := &i // `p` points to the address of `i`
	*p = 21 // Dereference `p` and assign the value 21 at the address of `i`
	fprint( i ) // "21"
	p = &j
	*p /= 37 // Divide `j` by 37 through the pointer
	fprint( j ) // "73"
}