package main

/*** Imports ***/
import (
	"fmt" //- Printing
	"strings"
	"math"
	//"os" // Terminal args
)

/*** Aliases ***/
var fprint = fmt.Println // Alias for PrintLine
var fprintf = fmt.Printf //- Alias for PrintFormat

/***** Utility Functions *******************************************/

func printSlice( s []int ){
	// Print info about a slice
	fmt.Printf( "len=%d cap=%d %v\n", len(s), cap(s), s )
}

/********** PART 3: Structs and Memory ***********************************************************/
// https://tour.golang.org/moretypes/1

// 27. Struct
type Vertex struct {
	x int
	y int
}

type Loc struct {
	Lat, Long  float64
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

	// 28. Instantiate a struct
	fprint( Vertex{1,2} ) // "{1 2}"

	// 29. Access a struct
	vtx := Vertex{3,4}
	fprint( "X:" , vtx.x , ", Y:" , vtx.y )

	// 30. Pointers to structs
	v_ptr := &vtx
	v_ptr.x = 5
	v_ptr.y = 6
	fprint( "X:" , vtx.x , ", Y:" , vtx.y )

	// 31. Partially-initialized struct
	vtx2 := Vertex{ x: 7 } // y:0 is implicit
	fprint( "X:" , vtx2.x , ", Y:" , vtx2.y )

	// 32. String Array
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fprint( a[0] , a[1] , "!" ) // Hello World !

	// 33. Int Array Initialization
	arr2 := [6]int { 1, 2, 3, 4, 5, 6 }
	fprint( arr2 ) // "[1 2 3 4 5 6]"

	// 34. Slices
	// arr2 = [1 2 3 4 5 6]
	// Index:  0 1 2 3 4 5
	var slc1 []int = arr2[1:4];  fprint( slc1 ) // [2 3 4]
	var slc2 []int = arr2[ :3];  fprint( slc2 ) // [1 2 3]
	var slc3 []int = arr2[3: ];  fprint( slc3 ) // [4 5 6]
	// [ <Beginning Index> : <Index AFTER the End> ]
	// var slc4 []int = arr2[ :-2];  fprint( slc4 ) // CANNOT use Python negative indices

	// 35. Slices in Memory
	// A slice does not store any data, it just describes a section of an underlying array.
	half1 := arr2[ :3];  half1[1] = 7
	half2 := arr2[3: ];  half2[1] = 8
	fprint( arr2 ) /* "[1 7 3 4 8 6]"  // Note that the original array was CHANGED
						  ^		^   */
	
	// 36. Array Literals -vs- Slice Literals
	// This is an array literal:
	q := [3]bool{true, true, false}
	// And this creates the same array as above, then builds a slice that references it:
	r := []bool{true, true, false}
	fprint( q )
	fprintf( "%p\n" , &p )
	fprint( r )
	fprintf( "%p\n" , &r )

	// 37. Initialize an array of structs
	strctArr := [4]struct{
		i int
		b bool
	}{
		{ 1, true  },
		{ 2, false },
		{ 3, true  },
		{ 4, false },
	}
	fprint( strctArr ) // [{1 true} {2 false} {3 true} {4 false}]

	// 38. Slice Length and Capacity
	/* A slice has both a length and a capacity.
	   The length of a slice is the number of elements it contains.
	   The capacity of a slice is the number of elements in the underlying array. */
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // "len=6 cap=6 [2 3 5 7 11 13]" // This is expected
	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s) // "len=0 cap=6 []" // It disappeared!
	// Extend its length.
	s = s[:4] // "len=4 cap=6 [2 3 5 7]" // It came back! The array itself is hiding out in memory
	printSlice(s)
	// NOTE: We do not manipulate the array through `s`, we manipulate the slice
	//       The array `[2 3 5 7 11 13]` does NOT have a name, but Go is still counting refs to it!
	
	// 39. `nil` slices
	var s1 []int // A nil slice has a length and capacity of 0 and has no underlying array.
	fprint( s1, len(s1), cap(s1) ) // [] 0 0
	if s1 == nil {  fprint("nil!")  } // "nil!"

	// 40. Dynamic Arrays
	t1 := make( []int, 5 )
	printSlice( t1 ) // len=5 cap=5 [0 0 0 0 0]
	/**/
	t2 := make( []int, 0, 5 )
	printSlice( t2 ) // len=0 cap=5 []
	/**/
	t3 := t2[ :2]
	printSlice( t3 ) // len=2 cap=5 [0 0]
	/**/
	t4 := t3[2:5]
	printSlice( t4 ) // len=3 cap=3 [0 0 0]

	// 41. Slices of Slices
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	/**/
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
	/*  X _ X
	    O _ X
		_ _ O  */

	// 42. Appending to a Slice
	var u []int
	printSlice( u ) // len=0 cap=0 []
	// append works on nil slices.
	u = append( u, 0 )
	printSlice( u ) // len=1 cap=1 [0]
	// The slice grows as needed.
	u = append( u, 1 )
	printSlice( u ) // len=2 cap=2 [0 1]
	// We can add more than one element at a time.
	u = append( u, 2, 3, 4 )
	printSlice( u ) // len=5 cap=6 [0 1 2 3 4]

	// 43. Iterating over Array / Slice w/ `range`
	var pow2 = []int{ 1, 2, 4, 8, 16, 32, 64, 128 }
	for i, v := range pow2 { // works like Python `enumerate`
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// 44. "Don't Care" with `range`
	pow := make([]int, 10)
	// A. Using the index only
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	// B. Using the value only
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	// 45. Maps: maps keys to values
	var m map[ string ] Loc
	//    map[ <key type> ] <value type>
	fprintf( "%#v\n" , m ) // map[string]main.Vertex(nil)
	// Right now the map is nil, you cannot add keys
	// The `make` function returns a map of the given type, initialized and ready for use.
	m = make( map[ string ] Loc )
	// Now we can add a key-val pair
	m["Bell Labs"] = Loc{ 40.68433, -74.39967 }
	fprint( m["Bell Labs"] ) // "{40.68433 -74.39967}"

	// 46. Map literals
	var n = map[ string ] Loc {
		"Bell Labs": Loc{
			40.68433, -74.39967,
		},
		"Google": Loc{
			37.42202, -122.08408,
		},
	}
	fprint( n ) // "map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]"
	// If the top-level type is just a type name, you can omit it from the elements of the literal.
	var o = map[ string ] Loc {
		"Bell Labs": {40.68433, -74.39967},
		"Google":    {37.42202, -122.08408},
	}
	fprint( o ) // "map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]"

	// 47. Modifying maps
	w := make( map[ string ] int )
	/* Insert or update an element in map m:  m[key] = elem   */
	w["Answer"] = 42
	fmt.Println("The value:", w["Answer"])
	w["Answer"] = 48
	fmt.Println("The value:", w["Answer"])

	// 48. Delete from map by key
	delete(w, "Answer")
	fmt.Println("The value:", w["Answer"])

	// 49. Test for existence in map
	v, ok := w["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

	// 52. Initialise a variable to a lambda function
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fprint( hypot( 3, 4 ) )
	fprint( compute( hypot ) )

	// 53. Closures (Function Factory)
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

}

// 50. We can define functions AFTER main
// 51. Functions as Parameters
func compute( // Function name `compute`
	fn func(float64, float64) float64, // Paramer: `fn`, Type: Function that takes 2 floats and returns 1 float
) float64 { // `compute` returns a float
	return fn( 3, 4 ) // Pass 3 and 4 to `fn` and return the result
}

// 53. Closures (Function Factory)
func adder() ( func(int) int ) {
	//         ^^^^^^^^^^^^^-- Return type is a function that takes an int and returns an int
	sum := 0 // <-- *Enclosed* variable
	// 54. Return a lambda function
	return func(x int) int {
		sum += x
		return sum
	}
}