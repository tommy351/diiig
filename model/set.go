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

func (s *Set) Add(key string) {
	if s.Get(key) == nil {
		s.Insert(&Element{
			Key: key,
		})
	}
}

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

func (s *Set) Range(start, end int) (elements []Element) {
	if start < 0 {
		start = s.Length() + start
	}

	if end < 0 {
		end = s.Length() + end
	}

	i := 0
	enum := s.Each()

	for enum.Next() && i <= end {
		if i >= start {
			elements = append(elements, *enum.Value())
		}

		i++
	}

	return
}
