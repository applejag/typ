// SPDX-FileCopyrightText: 2025 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package sync2

import (
	"fmt"
	"sort"
	"testing"
)

func TestMapValues(t *testing.T) {
	m := new(Map[int, string])
	m.Store(1, "1")
	m.Store(2, "2")
	m.Store(3, "3")

	t.Run("loop all", func(t *testing.T) {
		iter := MapValues(m)
		var values []string
		iter(func(v string) bool {
			values = append(values, v)
			return true
		})

		if len(values) != 3 {
			t.Errorf("want len() == 3, got len() == %d", len(values))
		}
		sort.Strings(values)
		got := fmt.Sprintf("%#v", values)
		want := `[]string{"1", "2", "3"}`
		if got != want {
			t.Errorf("wrong key-value pairs\nwant: %s\ngot:  %s", want, got)
		}
	})

	t.Run("break", func(t *testing.T) {
		iter := MapValues(m)
		count := 0
		iter(func(v string) bool {
			count++
			if count == 2 {
				return false
			}
			return true
		})

		if count != 2 {
			t.Errorf("want len() == 2, got len() == %d", count)
		}
	})
}

func TestMapKeys(t *testing.T) {
	m := new(Map[int, string])
	m.Store(1, "1")
	m.Store(2, "2")
	m.Store(3, "3")

	t.Run("loop all", func(t *testing.T) {
		iter := MapKeys(m)
		var values []int
		iter(func(v int) bool {
			values = append(values, v)
			return true
		})

		if len(values) != 3 {
			t.Errorf("want len() == 3, got len() == %d", len(values))
		}
		sort.Ints(values)
		got := fmt.Sprintf("%#v", values)
		want := "[]int{1, 2, 3}"
		if got != want {
			t.Errorf("wrong key-value pairs\nwant: %s\ngot:  %s", want, got)
		}
	})

	t.Run("break", func(t *testing.T) {
		iter := MapKeys(m)
		count := 0
		iter(func(v int) bool {
			count++
			if count == 2 {
				return false
			}
			return true
		})

		if count != 2 {
			t.Errorf("want len() == 2, got len() == %d", count)
		}
	})
}

func TestMapAll(t *testing.T) {
	m := new(Map[int, string])
	m.Store(1, "1")
	m.Store(2, "2")
	m.Store(3, "3")

	type Pair struct {
		Key   int
		Value string
	}

	t.Run("loop all", func(t *testing.T) {
		iter := MapAll(m)
		var pairs []Pair
		iter(func(k int, v string) bool {
			pairs = append(pairs, Pair{k, v})
			return true
		})

		if len(pairs) != 3 {
			t.Errorf("want len() == 3, got len() == %d", len(pairs))
		}
		sort.Slice(pairs, func(i, j int) bool { return pairs[i].Key < pairs[j].Key })
		got := fmt.Sprintf("%#v", pairs)
		want := `[]sync2.Pair{sync2.Pair{Key:1, Value:"1"}, sync2.Pair{Key:2, Value:"2"}, sync2.Pair{Key:3, Value:"3"}}`
		if got != want {
			t.Errorf("wrong key-value pairs\nwant: %s\ngot:  %s", want, got)
		}
	})

	t.Run("break", func(t *testing.T) {
		iter := MapAll(m)
		count := 0
		iter(func(k int, v string) bool {
			count++
			if count == 2 {
				return false
			}
			return true
		})

		if count != 2 {
			t.Errorf("want len() == 2, got len() == %d", count)
		}
	})
}

func TestMapInsert(t *testing.T) {
	type Pair struct {
		Key   int
		Value string
	}
	iter := func(yield func(int, string) bool) {
		pairs := []Pair{{1, "1"}, {3, "3"}}
		for _, pair := range pairs {
			if !yield(pair.Key, pair.Value) {
				break
			}
		}
	}

	m := new(Map[int, string])
	m.Store(0, "not from iter")
	MapInsert(m, iter)

	var pairs []Pair
	m.Range(func(key int, value string) bool {
		pairs = append(pairs, Pair{key, value})
		return true
	})
	if len(pairs) != 3 {
		t.Errorf("want len() == 3, got len() == %d", len(pairs))
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].Key < pairs[j].Key })
	got := fmt.Sprintf("%#v", pairs)
	want := `[]sync2.Pair{sync2.Pair{Key:0, Value:"not from iter"}, sync2.Pair{Key:1, Value:"1"}, sync2.Pair{Key:3, Value:"3"}}`
	if got != want {
		t.Errorf("wrong key-value pairs\nwant: %s\ngot:  %s", want, got)
	}
}

func TestMapCollect(t *testing.T) {
	type Pair struct {
		Key   int
		Value string
	}
	iter := func(yield func(int, string) bool) {
		pairs := []Pair{{1, "1"}, {3, "3"}}
		for _, pair := range pairs {
			if !yield(pair.Key, pair.Value) {
				break
			}
		}
	}

	m := MapCollect(iter)

	var pairs []Pair
	m.Range(func(key int, value string) bool {
		pairs = append(pairs, Pair{key, value})
		return true
	})
	if len(pairs) != 2 {
		t.Errorf("want len() == 2, got len() == %d", len(pairs))
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].Key < pairs[j].Key })
	got := fmt.Sprintf("%#v", pairs)
	want := `[]sync2.Pair{sync2.Pair{Key:1, Value:"1"}, sync2.Pair{Key:3, Value:"3"}}`
	if got != want {
		t.Errorf("wrong key-value pairs\nwant: %s\ngot:  %s", want, got)
	}
}
