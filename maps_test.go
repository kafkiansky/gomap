package gomap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapMap(t *testing.T) {
	m := From(map[string]int{"x": 1, "y": 2, "z": 3})

	assert.Equal(t, map[string]int{"x": 1, "y": 2, "z": 3}, m.Map())
}

func TestMapFilterKeys(t *testing.T) {
	m := From(map[string]int{"x": 1, "yy": 2, "zz": 3})

	assert.Equal(t, map[string]int{"x": 1}, m.FilterKeys(func(k string) bool {
		return len(k) == 1
	}).Map())
}

func TestMapFilterValues(t *testing.T) {
	m := From(map[string]int{"x": 1, "y": 2, "z": 3})

	assert.Equal(t, map[string]int{"z": 3}, m.FilterValues(func(v int) bool {
		return v > 2
	}).Map())
}

func TestMapFilter(t *testing.T) {
	m := From(map[string]int{"x": 1, "yy": 2, "zzz": 3})

	assert.Equal(t, map[string]int{"zzz": 3}, m.Filter(func(k string, v int) bool {
		return len(k) == 3 && v > 2
	}).Map())
}

func TestMapExists(t *testing.T) {
	m := From(map[string]int{"x": 1})
	assert.True(t, m.Exists("x"))
	assert.False(t, m.Exists("y"))
}

func TestMapGet(t *testing.T) {
	m := From(map[string]int{"x": 1})
	v, ok := m.Get("x")
	assert.True(t, ok)
	assert.Equal(t, 1, v)

	v, ok = m.Get("y")
	assert.False(t, ok)
	assert.Equal(t, 0, v)
}

func TestMapAdd(t *testing.T) {
	m := From(map[string]int{"x": 1})
	m.Add("y", 2)

	v, ok := m.Get("y")
	assert.True(t, ok)
	assert.Equal(t, 2, v)
}

func TestMapDelete(t *testing.T) {
	m := From(map[string]int{"x": 1})
	assert.True(t, m.Delete("x"))
	assert.False(t, m.Exists("x"))
}

func TestMapLen(t *testing.T) {
	m := From(map[string]int{"x": 1, "z": 2})
	assert.Equal(t, 2, m.Len())
}

func TestMapJoin(t *testing.T) {
	m := From(map[string]int{"x": 1})
	m = m.Join(From(map[string]int{"y": 2}))
	assert.Equal(t, map[string]int{"x": 1, "y": 2}, m.Map())
}

func TestMapDiff(t *testing.T) {
	m := M(map[string]int{"x": 1, "y": 2})
	m = m.Diff(M(map[string]int{"x": 1, "z": 3}))
	assert.Equal(t, map[string]int{"y": 2}, m.Map())
}

func TestEach(t *testing.T) {
	m := Each(
		M(map[string]int{"x": 1, "y": 2}),
		func(v int) int64 {
			return int64(v * 2)
		},
	)

	assert.Equal(t, map[string]int64{"x": 2, "y": 4}, m.Map())
}

func TestEachOnMap(t *testing.T) {
	m := M(map[string]int{"x": 1, "y": 2})

	assert.Equal(t, map[string]int{"x": 3, "y": 6}, m.Each(func(v int) int { return v * 3 }).Map())
}

func TestMapOnly(t *testing.T) {
	m := M(map[string]int{"x": 1, "y": 2, "z": 3, "q": 4})
	m = m.Only("z", "q")
	assert.Equal(t, map[string]int{"z": 3, "q": 4}, m.Map())
}

func TestFromSlice(t *testing.T) {
	m := FromSlice([]string{"x", "y", "z"})
	assert.Equal(t, map[int]string{0: "x", 1: "y", 2: "z"}, m.Map())
}

func TestMapChunk(t *testing.T) {
	m := M(map[string]int{"x": 1, "y": 2, "z": 3})
	maps := m.Chunk(2)
	assert.Equal(t, 2, len(maps))
	assert.Equal(t, map[string]int{"x": 1, "y": 2, "z": 3}, Join(maps...).Map())
}
