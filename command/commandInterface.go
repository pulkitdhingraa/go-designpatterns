package command

type Command interface {
	Execute()
	Undo()
}