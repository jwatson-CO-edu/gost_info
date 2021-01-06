/* https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/socket/tcp_sockets.html
[ ] Track number of open connections
[ ] Report number og running goroutines
*/

package main

import (
    "net"
    "os"
    "fmt"
    "io/ioutil"
)

/*** Aliases ***/
var fprint  = fmt.Println // Alias for PrintLine
var fprintf = fmt.Printf //- Alias for PrintFormat


/********** Networking Functions *****************************************************************/

func connect_for_a_while( srvrPort string, clntPortBgn string, duration_s int, rateHz int ){
	// Connect, send some stuff, then disconnect

	/* Establish a connection */

	/* While there is time remaining, send messages at the specified rate */

	// https://stackoverflow.com/a/56336811
}

func client_report() bool {
	// https://golangbyexample.com/number-currently-running-active-goroutines/
	return true
}


/********** Main *********************************************************************************/

func main() {
	// 1. Warn the user if they have not provided a port
    if len( os.Args ) != 2 {
        fmt.Fprintf( os.Stderr, "Usage: %s host:port ", os.Args[0] )
        os.Exit(1)
    }
    service := os.Args[1]

	// 2. Get the address of the endpoint
    tcpAddr, err := net.ResolveTCPAddr( "tcp4", service )
    checkError( err )

	// 3. Attempt to establish a connection
    conn, err := net.DialTCP( "tcp", nil, tcpAddr )
    checkError( err )

	// 4. Write some stuff 
    _, err = conn.Write(  []byte( "HEAD / HTTP/1.0\r\n\r\n" )  )
    checkError( err )

    // 5. Read anything that comes back
    result, err := ioutil.ReadAll( conn )
    checkError( err )

	// 6. Print what comes back
    fmt.Println(  string( result )  )

	// 7. Quit
    os.Exit(0)
}

func checkError( err error ) {
	// If there was an error, then print it
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}