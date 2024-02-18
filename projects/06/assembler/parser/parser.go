package parser

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const (
	CmdTypeA = "a-instruction"
	CmdTypeC = "c-instruction"
	CmdTypeL = "label"
	eol      = "\n"
)

type Parser struct {
	file    *os.File
	scanner *bufio.Scanner
}

func New(inputFile string) *Parser {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	return &Parser{
		file:    file,
		scanner: bufio.NewScanner(file),
	}
}

func (p *Parser) Close() {
	p.file.Close()
}

func IsVariable(s string) bool {
	_, err := strconv.ParseInt(s, 10, 16)
	return err != nil
}

func CommandType(command string) string {
	if strings.HasPrefix(command, "@") {
		return CmdTypeA
	} else if strings.HasPrefix(command, "(") && strings.HasSuffix(command, ")") {
		return CmdTypeL
	} else {
		return CmdTypeC
	}
}


func CommandArgs(s string) (dest, comp, jump string) {
	dest,comp,jump = "", "", ""

	switch CommandType(s) {
	case CmdTypeL:
		dest = s[1 : len(s)-1]
	case CmdTypeA:
		dest = s[1:]
	case CmdTypeC:
		compInd := strings.Index(s, "=")
		jumpInd := strings.Index(s, ";")

		if jumpInd != -1 {
			jump = s[jumpInd+1:]
		} else {
			jumpInd = len(s)
		}

		if compInd == -1 {
			comp = s[:jumpInd]
		} else {
			dest = s[:compInd]
			comp = s[compInd+1 : jumpInd]
		}
	}
	return
}

func (p *Parser) Parse() (string, bool) {
	nextLine := ""

	ok := p.scanner.Scan()

	if !ok {
		return nextLine, false
	}

	nextLine = p.scanner.Text()

	if comment := strings.Index(nextLine, "//"); comment > -1 {
		nextLine = strings.TrimSpace(nextLine[:comment])
	} else {
		nextLine = strings.TrimSpace(nextLine)
	}

	if len(nextLine) == 0 {
		return p.Parse()
	}

	return nextLine, true
}
