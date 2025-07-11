package main

import (
    "fmt"
    "os"
    "sort"
    "time"
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

    linesInDay := findAndCollectDay(string(content), date)
    processedLines, durationsByTag := processLines(linesInDay)
    summary, total := summarizeDurations(durationsByTag)

    printProcessedLines(processedLines)
    printSummary(summary)
    printTotalAndDiff(total)
}

func printProcessedLines(processedLines []ProcessedLine) {
    fmt.Println("----")
    for _, pLine := range processedLines {
        fmt.Printf("%s %s\n", formatDuration(pLine.Duration), pLine.Line)
    }
}

func printSummary(summary map[string]time.Duration) {
    fmt.Println("----")

    var tags []string
    for tag := range summary {
        tags = append(tags, tag)
    }
    sort.Strings(tags)

    for _, tag := range tags {
        fmt.Printf("%s %s\n", formatDuration(summary[tag]), tag)
    }
}

func printTotalAndDiff(total time.Duration) {
    fmt.Println("----")
    fullDay := time.Hour*7 + time.Minute*30
    diff := total - fullDay

    if diff == 0 {
        fmt.Println(formatDuration(total))
    } else {
        fmt.Printf("%s %s\n", formatDuration(total), formatDiff(diff))
    }
}

func formatDuration(d time.Duration) string {
    d = d.Round(time.Minute)
    h := d / time.Hour
    d -= h * time.Hour
    m := d / time.Minute
    return fmt.Sprintf("%02d:%02d", h, m)
}

func formatDiff(d time.Duration) string {
    if d < 0 {
        return fmt.Sprintf("-%s", formatDuration(-d))
    }
    return fmt.Sprintf("+%s", formatDuration(d))
}
