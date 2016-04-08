package main

import "fmt"
import "math"

const s string = "hard constant"


func main(){

	//printing the constant
	fmt.Println(s)

	const n = 500000000


	const d = 3e20/n
	fmt.Println(d)

	// a numeric constant has no type untill it's given one explicit or figured out by context ( function arg )
	fmt.Println(int64(d))

	fmt.Println(math.Sin(d))
}
