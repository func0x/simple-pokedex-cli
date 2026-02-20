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
			key: "https://pokeapi.co/api/v2/pokemon/35",
			val: []byte(`{
			  "userId": 1,
			  "id": 2,
			  "title": "quis ut nam facilis et officia qui",
			  "completed": false
			}`),
		},
		{
			key: "https://pokeapi.co/api/v2/pokemon/100",
			val: []byte(`{
			  "userId": 1,
			  "id": 1,
			  "title": "delectus aut autem",
			  "completed": false
			}`),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expeceted to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}
