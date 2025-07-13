package main

import (
	"fmt"
	"os"

	"github.com/istonikula/hrs-go/internal/hrs"
	"github.com/istonikula/hrs-go/internal/render"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: %s <file> <date>\n", os.Args[0])
		os.Exit(1)
	}

	filePath := os.Args[1]
	date := os.Args[2]

	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not read file \"%s\": %v\n", filePath, err)
		os.Exit(1)
	}

	linesInDay := hrs.FindLinesInDay(string(content), date)
	processedLines, durationsByTag := hrs.ProcessLines(linesInDay)
	summary, total := hrs.SummarizeDurations(durationsByTag)

	render.Lines(processedLines)
	render.Summary(summary)
	render.Total(total)
}
