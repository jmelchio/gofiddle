package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "vetlab"
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

func databaseStuff() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("We connected to the database!")
	return nil
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
	err = databaseStuff()
	errHandler(err)

	if err != nil {
		panic(err)
	}
}

// That's All Folks !!
