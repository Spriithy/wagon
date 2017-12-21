package wast

import (
	"fmt"
	"testing"
)

func TestScanner(t *testing.T) {
	s := NewScanner("test.wast")
	if len(s.Errors) > 0 {
		fmt.Println(s.Errors[0])
		return
	}

	var tok *Token
	tok = s.NextToken()
	for tok.Kind != EOF {
		fmt.Printf("%d:%d %s\n", tok.Line, tok.Column, tok.String())
		tok = s.NextToken()
	}
	fmt.Printf("%d:%d %s\n", tok.Line, tok.Column, tok.String())

	for _, err := range s.Errors {
		fmt.Print(err)
	}

	if len(s.Errors) > 0 {
		t.Errorf("wast: failed with %d errors", len(s.Errors))
	}
}
