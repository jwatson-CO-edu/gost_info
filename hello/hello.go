package main

/*** Imports ***/
import ( "fmt" ; "rsc.io/quote" ; "time" ; "math/rand" ; "math" ; "math/cmplx" )

/*** Aliases ***/
var fprint  = fmt.Println
var fprintf = fmt.Printf

/*** Global vars ***/
var c, python, java bool // Initialized to false
var j, k int = 1, 2 // Initialized to values

/*** Mass-declare Variables ***/
var (
	ToBe   bool /*-*/ = false
	MaxInt uint64     = 1 << 64 - 1 // Left-shift
	z      complex128 = cmplx.Sqrt( -5 + 12i )
	piRn   rune       = '\u03C0' // Unicode char
)

/*** Constants ***/
/* Constants are declared like variables, but with the const keyword.
   Constants can be character, string, boolean, or numeric values.
   Constants cannot be declared using the := syntax. 
   An untyped constant takes the type needed by its context. */
const (
	one   = 1
	two   = 2
	three = 3
) 

/********** Functions ****************************************************************************/

// A. Define a function
func add( x int , y int ) int {
	return x + y
}

// B. Trailing type ||| means all the same type
//                  vvv 
func add2( x , y    int ) int {
	return x + y
}

// C. Multiple return
func swap_str( x , y string ) ( string , string ){
	return y , x
}

// D. Named return, Names return vars behave like local vars, you do not have to list them in the return statement
func split( sum int ) ( x, y int ) {
	x = sum * 4 / 9
	y = sum - x
	return // A "naked return" returns the named return variables
	// Suggestion: Only use naked return in short functions where you won't forget what was returned
}

/********** MAIN *********************************************************************************/

func main(){

	//  1. Print fortune
	fprint( quote.Go() )

	//  2. Print time
	fprint( "The time is" , time.Now() )

	//  3. Choose a random int from 0 to N
	fprint( "My favorite number is" , rand.Intn( 100 ) )

	//  4. Math library
	fprintf( "Now you have %g problems.\n" , math.Sqrt( 7 ) )

	//  5. Unicode
	fprint( "\u03C0 is exactly" , math.Pi )

	//  6. Function call
	fprint( add(2,3) )

	//  7. Define multiple variables with function output
	a, b := swap_str( "Hello" , "World" )
	fprint( a , b )

	//  8. Local var
	var i int // Initialized to 0
	fprint( i , c , python , java ) // 0 false false false

	//  9. Initialized local var
	var l int = 7
	fprint( "My second-favorite number is" , l )

	// 10. Declare variable with implicit type
	m := 88
	fprint( "My third-favorite number is" , m )

	// 11. Type conversions
	o := 42
	f := float64(o)
	u := uint(f)
	fprint( u , "is a nice number as well" )

	// 12. Type inferences
	q := 42           ; fprintf( "q has type %T\n" , q ) // int
	r := 3.1416       ; fprintf( "r has type %T\n" , r ) // float64
	s := 0.867 + 0.5i ; fprintf( "s has type %T\n" , s ) // complex128
	// Typename format char: '%T'

	// 13. Constants
	const result bool = true
	fprint( "Is it true?" , result )

	// FIXME: START HERE:  https://tour.golang.org/flowcontrol/1

}

/********** Package / Implementation Facts ********************************************************

* "go.sum" contains the expected cryptographic checksums of the content of specific module versions.

* The go.mod file defines the module’s module path, which is also the import path used for the root directory, 
  and its dependency requirements, which are the other modules needed for a successful build. 
  Each dependency requirement is written as a module path and a specific semantic version.


*/
