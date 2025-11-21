
package lru_simple

import "testing"

func TestLRUCacheBasic(t *testing.T) {
    lru := NewLRU(2)
    lru.Put(1, 10)
    lru.Put(2, 20)
    v, ok := lru.Get(1)
    if !ok || v != 10 {
        t.Fatalf("expected 10, got %v", v)
    }
}

func TestLRUEviction(t *testing.T) {
    lru := NewLRU(2)
    lru.Put(1, 10)
    lru.Put(2, 20)
    lru.Put(3, 30)
    if _, ok := lru.Get(1); ok {
        t.Fatalf("expected eviction of 1")
    }
}

func TestLRUUpdatesMoveToFront(t *testing.T) {
    lru := NewLRU(2)
    lru.Put(1, 10)
    lru.Put(2, 20)
    lru.Get(1)
    lru.Put(3, 30)
    if _, ok := lru.Get(2); ok {
        t.Fatalf("expected eviction of 2")
    }
}

func TestLRUUpdateValue(t *testing.T) {
    lru := NewLRU(1)
    lru.Put(1, 10)
    lru.Put(1, 99)
    v, ok := lru.Get(1)
    if !ok || v != 99 {
        t.Fatalf("expected 99, got %v", v)
    }
}
