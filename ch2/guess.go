package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {

	target := rand.Intn(100) + 1
	fmt.Println(target)
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Print("Can you guess it? ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	guess, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil {
		log.Fatal(err)
	}

	if guess < target {
		fmt.Println("Oops. Your guess was LOW.")
	} else if guess > target {
		fmt.Println("Oops. Your guess was HIGH.")
	} else {
		fmt.Println("Good job! You guessed it!")
	}

}