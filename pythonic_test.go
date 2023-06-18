package pythonic

import (
	"testing"
)

func TestIN__String(t *testing.T) {
	testSlice := []string{"t", "A", "d"}
	if !IN(testSlice, "t") {
		t.Error("`t` is in slice and should be found.")
	}
	if IN(testSlice, "e") {
		t.Error("`e` is not in slice and shouldn't be found.")
	}
}

func TestIN__Int(t *testing.T) {
	testSlice := []int{1, 9, 10}
	if !IN(testSlice, 1) {
		t.Error("`1` is in slice and should be found.")
	}
	if IN(testSlice, 8) {
		t.Error("`8` is not in slice and shouldn't be found.")
	}
}

func TestSet__String(t *testing.T) {
	testSlice := []string{"t", "t", "d", "e", "a", "d"}
	output := Set(testSlice)
	if len(output) != 4 {
		t.Error("Not all duplicates were removed from slice", output)
	}
}

func TestSet__Int(t *testing.T) {
	testSlice := []int{1, 3, 4, 2, 1, 4, 5, 6}
	output := Set(testSlice)
	if len(output) != 6 {
		t.Error("Not all duplicates were removed from slice", output)
	}
}

func TestRemove__String(t *testing.T) {
	testSlice := []string{"t", "t", "d", "e", "a", "d"}
	output := Remove(testSlice, "a")
	if IN(output, "a") {
		t.Error("`a` should not be present in slice anymore", testSlice, output)
	}
}

func TestRemove__Int(t *testing.T) {
	testSlice := []int{1, 3, 4, 2, 1, 4, 5, 6}
	output := Remove(testSlice, 5)
	if IN(output, 5) {
		t.Error("`5` should not be present in slice anymore", testSlice, output)
	}
}

func TestCounter__String(t *testing.T) {
	testSlice := []string{"T", "T", "a", "a", "a", "es"}
	output := Counter(testSlice)
	if output["T"] != 2 || output["a"] != 3 || output["es"] != 1 {
		t.Error("We expected other count of items", testSlice, output)
	}
}
