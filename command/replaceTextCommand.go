package command

type ReplaceTextCommand struct {
	editor Editor
	pos    int
	len    int
	text   string
	oldText string
}

func (r *ReplaceTextCommand) Execute() {
	r.oldText = r.editor.GetContent()[r.pos:r.pos+r.len]
	r.editor.Delete(r.pos,r.len)
	r.editor.Insert(r.pos,r.text)
}

func (r *ReplaceTextCommand) Undo() {
	r.editor.Delete(r.pos, len(r.text))
	r.editor.Insert(r.pos, r.oldText)
}
