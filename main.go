package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	info, _ := os.Stdin.Stat()

	fmt.Println(info)
	fmt.Println(info.Mode())

	if info.Mode().String()[0] != os.ModeNamedPipe.String()[0] {
		return
	}

	input, errStdin := io.ReadAll(os.Stdin)

	if errStdin != nil {
		fmt.Println("Error: " + errStdin.Error())
		return
	}

	counter := 0
	inSize := len(input)
	for counter < inSize && input[counter] != '\n' {
		counter++
	}

	if counter >= inSize {
		fmt.Println("Not a quad function")
		return
	}

	abs := counter
	ord := 0
	for counter < inSize {
		if input[counter] == '\n' {
			if counter-abs*ord-ord == abs {
				ord++
			} else {
				fmt.Println("Not a quad function")
				return
			}
		}
		counter++
	}

	quads := []string{}

	cmdA := exec.Command("./quadA", itoa(abs), itoa(ord))

	contentA, _ := cmdA.Output()

	cmdA.Run()

	if inSize != len(contentA) {
		fmt.Println("Not a quad function")
		return
	}

	position := 0
	for position < inSize && input[position] == contentA[position] {
		position++
	}

	if position >= inSize {
		quads = append(quads, "quadA")
	}

	cmdB := exec.Command("./quadB", itoa(abs), itoa(ord))

	contentB, _ := cmdB.Output()

	cmdB.Run()

	if inSize != len(contentB) {
		fmt.Println("Not a quad function")
		return
	}

	position = 0
	for position < inSize && input[position] == contentB[position] {
		position++
	}

	if position >= inSize {
		quads = append(quads, "quadB")
	}

	cmdC := exec.Command("./quadC", itoa(abs), itoa(ord))

	contentC, _ := cmdC.Output()

	cmdC.Run()

	if inSize != len(contentC) {
		fmt.Println("Not a quad function")
		return
	}

	position = 0
	for position < inSize && input[position] == contentC[position] {
		position++
	}

	if position >= inSize {
		quads = append(quads, "quadC")
	}

	cmdD := exec.Command("./quadD", itoa(abs), itoa(ord))

	contentD, _ := cmdD.Output()

	cmdD.Run()

	if inSize != len(contentD) {
		fmt.Println("Not a quad function")
		return
	}

	position = 0
	for position < inSize && input[position] == contentD[position] {
		position++
	}

	if position >= inSize {
		quads = append(quads, "quadD")
	}

	cmdE := exec.Command("./quadE", itoa(abs), itoa(ord))

	contentE, _ := cmdE.Output()

	cmdE.Run()

	if inSize != len(contentE) {
		fmt.Println("Not a quad function")
		return
	}

	position = 0
	for position < inSize && input[position] == contentE[position] {
		position++
	}

	if position >= inSize {
		quads = append(quads, "quadE")
	}

	if len(quads) == 0 {
		fmt.Println("Not a quad function")
		return
	}

	for index, quad := range quads {
		fmt.Print("[" + quad + "]" + " [" + itoa(abs) + "] " + "[" + itoa(ord) + "]")
		if index+1 < len(quads) {
			fmt.Print(" || ")
		}
	}
	fmt.Println("")
}

func itoa(number int) string {
	if number == 0 {
		return "0"
	}

	keep := number
	digits := []byte{}

	for keep != 0 {
		if number < 0 {
			digits = append([]byte{byte('0' - keep%10)}, digits...)
		} else {
			digits = append([]byte{byte('0' + keep%10)}, digits...)
		}
		keep /= 10
	}

	if number < 0 {
		digits = append([]byte{'-'}, digits...)
	}

	return string(digits)
}
