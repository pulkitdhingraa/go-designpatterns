package command

import "fmt"

type EditorInvoker struct {
	history []Command
}

func (ed *EditorInvoker) ExecuteCommand(cmd Command) {
	cmd.Execute()
	ed.history = append(ed.history, cmd)
}

func (ed *EditorInvoker) UndoLast() {
	if len(ed.history) == 0 {
		fmt.Println("No prev command available to undo")
	}

	lastCmd := ed.history[len(ed.history)-1]
	lastCmd.Undo()
	ed.history = ed.history[:len(ed.history)-1]
}
