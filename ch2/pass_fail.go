package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	grade, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		log.Fatal(err)
	}

	var status string

	if grade >= 60.0 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println(status)

}
