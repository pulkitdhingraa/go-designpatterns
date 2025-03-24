package command

type DeleteTextCommand struct {
	editor Editor
	pos int
	len int
	deleted string
}

func (d *DeleteTextCommand) Execute() {
	d.deleted = d.editor.GetContent()[d.pos:d.pos+d.len]
	d.editor.Delete(d.pos, d.len)
}

func (d *DeleteTextCommand) Undo() {
	d.editor.Insert(d.pos, d.deleted)
}