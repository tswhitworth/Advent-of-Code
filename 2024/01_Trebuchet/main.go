package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"unicode"
)

func p2() {
	var total int32

	dw := [][]byte{
		[]byte("one"),
		[]byte("two"),
		[]byte("three"),
		[]byte("four"),
		[]byte("five"),
		[]byte("six"),
		[]byte("seven"),
		[]byte("eight"),
		[]byte("nine"),
	}

	f, s := getScanner()
	defer f.Close()

	for s.Scan() {
		lineValue := processLine_2(dw, s)
		total += lineValue
	}
	fmt.Println("p2 Total: ", total)
}

func processLine_2(dw [][]byte, scanner *bufio.Scanner) int32 {
	var digits []int32
	var radix int32 = 10

	b := scanner.Bytes()

	for len(b) > 0 {
		if isDigit(b[0]) {
			digits = append(digits, int32(b[0]-'0'))
		} else {
			for i, word := range dw {
				if bytes.HasPrefix(b, word) {
					digits = append(digits, int32(i)+1)

					// Move pointer forward by the length of the matched word
					b = b[len(word)-3:]
					break
				}
			}
		}
		// move pointer forward by one byte
		b = b[1:]
	}
	return digits[0]*radix + digits[len(digits)-1]

}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func p1() {
	var total int32

	f, s := getScanner()
	defer f.Close()

	for s.Scan() {
		line := s.Text()
		lineValue := processLine_1(line)
		total += lineValue
	}
	fmt.Println("p1 Total: ", total)
}

func getScanner() (*os.File, *bufio.Scanner) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)

	return f, scanner
}

func processLine_1(line string) int32 {
	var digits []int32
	var radix int32 = 10

	for _, ch := range line {
		if unicode.IsDigit(ch) {
			digits = append(digits, ch-0x30)
		}
	}

	return digits[0]*radix + digits[len(digits)-1]
}

func main() {
	p1()
	p2()
}
