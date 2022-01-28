// SPDX-FileCopyrightText: 2022 Kalle Fagerberg
//
// SPDX-License-Identifier: MIT

package typ

import (
	"fmt"
	"testing"
)

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
		t.Errorf("want group[1].Key = 'P', got %q", got[1].Key)
	}
	assertSlice(t, "group[1]", []string{"Hamburger", "Hummus"}, got[1].Values)
	if got[2].Key != 'T' {
		t.Errorf("want group[2].Key = 'P', got %q", got[2].Key)
	}
	assertSlice(t, "group[2]", []string{"Toast"}, got[2].Values)
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
