package omap

type OMap[K comparable, V any] struct {
	map_   map[K]V
	slice_ []K
}

func New[K comparable, V any]() *OMap[K, V] {
	return &OMap[K, V]{
		map_:   make(map[K]V),
		slice_: make([]K, 0),
	}
}

func (om *OMap[K, V]) At(key K) (V, bool) {
	value, ok := om.map_[key]
	return value, ok
}

func (om *OMap[K, V]) Add(key K, value V) {
	if _, ok := om.At(key); !ok {
		om.slice_ = append(om.slice_, key)
	}
	om.map_[key] = value

}

type KV[K comparable, V any] struct {
	Key   K
	Value V
}

func (om *OMap[K, V]) Iter() chan KV[K, V] {
	ch := make(chan (KV[K, V]), len(om.slice_))
	go func() {
		defer close(ch)
		for _, key := range om.slice_ {
			ch <- KV[K, V]{Key: key, Value: om.map_[key]}
		}
	}()

	return ch
}
