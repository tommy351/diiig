package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeLinkedList(elements []Element) *LinkedList {
	list := new(LinkedList)

	for i := range elements {
		list.Insert(&elements[i])
	}

	return list
}

func linkedListToSlice(list *LinkedList) (elements []Element) {
	iter := list.Each()

	for iter.Next() {
		elements = append(elements, *iter.Value())
	}

	return
}

func TestLinkedList_Insert(t *testing.T) {
	t.Run("Normal", func(t *testing.T) {
		elements := []Element{
			{Key: "a", Value: 1},
			{Key: "b", Value: 3},
			{Key: "c", Value: 2},
			{Key: "d", Value: 1},
			{Key: "e", Value: 4},
		}
		list := makeLinkedList(elements)

		assert.Equal(t, []Element{elements[4], elements[1], elements[2], elements[3], elements[0]}, linkedListToSlice(list))
		assert.Equal(t, len(elements), list.Length())
	})

	t.Run("Ordered insertion", func(t *testing.T) {
		elements := []Element{
			{Key: "a", Value: 3},
			{Key: "b", Value: 2},
			{Key: "c", Value: 1},
		}
		list := makeLinkedList(elements)

		assert.Equal(t, elements, linkedListToSlice(list))
		assert.Equal(t, len(elements), list.Length())
	})
}

func TestLinkedList_Get(t *testing.T) {
	elements := []Element{
		{Key: "a", Value: 1},
		{Key: "b", Value: 3},
	}
	list := makeLinkedList(elements)

	t.Run("Exist", func(t *testing.T) {
		for _, elem := range elements {
			assert.Equal(t, &elem, list.Get(elem.Key))
		}
	})

	t.Run("Not exist", func(t *testing.T) {
		assert.Nil(t, list.Get("foo"))
	})
}

func TestLinkedList_Each(t *testing.T) {
	t.Run("Not empty", func(t *testing.T) {
		elements := []Element{
			{Key: "a", Value: 3},
			{Key: "b", Value: 2},
			{Key: "c", Value: 1},
		}
		list := makeLinkedList(elements)
		iter := list.Each()
		i := 0

		for iter.Next() {
			assert.Equal(t, &elements[i], iter.Value())
			assert.Equal(t, i, iter.Index())
			i++
		}

		assert.Equal(t, len(elements), i)
	})

	t.Run("Empty", func(t *testing.T) {
		list := new(LinkedList)
		iter := list.Each()
		assert.False(t, iter.Next())
		assert.Nil(t, iter.Value())
	})
}

func TestLinkedList_Remove(t *testing.T) {
	elements := []Element{
		{Key: "a", Value: 3},
		{Key: "b", Value: 2},
		{Key: "c", Value: 1},
	}

	tests := []struct {
		Name     string
		Key      string
		Expected []int
	}{
		{
			Name:     "head",
			Key:      "a",
			Expected: []int{1, 2},
		},
		{
			Name:     "middle",
			Key:      "b",
			Expected: []int{0, 2},
		},
		{
			Name:     "tail",
			Key:      "c",
			Expected: []int{0, 1},
		},
		{
			Name:     "not exist",
			Key:      "foo",
			Expected: []int{0, 1, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			list := makeLinkedList(elements)
			list.Remove(test.Key)

			var expected []Element

			for _, i := range test.Expected {
				expected = append(expected, elements[i])
			}

			assert.Equal(t, expected, linkedListToSlice(list))
			assert.Equal(t, len(expected), list.Length())
		})
	}

	t.Run("multi", func(t *testing.T) {
		elements := []Element{
			{Key: "a", Value: 1},
			{Key: "a", Value: 1},
			{Key: "b", Value: 2},
		}
		list := makeLinkedList(elements)
		list.Remove("a")

		assert.Equal(t, elements[2:], linkedListToSlice(list))
		assert.Equal(t, 1, list.Length())
	})
}
