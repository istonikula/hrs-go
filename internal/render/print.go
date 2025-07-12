package render

import (
	"fmt"
	"sort"
	"time"

	"github.com/istonikula/hrs-go/internal/hrs"
)

func ProcessedLines(lines []hrs.ProcessedLine) {
	fmt.Println("----")
	for _, line := range lines {
		fmt.Printf("%s %s\n", HumanDuration{line.Duration}.Line(), line.Line)
	}
}

func Summary(summary map[string]time.Duration) {
	fmt.Println("----")

	var tags []string
	for tag := range summary {
		tags = append(tags, tag)
	}
	sort.Strings(tags)

	for _, tag := range tags {
		fmt.Printf("%s %s\n", HumanDuration{summary[tag]}.Tag(), tag)
	}
}

func TotalAndDiff(total time.Duration) {
	fmt.Println("----")
	fullDay := time.Hour*7 + time.Minute*30
	diff := total - fullDay

	if diff == 0 {
		fmt.Println(HumanDuration{total}.Total())
	} else {
		fmt.Printf("%s %s\n", HumanDuration{total}.Total(), HumanDuration{diff}.Diff())
	}
}
