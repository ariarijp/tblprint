package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func main() {
	optSep := flag.String("sep", "\t", "Separator")
	optHeader := flag.Bool("header", true, "First line as header")
	flag.Parse()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)

	scanner := bufio.NewScanner(os.Stdin)

	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()

		if *optHeader && lineNum == 1 {
			table.SetHeader(strings.Split(line, *optSep))
		} else {
			table.Append(strings.Split(line, *optSep))
		}

		lineNum++
	}

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	table.Render()
}
