package utility

import (
	"reflect"
	"testing"
)

func TestSliceStringDifference(t *testing.T) {
	tests := []struct {
		a    []string
		b    []string
		want []string
	}{
		{
			[]string{"a", "b"},
			[]string{"b", "c"},
			[]string{"a"},
		},
		{
			[]string{"1", "2", "3"},
			[]string{"1", "2", "3"},
			[]string{},
		},
		{
			[]string{"1", "2", "3"},
			[]string{"a", "b", "c"},
			[]string{"1", "2", "3"},
		},
		{
			[]string{},
			[]string{},
			[]string{},
		},
		{
			[]string{"aa", "bb", "cc"},
			[]string{},
			[]string{"aa", "bb", "cc"},
		},
	}
	for _, test := range tests {
		result := SliceStringDifference(test.a, test.b)
		if !reflect.DeepEqual(result, test.want) {
			t.Errorf("Difference of %s & %s was incorrect, got: %s, wanted: %s", test.a, test.b, result, test.want)
		}
	}
}
