package maps_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/typ.v3/pkg/maps"
	"gopkg.in/typ.v3/pkg/sets"
)

func ExampleSet() {
	set := make(maps.Set[string], 0)
	set.Add("A")
	set.Add("B")
	set.Add("C")

	set.Range(func(value string) bool {
		fmt.Println("Value:", value)
		return true
	})

	// Unordered output:
	// Value: A
	// Value: B
	// Value: C
}

func ExampleSet_setOperations() {
	set1 := make(maps.Set[string], 0)
	set1.Add("A")
	set1.Add("B")
	set1.Add("C")
	fmt.Println("set1:", set1) // {A B C}

	set2 := make(maps.Set[string], 0)
	set2.Add("B")
	set2.Add("C")
	set2.Add("D")
	fmt.Println("set2:", set2) // {B C D}

	fmt.Println("union:", set1.Union(set2))         // {A B C D}
	fmt.Println("intersect:", set1.Intersect(set2)) // {B C}
	fmt.Println("set diff:", set1.SetDiff(set2))    // {A}
	fmt.Println("sym diff:", set1.SymDiff(set2))    // {A D}

	// Please note: the Set.String() output is not sorted!
}

func TestSet_SetOperations(t *testing.T) {
	setA := newSetA()
	assert.ElementsMatch(t, setA.Slice(), []string{"A", "B", "C"}, "setA values")
	setB := newSetB()
	assert.ElementsMatch(t, setB.Slice(), []string{"B", "C", "D"}, "setB values")

	assert.ElementsMatch(t, setA.Union(setB).Slice(), []string{"A", "B", "C", "D"}, "union")
	assert.ElementsMatch(t, setA.Intersect(setB).Slice(), []string{"B", "C"}, "intersect")
	assert.ElementsMatch(t, setA.SetDiff(setB).Slice(), []string{"A"}, "set diff")
	assert.ElementsMatch(t, setA.SymDiff(setB).Slice(), []string{"A", "D"}, "sym diff")

	setAClone := setA.Clone()
	setAClone.Remove("A") // B C
	assert.ElementsMatch(t, setA.Slice(), []string{"A", "B", "C"}, "removing from clone doesn't affect original")

	setAClone.Add("D") // B C D
	assert.ElementsMatch(t, setA.Slice(), []string{"A", "B", "C"}, "adding to clone doesn't affect original")

	numRemoved := setAClone.RemoveSet(setB)
	assert.Equal(t, 3, numRemoved, "remove set works")

	numRemoved2 := setAClone.RemoveSet(setB)
	assert.Equal(t, 0, numRemoved2, "remove set removed count excludes non-existing")

	numAdded := setAClone.AddSet(setB)
	assert.Equal(t, 3, numAdded, "add set works")

	numAdded2 := setAClone.AddSet(setB)
	assert.Equal(t, 0, numAdded2, "add set added count excludes existing")

	var slice1 []string
	setAClone.Range(func(value string) bool {
		slice1 = append(slice1, value)
		return true
	})
	assert.Len(t, slice1, len(setAClone.Slice()), "range works, no interrupt")
	var slice2 []string
	setAClone.Range(func(value string) bool {
		slice2 = append(slice2, value)
		return len(slice2) != 2
	})
	assert.Len(t, slice2, 2, "range works, interrupt at length=2")
}

func TestSet_String(t *testing.T) {
	set1 := make(maps.Set[string], 0)
	set1.Add("A")
	assert.Equal(t, "{A}", set1.String(), "one value")

	set1.Add("B")
	str := set1.String()
	assert.True(t, str == "{A B}" || str == "{B A}", "two values")
}

func Test_CartesianProduct(t *testing.T) {
	setA := newSetA()
	setB := newSetB()

	setAB := maps.NewSetFromKeys(sets.CartesianProduct(setA, setB))

	expected := []sets.Product[string, string]{
		{A: "A", B: "B"},
		{A: "A", B: "C"},
		{A: "A", B: "D"},
		{A: "B", B: "B"},
		{A: "B", B: "C"},
		{A: "B", B: "D"},
		{A: "C", B: "B"},
		{A: "C", B: "C"},
		{A: "C", B: "D"},
	}

	assert.Len(t, setAB.Slice(), len(setA.Slice())*len(setB.Slice()), "length is product of sets' length")
	for _, v := range expected {
		assert.True(t, setAB.Has(v), "setAB has expected values")
	}
}

func newSetA() sets.Set[string] {
	setA := make(maps.Set[string], 0)
	setA.Add("A")
	setA.Add("B")
	setA.Add("C")
	return setA
}

func newSetB() sets.Set[string] {
	setB := make(maps.Set[string], 0)
	setB.Add("B")
	setB.Add("C")
	setB.Add("D")
	return setB
}
