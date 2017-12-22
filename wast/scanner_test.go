package wast

import (
	"flag"
	"fmt"
	"testing"
)

var wastFile = flag.String("wast-file", "", "the wast file to test")

func TestScanner(t *testing.T) {
	if *wastFile == "" {
		t.Errorf("error: no input file")
	}

	s := NewScanner(*wastFile)
	if len(s.Errors) > 0 {
		fmt.Println(s.Errors[0])
		return
	}

	var tok *Token
	tok = s.Next()
	for tok.Kind != EOF {
		fmt.Printf("%d:%d %s\n", tok.Line, tok.Column, tok.String())
		tok = s.Next()
	}
	fmt.Printf("%d:%d %s\n", tok.Line, tok.Column, tok.String())

	for _, err := range s.Errors {
		fmt.Print(err)
	}

	if len(s.Errors) > 0 {
		t.Errorf("wast: failed with %d errors", len(s.Errors))
	}
}
