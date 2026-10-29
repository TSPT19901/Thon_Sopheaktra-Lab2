package main

import "fmt"

func main() {
   	var a, b int
	fmt.Printf("Enter value of a: ")
	fmt.Scan(&a)

	fmt.Printf("Enter value of b: ")
	fmt.Scan(&b)

	bothpositive:= (a>0) && (b>0)
	fmt.Printf("Both a and b are positive:  %t\n" , bothpositive)

	onegreater:= (a>b) || (b>a)
	fmt.Printf("One greater than the other:  %t\n" , onegreater	)

	notequal:= !(a==b)
	fmt.Printf("a and b are not equal:  %t\n"  , notequal)
}