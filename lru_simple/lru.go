
package lru_simple

import "container/list"

type LRUCache struct {
    capacity int
    list     *list.List
    cache    map[int]*list.Element
}

type entry struct {
    key   int
    value int
}

func NewLRU(cap int) *LRUCache {
    return &LRUCache{
        capacity: cap,
        list:     list.New(),
        cache:    make(map[int]*list.Element),
    }
}

func (l *LRUCache) Get(key int) (int, bool) {
    if elem, ok := l.cache[key]; ok {
        l.list.MoveToFront(elem)
        return elem.Value.(*entry).value, true
    }
    return 0, false
}

func (l *LRUCache) Put(key, value int) {
    if elem, ok := l.cache[key]; ok {
        l.list.MoveToFront(elem)
        elem.Value.(*entry).value = value
        return
    }
    if l.list.Len() == l.capacity {
        back := l.list.Back()
        evicted := back.Value.(*entry).key
        delete(l.cache, evicted)
        l.list.Remove(back)
    }
    elem := l.list.PushFront(&entry{key, value})
    l.cache[key] = elem
}
