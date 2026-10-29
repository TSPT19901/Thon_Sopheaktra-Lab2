package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var text string

	fmt.Print("Enter text: ")
	fmt.Scan(&text)

	//to convert to hex, we convert by division decimal by 2
	fmt.Print("Binary: ")

	for i := 0; i < len(text); i++ {
		n := int(text[i])
		binary := ""

		for n > 0 {
			remainder := n % 2
			binary = fmt.Sprint(remainder) + binary
			n = n / 2
		}

		for len(binary) < 8 {
			binary = "0" + binary
		}

		fmt.Print(binary, " ")
	}
	fmt.Println()

	//to convert to hex, we convert by division decimal by 16
	fmt.Print("Hexadecimal: ")

	for i := 0; i < len(text); i++ {
		n := int(text[i])
		hex := ""

		for n > 0 {
			remainder := n % 16
			digit := fmt.Sprintf("%X", remainder)
			hex = digit + hex
			n = n / 16
		}

		for len(hex) < 2 {
			hex = "0" + hex
		}

		fmt.Print(hex, " ")
	}
	fmt.Println()

	//base64
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	fmt.Println("Base64:", encoded)
}

