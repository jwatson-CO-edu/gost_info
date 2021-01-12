package main
/***** MAIN_TEMPLATE.go *****
A_ONE_LINE_DESCRIPTION_OF_THE_FILE
James Watson, 202X-MM

***** DEV PLAN *****

** Basic Implementation **
[ ] Gather && Scan DL class notes --> Add to repo
[ ] Forward pass
[ ] Backwards pass
[ ] Layer interface
[ ] MNIST test 1
[ ] Multiple, fully-connected layers
	{ } Send
	{ } Receive

** Arbitrary Connections **
[ ] Unit "Grid" / Ganglion
	- Option for grid, begin with this
	- Supports arbitrary 
[ ] Connection agent
	* Operate On a Grid *
	[ ] Connect
	[ ] Prune
	* Grow Tissue *
	[ ] Spawn layer: When does a layer represent too much?
	[ ] Erase layer: When does a layer represent too little?
[ ] Ganglion export
	

Template Ver. 2020-12
*/

/*** Imports ***/
import (
	"fmt" //- Printing
	//"os" // Terminal args
)

/*** Aliases ***/
var fprint  = fmt.Println // Alias for PrintLine
var fprintf = fmt.Printf //- Alias for PrintFormat

/********** Perceptron ***************************************************************************/

type PerceptronLayer struct {
	// The simplest perceptron layer with arbitrary connections
	inRefs  []*PerceptronLayer // Incoming connections
	i /*-*/ []float64 // -------- Input array
	w /*-*/ [][]float64 // ------ Layer refs
	o /*-*/ []float64 // -------- Output array
	outRefs []*PerceptronLayer // Outgoing connections
}

func ( PL *PerceptronLayer ) forward(){
	// Forward pass, predict
}

func ( PL *PerceptronLayer ) reverse(){
	// Backprop
}

func ( PL *PerceptronLayer ) transmit_fwd(){
	// Backprop
}

/********** MAIN *********************************************************************************/
func main(){ /*Terminal args*/ //progArgs := os.Args[1:]

}