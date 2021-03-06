package main
/***** MAIN_TEMPLATE.go *****
A_ONE_LINE_DESCRIPTION_OF_THE_FILE
James Watson, 202X-MM

***** DEV PLAN *****
[ ] Stage 1
[ ] Stage 2

Template Ver. 2020-12
*/

/*** Imports ***/
import (
	"fmt" //- Printing
	"time"
	"runtime"
	"sync"
	//"os" // Terminal args
)

/*** Aliases ***/
var fprint  = fmt.Println // Alias for PrintLine
var fprintf = fmt.Printf //- Alias for PrintFormat

/********** PART 5: Concurrency ******************************************************************/
// https://tour.golang.org/concurrency/1

/* 71. Channels
Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
		   // assign value to v.
Like maps and slices, channels must be created before use:
	ch := make(chan int)
By default, sends and receives *BLOCK* until the other side is ready. 
This allows goroutines to synchronize without explicit locks or condition variables. */

func sum( // Parameters
	s []int, //--- int slice
	c chan int, // int channel
){ // return void
	// Sum all the values in the array and send the int sum to the int channel
	sum := 0
	for _, v := range s {  sum += v }
	c <- sum // Send the int `sum` to the int channel `c`
}

// 73. Closing channels
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// 74. Select
func fibonacci2( c, quit chan int ) { // Takes 2 int channels: `c` and `quit`
	/* The select statement lets a goroutine wait on multiple communication operations. 
	A select blocks until one of its cases can run, then it executes that case. 
	It chooses one at random if multiple are ready.*/
	x, y := 0, 1
	for { // Infinite while
		select { // Do the first one possible
			case c <- x: // Push x onto the `c` channel
				x, y = y, x+y // Calc the next fib num
			case <- quit: // Pop a value from `quit` channel
				fprint( "quit" ) // Print "quit" and return
				return
		}
	}
}


/***** 77. Mutex ********************/

type SafeCounter struct {
	// SafeCounter is safe to use concurrently.
	mu sync.Mutex
	v  map[ string ] int
}

func ( c *SafeCounter ) Inc(key string) {
	// Inc increments the counter for the given key.
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

func ( c *SafeCounter ) Value( key string ) int {
	// Value returns the current value of the counter for the given key.
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock() // This runs AFTER the function returns, handy!
	return c.v[key]
}

/********** MAIN *********************************************************************************/
func main(){ /*Terminal args*/ //progArgs := os.Args[1:]

	// 68. Coroutines first try
	go say("world")
	say("hello")
	// world
	// hello
	// world
	// hello
	// hello
	// world
	// world
	// hello
	// hello

	// 69. System Parallelization Capcity & Golang Info, Nice
	fprint( "Num. CPU: _____" , runtime.NumCPU() )
	fprint( "Num. Goroutine:" , runtime.NumGoroutine() )
	fprint( "$GOROOT: ______" , runtime.GOROOT() )
	fprint( "Version: ______" , runtime.Version() )

	// 70. Goroutine Perf Test
	N := 10000000
	M :=        4
	bgn := time.Now()
	for i := 0 ; i < M ; i++{  pi_CPU_heater( N )  }
	fprint( "Serial Running Time:", time.Since( bgn ) )
	bgn  = time.Now()
	for i := 0 ; i < M ; i++{  go pi_CPU_heater( N )  }
	fprint( "Goroutine Running Time:", time.Since( bgn ) )
	// The second test returns immediately, but task manager does not show parallelism

	// 71. Channels
	// Create an int slice
	sss := []int{ 1, 2, 3, 4, 5, 6 }
	// Create an int channel
	ccc := make( chan int )
	hlfLen := len(sss)/2
	go sum(  sss[       :hlfLen ], ccc  )
	go sum(  sss[ hlfLen:       ], ccc  )
	x, y := <-ccc, <-ccc // Receive from the channel
	fprint( x, y )

	// 72. Buffered Channels
	chch := make(chan int, 2) // ( chan <type>, <buffer length> )
	chch <- 1
	chch <- 2
	fprint( <- chch ) // "1"
	fprint( <- chch ) // "2"

	// 73. Closing channels
	ddd := make( chan int, 10 ) // Create an int channel with a buffer size of 10
	go fibonacci( cap(ddd), ddd ) // Load the first 10 fib numbers in the channel
	for i := range ddd {  fprintf( "%d, ", i )  }
	fprint("")

	// 74. Select
	// Create the channels that `fibonacci2` is looking for
	eee  := make( chan int ) // Note that we did not specify a buffer size
	quit := make( chan int ) // So each of these channels can only hold one message
	// Run a lambda function in a goroutine
	go func() {
		// Fetch 10 elements from the channel
		for i := 0; i < 10; i++ {  fprint( <- eee )  }
		// Then send something to the `quit` channel
		quit <- 0
	}()
	// Run the fib num function
	fibonacci2( eee, quit )


	// 75. Capacity and length of channels
	fff := make( chan int )
	ggg := make( chan int, 10 )
	go func(){	
		fff <- 1
		// fff <- 2 // Causes this goroutine to HANG, prints below do not work
		// Therefore an unbuffered channel has an effective capacity of 1, and only first push appears
		for i := 0 ; i < 5 ; i++ {  ggg <- i  }
		// for i := 5 ; i < 12 ; i++ {  ggg <- i  } // Capacity exceeded, same story as above
		// 											   A push that exceeds capacity will hang
		fprintf( "fff, cap: %d, len: %d\n" , cap( fff ) , len( fff ) )
		fprintf( "ggg, cap: %d, len: %d\n" , cap( ggg ) , len( ggg ) )
	}()
	fprint( "Found int:" , <- fff ) // Force main goroutine to wait

	/* fff, cap: 0, len: 0  // The actual capacity is "1"
	   ggg, cap: 10, len: 5
	   Found int: 1 */

	/* Your program will exit when the main() function finishes. 
	This is likely to happen before your goroutine has time to run and print its output. 
	Write data to a channel ch at the end of goroutine and read data from ch out of goroutine 
	can make the main function waiting for goroutine print message.
	( https://stackoverflow.com/questions/16228887/why-does-fmt-println-inside-a-goroutine-not-print-a-line ) */
	

	/* 76. Select Default
	The default case in a select is run if no other case is ready.
	Use a default case to try a send or receive without blocking.  */
	func(){
		tick := time.Tick(  100 * time.Millisecond )
		boom := time.After( 500 * time.Millisecond )
		for { // Infinite While
			select { // Do the possible one
				case <-tick:
					fmt.Println( "tick." )
				case <-boom:
					fmt.Println( "BOOM!" )
					return
				default:
					fmt.Println("    .")
					time.Sleep(50 * time.Millisecond)
			}
		}
	}()

	// 77. Using Mutex
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey") // Schedule 1000 goroutines to add to the same key
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

/********** Utility Functions ********************************************************************/

func say( s string ) {
	for i := 0; i < 5; i++ {
		time.Sleep( 100 * time.Millisecond )
		fprint(s)
	}
}

func get_pi_digit_calculator() ( func() int ){
	// Return a function that returns the digits of pi one at a time
	// https://rosettacode.org/wiki/Pi#Python

	/* Enclosed Variables */
	q, r, t, k, n, l := 1, 0, 1, 1, 3, 3
	var (
		nn, nr int
		rdy    bool
	) 
	rdy = true
	
	return func() ( int ) {
		for {
			if 4*q+r-t < n*t {
				if rdy {  
					rdy = false
					return n  
				}
				nr  = 10*(r-n*t)
				n   = ((10*(3*q+r))/t)-10*n
				q  *= 10
				r   = nr
				rdy = true
			}else{
				nr = (2*q+r)*l
				nn = (q*(7*k)+2+(r*l))/(t*l)
				q  *= k
				t  *= l
				l  += 2
				k += 1
				n  = nn
				r  = nr
			}
		}
	}
}

func pi_CPU_heater( n int ){
	// Calculate `n` digits of pi
	pi_digit := get_pi_digit_calculator()
	for i := 0 ; i < n ; i++ {  pi_digit()  }
}