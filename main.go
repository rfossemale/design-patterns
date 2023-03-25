package main

import (
	"fmt"
	"patterns/abstractfactory"
	"patterns/builder"
)

func main() {
	fmt.Println("hello")
	abstractfactory.Run()
	builder.Run()
}
