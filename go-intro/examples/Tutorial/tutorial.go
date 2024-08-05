package main

import (
	"fmt"
	"math"
)

const s string = "constant"



func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f:= "apple"
	fmt.Println(f)

	fmt.Println("-----------------------------------------------------------------------------------------")


	fmt.Println()

	fmt.Println(s)

	const n = 500000000

	const x = 3e20 / n

	fmt.Println(x)

	fmt.Println(int64(x))

	fmt.Println(math.Sin(n))

	fmt.Println("-----------------------------------------------------------------------------------------")

	i:=1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}
	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
