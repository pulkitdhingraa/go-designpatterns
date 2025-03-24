package command

type Editor interface {
	Insert(pos int, text string)
	Delete(pos int, len int)
	Replace(pos int, len int, text string)
	GetContent() string
}
