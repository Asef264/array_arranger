package go_slice

/*
 Package go-slice provides utility functions to work with Go slices.
 Author: Khaled Moazedi
 Email: Xhmoazedi@gmail.com
*/

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShiftElements_ZiroShift(t *testing.T) {
	//ZiroShiftTest
	sa := []string{"khaled", "hamzah", "saeed", "sara"}
	arg := 0
	want := sa
	testy := NewStringArray(sa)
	result := testy.ShiftElements(arg)
	assert.Equal(t, want, result)
}
func TestShiftElements_OneShiftToTheRight(t *testing.T) {
	//OneShiftToTheRight
	sa1 := []string{"khaled", "hamzah", "saeed", "sara"}
	arg1 := -1
	want1 := []string{"sara", "khaled", "hamzah", "saeed"}
	testy1 := NewStringArray(sa1)
	result1 := testy1.ShiftElements(arg1)
	assert.Equal(t, want1, result1)
}

func TestDeleteOneElement(t *testing.T) {
	//DeleteOneElement
	Element := "khaled"
	sa := []string{"khaled", "hamzah", "saeed", "sara"}
	want1 := []string{"hamzah", "saeed", "sara"}
	testy := NewStringArray(sa)
	result := testy.DeleteElement(Element)
	assert.Equal(t, want1, result)
}

func TestDeleteTwoElements(t *testing.T) {
	//DeleteOneElement
	Element := "khaled"
	sa := []string{"khaled", "hamzah", "saeed", "khaled", "sara"}
	want1 := []string{"hamzah", "saeed", "sara"}
	testy := NewStringArray(sa)
	result := testy.DeleteElement(Element)
	assert.Equal(t, want1, result)
}

func TestAppendElementToTheFirst(t *testing.T) {
	Element := "pari"
	sa := []string{"khaled", "hamzah", "saeed", "khaled", "sara"}
	want := []string{"pari", "khaled", "hamzah", "saeed", "khaled", "sara"}
	testy := NewStringArray(sa)
	result := testy.AppendElementToTheFirst(Element)
	assert.Equal(t, want, result)
}

func TestAppendSlices(t *testing.T) {
	first := []string{"khaled", "hamzah", "saeed", "khaled", "sara"}
	second := []string{"man", "to"}
	want1 := []string{"khaled", "hamzah", "saeed", "khaled", "sara", "man", "to"}
	result := AppendSlices(first, second)
	assert.Equal(t, want1, result)
}

func TestToDigits(t *testing.T) {
	first := []string{"0", "43", "3", "2", "1"}
	want1 := []int{0, 43, 3, 2, 1}
	testy := NewStringArray(first)
	result := testy.ToDigits()
	assert.Equal(t, want1, result)
}

func TestToDigitsError(t *testing.T) {
	first := []string{"0", "error", "3", "2", "1"}
	testy := NewStringArray(first)
	result := testy.ToDigits()
	assert.Error(t, errors.ErrUnsupported, result)
}

func TestReverseRangeStringArray(t *testing.T) {
	first := []string{"khaled", "hamzah", "saeed", "sara"}
	second := []string{"sara", "saeed", "hamzah", "khaled"}
	testy := NewStringArray(first)
	index := 0
	v := ""
	out := []string{}
	for {
		if index == len(first) {
			break
		}
		_, v = testy.ReverseRange(index)
		index++
		out = append(out, v)
	}
	assert.Equal(t, second, out)

}

func TestReverse(t *testing.T) {
	first := []string{"khaled", "hamzah", "saeed", "sara"}
	second := []string{"sara", "saeed", "hamzah", "khaled"}

	firstInt := []int{1, 2, 3, 4}
	secondInt := []int{4, 3, 2, 1}

	out := Reverse(first)
	outInt := Reverse(firstInt)
	assert.Equal(t, second, out)
	assert.Equal(t, secondInt, outInt)
}

func TestTypeOf(t *testing.T) {
	first := []string{"khaled", "hamzah", "saeed", "sara"}
	arrayType := TypeOf(first)
	assert.Equal(t, "string", arrayType)

	firstInt := []int{1, 2, 3, 4}

	arrayType1 := TypeOf(firstInt)
	assert.Equal(t, "int", arrayType1)

}

func TestSum(t *testing.T) {
	firstInt := []int{1, 2, 3, 4}
	intSlice := NewIntSlice(firstInt)
	out := intSlice.SumInt()
	assert.Equal(t, int64(10), out)
}

func TestAverage(t *testing.T) {
	firstInt := []int{1, 2, 3, 6}
	intSlice := NewIntSlice(firstInt)
	out := intSlice.Average()
	assert.Equal(t, 3., out)
}

func TestSortAsc(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 3, 8, 1, 2}, []int{1, 2, 3, 5, 8}},
		{[]int{10, 7, 8, 9, 1}, []int{1, 7, 8, 9, 10}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		input := make([]int, len(test.input))
		copy(input, test.input)
		SortAsc(input)
		assert.Equal(t, test.expected, input)
	}
}

func TestSortDesc(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 3, 8, 1, 2}, []int{8, 5, 3, 2, 1}},
		{[]int{10, 7, 8, 9, 1}, []int{10, 9, 8, 7, 1}},
		{[]int{3, 2, 1}, []int{3, 2, 1}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		input := make([]int, len(test.input))
		copy(input, test.input)
		SortDesc(input)
		assert.Equal(t, test.expected, input)
	}
}
