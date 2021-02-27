// https://tutorialedge.net/golang/go-webassembly-tutorial/

// To complile to WASM
// GOARCH=wasm GOOS=js go build -o lib.wasm test03.go
//                                 ^output  ^input

package main

import (
	// "flag"
	// "log"
	// "net/http"
	"strconv"
	"syscall/js"
)

func get_js_func_arg( arr []js.Value, i int ) js.Value {
    // Get the `i`th value from argument array `arr`
    return js.Global().Get("document").Call(
        "getElementById", 
        arr[i].String(),
    ).Get("value")
}

func set_elem_value( elem string, val interface{} ) {
    // Get the `elem` and set its 'value' member to `val`
    // https://medium.com/tech-lah/webassembly-part-ii-b-golang-with-wasm-8b3c690221b4
    // https://golang.org/pkg/syscall/js/#ValueOf
    js.Global().Get("document").Call(
        "getElementById", 
        elem,
    ).Set("value", js.ValueOf( val ) )
}

func add(v js.Value, i []js.Value) interface {} {
    println( "About to add ..." )
    value1, _ := strconv.Atoi(  get_js_func_arg(i,0).String()  )
    value2, _ := strconv.Atoi(  get_js_func_arg(i,1).String()  )
    result := value1+value2
    set_elem_value( "a_result", result )
    println( "Done:", result )
    return nil
}

func subtract(v js.Value, i []js.Value) interface {} {
    println( "About to subtract ..." )
    value1, _ := strconv.Atoi(  get_js_func_arg(i,0).String()  )
    value2, _ := strconv.Atoi(  get_js_func_arg(i,1).String()  )
    result := value1-value2
    set_elem_value( "s_result", result )
    println( "Done:", result )
    return nil
}

func registerCallbacks() {
    js.Global().Set("add", js.FuncOf(add))
    js.Global().Set("subtract", js.FuncOf(subtract))
}

func main() {
    
    

    println("WASM Go Initialized")
    // register functions
    registerCallbacks()
    
    // Prevent the function from returning, which is required in a wasm module
	select {} // https://withblue.ink/2020/10/03/go-webassembly-http-requests-and-promises.html
}