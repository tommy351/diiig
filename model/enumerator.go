package model

// Enumerator represents a way of iteration.
type Enumerator interface {
	Next() bool
	Value() *Element
}
