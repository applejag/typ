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
