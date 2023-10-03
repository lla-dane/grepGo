package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(" Not enough arguments")
		return
	}

	query := os.Args[1]
	filePath := os.Args[2]

	//open file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening the file : %v", err)
		return
	}
	//defer means file.Close() will be called when main function will be executed completely
	defer file.Close()

	//create a reader to read lines from the file
	scanner := bufio.NewScanner(file)

	//Iterate through each line and check if it contains the query
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, query) {
			fmt.Printf(" %s %d %s \n", filePath, lineNum, line)
		}
		lineNum++
	}

	//Check for errors while scanning the  file
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning the file: %v \n", err)
		return
	}

}
