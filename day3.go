package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	NTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	N := int32(NTemp)
	//var isOdd = false
	if N%2 > 0 {
		fmt.Printf("Weird\n")
		return
	} else {
		switch N {
		case 2, 4:
			fmt.Printf("Not Weird\n")
		case 6, 8, 10, 12, 14, 16, 18, 20:
			fmt.Printf("Weird\n")
		default:
			fmt.Printf("Not Weird\n")
		}
	}
}
func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
