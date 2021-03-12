package art

import (
	"image"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestImageCache(t *testing.T) {
	const count = 5
	imgs := &images{
		rw:    &sync.RWMutex{},
		cache: make([]image.Image, 0, count),
		cl:    http.Client{Timeout: 10 * time.Second},
		src:   "https://source.unsplash.com/random/160x90",
	}

	if err := imgs.populate(count); err != nil {
		t.Fatal(err)
	}

	if len(imgs.cache) != count {
		t.Fatal("unexpected count. wanted:", count, "got:", len(imgs.cache))
	}
}
