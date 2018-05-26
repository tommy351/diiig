package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet_Add(t *testing.T) {
	s := &Set{
		Collection: new(LinkedList),
	}
	expected := &Element{
		Key:   "foo",
		Value: 0,
	}

	// Do twice to make sure elements is not duplicated.
	for i := 0; i < 2; i++ {
		s.Add("foo")
		assert.Equal(t, expected, s.Get("foo"))
		assert.Equal(t, 1, s.Length())
	}
}

func TestSet_Increment(t *testing.T) {
	t.Run("Exist", func(t *testing.T) {
		s := &Set{
			Collection: makeLinkedList([]Element{
				{Key: "foo", Value: 3},
			}),
		}

		s.Increment("foo", -5)
		assert.Equal(t, &Element{Key: "foo", Value: -2}, s.Get("foo"))
		assert.Equal(t, 1, s.Length())
	})

	t.Run("Not exist", func(t *testing.T) {
		s := &Set{
			Collection: new(LinkedList),
		}

		s.Increment("foo", 42)
		assert.Equal(t, &Element{Key: "foo", Value: 42}, s.Get("foo"))
	})
}

func TestSet_Range(t *testing.T) {
	elements := []Element{
		{Key: "a", Value: 3},
		{Key: "b", Value: 2},
		{Key: "c", Value: 1},
	}
	s := &Set{
		Collection: makeLinkedList(elements),
	}

	tests := []struct {
		Name     string
		Start    int
		End      int
		Expected []int
	}{
		{
			Name:     "Positive indices",
			Start:    0,
			End:      2,
			Expected: []int{0, 1, 2},
		},
		{
			Name:     "Negative indices",
			Start:    -2,
			End:      -1,
			Expected: []int{1, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var expected []Element

			for _, i := range test.Expected {
				expected = append(expected, elements[i])
			}

			assert.Equal(t, expected, s.Range(test.Start, test.End))
		})
	}
}
