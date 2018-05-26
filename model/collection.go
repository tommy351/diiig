package model

type Collection interface {
	Insert(element *Element)
	Get(key string) *Element
	Remove(key string)
	Each() Enumerator
	Length() int
}
