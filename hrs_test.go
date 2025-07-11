package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFindLinesInDay(t *testing.T) {
	input := `


27.2
--
foo


28.2 (day info)
--
bar

1.3  day info
--
baz
`

	assert.Equal(t, []string{"27.2", "--", "foo"}, findLinesInDay(input, "27.2"))
	assert.Equal(t, []string{"28.2 (day info)", "--", "bar"}, findLinesInDay(input, "28.2"))
	assert.Equal(t, []string{"1.3  day info", "--", "baz"}, findLinesInDay(input, "1.3"))
}

func TestProcessLines(t *testing.T) {
	lines, durationsByTag := processLines([]string{
		"8-9 desc without tag 1",
		"9-9.30 [tag1] desc",
		`9.45-10 -"-`,
		"10-10.30 desc without tag 2",
		`10.45-11 -"-`,
		"12-14.15 desc without tag 1",
		"14.15-16 [tag1] desc, with some additinal info",
		"16-17   [tag1] NOTE: whitespace before tag",
	})

	assert.Equal(t, []ProcessedLine{
		{Duration: time.Hour, Line: "8-9 desc without tag 1"},
		{Duration: 30 * time.Minute, Line: "9-9.30 [tag1] desc"},
		{Duration: 15 * time.Minute, Line: `9.45-10 -"-`},
		{Duration: 30 * time.Minute, Line: "10-10.30 desc without tag 2"},
		{Duration: 15 * time.Minute, Line: `10.45-11 -"-`},
		{Duration: 2*time.Hour + 15*time.Minute, Line: "12-14.15 desc without tag 1"},
		{Duration: time.Hour + 45*time.Minute, Line: "14.15-16 [tag1] desc, with some additinal info"},
		{Duration: time.Hour, Line: "16-17   [tag1] NOTE: whitespace before tag"},
	}, lines)

	assert.Equal(t, map[string][]time.Duration{
		"desc without tag 1": {time.Hour, 2*time.Hour + 15*time.Minute},
		"[tag1]":             {30 * time.Minute, 15 * time.Minute, time.Hour + 45*time.Minute, time.Hour},
		"desc without tag 2": {30 * time.Minute, 15 * time.Minute},
	}, durationsByTag)
}

func TestSummarizeDurations(t *testing.T) {
	summary, total := summarizeDurations(map[string][]time.Duration{
		"desc without tag 2": {30 * time.Minute, 15 * time.Minute},
		"[tag1]":             {30 * time.Minute, 15 * time.Minute, 105 * time.Minute},
		"desc without tag 1": {time.Hour, 135 * time.Minute},
		"[tag2]":             {time.Hour, 45 * time.Minute},
	})

	assert.Equal(t, map[string]time.Duration{
		"[tag1]":             30*time.Minute + 15*time.Minute + 105*time.Minute,
		"[tag2]":             time.Hour + 45*time.Minute,
		"desc without tag 1": time.Hour + 135*time.Minute,
		"desc without tag 2": 30*time.Minute + 15*time.Minute,
	}, summary)

	var expectedTotal time.Duration
	for _, d := range summary {
		expectedTotal += d
	}
	assert.Equal(t, expectedTotal, total)
}
