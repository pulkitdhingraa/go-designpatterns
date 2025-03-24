package command

type TextEditor struct {
	content string
}

func (t *TextEditor) Insert(pos int, text string){
	t.content = t.content[:pos] + text + t.content[pos:]
}

func (t *TextEditor) Delete(pos int, len int) {
	t.content = t.content[:pos] + t.content[pos+len:]
}

func (t *TextEditor) Replace(pos int, len int, text string) {
	t.Delete(pos, len)
	t.Insert(pos, text)
}

func (t *TextEditor) GetContent() string{
	return t.content
}