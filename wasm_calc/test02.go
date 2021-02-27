// https://tutorialedge.net/golang/go-webassembly-tutorial/

// To complile to WASM
// GOARCH=wasm GOOS=js go build -o lib.wasm test02.go
//                                 ^output  ^input

package main

import (
	// "flag"
	// "log"
	// "net/http"
    "syscall/js"
)

func add(v js.Value, i []js.Value) interface {} {
    js.Global().Set("output", js.ValueOf(i[0].Int()+i[1].Int()))
    println(js.ValueOf(i[0].Int() + i[1].Int()).String())
    return nil
}

func subtract(v js.Value, i []js.Value) interface {} {
    js.Global().Set("output", js.ValueOf(i[0].Int()-i[1].Int()))
    println(js.ValueOf(i[0].Int() - i[1].Int()).String())
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