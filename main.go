package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func writeSomeStuff(fileName string, text string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	errHandler(err)

	defer deferredFileClose(file)

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(text)
	errHandler(err)
	err = writer.Flush()
	errHandler(err)

	return err
}

func readAndPrint(fileName string) error {
	file, err := os.Open(fileName)
	errHandler(err)

	defer deferredFileClose(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	return err
}

func deferredFileClose(file *os.File) {
	err := file.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("Closing file")
}

func errHandler(err error) {
	if err != nil {
		fmt.Println("Error occurred: ", err.Error())
	}
}

func main() {
	text := `And now for something completely different:
The Larch!
This is some sort of tree
Not to be confused with a shrubbery!
`

	err := writeSomeStuff("./readfile.txt", text)
	errHandler(err)
	err = readAndPrint("./readfile.txt")
	errHandler(err)

	if err != nil {
		panic(err)
	}
}

// That's All Folks !!
