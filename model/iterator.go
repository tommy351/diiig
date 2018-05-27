package model

// Iterator is an interface which is iterabled.
type Iterator interface {
	Next() bool
	Value() *Element
	Index() int
}
