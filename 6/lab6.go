package main

import "fmt"

func xorEncrypt(text string, key byte) string {
	result := ""

	for i := 0; i < len(text); i++ {
		result += string(text[i] ^ key)
	}

	return result
}

func main() {
	var choice int
	var text string
	var key int

	fmt.Println("===== XOR Encryption/Decryption =====")
	fmt.Println("1. Encrypt")
	fmt.Println("2. Decrypt")
	fmt.Print("Choose an option (1 or 2): ")
	fmt.Scan(&choice)

	if choice != 1 && choice != 2 {
		fmt.Println("Invalid choice!")
		return
	}

	fmt.Print("Enter text: ")
	fmt.Scan(&text)

	fmt.Print("Enter key (0-255): ")
	fmt.Scan(&key)

	if key < 0 || key > 255 {
		fmt.Println("Invalid key! Please enter a number from 0 to 255.")
		return
	}

	if choice == 1 {
		result := xorEncrypt(text, byte(key))
		fmt.Println("Encrypted text:", result)
	} else {
		result := xorEncrypt(text, byte(key))
		fmt.Println("Decrypted text:", result)
	}
}