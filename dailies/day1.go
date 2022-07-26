package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	// Declare second integer, double, and String variables.
	var a int64 = 0
	var b float64 = 0.0
	var c string = ""

	// Read and save an integer, double, and String to your variables.
	//reader := bufio.NewReader(os.Stdin)

	//input := make([]string, 0)

	//fmt.Print("Enter an Int: ")
	scanner.Scan()
	a, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	//input = append(input, a)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "shouldn't see an error scanning a string")
	}
	//fmt.Print("Enter an float: ")
	scanner.Scan()
	b, _ = strconv.ParseFloat(scanner.Text(), 32)

	//fmt.Print("Enter an String: ")
	scanner.Scan()
	c = scanner.Text()

	// Print the sum of both integer variables on a new line.
	var sumI = a + int64(i)
	fmt.Printf("%d\n", sumI)
	// Print the sum of the double variables on a new line.
	var sumD = d + b
	fmt.Printf("%.2f\n", sumD)

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Printf("%s%s", s, c)
}
