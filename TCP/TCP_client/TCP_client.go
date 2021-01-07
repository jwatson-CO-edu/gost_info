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

func connect_for_a_while( srvrPort string, clntPortBgn string, duration_s float64, rateHz int ) string {
	// Connect, send some stuff, then disconnect

    var (
        conn /**/ *net.TCPConn
        err /*-*/ error
        tcpAddr   *net.TCPAddr
        connected bool
        result    []byte
    )

    /* Establish a connection */
    connected = false
    
    // 0. While not connected, try ports until one connects
    for {
        // 1. Get the address of the endpoint
        tcpAddr, err = net.ResolveTCPAddr( "tcp4", srvrPort )
        if !checkError( err ) {
            // 3. Attempt to establish a connection
            conn, err = net.DialTCP( "tcp", nil, tcpAddr )
            connected = !checkError( err ) 
        }

        

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
    elapsed := get_elapsed_timer()
    for {

        // 6. Write some stuff 
        _, err = conn.Write(  []byte( "HEAD / HTTP/1.0\r\n\r\n" )  )
        checkError( err )

        // 7. Read anything that comes back
        result, err = ioutil.ReadAll( conn )
        checkError( err )

        // 8. Print what comes back
        fmt.Println(  string( result )  )


        // 10. If the timer has run out, then stop send/recv
        if elapsed() > duration_s {  break  }
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

	

	

	

    

	

	// 7. Quit
    os.Exit(0)
}

func checkError( err error ) bool {
	// If there was an error, then return true, otherwise return false
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        return true
    }
    return false
}


func get_elapsed_timer() ( func() float64 ) {
    // Return a function that returns the number of seconds since the closure was created
    bgn := time.Now() // <-- *Enclosed* variable
    // Return a lambda function that returns time since `get_elapsed_timer` returned
	return func() float64 {
		return float64(  time.Since( bgn )  )
	}
}