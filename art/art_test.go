package art

import "testing"

func TestPics(t *testing.T) {
	entries, err := pics.ReadDir("pics")
	if err != nil {
		t.Fatal(err)
	}

	for _, e := range entries {
		t.Log(e.Name(), e.IsDir())
	}
}
