package main

import "fmt"

func main() {
    var a, b int
	fmt.Printf("Enter value of a: ")
	fmt.Scan(&a)

	fmt.Printf("Enter value of b: ")
	fmt.Scan(&b)

	result := a 
	fmt.Printf("After set result = a value:  %d\n", result ) 

	result -=b
	fmt.Printf("After result +=b:  %d\n" , result)

	result +=b
	fmt.Printf("After result -=b:  %d\n", result )

	result *= b 
	fmt.Printf("After result *=b:  %d\n" , result)

	if b!=0{
		result /=b 
	fmt.Printf("After result /=b :  %d\n" , result)

	result %=b
	fmt.Printf("After result %%=b:  %d\n" , result)
	} else {
		fmt.Printf("Cannot do /= and %= here because b is zero ")
	}
}