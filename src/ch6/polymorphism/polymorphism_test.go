package polymorphism

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

type JavaProgrammer struct {
}

func (gp *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

func (jp *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.println(\"Hello,World\")"
}

func WriteFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestClient(t *testing.T) {
	goProg := new(GoProgrammer)
	javaProg := new(JavaProgrammer)

	WriteFirstProgram(goProg)
	WriteFirstProgram(javaProg)
}
