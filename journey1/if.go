package main

import "fmt"

func main(){
	if 7 % 2 == 0 {
		fmt.Println("7 is even")

	} else {
		fmt.Println("7 is odd")

	}

	if 8 % 2 == 0 {
		fmt.Println("8 is divisible by 4.")

	}

	if num := 10 ; num < 0 {
		fmt.Println("Num is Negative")

	}else if pnum:=20 ; pnum < num {
		fmt.Println("The PNUM available here is less than the num.")
	}else {
		fmt.Println("Has Multiple Digits.")
	}
}
