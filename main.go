package main

import (
	"fmt"
	"os"

	"github.com/cron-parser/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Invalid cron expression")
		return
	}
	input := os.Args[1]

	p, err := parser.NewCronParser(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	parser.Init(p)

	err = parser.GetService().Parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	parser.GetService().Print()
}
