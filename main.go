package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func main() {
	SEP := "\t"

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)

	scanner := bufio.NewScanner(os.Stdin)
	header := false

	for scanner.Scan() {
		line := scanner.Text()

		if !header {
			table.SetHeader(strings.Split(line, SEP))
			header = true
			continue
		}

		table.Append(strings.Split(line, SEP))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}

	table.Render()
}
