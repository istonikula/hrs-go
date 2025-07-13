package render

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

type Duration struct {
	time.Duration
}

func (d Duration) String() string {
	dd := d.Duration
	if dd < 0 {
		dd = -dd
	}

	minutes := int64(dd.Minutes())
	hours := minutes / 60
	mins := minutes % 60

	return fmt.Sprintf("%02d:%02d", hours, mins)
}

func (d Duration) Plain() string {
	return d.String()
}

func (d Duration) Line() string {
	c := color.New(color.FgGreen, color.Bold)
	return c.Sprint(d.String())
}

func (d Duration) Tag() string {
	c := color.New(color.FgBlue, color.Bold)
	return c.Sprint(d.String())
}

func (d Duration) Total() string {
	c := color.New(color.FgWhite, color.Bold)
	return c.Sprint(d.String())
}

func (d Duration) Diff() string {
	if d.Duration < 0 {
		c := color.New(color.FgRed)
		return c.Sprintf("-%s", d.String())
	}

	c := color.New(color.FgGreen)
	return c.Sprintf("+%s", d.String())
}
