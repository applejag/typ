// SPDX-FileCopyrightText: 2022 Per Alexander Fougner
//
// SPDX-License-Identifier: MIT

package sync2_test

import (
	"fmt"
	"testing"

	"gopkg.in/typ.v3/internal/assert"
	"gopkg.in/typ.v3/pkg/sets"
	"gopkg.in/typ.v3/pkg/sync2"
)

func TestSet_Add(t *testing.T) {
	setA := newSetABC()
	assert.ElementsMatch(t, setA.Slice(), []string{"A", "B", "C"})
}

func TestSet_Remove(t *testing.T) {
	setA := newSetABC()
	setA.Remove("B")
	assert.ElementsMatch(t, setA.Slice(), []string{"A", "C"})
}

func TestSet_Union(t *testing.T) {
	setABC, setBCD := newSetABC(), newSetBCD()
	assert.ElementsMatch(t, []string{"A", "B", "C", "D"}, setABC.Union(setBCD).Slice())
}

func TestSet_Intersect(t *testing.T) {
	setABC, setBCD := newSetABC(), newSetBCD()
	assert.ElementsMatch(t, []string{"B", "C"}, setABC.Intersect(setBCD).Slice())
}

func TestSet_SetDiff(t *testing.T) {
	setABC, setBCD := newSetABC(), newSetBCD()
	assert.ElementsMatch(t, []string{"A"}, setABC.SetDiff(setBCD).Slice())
}

func TestSet_SymDiff(t *testing.T) {
	setABC, setBCD := newSetABC(), newSetBCD()
	assert.ElementsMatch(t, []string{"A", "D"}, setABC.SymDiff(setBCD).Slice())
}

func TestSet_CartesianProduct(t *testing.T) {
	testCases := []struct {
		name string
		A    sets.Set[string]
		B    sets.Set[string]
		want []string
	}{
		{
			name: "ABC x BCD",
			A:    newSetABC(),
			B:    newSetBCD(),
			want: []string{"AB", "AC", "AD", "BB", "BC", "BD", "CB", "CC", "CD"},
		},
		{
			name: "BCD x ABC",
			A:    newSetBCD(),
			B:    newSetABC(),
			want: []string{"BA", "BB", "BC", "CA", "CB", "CC", "DA", "DB", "DC"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotSet := sets.CartesianProduct(tc.A, tc.B)
			var got []string
			for _, v := range gotSet {
				got = append(got, fmt.Sprintf("%s%s", v.A, v.B))
			}

			assert.ElementsMatch(t, tc.want, got)
		})
	}
}

func TestSet_AddSet(t *testing.T) {
	testCases := []struct {
		name       string
		A          sets.Set[string]
		B          sets.Set[string]
		wantAmount int
		want       []string
	}{
		{
			name:       "add set added count excludes existing",
			A:          newSetABC(),
			B:          newSetBCD(),
			wantAmount: 1,
			want:       []string{"A", "B", "C", "D"},
		},
		{
			name:       "add empty set",
			A:          newSetABC(),
			B:          &sync2.Set[string]{},
			wantAmount: 0,
			want:       []string{"A", "B", "C"},
		},
		{
			name:       "add to empty set",
			A:          &sync2.Set[string]{},
			B:          newSetABC(),
			wantAmount: 3,
			want:       []string{"A", "B", "C"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			aClone := tc.A.Clone()
			bClone := tc.B.Clone()
			gotAmount := aClone.AddSet(bClone)
			assert.Comparable(t, "amount", tc.wantAmount, gotAmount)
			assert.ElementsMatch(t, tc.want, aClone.Slice())
		})
	}
}

func TestSet_RemoveSet(t *testing.T) {
	testCases := []struct {
		name       string
		A          sets.Set[string]
		B          sets.Set[string]
		wantAmount int
		want       []string
	}{
		{
			name:       "remove set removed count excludes non-existing",
			A:          newSetABC(),
			B:          newSetBCD(),
			wantAmount: 2,
			want:       []string{"A"},
		},
		{
			name:       "remove empty set",
			A:          newSetABC(),
			B:          &sync2.Set[string]{},
			wantAmount: 0,
			want:       []string{"A", "B", "C"},
		},
		{
			name:       "remove from empty set",
			A:          &sync2.Set[string]{},
			B:          newSetABC(),
			wantAmount: 0,
			want:       []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			aClone := tc.A.Clone()
			bClone := tc.B.Clone()
			gotAmount := aClone.RemoveSet(bClone)
			assert.Comparable(t, "amount", tc.wantAmount, gotAmount)
			assert.ElementsMatch(t, tc.want, aClone.Slice())
		})
	}
}

func TestSet_OriginalNotAffectedByRemoveOnClone(t *testing.T) {
	setABC := newSetABC()
	setABCClone := setABC.Clone()
	setABCClone.Remove("A")
	assert.ElementsMatch(t, []string{"B", "C"}, setABCClone.Slice())
	assert.ElementsMatch(t, []string{"A", "B", "C"}, setABC.Slice())
}

func TestSet_OriginalNotAffectedByAddOnClone(t *testing.T) {
	setABC := newSetABC()
	setABCClone := setABC.Clone()
	setABCClone.Add("D")
	assert.ElementsMatch(t, []string{"A", "B", "C", "D"}, setABCClone.Slice())
	assert.ElementsMatch(t, []string{"A", "B", "C"}, setABC.Slice())
}

func TestSet_Range(t *testing.T) {
	setABC := newSetABC()

	var slice1 []string
	setABC.Range(func(value string) bool {
		slice1 = append(slice1, value)
		return true
	})
	assert.Comparable(t, "no interrupt", len(slice1), len(setABC.Slice()))

	var slice2 []string
	setABC.Range(func(value string) bool {
		slice2 = append(slice2, value)
		return len(slice2) != 2
	})

	assert.Comparable(t, "interrupts at length=2", len(slice2), 2)
}

func TestSet_String(t *testing.T) {
	set1 := &sync2.Set[string]{}
	set1.Add("A")
	assert.Comparable(t, "one value", "{A}", set1.String())

	set1.Add("B")
	str := set1.String()

	assert.Comparable(t, "strings is one of two possible", true, str == "{A B}" || str == "{B A}")
}

func newSetABC() sets.Set[string] {
	setA := &sync2.Set[string]{}
	setA.Add("A")
	setA.Add("B")
	setA.Add("C")
	return setA
}

func newSetBCD() sets.Set[string] {
	setB := &sync2.Set[string]{}
	setB.Add("B")
	setB.Add("C")
	setB.Add("D")
	return setB
}
