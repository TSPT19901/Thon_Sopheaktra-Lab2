package main

import "fmt"

// func getinfo(){

// 	var a, b int
// 	fmt.Printf("Enter value of a: ")
// 	fmt.Scan(&a)

// 	fmt.Printf("Enter value of b: ")
// 	fmt.Scan(&b)
// }

func main() {
    var a, b, c int
	fmt.Println("===== Mini Calculator =====")

	for true{
	fmt.Printf("1) Add 2) Sub 3) Mul 4) Div 5) Mod 6) Exit\n")
	fmt.Printf("Choose:")
	fmt.Scan(&c)
	if c == 6 {
			fmt.Println("Goodbye!")
			break
		}
	if c < 1 || c > 5 {
		fmt.Println("Invalid choice! Please choose 1-6.")
		continue
	}
	fmt.Printf("Enter value of a: ")
	fmt.Scan(&a)

	fmt.Printf("Enter value of b: ")
	fmt.Scan(&b)
	
	result:=a

	if (c>=1 && c<=5){
	switch c{
	case 1:
		result+=b
	case 2:
		result-=b
	case 3:
		result*=b
	case 4:
		if b!=0{
			result /=b 
		} else {
			fmt.Printf("Cannot do /= here because b is zero ")
			continue
		}
	case 5:
		if b!=0{
			result %=b 
		} else {
			fmt.Printf("Cannot do %= here because b is zero ")
			continue
		}
	}
	fmt.Println("Result:" , result)
	}
	}
}