package main

import (
	"fmt"
	"gopkg.in/pipe.v2"
	"strings"
)

func test(s string)  {
	p1 := pipe.Line(
		pipe.ReadFile(fmt.Sprintf("input%s.txt", s)),
		pipe.Exec("./palindrome"),
	)
	p2 := pipe.Line(
		pipe.ReadFile(fmt.Sprintf("output%s.txt", s)),
	)
	got, err := pipe.CombinedOutput(p1)
	expected, err := pipe.CombinedOutput(p2)

	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("got=%s expected=%s\n", strings.TrimSpace(string(got)), strings.TrimSpace(string(expected)))
}

func main() {
	for i := 0; i < 3; i++ {
		test(fmt.Sprintf("00%d", i))
	}
}
