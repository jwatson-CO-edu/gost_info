package main

import (
    "github.com/jwatson-CO-edu/httpcanvas"
    "math"
    "time"
    "fmt"
)

func app(context *httpcanvas.Context) {

    fmt.Println( "About to start app ..." )
    centerX := context.Width / 2
    centerY := context.Height / 2

    fmt.Println( "About to draw ..." )
    context.BeginPath()
    context.Arc(centerX, centerY, 70, 0, 2*math.Pi, false)
    context.FillStyle("green")
    context.Fill()

    fmt.Println( "About to sleep ..." )
    time.Sleep(5 * time.Second)

    context.LineWidth(5)
    context.StrokeStyle("#003300")
    context.Stroke()

    fmt.Println( "About to show ..." )
    context.ShowFrame()

    fmt.Println( "Ended!" )
}

func main() {
    httpcanvas.ListenAndServe("127.0.0.1:8080", app)
}