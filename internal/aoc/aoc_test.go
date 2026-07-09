package aoc

import (
	"reflect"
	"testing"
)

func TestInts(t *testing.T) {
	got := Ints("x=12, y=-3 and 0")
	want := []int{12, -3, 0}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Ints = %v, want %v", got, want)
	}
}

func TestAbs(t *testing.T) {
	if Abs(-5) != 5 || Abs(5) != 5 {
		t.Fatal("Abs broken")
	}
}

func TestSum(t *testing.T) {
	if Sum([]int{1, 2, 3}) != 6 {
		t.Fatal("Sum broken")
	}
}
