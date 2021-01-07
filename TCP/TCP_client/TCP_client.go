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
    "strconv"
    "time"
)

/*** Aliases ***/
var fprint  = fmt.Println // Alias for PrintLine
var fprintf = fmt.Printf //- Alias for PrintFormat


/********** Networking Functions *****************************************************************/

func connect_for_a_while( srvrPort string, clntPortBgn string, duration_s int, rateHz int ) string {
	// Connect, send some stuff, then disconnect

    /* Establish a connection */
    connected := false
    
    // 0. While not connected, try ports until one connects
    for {
        // 1. Get the address of the endpoint
        tcpAddr, err := net.ResolveTCPAddr( "tcp4", srvrPort )
        checkError( err ) // FIXME: Need an error check that does NOT crash the program!

        // 3. Attempt to establish a connection
        conn, err := net.DialTCP( "tcp", nil, tcpAddr )
        checkError( err ) // FIXME: Need an error check that does NOT crash the program!
        connected = true

        // 4. If the connection was successful, then break out of the loop
        if connected {  
            break  
        // 5. Else, could not make connection
        }else{
            i, err := strconv.Atoi( srvrPort )
            i++
            srvrPort = strconv.Itoa(i)
        }
    }
    
    /* While there is time remaining, send messages at the specified rate */
    // NOTE: We assume there are not network errors if we have reached this point
    for {

    }

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


func get_elapsed_timer() ( func() float64 ) {
    // Return a function that returns the number of seconds since the closure was created
    bgn := time.Now() // <-- *Enclosed* variable
    // Return a lambda function that returns time since `get_elapsed_timer` returned
	return func() float64 {
		return float64(  time.Since( bgn )  )
	}
}