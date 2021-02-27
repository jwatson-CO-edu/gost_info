// https://tutorialedge.net/golang/go-webassembly-tutorial/
// trm:<dir>$ go run sever.go

package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

var (
	port    = 8080
	portNam = ":" + strconv.FormatInt( int64( port ) , 10)
    listen  = flag.String("listen", portNam, "listen address")
    dir     = flag.String("dir", ".", "directory to serve")
)

func main() {
    flag.Parse()
    log.Printf("listening on %q...", *listen)
    log.Fatal(http.ListenAndServe(*listen, http.FileServer(http.Dir(*dir))))
}