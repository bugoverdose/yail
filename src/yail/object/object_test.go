package object

import (
	"fmt"
	"testing"
)

func TestHashKeyComparison(t *testing.T) {
	tests := []struct {
		value1   Hashable
		value2   Hashable
		expected bool
	}{
		{NewString("Hello"), NewString("Hello"), true},
		{NewString("Hello"), NewString("World"), false},
		{TRUE, TRUE, true},
		{FALSE, FALSE, true},
		{TRUE, FALSE, false},
		{NewInteger(1), NewInteger(1), true},
		{NewInteger(1), NewInteger(2), false},
		{NewString("Hello"), TRUE, false},
		{NewString("Hello"), NewInteger(1), false},
		{NewInteger(1), FALSE, false},
	}
	for i, tt := range tests {
		actual := tt.value1.HashKey() == tt.value2.HashKey()
		fmt.Println(tt.value1.HashKey(), tt.value2.HashKey())
		if actual != tt.expected {
			t.Errorf("test %d: expected %+v to be %+v", i+1, actual, tt.expected)
		}
	}
}
