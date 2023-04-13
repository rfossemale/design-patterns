package main

import (
	"fmt"
	"patterns/abstractfactory"
	"patterns/adapter"
	"patterns/builder"
	"patterns/prototype"
)

func main() {
	fmt.Println("######")
	abstractfactory.Run()
	fmt.Println("######")
	builder.Run()
	fmt.Println("######")
	prototype.Run()
	fmt.Println("######")
	Adapter.Run()
}

// go mod edit -replace example.com/greetings=../greetings
