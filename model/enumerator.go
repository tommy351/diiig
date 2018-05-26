package model

type Enumerator interface {
	Next() bool
	Value() *Element
}
