package hrs

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type ProcessedLine struct {
	Duration time.Duration
	Line     string
}

func FindLinesInDay(content, date string) []string {
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

func ProcessLines(lines []string) ([]ProcessedLine, map[string][]time.Duration) {
	durationsByTag := make(map[string][]time.Duration)
	var processedLines []ProcessedLine
	var tag string

	for _, line := range lines {
		duration, tag := processLine(line, &tag)
		if duration > 0 {
			processedLines = append(processedLines, ProcessedLine{Duration: duration, Line: line})
			durationsByTag[tag] = append(durationsByTag[tag], duration)
		}
	}
	return processedLines, durationsByTag
}

var lineRe = regexp.MustCompile(`^([0-9\.]{1,5})-([0-9\.]{1,5})\s+(\[.*?\])?.*$`)

func processLine(line string, prevTag *string) (time.Duration, string) {
	caps := lineRe.FindStringSubmatch(line)
	if caps == nil {
		return 0, ""
	}

	startStr := caps[1]
	endStr := caps[2]
	tag := caps[3]
	if tag == "" {
		tag = line[len(fmt.Sprintf("%s-%s ", startStr, endStr)):]
	}

	if strings.HasPrefix(tag, `-"-`) {
		tag = *prevTag
	} else {
		*prevTag = tag
	}

	withMins := func(t string) string {
		if !strings.Contains(t, ".") {
			return t + ".00"
		}
		return t
	}

	layout := "15.04"
	start, err := time.Parse(layout, withMins(startStr))
	if err != nil {
		return 0, ""
	}
	end, err := time.Parse(layout, withMins(endStr))
	if err != nil {
		return 0, ""
	}

	return end.Sub(start), tag
}

func SummarizeDurations(durationsByTag map[string][]time.Duration) (map[string]time.Duration, time.Duration) {
	summary := make(map[string]time.Duration, len(durationsByTag))
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
