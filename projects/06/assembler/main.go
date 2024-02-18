package main

import (
	"os"
	"strconv"

	"github.com/KurobaneShin/nand2tetris-parser-golang/code"
	"github.com/KurobaneShin/nand2tetris-parser-golang/parser"
	"github.com/KurobaneShin/nand2tetris-parser-golang/symbletable"
)

func main() {
	asmSource := os.Args[1]
	hackOutput := os.Args[2]

	println("Assembling " + asmSource + " to " + hackOutput)
  Compile(asmSource, hackOutput)
}


func Compile(src string, out string) {
	var lines []string        // lines slice
	st := symbletable.New()   // new symbol table
	p := parser.New(src)      // new parser
	cw := code.New(out)

	defer p.Close()
	defer cw.Close()

	// First run:
	hasMore := true
	for hasMore {
		c, ok := p.Parse()
		hasMore = ok
		if ok {
			if parser.CommandType(c) != parser.CmdTypeL {
				lines = append(lines, c)
			} else {
				label, _, _ := parser.CommandArgs(c)
				st[label] = len(lines)
			}
		}
	}

	// Second run:
	defaultCustomVarLen := 16
	for _, l := range lines {
		cmdType := parser.CommandType(l)
		dest, comp, jump := parser.CommandArgs(l)

		switch cmdType {
		
		case parser.CmdTypeA:
			if parser.IsVariable(dest) {
				// if variable, use from table
				_, found := st[dest]
				if !found { // if not declared yet, put to table
					st[dest] = defaultCustomVarLen
					defaultCustomVarLen++
				}
				code.WriteAInstruction(cw,st[dest])
			} else {
				// if value, use directly
				aInt, err := strconv.Atoi(dest)
				if err != nil {
          panic(err)
        }
				code.WriteAInstruction(cw,aInt)
			}
		// process C instruction
		case parser.CmdTypeC:
			code.WriteCInstruction(cw,dest, comp, jump)
		}
	}

  println("Assembling done")

}
