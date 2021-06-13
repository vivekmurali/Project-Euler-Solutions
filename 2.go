package main

import "fmt"

func main() {
	lim := 4_000_000
	_ = lim
	sum := 0
	a := 0
	b := 1
	for sum <= lim {
		temp := a
		a = b
		b = temp + a
		if b%2 == 0 {
			sum += b
		}
	}
	fmt.Println(sum)
}
