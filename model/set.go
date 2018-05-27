package model

// Element is a key-value structure which put in a set.
type Element struct {
	Key   string
	Value int
}

// Set is a collection which contains unique elements.
type Set struct {
	Collection
}

// Add inserts the key to the set if the key does not exist yet.
func (s *Set) Add(key string) {
	if s.Get(key) == nil {
		s.Insert(&Element{
			Key: key,
		})
	}
}

// Increment adds the value of the element in the set.
func (s *Set) Increment(key string, n int) {
	elem := s.Get(key)

	if elem == nil {
		elem = &Element{
			Key:   key,
			Value: n,
		}

		s.Insert(elem)
	} else {
		s.Remove(key)
		s.Insert(&Element{
			Key:   key,
			Value: elem.Value + n,
		})
	}
}

// Range returns a slice of elements within the specified range.
func (s *Set) Range(start, end int) (elements []Element) {
	if start < 0 {
		start = s.Length() + start
	}

	if end < 0 {
		end = s.Length() + end
	}

	i := 0
	enum := s.Each()

	for i <= end && enum.Next() {
		if i >= start {
			elements = append(elements, *enum.Value())
		}

		i++
	}

	return
}
