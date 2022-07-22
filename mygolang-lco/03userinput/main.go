package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza: ")

	// variable, action error := action
	input, _ := reader.ReadString('\n')
	fmt.Println("You rating was: ", input)
	fmt.Printf("Rating type: %T \n", input)

}
