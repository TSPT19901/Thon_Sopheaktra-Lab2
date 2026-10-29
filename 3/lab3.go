package main 

import "fmt"


func useAND(a , b int){
	fmt.Println("AND of a and b:" , a&b)
}

func useOR(a , b int){
	fmt.Println("OR of a and b:" , a|b)
}

func useXOR(a , b int){
	fmt.Println("XOR of a and b:" , a^b)
}

func useNOT(a , b int){
	fmt.Println("NOT of a:" , ^a)
	fmt.Println("NOT of b:" , ^b)
}


func main(){
   a:=10
   b:=10

   useAND(a, b)
   useNOT(a,b)
   useOR(a,b)
   useXOR(a,b)

   //Assigment operator
	c:=15
	c>>=2
	fmt.Println("c>>=2:" , c)

	c<<=2
	fmt.Println("c<<=2:" , c)

	c&=2
	fmt.Println("c&=2:" , c)

	c|=2
	fmt.Println("c|=2:" , c)

	c^=2
	fmt.Println("c^=2:" , c)

}