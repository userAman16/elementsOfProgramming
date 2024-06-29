package main

import (
	"fmt"
	"sync"
)

type LRUCache struct {
	capacity int
	cache    map[int]*DLLNode
	head     *DLLNode
	tail     *DLLNode
	mu       sync.Mutex
}

type DLLNode struct {
	key   int
	value int
	prev  *DLLNode
	next  *DLLNode
}

var instance *LRUCache
var once sync.Once

func GetInstance(capacity int) *LRUCache {
	once.Do(func() {
		instance = NewLRUCache(capacity)
	})
	return instance
}

func NewLRUCache(capacity int) *LRUCache {
	cache := make(map[int]*DLLNode, capacity)
	return &LRUCache{
		capacity: capacity,
		cache:    cache,
		head:     nil,
		tail:     nil,
	}
}

func (c *LRUCache) Get(key int) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	node, ok := c.cache[key]
	if !ok {
		return -1
	}
	c.moveToHead(node)
	return node.value
}

func (c *LRUCache) Put(key int, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	node, ok := c.cache[key]
	if ok {
		node.value = value
		c.moveToHead(node)
		return
	}
	if len(c.cache) >= c.capacity {
		c.removeTail()
	}
	node = &DLLNode{
		key:   key,
		value: value,
	}
	c.addToHead(node)
	c.cache[key] = node
}

func (c *LRUCache) addToHead(node *DLLNode) {
	if c.head == nil {
		c.head = node
		c.tail = node
	} else {
		node.next = c.head
		c.head.prev = node
		c.head = node
	}
}

func (c *LRUCache) removeTail() {
	if c.tail == nil {
		return
	}
	if c.tail == c.head {
		c.tail = nil
		c.head = nil
	} else {
		c.tail = c.tail.prev
		c.tail.next = nil
	}
	delete(c.cache, c.tail.key)
}

func (c *LRUCache) moveToHead(node *DLLNode) {
	if node == c.head {
		return
	}
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		c.tail = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	node.prev = nil
	node.next = c.head
	c.head.prev = node
	c.head = node
}

func LruCacheTest() {
	cache := GetInstance(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1)) // returns 1
	cache.Put(3, 3)           // evicts key 2
	fmt.Println(cache.Get(2)) // returns -1 (not found)
	cache.Put(4, 4)           // evicts key 1
	fmt.Println(cache.Get(1)) // returns -1 (not found)
	fmt.Println(cache.Get(3)) // returns 3
	fmt.Println(cache.Get(4)) // returns 4

}
