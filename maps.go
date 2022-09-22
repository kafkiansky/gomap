package gomap

import "sync"

// Map is a concurrency safe data structure, which represents a generic builtin map as a Map[K, V].
type Map[K comparable, V any] struct {
	mutex    *sync.RWMutex
	innerMap map[K]V
}

func newMap[K comparable, V any](m map[K]V) Map[K, V] {
	return Map[K, V]{
		mutex:    &sync.RWMutex{},
		innerMap: m,
	}
}

// Add adds the element to Map[K, V].
func (m Map[K, V]) Add(k K, v V) Map[K, V] {
	m.mutex.Lock()
	m.innerMap[k] = v
	m.mutex.Unlock()

	return m
}

// Delete delete the element from Map[K, V] using key.
func (m Map[K, V]) Delete(k K) bool {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, exists := m.innerMap[k]; exists {
		delete(m.innerMap, k)
		return true
	}

	return false
}

// Get return the V and the true, if element by K exists in Map[K, V]. Otherwise, the zero value of V and false will return.
func (m Map[K, V]) Get(k K) (V, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if v, exists := m.innerMap[k]; exists {
		return v, true
	}

	var v V
	return v, false
}

// Len return the actual len of inner map.
func (m Map[K, V]) Len() int {
	return len(m.innerMap)
}

// Exists check if value by key exists in Map[K, V].
func (m Map[K, V]) Exists(k K) bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if _, exists := m.innerMap[k]; exists {
		return true
	}

	return false
}

// From creates the generic Map[K, V] from builtin map.
func From[K comparable, V any](m map[K]V) Map[K, V] {
	return newMap(m)
}

// FromSlice creates Map[int, V] from slice []V.
func FromSlice[V any](values []V) Map[int, V] {
	newmap := make(map[int]V, len(values))

	for i, v := range values {
		newmap[i] = v
	}

	return newMap(newmap)
}

// M alias to From.
func M[K comparable, V any](m map[K]V) Map[K, V] {
	return newMap(m)
}

// Filter filters both key and value of generic Map[K, V].
func (m Map[K, V]) Filter(filter func(K, V) bool) Map[K, V] {
	newmap := make(map[K]V, len(m.innerMap))

	for k, v := range m.innerMap {
		if filter(k, v) {
			newmap[k] = v
		}
	}

	return newMap(newmap)
}

// FilterValues filters only values of generic Map[K, V].
func (m Map[K, V]) FilterValues(filter func(V) bool) Map[K, V] {
	newmap := make(map[K]V, len(m.innerMap))

	for k, v := range m.innerMap {
		if filter(v) {
			newmap[k] = v
		}
	}

	return newMap(newmap)
}

// FilterKeys filters only keys of generic Map[K, V].
func (m Map[K, V]) FilterKeys(filter func(K) bool) Map[K, V] {
	newmap := make(map[K]V, len(m.innerMap))

	for k, v := range m.innerMap {
		if filter(k) {
			newmap[k] = v
		}
	}

	return newMap(newmap)
}

// Chunk creates slice of Map[K, V] with provided size.
func (m Map[K, V]) Chunk(size uint) []Map[K, V] {
	var maps []Map[K, V]

	chunk := make(map[K]V, size)
	for k, v := range m.innerMap {
		chunk[k] = v

		if uint(len(chunk)) >= size {
			maps = append(maps, newMap(chunk))
			chunk = make(map[K]V, size)
		}
	}

	if len(chunk) > 0 {
		maps = append(maps, newMap(chunk))
	}

	return maps
}

// Diff the items in the Map[K, V] that are not present in the other and return them as new Map[K, V].
func (m Map[K, V]) Diff(other Map[K, V]) Map[K, V] {
	differ := make(map[K]V, m.Len())

	for k, v := range m.innerMap {
		if !other.Exists(k) {
			differ[k] = v
		}
	}

	return newMap(differ)
}

// Join joins the target Map[K, V] with the others ...Map[K, V].
func (m Map[K, V]) Join(others ...Map[K, V]) Map[K, V] {
	mapLen := m.Len()

	for _, other := range others {
		mapLen += other.Len()
	}

	joined := make(map[K]V, mapLen)
	others = append(others, m)

	for _, other := range others {
		for k, v := range other.innerMap {
			joined[k] = v
		}
	}

	return newMap(joined)
}

// Join maps together to Map[K, V].
func Join[K comparable, V any](others ...Map[K, V]) Map[K, V] {
	return M(map[K]V{}).Join(others...)
}

// Only return Map[K, V] which contains values only for given keys.
func (m Map[K, V]) Only(keys ...K) Map[K, V] {
	newmap := make(map[K]V, len(keys))

	for _, key := range keys {
		if v, exists := m.Get(key); exists {
			newmap[key] = v
		}
	}

	return newMap(newmap)
}

// Each iterate the Map[K, V] and apply the mapper function to each element and output the modified Map[K, V].
func (m Map[K, V]) Each(mapper func(V) V) Map[K, V] {
	newmap := make(map[K]V, len(m.innerMap))

	for k, v := range m.innerMap {
		newmap[k] = mapper(v)
	}

	return newMap(newmap)
}

// Each iterate the Map[K, V] and apply the mapper function to each element of map and output the new Map[K, E].
func Each[K comparable, V, E any](m Map[K, V], mapper func(V) E) Map[K, E] {
	newmap := make(map[K]E, len(m.innerMap))

	for k, v := range m.innerMap {
		newmap[k] = mapper(v)
	}

	return newMap(newmap)
}

// Map return the Map[K, V] as builtin map[K]V.
func (m Map[K, V]) Map() map[K]V {
	return m.innerMap
}
