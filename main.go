package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
)

func main() {
	wp := makeWorkplace()
	handler := MakeWorkplaceCommandHandler(wp)
	ctx := context.Background()

	cliArgs := os.Args[1:]

	if len(cliArgs) == 0 {
		fmt.Println("Please provide the input file path")

		return
	}

	filePath := cliArgs[0]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening the input file")

		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		handler.Handle(ctx, os.Stdout, scanner.Text())
	}
}
