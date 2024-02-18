package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
)

// filterNewLines filters out new line chars
func filterNewLines(s string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case 0x000A, 0x000B, 0x000C, 0x000D, 0x0085, 0x2028, 0x2029:
			return -1
		default:
			return r
		}
	}, s)
}

// CompareStrings two strings ignoring whitespaces,
// returns true if equal
func CompareStrings(a, b string) bool {
	return filterNewLines(a) == filterNewLines(b)
}

// For each file in /asm dir compile hack file
// and comapre with corresponding one in /compare dir

func TestOuput(t *testing.T) {
	var wg sync.WaitGroup

	asmFiles, err := filepath.Glob("test/asm/**/*")
	if err != nil {
		t.Error(err)
	}
	compareFiles, err := filepath.Glob("test/compare/**/*")
	if err != nil {
		t.Error(err)
	}

	for i, f := range asmFiles {
		wg.Add(1)
		go func(i int, f string) {
			defer wg.Done()
			outFile := fmt.Sprintf("test/out_test_%d.hack", i)
			defer os.Remove(outFile)

			Compile(f, outFile)

			a, err := os.ReadFile(compareFiles[i])
			if err != nil {
				t.Error(err)
			}
			b, err := os.ReadFile(outFile)
			if err != nil {
				t.Error(err)
			}

			if !CompareStrings(string(a), string(b)) {
				t.Error("Error processing file", f)
			}
		}(i, f)
	}

	wg.Wait()
}
