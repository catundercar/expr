package expr

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	tokens, err := Parse("(1+1)+22*3/4")
	if err != nil {
		t.Fatal(err)
	}
	for _, tk := range tokens {
		fmt.Println(tk)
	}
}
