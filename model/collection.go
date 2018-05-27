package model

// Collection represents a group of elements.
type Collection interface {
	Insert(element *Element)
	Get(key string) *Element
	Remove(key string)
	Each() Enumerator
	Length() int
}
