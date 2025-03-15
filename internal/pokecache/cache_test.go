// internal/pokecache/cache_test.go
package pokecache

import (
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		name     string
		key      string
		value    []byte
		expected []byte
	}{
		{
			name:     "basic string",
			key:      "test-key",
			value:    []byte("test-value"),
			expected: []byte("test-value"),
		},
		{
			name:     "empty string",
			key:      "empty",
			value:    []byte(""),
			expected: []byte(""),
		},
		{
			name:     "binary data",
			key:      "binary",
			value:    []byte{0, 1, 2, 3},
			expected: []byte{0, 1, 2, 3},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cache := NewCache(5 * time.Minute)

			cache.Add(c.key, c.value)

			actual, found := cache.Get(c.key)

			if !found {
				t.Errorf("key %s was not found in cache", c.key)
				return
			}

			if len(actual) != len(c.expected) {
				t.Errorf("values don't match: got %v, expected %v", actual, c.expected)
				return
			}

			for i := range actual {
				if actual[i] != c.expected[i] {
					t.Errorf("cache.Add(%q, %v) => %v, expected %v",
						c.key, c.value, actual, c.expected)
				}
			}
		})
	}
}

func TestExpiration(t *testing.T) {
	cases := []struct {
		name            string
		cacheInterval   time.Duration
		waitTime        time.Duration
		shouldBeRemoved bool
	}{
		{
			name:            "not expired",
			cacheInterval:   200 * time.Millisecond,
			waitTime:        100 * time.Millisecond,
			shouldBeRemoved: false,
		},
		{
			name:            "expired",
			cacheInterval:   100 * time.Millisecond,
			waitTime:        200 * time.Millisecond,
			shouldBeRemoved: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cache := NewCache(c.cacheInterval)

			testKey := "expiration-test"
			testValue := []byte("test-value")
			cache.Add(testKey, testValue)

			time.Sleep(c.waitTime)

			cache.reap()

			_, found := cache.Get(testKey)

			if found && c.shouldBeRemoved {
				t.Errorf("item should have been removed after %v but was still in cache",
					c.waitTime)
			}

			if !found && !c.shouldBeRemoved {
				t.Errorf("item should not have been removed after %v but was gone",
					c.waitTime)
			}
		})
	}
}

func TestGet(t *testing.T) {
	cases := []struct {
		name       string
		key        string
		value      []byte
		lookupKey  string
		shouldFind bool
	}{
		{
			name:       "existing key",
			key:        "test-key",
			value:      []byte("test-value"),
			lookupKey:  "test-key",
			shouldFind: true,
		},
		{
			name:       "non-existent key",
			key:        "test-key",
			value:      []byte("test-value"),
			lookupKey:  "wrong-key",
			shouldFind: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cache := NewCache(5 * time.Minute)

			cache.Add(c.key, c.value)

			value, found := cache.Get(c.lookupKey)

			if found != c.shouldFind {
				t.Errorf("cache.Get(%q) found = %v, expected %v",
					c.lookupKey, found, c.shouldFind)
			}

			if found && c.shouldFind {
				if string(value) != string(c.value) {
					t.Errorf("cache.Get(%q) = %q, expected %q",
						c.lookupKey, string(value), string(c.value))
				}
			}
		})
	}
}
