// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package slices

import (
	"fmt"
	"testing"

	"gopkg.in/typ.v4/internal/assert"
)

func TestFill(t *testing.T) {
	slice := make([]int, 127)
	Fill(slice, 42)
	for i, v := range slice {
		if v != 42 {
			t.Errorf("index %d: want 42, got %d", i, v)
		}
	}
}

func TestInsert(t *testing.T) {
	const insertion = '_'
	testCases := []struct {
		name  string
		slice string
		index int
		want  string
	}{
		{
			name:  "empty",
			slice: "",
			index: 0,
			want:  "_",
		},
		{
			name:  "start",
			slice: "start",
			index: 0,
			want:  "_start",
		},
		{
			name:  "end",
			slice: "end",
			index: 3,
			want:  "end_",
		},
		{
			name:  "middle",
			slice: "middle",
			index: 3,
			want:  "mid_dle",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			slice := []byte(tc.slice)
			Insert(&slice, tc.index, insertion)
			gotStr := string(slice)
			if gotStr != tc.want {
				t.Errorf("want %q, got %q", tc.want, gotStr)
			}
		})
	}
}

func TestInsertSlice(t *testing.T) {
	const insertion = "123"
	testCases := []struct {
		name  string
		slice string
		index int
		want  string
	}{
		{
			name:  "empty",
			slice: "",
			index: 0,
			want:  "123",
		},
		{
			name:  "start",
			slice: "start",
			index: 0,
			want:  "123start",
		},
		{
			name:  "end",
			slice: "end",
			index: 3,
			want:  "end123",
		},
		{
			name:  "middle",
			slice: "middle",
			index: 3,
			want:  "mid123dle",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			slice := []byte(tc.slice)
			InsertSlice(&slice, tc.index, []byte(insertion))
			gotStr := string(slice)
			if gotStr != tc.want {
				t.Errorf("want %q, got %q", tc.want, gotStr)
			}
		})
	}
}

func TestRemove2(t *testing.T) {
	testCases := []struct {
		name  string
		slice string
		index int
		want  string
	}{
		{
			name:  "start",
			slice: "_start",
			index: 0,
			want:  "start",
		},
		{
			name:  "end",
			slice: "end_",
			index: 3,
			want:  "end",
		},
		{
			name:  "middle",
			slice: "mid_dle",
			index: 3,
			want:  "middle",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			slice := []byte(tc.slice)
			Remove(&slice, tc.index)
			gotStr := string(slice)
			if gotStr != tc.want {
				t.Errorf("want %q, got %q", tc.want, gotStr)
			}
		})
	}
}

func TestRemoveSlice(t *testing.T) {
	const length = 3
	testCases := []struct {
		name  string
		slice string
		index int
		want  string
	}{
		{
			name:  "start",
			slice: "123start",
			index: 0,
			want:  "start",
		},
		{
			name:  "end",
			slice: "end123",
			index: 3,
			want:  "end",
		},
		{
			name:  "middle",
			slice: "mid123dle",
			index: 3,
			want:  "middle",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			slice := []byte(tc.slice)
			RemoveSlice(&slice, tc.index, length)
			gotStr := string(slice)
			if gotStr != tc.want {
				t.Errorf("want %q, got %q", tc.want, gotStr)
			}
		})
	}
}

func TestFold(t *testing.T) {
	testCases := []struct {
		name  string
		slice []string
		seed  string
		want  string
	}{
		{
			name:  "values",
			slice: []string{"a", "b", "c"},
			seed:  "",
			want:  "abc",
		},
		{
			name:  "nil",
			slice: nil,
			seed:  "seed",
			want:  "seed",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotStr := Fold(tc.slice, tc.seed, func(state string, seed string) string {
				return state + seed
			})
			if gotStr != tc.want {
				t.Errorf("want %q, got %q", tc.want, gotStr)
			}
		})
	}
}

func TestFoldReverse(t *testing.T) {
	testCases := []struct {
		name  string
		slice []string
		seed  string
		want  string
	}{
		{
			name:  "values",
			slice: []string{"a", "b", "c"},
			seed:  "",
			want:  "cba",
		},
		{
			name:  "nil",
			slice: nil,
			seed:  "seed",
			want:  "seed",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotStr := FoldReverse(tc.slice, tc.seed, func(state string, seed string) string {
				return state + seed
			})
			if gotStr != tc.want {
				t.Errorf("want %q, got %q", tc.want, gotStr)
			}
		})
	}
}

func TestConcat(t *testing.T) {
	testCases := []struct {
		a    string
		b    string
		want string
	}{
		{
			a:    "abc",
			b:    "def",
			want: "abcdef",
		},
		{
			a:    "abc",
			b:    "",
			want: "abc",
		},
		{
			a:    "",
			b:    "def",
			want: "def",
		},
		{
			a:    "",
			b:    "",
			want: "",
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s_%s", tc.a, tc.b), func(t *testing.T) {
			c := Concat([]byte(tc.a), []byte(tc.b))
			got := string(c)
			if tc.want != got {
				t.Errorf("want %q, got %q", tc.want, got)
			}
		})
	}
}

func TestGroupBy(t *testing.T) {
	in := []string{
		"Potatoes",
		"Hamburger",
		"Pizza",
		"Toast",
		"Hummus",
		"Pancake",
	}
	got := GroupBy(in, func(value string) byte {
		return value[0]
	})
	if len(got) != 3 {
		t.Fatalf("want 3 groups, got %d: %v", len(got), got)
	}
	if got[0].Key != 'P' {
		t.Errorf("want group[0].Key = 'P', got %q", got[0].Key)
	}
	assertSlice(t, "group[0]", []string{"Potatoes", "Pizza", "Pancake"}, got[0].Values)
	if got[1].Key != 'H' {
		t.Errorf("want group[1].Key = 'H', got %q", got[1].Key)
	}
	assertSlice(t, "group[1]", []string{"Hamburger", "Hummus"}, got[1].Values)
	if got[2].Key != 'T' {
		t.Errorf("want group[2].Key = 'T', got %q", got[2].Key)
	}
	assertSlice(t, "group[2]", []string{"Toast"}, got[2].Values)
}

func TestCountBy(t *testing.T) {
	in := []string{
		"Potatoes",
		"Hamburger",
		"Pizza",
		"Toast",
		"Hummus",
		"Pancake",
	}
	got := CountBy(in, func(value string) byte {
		return value[0]
	})
	if len(got) != 3 {
		t.Fatalf("want 3 groups, got %d: %v", len(got), got)
	}
	if got[0].Key != 'P' {
		t.Errorf("want group[0].Key = 'P', got %q", got[0].Key)
	}
	assert.Comparable(t, "group[0]", 3, got[0].Count)
	if got[1].Key != 'H' {
		t.Errorf("want group[1].Key = 'H', got %q", got[1].Key)
	}
	assert.Comparable(t, "group[1]", 2, got[1].Count)
	if got[2].Key != 'T' {
		t.Errorf("want group[2].Key = 'T', got %q", got[2].Key)
	}
	assert.Comparable(t, "group[2]", 1, got[2].Count)
}

func TestPairs(t *testing.T) {
	in := []byte("abcdefg")
	got := Pairs(in)
	if len(got) != 6 {
		t.Fatalf("want len=6, got len=%d", len(got))
	}
	want := []string{
		"ab", "bc", "cd", "de", "ef", "fg",
	}
	for i := range got {
		assert.Comparable(t, fmt.Sprintf("got[%d]", i), want[i], string(got[i][:]))
	}
}

func TestPairsFunc(t *testing.T) {
	in := []byte("abcdefg")
	var got [][]byte
	PairsFunc(in, func(a, b byte) {
		got = append(got, []byte{a, b})
	})
	if len(got) != 6 {
		t.Fatalf("want len=6, got len=%d", len(got))
	}
	want := []string{
		"ab", "bc", "cd", "de", "ef", "fg",
	}
	for i := range got {
		assert.Comparable(t, fmt.Sprintf("got[%d]", i), want[i], string(got[i]))
	}
}

func TestWindowed(t *testing.T) {
	in := []byte("abcdefg")
	got := Windowed(in, 3)
	if len(got) != 5 {
		t.Fatalf("want len=5, got len=%d", len(got))
	}
	want := []string{
		"abc", "bcd", "cde", "def", "efg",
	}
	for i := range got {
		assert.Comparable(t, fmt.Sprintf("got[%d]", i), want[i], string(got[i]))
	}
}

func TestWindowedFunc(t *testing.T) {
	in := []byte("abcdefg")
	var got [][]byte
	WindowedFunc(in, 3, func(window []byte) {
		got = append(got, window)
	})
	if len(got) != 5 {
		t.Fatalf("want len=5, got len=%d", len(got))
	}
	want := []string{
		"abc", "bcd", "cde", "def", "efg",
	}
	for i := range got {
		assert.Comparable(t, fmt.Sprintf("got[%d]", i), want[i], string(got[i]))
	}
}

func TestChunk(t *testing.T) {
	in := []byte("abcdefg")
	got := Chunk(in, 3)
	if len(got) != 3 {
		t.Fatalf("want len=3, got len=%d", len(got))
	}
	want := []string{
		"abc", "def", "g",
	}
	for i := range got {
		assert.Comparable(t, fmt.Sprintf("got[%d]", i), want[i], string(got[i]))
	}
}

func TestChunkFunc(t *testing.T) {
	in := []byte("abcdefg")
	var got [][]byte
	ChunkFunc(in, 3, func(chunk []byte) {
		got = append(got, chunk)
	})
	if len(got) != 3 {
		t.Fatalf("want len=3, got len=%d", len(got))
	}
	want := []string{
		"abc", "def", "g",
	}
	for i := range got {
		assert.Comparable(t, fmt.Sprintf("got[%d]", i), want[i], string(got[i]))
	}
}

func assertSlice[T comparable](t *testing.T, name string, want, got []T) {
	if len(want) != len(got) {
		t.Errorf("%s: want len=%d, got len=%d", name, len(want), len(got))
		return
	}
	for i := range want {
		if want[i] != got[i] {
			t.Errorf(`%s: index %d: want "%v", got "%v"`, name, i, want[i], got[i])
		}
	}
}
