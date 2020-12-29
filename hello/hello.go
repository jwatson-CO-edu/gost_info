/***** Tour of Go *****
* https://tour.golang.org/list
 */

package main

/********** PART 1: Basics ***********************************************************************/
// https://tour.golang.org/basics/1

/*** Imports ***/
import (
	"fmt"        // -------- Fancy printing
	"math"       // ------- Basic math functions
	"math/cmplx" // - Complex numbers
	"math/rand"  // -- Pseudo-random number generation
	"runtime"    // ---- Interact with Go's runtime system
	"time"       // ------- Time stamps and calc
	"rsc.io/quote" // Fortune Cookie
)

/*** Aliases ***/
var fprint  = fmt.Println // Alias for PrintLine
var fprintf = fmt.Printf //- Alias for PrintFormat

/*** Global vars ***
var <varName> <typeName> */
var foo float64
var c, python, java bool // Initialized to false
var j, k int = 1, 2      // Initialized to values

/*** Mass-declare Variables ***/
var (
	ToBe   bool       = /*-*/ false
	MaxInt uint64     = 1<<64 - 1 // Left-shift
	z      complex128 = cmplx.Sqrt(-5 + 12i)
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
func add(x int, y int) int {
	return x + y
}

func give_two() int { return 2 }

//// B. Trailing type ||| means all the same type
////                  vvv
func add2(x, y int) int {
	return x + y
}

// C. Multiple return
func swap_str(x, y string) (string, string) {
	return y, x
}

// D. Named return, Names return vars behave like local vars,
//    you do not have to list them in the return statement
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // A "naked return" returns the named return variables
	// Suggestion: Only use naked return in short functions where you won't forget what was returned
}

/********** MAIN *********************************************************************************/

func main() {

	//  1. Print fortune
	fprint(quote.Go())

	//  2. Print time
	fprint("The time is", time.Now())

	//  3. Choose a random int from 0 to N
	fprint("My favorite number is", rand.Intn(100))

	//  4. Math library
	fprintf("Now you have %g problems.\n", math.Sqrt(7))

	//  5. Unicode
	fprint("\u03C0 is exactly", math.Pi)

	//  6. Function call
	fprint(add(2, 3))

	//  7. Define multiple variables with function output
	a, b := swap_str("Hello", "World")
	fprint(a, b)

	//  8. Local var
	var i int                  // ------------------ Initialized to 0
	fprint(i, c, python, java) // "0 false false false"

	//  9. Initialized local var
	var l int = 7
	fprint("My second-favorite number is", l)

	// 10. Declare variable with implicit type
	m := 88
	fprint("My third-favorite number is", m)

	// 11. Type conversions
	o := 42
	f := float64(o)
	u := uint(f)
	fprint(u, "is a nice number as well")

	// 12. Type inferences
	q := 42
	fprintf("q has type %T\n", q) // int
	r := 3.1416
	fprintf("r has type %T\n", r) // float64
	s := 0.867 + 0.5i
	fprintf("s has type %T\n", s) // complex128
	// Typename format char: '%T'

	// 13. Constants
	const result bool = true
	fprint("Is it true?", result)

	/********** PART 2: Flow Control *************************************************************/
	// https://tour.golang.org/flowcontrol/1

	// 14. For Loop
	sum := 0
	// CANNOT use parens around loop statements
	//>( nope ; not ; used )<
	for i := 0; i < 10; i++ { // Curly braces REQUIRED for all blocks
		sum += i
	}
	fprint("Sum =", sum)

	// 15. For Init and Post statements are optional
	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fprint("Sum =", sum)

	// 16. "While" is just `for` without Init and Post statements
	sum = 1
	for sum < 2000 {
		sum += sum
	}
	fprint("Sum =", sum)

	// 17. Infinite While can be achieved with a for loop with no loop statements
	sum = 1
	for {
		sum += sum
		if sum > 4000 {
			break
		} // 18. If statement, Parens optional, Curly braces required
	}
	fprint("Sum =", sum)

	// 19. You can add an Init Statement to If
	if v := 4; v < 5 {
		fprint(v, "is less than 5")
	} else { // 20. Else
		fprint(v, "is greater than 5") // `v` has scope here too
	}
	// fprint( v , "is less than 5" ), ERROR: Init var `v` only has scope INSIDE the If statement

	// 21. switch
	//	   Go only runs the selected case, not all the cases that follow. No `break` required

	fprintf("Go runs on ")

	switch os := runtime.GOOS; os {
	case "darwin":
		fprint("OS X.")
	case "linux":
		fprint("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fprint("%s.\n", os)
	}

	// 22. Eval switch cases
	fprint("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0: // A case can be any statement
		fprint("Today.")
	case today + 1:
		fprint("Tomorrow.")

	case today + time.Weekday( // Explict cast
		give_two(), // Can also put function calls here
	): // NOTE: When arg appears on its own line, line MUST end in a comma
		fprint("In two days.")

	default:
		fprint("Too far away.")
	}

	// 23. Switch with No Condition
	t := time.Now()
	fprintf("The time is now %s, ", t.String())
	switch { // No condition
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	/* 24. `defer`
	   A defer statement defers the execution of a function until the surrounding function returns.
	   The deferred call's arguments are evaluated immediately,
	   but the function call is not executed until the surrounding function returns. */
	defer fprint("Ooops, I was `defer`ed!")
	fprint("Hello!")

	/* 25. Stacked `defer`
	   Deferred function calls are pushed onto a stack.
	   When a function returns, its deferred calls are executed in last-in-first-out order. */
	for i := 0; i < 10; i++ {
		defer fprintf("%d, ", i) // Note that this counts *backwards*
	}
}

/********** Package / Implementation Facts ********************************************************

* "go.sum" contains the expected cryptographic checksums of the content of specific module versions.

* The go.mod file defines the module’s module path,
  which is also the import path used for the root directory,
  and its dependency requirements, which are the other modules needed for a successful build.
  Each dependency requirement is written as a module path and a specific semantic version.


*/
