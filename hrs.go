package main

import (
	"bufio"
	"regexp"
	"strings"
	"time"
)

type ProcessedLine struct {
	Duration time.Duration
	Line     string
}

func findAndCollectDay(content, date string) []string {
	var linesInDay []string
	inDay := false

	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		if !inDay {
			if line == date || strings.HasPrefix(line, date+" ") {
				inDay = true
				linesInDay = append(linesInDay, line)
			}
		} else {
			if line == "" {
				break
			}
			linesInDay = append(linesInDay, line)
		}
	}
	return linesInDay
}

func processLines(lines []string) ([]ProcessedLine, map[string][]time.Duration) {
	durationsByTag := make(map[string][]time.Duration)
	var processedLines []ProcessedLine
	var prevTag string

	// Regex to capture start, end, and the entire description part
	lineRe := regexp.MustCompile(`^([0-9\.]+)-([0-9\.]+)\s+(.*)$`)

	for _, line := range lines {
		duration, tag := processLine(line, &prevTag, lineRe)
		if duration > 0 {
			processedLines = append(processedLines, ProcessedLine{Duration: duration, Line: line})
			durationsByTag[tag] = append(durationsByTag[tag], duration)
		}
	}
	return processedLines, durationsByTag
}

func processLine(line string, prevTag *string, lineRe *regexp.Regexp) (time.Duration, string) {
	caps := lineRe.FindStringSubmatch(line)
	if caps == nil {
		return 0, ""
	}

	startStr := caps[1]
	endStr := caps[2]
	fullDescription := caps[3] // This is the entire description part, e.g., "[TAG-1] desc" or "tagless desc" or "-"-"

	var currentTag string

	// Try to extract a bracketed tag from fullDescription
	bracketedTagRe := regexp.MustCompile(`^\[(.*?)\]`)
	bracketedTagCaps := bracketedTagRe.FindStringSubmatch(strings.TrimSpace(fullDescription))

	if bracketedTagCaps != nil {
		currentTag = bracketedTagCaps[0] // Use the bracketed tag including brackets
	} else {
		currentTag = strings.TrimSpace(fullDescription) // Otherwise, the entire description is the tag
	}

	// Handle the -"-" convention
	if strings.HasPrefix(currentTag, `-"-`) {
		if *prevTag != "" {
			currentTag = *prevTag
		}
		// If *prevTag is empty, currentTag remains "-"-" which is correct for the first "-"-" line
	} else {
		// If it's not "-"-" or if it's the first line, update prevTag
		*prevTag = currentTag
	}

	layout := "15.04"
	start, err1 := time.Parse(layout, withMins(startStr))
	end, err2 := time.Parse(layout, withMins(endStr))

	if err1 != nil || err2 != nil {
		return 0, ""
	}

	return end.Sub(start), currentTag
}

func withMins(timeStr string) string {
	if !strings.Contains(timeStr, ".") {
		return timeStr + ".00"
	}
	return timeStr
}

func summarizeDurations(durationsByTag map[string][]time.Duration) (map[string]time.Duration, time.Duration) {
	summary := make(map[string]time.Duration)
	var total time.Duration

	for tag, durations := range durationsByTag {
		var tagTotal time.Duration
		for _, d := range durations {
			tagTotal += d
		}
		summary[tag] = tagTotal
		total += tagTotal
	}

	return summary, total
}