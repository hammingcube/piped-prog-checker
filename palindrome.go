package main

import "fmt"

func memoized() func(string) int {
	cache := make(map[string]int)
	var fn func(string) int
	fn = func(s string) int {
		if _, ok := cache[s]; ok {
			return cache[s]
		}
		if len(s) < 2 {
			return 0
		}
		a := fn(s[1:]) + 1
		b := fn(s[:len(s)-1]) + 1
		c := len(s)
		if s[0] == s[len(s)-1] {
			c = fn(s[1:len(s)-1])
		}
		min := func(x, y int) int {
			if x < y {
				return x
			} else {
				return y
			}
		}
		m := min(a,b)
		m = min(m,c)
		cache[s] = m
		return m
	}
	return fn
}

func main() {
	f := memoized()
	var s string
	fmt.Scanf("%s", &s)
	fmt.Println(f(s))
}
