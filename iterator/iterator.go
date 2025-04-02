package iterator

type Iterator interface {
	Next() *Song
	Prev() *Song
	HasNext() bool
	HasPrev() bool
	Reset()
}