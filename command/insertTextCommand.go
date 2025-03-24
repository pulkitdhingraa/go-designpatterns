package command

type InsertTextCommand struct {
	editor Editor
	pos    int
	text   string
}

func (i *InsertTextCommand) Execute() {
	i.editor.Insert(i.pos, i.text)
}

func (i *InsertTextCommand) Undo() {
	i.editor.Delete(i.pos, len(i.text))
}