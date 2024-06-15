package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {

	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "http://example.com",
			val: []byte("testdata"),
		},
		{
			key: "http:example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.AddEntry(c.key, c.val)
			val, ok := cache.GetEntry(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}

}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond

	cache := NewCache(baseTime)
	cache.AddEntry("http://example.com", []byte("testdata"))
	_, ok := cache.GetEntry("http://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.GetEntry("https://exmaple.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
