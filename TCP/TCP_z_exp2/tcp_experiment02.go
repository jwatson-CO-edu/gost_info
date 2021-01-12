package main
/***** MAIN_TEMPLATE.go *****
A_ONE_LINE_DESCRIPTION_OF_THE_FILE
James Watson, 202X-MM

***** DEV PLAN *****
[ ] Server 
[ ] Client
[ ] Test Spam one msg
[ ] Test send/recv loop

Template Ver. 2020-12
*/

/*** Imports ***/
import (
	"fmt" //- Printing
	"net"
	//"os" // Terminal args
)

/*** Aliases ***/
var fprint  = fmt.Println // Alias for PrintLine
var fprintf = fmt.Printf //- Alias for PrintFormat


/********** Structs ******************************************************************************/

/***** NetAgent *****/
type NetAgent struct {
	// Container struct for networking info
	// NOTE: For simplicity, there are no locks and it is assumed there will be no races
	/* Connection */
	conn /*-*/ *net.Conn // Network interface
	addrPtr    *net.Addr // Address object
	host /*-*/ string // -- IP Address
	port /*-*/ string // -- Port number
	vers /*-*/ string // -- Protocol version
	cnnctd     bool // ---- Is connected?
	maxAttempt int // ----- Max number of attempts without giving up
	/* TX/RX */
	sendData []byte //- Data to send
	recvData []byte //- Data received
	netCode  int // --- Error code   from network ops
	err /**/ error // - Error object from network ops
	/* Timing */
	duration_s float64 // How long to keep the connection alive in seconds
	rateHz     float64 // How often to send/recv in Hz
}

func default_connection_spec() NetAgent {
	// Return a `NetAgent` pre-populated with default connection data
	return NetAgent{
		host: "127.0.0.1",
		port: "8000",
		vers: "tcp6",
	}
}

/***** Server *****/

func ( ntwkAgnt *NetAgent ) Listen() {
	// Open the `NetAgent` port for listening
}

func ( ntwkAgnt *NetAgent ) Accept() {
	// Accept one connection (blocking)
}

func ( ntwkAgnt *NetAgent ) srvr_handle_conn() {
	// 
}

/***** Client *****/

func ( ntwkAgnt *NetAgent ) clnt_resolve_and_dial() {
	
}

func ( ntwkAgnt *NetAgent ) clnt_handle_conn() {
	
}


/********** MAIN *********************************************************************************/
func main(){ /*Terminal args*/ //progArgs := os.Args[1:]

	// 1. Init server

	// 2. Begin server loop

	// 3. Init client

	// 4. Begin client loop

}