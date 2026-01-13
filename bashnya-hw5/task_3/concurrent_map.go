package concurrent_map

import "sync"

type ConcurrentMap[K comparable, V any] struct {
	myMap map[K]V
	mu    sync.RWMutex
}

func New[K comparable, V any]() *ConcurrentMap[K, V] {
	return &ConcurrentMap[K, V]{
		myMap: make(map[K]V),
	}
}

func NewFromMap[K comparable, V any](initial map[K]V) *ConcurrentMap[K, V] {
	return &ConcurrentMap[K, V]{
		myMap: initial,
	}
}

func (cm *ConcurrentMap[K, V]) Get(key K) (V, bool) {
	cm.mu.RLock()
	defer cm.mu.RUnlock()
	value, ok := cm.myMap[key]

	return value, ok
}

func (cm *ConcurrentMap[K, V]) Add(key K, value V) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	cm.myMap[key] = value
}

func (cm *ConcurrentMap[K, V]) Remove(key K) {
	cm.mu.Lock()
	defer cm.mu.Unlock()
	delete(cm.myMap, key)
}
