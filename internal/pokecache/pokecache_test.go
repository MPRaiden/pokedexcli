package pokecache

import (
	"testing"
	"time"
)

func TestAddCache(t *testing.T) {
	cache := NewCache(5 * time.Second)
	cache.Add("key", []byte("value"))

	val, ok := cache.Get("key")
	if !ok {
		t.Fatalf("Expected to find key in cache")
	}

	if string(val) != "value" {
		t.Fatalf("Expected value to be 'value', got '%s'", string(val))
	}
}

func TestGetCache(t *testing.T) {
	cache := NewCache(5 * time.Second)
	cache.Add("key", []byte("value"))

	val, ok := cache.Get("key")
	if !ok {
		t.Fatalf("Expected to find key in cache")
	}

	if string(val) != "value" {
		t.Fatalf("Expected value to be 'value', got '%s'", string(val))
	}
}

