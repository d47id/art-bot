package art

import (
	"embed"
	"html/template"
	"image"
	"math/rand"
	"net/http"
	"sync"
	"time"
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

	imgs := &images{
		rw:    sync.RWMutex{},
		cache: make([]image.Image, 0, 5),
		cl:    http.Client{Timeout: 5 * time.Second},
		min:   5,
		max:   25,
	}

	return &Bot{tpl: tpl, imgs: imgs}, nil
}

type images struct {
	rw    sync.RWMutex
	cache []image.Image
	cl    http.Client
	min   int
	max   int
}

func (i *images) GetImage() (image.Image, error) {
	if len(i.cache) < i.min {
		// add image synchronously
		if err := i.add(); err != nil {
			return nil, err
		}
	}

	if len(i.cache) >= i.min && len(i.cache) < i.max {
		// add image asynchronously
		go i.add()
	}

	i.rw.RLock()
	defer i.rw.RUnlock()
	return i.cache[rand.Intn(len(i.cache))], nil
}

func (i *images) add() error {
	// get new image
	res, err := i.cl.Get("https://source.unsplash.com/random/160x90")
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
