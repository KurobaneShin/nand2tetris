package code

import (
	"fmt"
	"os"
)

type Code struct {
  file *os.File
}

func New(outputFile string) *Code {
	file, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	return &Code{
    file: file,
  }
}

func (c *Code) Close() {
  c.file.Close()
}

func writeln(file *os.File, s string) {
 _, err :=	file.WriteString(s + "\n")
  if err != nil {
    panic(err)
  }
}

func (c *Code) WriteAInstruction( value int) {
	binaryToStringWithPad0 := fmt.Sprintf("%015b", value)

	trimmedBinary := binaryToStringWithPad0[len(binaryToStringWithPad0)-15:]

	hackInstructionWithPrefix0 := "0" + trimmedBinary
	writeln(c.file, hackInstructionWithPrefix0)
}

func (c *Code) WriteCInstruction( dest, comp, jump string) {
	if dest == "" {
		dest = "null"
	}

	if jump == "" {
		jump = "null"
	}

	hackInstruction := fmt.Sprintf("111%s%s%s", computations[comp], destinations[dest], jumps[jump])
	writeln(c.file, hackInstruction)
}
