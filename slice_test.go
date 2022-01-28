// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import (
	"fmt"
	"testing"
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

func TestInserted(t *testing.T) {
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
			got := Inserted([]byte(tc.slice), tc.index, insertion)
			gotStr := string(got)
			if gotStr != tc.want {
				t.Errorf("want %q, got %q", tc.want, got)
			}
		})
	}
}

func TestInsertedSlice(t *testing.T) {
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
			got := InsertedSlice([]byte(tc.slice), tc.index, []byte(insertion))
			gotStr := string(got)
			if gotStr != tc.want {
				t.Errorf("want %q, got %q", tc.want, got)
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
	assertComparable(t, "group[0]", 3, got[0].Count)
	if got[1].Key != 'H' {
		t.Errorf("want group[1].Key = 'H', got %q", got[1].Key)
	}
	assertComparable(t, "group[1]", 2, got[1].Count)
	if got[2].Key != 'T' {
		t.Errorf("want group[2].Key = 'T', got %q", got[2].Key)
	}
	assertComparable(t, "group[2]", 1, got[2].Count)
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
		assertComparable(t, fmt.Sprintf("got[%d]", i), want[i], string(got[i][:]))
	}
}

func TestPairsIter(t *testing.T) {
	in := []byte("abcdefg")
	var got [][]byte
	PairsIter(in, func(a, b byte) {
		got = append(got, []byte{a, b})
	})
	if len(got) != 6 {
		t.Fatalf("want len=6, got len=%d", len(got))
	}
	want := []string{
		"ab", "bc", "cd", "de", "ef", "fg",
	}
	for i := range got {
		assertComparable(t, fmt.Sprintf("got[%d]", i), want[i], string(got[i]))
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
		assertComparable(t, fmt.Sprintf("got[%d]", i), want[i], string(got[i]))
	}
}

func TestWindowedIter(t *testing.T) {
	in := []byte("abcdefg")
	var got [][]byte
	WindowedIter(in, 3, func(window []byte) {
		got = append(got, window)
	})
	if len(got) != 5 {
		t.Fatalf("want len=5, got len=%d", len(got))
	}
	want := []string{
		"abc", "bcd", "cde", "def", "efg",
	}
	for i := range got {
		assertComparable(t, fmt.Sprintf("got[%d]", i), want[i], string(got[i]))
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
		assertComparable(t, fmt.Sprintf("got[%d]", i), want[i], string(got[i]))
	}
}

func TestChunkIter(t *testing.T) {
	in := []byte("abcdefg")
	var got [][]byte
	ChunkIter(in, 3, func(chunk []byte) {
		got = append(got, chunk)
	})
	if len(got) != 3 {
		t.Fatalf("want len=3, got len=%d", len(got))
	}
	want := []string{
		"abc", "def", "g",
	}
	for i := range got {
		assertComparable(t, fmt.Sprintf("got[%d]", i), want[i], string(got[i]))
	}
}

func assertComparable[T comparable](t *testing.T, name string, want T, got T) {
	if want != got {
		t.Errorf(`%s: want "%v", got "%v"`, name, want, got)
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
