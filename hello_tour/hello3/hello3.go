package main
/***** MAIN_TEMPLATE.go *****
A_ONE_LINE_DESCRIPTION_OF_THE_FILE
James Watson, 202X-MM
Template Ver. 2020-12
*/

/*** Imports ***/
import (
	"fmt" //- Printing
	"math"
	"time"
	"io"
	"strings"
	"image"
	//"os" // Terminal args
)

/*** Aliases ***/
var fprint  = fmt.Println // Alias for PrintLine
var fprintf = fmt.Printf //- Alias for PrintFormat

/********** PART 4: Methods and Interfaces *******************************************************/
// https://tour.golang.org/methods/1


/********** Vertex **********/
/***** Struct *****/
type Vertex struct {
	x, y float64
}
/***** 54. Methods *****
A method is a function with a special receiver argument.
The receiver appears in its own argument list between the func keyword and the method name.
In this example, the Abs method has a receiver of type Vertex named v. 
       ********  <<------------------------------------''''''''''''''         */
func ( v Vertex ) Abs() float64 {
	// Return the Euclidean distance of the vert from the origin
	return math.Sqrt( math.Pow(v.x,2) + math.Pow(v.y,2) )
}

// 55. Type alias
type MyFloat float64

/* 56. Method of an aliased type
You can only declare a method with a receiver whose type is defined in the same package as the method. 
You cannot declare a method with a receiver whose type is defined in another package 
(which includes the built-in types such as int). */
func ( f MyFloat ) Abs() float64 {
	if f < 0 {  return float64( -f )  }
	return float64(f) 
}

/* 57. Value -vs- Pointer receivers
With a value receiver, the Scale method operates on a copy of the original Vertex value. 
(This is the same behavior as for any other function argument.) 
ALL NON-POINTER ARGUMENTS ARE PASS-BY-VALUE. 

FUNCTIONS that take a value argument *MUST* take a value of that specific type.
METHODS with value receivers can take *EITHER* a value or a pointer as the receiver.
*/

func ( v Vertex ) Scale1( f float64 ) {
	// Value receiver, no change to the struct
	v.x = v.x * f
	v.y = v.y * f
}

/* Methods with pointer receivers can modify the value to which the receiver points (as `Scale2` does here). 
Since methods often need to modify their receiver, pointer receivers are more common than value receivers. 
The `Scale2` method *MUST* have a pointer receiver to change the Vertex value declared in the main function.

FUNCTIONS with a pointer argument *MUST* take a pointer.
METHODS with pointer receivers can take *EITHER* a value or a pointer as the receiver.
*/

func ( v *Vertex ) Scale2( f float64 ) {
	// Pointer reciever, the struct is modified
	v.x = v.x * f
	v.y = v.y * f
}
/* There are two reasons to use a pointer receiver:
	1. So that the method can modify the value that its receiver points to.
	2. To avoid copying the value on each method call. 
	   This can be more efficient if the receiver is a large struct, for example. */

/* 58. Interfaces
An interface type is defined as a set of method signatures.
A value of interface type can hold any value that implements those methods. */
type Abser interface {
	Abs() float64 // All `Abser`s must implement an `Abs` function that takes nothing and returns float
}
/* A type implements an interface by implementing its methods. 
There is no explicit declaration of intent, no "implements" keyword.
Implicit interfaces decouple the definition of an interface from its implementation, 
which could then appear in any package without prearrangement. 

In Go it is common to write methods that gracefully handle being called with a nil receiver.
Note that an interface value that holds a nil concrete value is itself non-nil. */


// 59. Nil interface values
type I interface {
	M()
}

/* 64. String Representation (Stringer) Interface
A Stringer type must be able to describe itself as a string. 
*Many* look for this interface to print the string representation of the struct value. */
type Stringer interface { // NOTE: This is a built-in interface
    String() string
}

/* 65. Error Interface 
Functions often return an error value, 
and calling code should handle errors by testing whether the error equals nil. 
A nil error denotes success; a non-nil error denotes failure.*/
type error interface { // NOTE: This is a built-in interface
    Error() string
}

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}


/********** MAIN *********************************************************************************/
func main(){ /*Terminal args*/ //progArgs := os.Args[1:]

	// 54. Using methods
	vtx1 := Vertex{ 1, 2 }
	fprint( vtx1.Abs() )

	// 56. Method of aliased type
	f := MyFloat( -math.Sqrt(2) )
	fprint( f.Abs() )

	// 57. Value -vs- Pointer receivers
	vtx2 := Vertex{ 3, 4 }
	vtx2.Scale1( 2.0 )
	fprint( vtx2 ) // {3 4}, No change
	vtx2.Scale2( 2.0 )
	fprint( vtx2 ) // {6 8}, Vertex modified

	/* 58. Interfaces and Interface Values
	Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
	(value, type)
	An interface value holds a value of a specific underlying concrete type. */
	var absr Abser
	flt := MyFloat( -math.Sqrt(2) )
	vrt := Vertex{ 3, 4 }

	absr = flt
	fprint( absr.Abs() )
	absr = &vrt
	fprint( absr.Abs() )
	absr = vrt
	fprint( absr.Abs() )

	// 59. Nil interface values
	var i I // This interface value holds neither a value nor a type
	describe(i) // "(<nil>, <nil>)"
	// i.M() // panic: runtime error: invalid memory address or nil pointer dereference

	// 60. The empty interface
	var ii interface{}
	describe(ii) // "(<nil>, <nil>)"

	// 61. Type assertions
	var iii interface{} = "hello" // This can hold anything, but it holds a string
	/* This statement asserts that the interface value `iii` holds the concrete type `string` 
	and assigns the underlying `string` value to the variable `sss`. 
	If i does not hold a T, the statement will trigger a panic. */
	sss := iii.(string)
	fprint( sss ) // "hello"

	/* 62. To test whether an interface value holds a specific type, a type assertion can return two values: 
	the underlying value and a boolean value that reports whether the assertion succeeded. */
	ttt, ok := iii.(string)
	fprint( ttt, ok ) // "hello true"

	// 63. Type Switches
	switch iii.(type) {
		case int:
			fprint( "Var was int" )
		case string:
			fprint( "Var was string" )
		default:
			fprint( "I don't know what this is!" )
	} // "Var was string"

	// 65. Error interface usage
	if err := run(); err != nil {
		fprint(err) 
	} // "at 2020-12-29 23:30:47.9924235 -0700 MST m=+0.000581601, it didn't work"

	/* 66. Reader Interface
	The io package specifies the io.Reader interface, which represents the read end of a stream of data.
	The Go standard library contains many implementations of this interface, 
	including files, network connections, compressors, ciphers, and others.
	The io.Reader interface has a Read method:
		`func (T) Read(b []byte) (n int, err error)` */
	rrr := strings.NewReader("Hello, Reader!")
	bbb := make([]byte, 8)
	for { // While there is string left, read it in 8-byte chunks
		n, err := rrr.Read(bbb)
		fprint("n = %v err = %v b = %v\n", n, err, bbb)
		fprint("b[:n] = %q\n", bbb[:n])
		if err == io.EOF {  break  }
	}

	// 67. Image Interface
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())

}

func describe( i interface{} ) {
	// Describe an `I` interface: Print the type and the value that is held
	// Helper for: 59 et al
	fprintf( "(%v, %T)\n", i, i )
}

