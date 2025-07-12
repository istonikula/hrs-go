package render

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

type HumanDuration struct {
	time.Duration
}

func (hd HumanDuration) String() string {
	d := hd.Duration
	if d < 0 {
		d = -d
	}

	minutes := int64(d.Minutes())
	hours := minutes / 60
	mins := minutes % 60

	return fmt.Sprintf("%02d:%02d", hours, mins)
}

func (hd HumanDuration) Plain() string {
	return hd.String()
}

func (hd HumanDuration) Line() string {
	c := color.New(color.FgGreen, color.Bold)
	return c.Sprint(hd.String())
}

func (hd HumanDuration) Tag() string {
	c := color.New(color.FgBlue, color.Bold)
	return c.Sprint(hd.String())
}

func (hd HumanDuration) Total() string {
	c := color.New(color.FgWhite, color.Bold)
	return c.Sprint(hd.String())
}

func (hd HumanDuration) Diff() string {
	if hd.Duration < 0 {
		c := color.New(color.FgRed)
		return c.Sprintf("-%s", hd.String())
	}

	c := color.New(color.FgGreen)
	return c.Sprintf("+%s", hd.String())

}
