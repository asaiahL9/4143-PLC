package main

import "fmt"

func color_text() {
	fmt.Println("\033[31mRed text\033[0m")
	fmt.Println("\033[32mGreen text\033[0m")
}
