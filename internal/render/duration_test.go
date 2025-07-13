package render_test

import (
	"testing"
	"time"

	"github.com/fatih/color"
	"github.com/istonikula/hrs-go/internal/render"
	"github.com/stretchr/testify/assert"
)

func TestDuration(t *testing.T) {
	t.Run("formatting", func(t *testing.T) {
		assert.Equal(t, "00:01", render.Duration{time.Minute}.Plain())
		assert.Equal(t, "00:15", render.Duration{15 * time.Minute}.Plain())
		assert.Equal(t, "01:00", render.Duration{time.Hour}.Plain())
		assert.Equal(t, "02:15", render.Duration{135 * time.Minute}.Plain())
		assert.Equal(t, "10:00", render.Duration{10 * time.Hour}.Plain())
	})

	t.Run("attributes", func(t *testing.T) {
		noColorOrig := color.NoColor
		color.NoColor = false
		defer func() { color.NoColor = noColorOrig }()

		h := render.Duration{Duration: time.Hour}
		hNeg := render.Duration{Duration: -time.Hour}

		assert.Equal(t, "01:00", h.Plain())
		assert.Equal(t, color.New(color.FgGreen, color.Bold).Sprint("01:00"), h.Line())
		assert.Equal(t, color.New(color.FgBlue, color.Bold).Sprint("01:00"), h.Tag())
		assert.Equal(t, color.New(color.FgWhite, color.Bold).Sprint("01:00"), h.Total())
		assert.Equal(t, color.New(color.FgGreen).Sprint("+01:00"), h.Diff())
		assert.Equal(t, color.New(color.FgRed).Sprint("-01:00"), hNeg.Diff())
	})
}
