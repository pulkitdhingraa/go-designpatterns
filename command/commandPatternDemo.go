package command

import "fmt"

func Run() {
	// 1. Create receiver object
	editor := &TextEditor{}
	// 2. Create invoker object
	manager := &EditorInvoker{}

	// 3. Create commands

	insertCmd := &InsertTextCommand{
		editor: editor,
		pos:    0,
		text:   "Hello",
	}

	manager.ExecuteCommand(insertCmd)
	fmt.Println(editor.GetContent())

	insertCmd2 := &InsertTextCommand{
		editor: editor,
		pos:    5,
		text:   " Konichiwa",
	}
	manager.ExecuteCommand(insertCmd2)
	fmt.Println(editor.GetContent())

	manager.UndoLast()
	fmt.Println(editor.GetContent())
}