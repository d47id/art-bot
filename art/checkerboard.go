package art

import (
	"io"
	"math/rand"

	"github.com/d47id/art-bot/colors"
)

func makeRandomCheckerboard() []*rectangle {
	rects := make([]*rectangle, 0, 100)

	for x := 0; x < 100; x += 10 {
		for y := 0; y < 100; y += 10 {
			rects = append(rects, &rectangle{
				X:      x,
				Y:      y,
				Width:  10,
				Height: 10,
				Color:  colors.Random(),
			})
		}
	}

	return rects
}

func makeErodedCheckerboard() []*rectangle {
	rects := make([]*rectangle, 0, 100)
	limit := rand.Intn(100)
	gate := func() bool {
		return rand.Intn(100) < limit
	}

	for x := 0; x < 100; x += 10 {
		for y := 0; y < 100; y += 10 {
			if gate() {
				rects = append(rects, &rectangle{
					X:      x,
					Y:      y,
					Width:  10,
					Height: 10,
					Color:  colors.Random(),
				})
			}
		}
	}

	return rects
}

// WriteCheckerboard writes an eroded checkerboard svg+xml to w
func (b *Bot) WriteCheckerboard(w io.Writer) error {
	return b.tpl.ExecuteTemplate(w, "checkerboard.svg", makeErodedCheckerboard())
}
