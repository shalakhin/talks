// +build ignore,OMIT
package main

import (
	"fmt"
)

// interface OMIT
type SoftwareEngineer interface {
	Code()
}

func writeCode(s SoftwareEngineer) {
	s.Code()
}

// end interface OMIT

// example OMIT
// Pythonist ...
type Pythonist struct {
	pythonVersion string
}

// Code ...
func (p Pythonist) Code() {
	fmt.Println("import this")
}

// end example OMIT

// START main OMIT
func main() {
	pythonist := Pythonist{"3.4"}
	writeCode(pythonist)
}

// END main OMIT
