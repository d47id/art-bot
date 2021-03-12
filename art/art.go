package art

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"image"
	"image/color"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

//go:embed tpl/*
var tpl embed.FS

// Bot produces generative art
type Bot struct {
	tpl  *template.Template
	imgs *images
}

// New parses the embedded templates and returns a fresh Bot
func New() (*Bot, error) {
	tpl, err := template.ParseFS(tpl, "tpl/*")
	if err != nil {
		return nil, err
	}

	const count = 10
	imgs := &images{
		rw:    &sync.RWMutex{},
		cache: make([]image.Image, 0, count),
		cl:    http.Client{Timeout: 10 * time.Second},
		src:   "https://source.unsplash.com/random/160x90",
	}

	if err := imgs.populate(count); err != nil {
		return nil, err
	}

	return &Bot{tpl: tpl, imgs: imgs}, nil
}

type images struct {
	rw    *sync.RWMutex
	cache []image.Image
	cl    http.Client
	src   string
}

func (i *images) get() (image.Image, error) {
	i.rw.RLock()
	defer i.rw.RUnlock()

	r := rand.Intn(len(i.cache))
	return i.cache[r], nil
}

func (i *images) populate(count int) error {
	g, ctx := errgroup.WithContext(context.Background())
	for x := 0; x < count; x++ {
		g.Go(i.add(ctx))
	}
	return g.Wait()
}

func (i *images) add(ctx context.Context) func() error {
	return func() error {
		// create request
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, i.src, nil)
		if err != nil {
			return err
		}

		// get new image
		res, err := i.cl.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		img, _, err := image.Decode(res.Body)
		if err != nil {
			return err
		}

		// add new image to cache
		i.rw.Lock()
		defer i.rw.Unlock()
		i.cache = append(i.cache, img)

		return nil
	}
}

type rectangle struct {
	X      int
	Y      int
	Height int
	Width  int
	Color  string
}

func cssRGBA(c color.Color) string {
	r, g, b, _ := c.RGBA()
	r, g, b = r/0x101, g/0x101, b/0x101 // convert 16bit rgb values to 8-bit
	return fmt.Sprintf("rgba(%d, %d, %d, 1)", r, g, b)
}
