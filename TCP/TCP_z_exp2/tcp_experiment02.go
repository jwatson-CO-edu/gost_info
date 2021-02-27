package main
/***** MAIN_TEMPLATE.go *****
A_ONE_LINE_DESCRIPTION_OF_THE_FILE
James Watson, 202X-MM

***** DEV PLAN *****
[Y] Server - 2021-01-13, Packaged connection info in a single struct
[Y] Client - 2021-01-13, Packaged connection info in a single struct
[ ] Test Spam one msg
	{ } Server spams, Client reads
	{ } Client spams, Server reads
[ ] Test send/recv loop

Template Ver. 2020-12
*/

/*** Imports ***/
import (
	"fmt" //- Printing
	"net"
	"time"
	//"os" // Terminal args
)

/*** Aliases ***/
var fprint  = fmt.Println // Alias for PrintLine
var fprintf = fmt.Printf //- Alias for PrintFormat


/********** Structs ******************************************************************************/

/***** time.Duration *****/
// TODO: SEND TO GOSTBUSTR

func dur_subtract( dr1 *time.Duration, dr2 *time.Duration ) time.Duration {
	// Return a time.Duration that represents "dr1 - dr2"
	return time.Duration( dr1.Nanoseconds() - dr2.Nanoseconds() )
}

func get_expired_timer( duration_s float64 ) ( func() bool ) { // TODO: SEND TO GOSTBUSTR
	// Return a function that returns the number of seconds since the closure was created
	// *Enclosed* variables
	bgn     := time.Now() 
	timeOut := duration_s
    // Return a lambda function that returns time since `get_elapsed_timer` returned
	return func() float64 {
		return float64(  time.Since( bgn ).Seconds()  ) >= timeOut
	}
}

/***** Pulse *****/
// TODO: SEND TO GOSTBUSTR

type Pulse struct {
	// Attempt to keep a loop at `freq_Hz`, but no faster. ROS-style loop rate limiter
	freq_Hz	 float64 // - Max allowed frequency
	period_s time.Duration // - Sleep time that prevents us from exceeding frequency
	lastTime time.Time // Time when last checked
}

func GetPulseWithHz( rate_Hz float64 ) Pulse {
	// Factory method returns a Pulse that is ready to keep time
	return Pulse{
		freq_Hz:  rate_Hz, // --- Max loop rate requested by client code
		period_s: time.Duration(  int( 1.0/rate_Hz * 1e9  )  ), // Period to be used by sleep, inverse of frequency
		lastTime: time.Now(), //- Use creation time as initial loop time, (likely results in zero sleep first loop)
	}
}

func ( puls *Pulse ) Regulate() bool {
	// Sleep for an amount of time that maintains loop `freq_Hz` or *lower*
	// 0. Init calcs
	now     := time.Now() // ------------------------------------ The time this function was called
	slept   := false // ----------------------------------------- Did we have to sleep to regulate?
	elapsed := now.Sub( puls.lastTime ) // Duration since the last function call
	// 1. If the period has not elapsed, then we must sleep to prevent runaway freq
	if elapsed.Seconds() < puls.period_s.Seconds(){
		slpDur := dur_subtract( &(puls.period_s) , &elapsed )
		time.Sleep( slpDur )
		slept = true
	} // else there is no need to sleep, the period has already past and runaway not possible
	// 2. Mark time function was called
	puls.lastTime = now
	// 3. Tell client code if sleep was needed
	return slept
}

/***** NetAgent *****/
// Basic container struct with simple methods for server and client socket connections

type NetAgent struct {
	// Container struct for networking info
	// NOTE: For simplicity, there are no locks and it is assumed there will be no races
	/* Connection */
	conn /*-*/ net.Conn // Network interface
	addr /*-*/ net.Addr // Address object
	lstnr      net.Listener
	host /*-*/ string // -- IP Address
	port /*-*/ string // -- Port number
	vers /*-*/ string // -- Protocol version
	cnnctd     bool // ---- Is connected?
	maxAttempt int // ----- Max number of attempts without giving up
	/* TX/RX */
	sendData   []byte //- Data to send
	recvData   []byte //- Data received
	netCode    int // --- Error code   from network ops
	err /*--*/ error // - Error object from network ops
	lastMsgLen int
	/* Timing */
	duration_s float64 // How long to keep the connection alive in seconds
	rateHz     float64 // How often to send/recv in Hz
}

func ( ntwkAgnt NetAgent ) Get_Full_Addr() string {
	// Get the "HOST:PORT" string address where the connection will be made
	return ntwkAgnt.host + ":" + ntwkAgnt.port
}

func default_connection_spec() NetAgent {
	// Return a `NetAgent` pre-populated with default connection data
	return NetAgent{
		host: "127.0.0.1",
		port: "8000",
		vers: "tcp6",
	}
}

// TODO: Wrap read and write so that the total bytes received and sent are logged
// TODO: Wrap connect for both server and client to respect the number of allowed attempts

/***** Server *****/

func ( ntwkAgnt *NetAgent ) Listen() {
	// Open the `NetAgent` port for listening
	ntwkAgnt.lstnr, ntwkAgnt.err = net.Listen( ntwkAgnt.vers, ntwkAgnt.Get_Full_Addr() ) // Return the connection and an error code
}

func ( ntwkAgnt *NetAgent ) Accept() {
	// Accept a connection requested by the client (blocking)
	ntwkAgnt.conn, ntwkAgnt.err = ntwkAgnt.lstnr.Accept()
}

/** App-Specific **/

func ( ntwkAgnt *NetAgent ) srvr_handle_conn_TEST1( ) {
	// SERVER: Server spams, Client reads
	// NOTE: This function assumes that a connection has already been accepted
	/* Timing */
	timOut := get_expired_timer( ntwkAgnt.duration_s )
	puls   := GetPulseWithHz( ntwkAgnt.rateHz )
	msg    := []byte( "010101" )

	for ; !( timOut() ) ; {
		ntwkAgnt.lastMsgLen, ntwkAgnt.err = ntwkAgnt.conn.Write( msg )
		puls.Regulate()
	}
}

/***** Client *****/

func ( ntwkAgnt *NetAgent ) clnt_resolve_and_dial() {
	// Get connection info and attempt a connection
	// 1. Get the address of the endpoint && Attempt to establish a connection
	// ntwkAgnt.addr, ntwkAgnt.err = net.ResolveTCPAddr( "tcp", ntwkAgnt.Get_Full_Addr() )
	ntwkAgnt.conn, ntwkAgnt.err = net.Dial( ntwkAgnt.vers, ntwkAgnt.Get_Full_Addr() )
}

/** App-Specific **/

func ( ntwkAgnt *NetAgent ) clnt_handle_conn() {
	// Server spams, Client reads
	timOut := get_expired_timer( ntwkAgnt.duration_s )
	puls   := GetPulseWithHz( ntwkAgnt.rateHz )

	for ; !( timOut() ) ; {
		ntwkAgnt.lastMsgLen, ntwkAgnt.err = ntwkAgnt.conn.Read( ntwkAgnt.recvData )
		puls.Regulate()
	}
}


/********** MAIN *********************************************************************************/
func main(){ /*Terminal args*/ //progArgs := os.Args[1:]

	// 1. Init server
	srvr := default_connection_spec()
	srvr.Listen()

	// 3. Init client
	clnt := default_connection_spec()
	clnt.clnt_resolve_and_dial()

	// 2. Begin server loop
	srvr.srvr_handle_conn_TEST1()

	// 4. Begin client loop
	clnt.clnt_handle_conn()

}