package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write i  ", i, "as ")
	switch i {
	case 1 :
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's Weekend!")
	default :
		fmt.Println("It's a Week Day")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("It's before Noon")
	case t.Hour() > 12 :
		fmt.Println("It's After Noon")
	}

}
