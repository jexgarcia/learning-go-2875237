package main

import (
	"fmt"
<<<<<<< Updated upstream
=======
	"os"
	"strconv"
	"strings"
>>>>>>> Stashed changes
)

func main() {

	fmt.Println("Math")

	fmt.Print("Enter num: ")
	numInput, _ := reader.ReadString('\n')
	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Println("the error is", err)
	} else {
		fmt.Println("Value of number:", aFloat)
	}
	fmt.Println("You entered:", aFloat)
}
