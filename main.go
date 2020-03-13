package main

import (
	"fmt"
	// "time"
	// "github.com/SasukeBo/examples/usegorm"
	// "github.com/SasukeBo/examples/learnreflect"
)

func main() {
	// learnreflect.Example1()
	// learnreflect.Example2()
	// learnreflect.Example3()
	testMapNilField()
}

func testMapNilField() {
	m := make(map[string]bool)
	if !m["hello"] {
		fmt.Printf("false: %v\n", m["hello"])
	}
}
