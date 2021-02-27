/* https://opensource.com/article/18/5/building-concurrent-tcp-server-go
[ ] Track number of open connections
[ ] Report number og running goroutines
*/
package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const MIN =   1 // Min random number
const MAX = 100 // Max random number


/********** Networking Functions *****************************************************************/

func handleConnection( c net.Conn ){
	// Do the server work once the connection `c` has been established
	fmt.Printf( "Serving %s\n", c.RemoteAddr().String() ) // Print the connected address

	var (
		netData []byte
		netCode int
		err     error
	)

	for { // Infinite while

		//  1. Read from the connection until the next newline
		netCode, err = c.Read( netData )
		// netData, err := bufio.NewReader(c).ReadString('\n')
		fmt.Println( "Got:" , netData , "," , netCode )
		
		//  2. If there was an error, print it and return
		if err != nil {
			fmt.Println(err)
			return
		}

		//  3. Else, strip whitespace from the ends of the string
		temp := strings.TrimSpace( string( netData ) )
		
		//  4. If the client sent a "STOP" command, then break out of the loop
		if temp == "STOP" {  break  }
		
		//  5. Convert a random number to a string
		result := strconv.Itoa( random() ) + "\n"
		
		//  6. Convert a string to a bytestring and sent it back over the connection
		c.Write( []byte( string( result ) ) )
	}
	//  7. Close the connection when the loop exits
	c.Close()
}

func server_report() bool {
	// https://golangbyexample.com/number-currently-running-active-goroutines/
	return true
}


/********** Main *********************************************************************************/

func main() {
	// Run the server
	//  1. Fetch the terminal arguments
	arguments := os.Args
	//  2. Warn the user if there are too few arguments
	if len( arguments ) == 1 {
		fmt.Println( "Please provide a port number!" )
		return
	}
	//  3. Prepend a colon to the port
	PORT := ":" + arguments[1]
	//  4. Listen at the specified port
	l, err := net.Listen( "tcp", PORT ) // Return the connection and an error code
	//  5. If there was an error, report it and end program
	if err != nil {
		fmt.Println(err)
		return
	}
	//  6. After `main` exits, make sure to close the connection
	defer l.Close()
	//  7. Use the clock to 
	rand.Seed( time.Now().Unix() )
	//  8. Infinite while
	for {
		//  9. Accept a connection requested by the client
		c, err := l.Accept()
		// 10. If an error occurred, then print it
		if err != nil {
			fmt.Println( err )
			return
		}
		// 11. Handle the connection in a separate goroutine
		go handleConnection( c )
	}
}


/********** Utility Functions ********************************************************************/

func random() int { // TODO: SEND TO GOSTBUSTR
	// Get a random int between `MIN` and `MAX`
	return rand.Intn( MAX - MIN ) + MIN
}