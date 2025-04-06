package art

import (
	"embed"
	"fmt"
	"html/template"
	"image"
	"image/color"
	_ "image/jpeg"
	"math/rand"
	"path"
)

//go:embed tpl/*
var tpl embed.FS

//go:embed pics/*
var pics embed.FS

// Bot produces generative art
type Bot struct {
	tpl  *template.Template
	imgs images
}

// New parses the embedded templates and returns a fresh Bot
func New() (*Bot, error) {
	tpl, err := template.ParseFS(tpl, "tpl/*")
	if err != nil {
		return nil, err
	}

	imgs, err := initImages()
	if err != nil {
		return nil, err
	}

	return &Bot{tpl: tpl, imgs: imgs}, nil
}

type images []image.Image

func initImages() (images, error) {
	entries, err := pics.ReadDir("pics")
	if err != nil {
		return nil, err
	}

	imgs := make(images, 0, len(entries))

	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		f, err := pics.Open(path.Join("pics", e.Name()))
		if err != nil {
			return nil, err
		}

		img, _, err := image.Decode(f)
		if err != nil {
			return nil, err
		}

		imgs = append(imgs, img)
	}

	return imgs, nil
}

func (i images) get() image.Image {
	r := rand.Intn(len(i))
	return i[r]
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
