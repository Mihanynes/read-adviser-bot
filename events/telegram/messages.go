package telegram

const msgHelp = `I can save and keep you notes. Also I can offer you them to read.

In order to save the note, just choose command /todo in menu.

In order to get a random note from your list, send me command /rnd.
Caution! After that, this note will be removed from your list! You MUST do it)`

const msgHello = "Hi there! 👾\n\n" + msgHelp

const (
	msgUnknownCommand = "Unknown command, choose command in menu 🤔"
	msgNoSavedPages   = "You have no saved notes 🙊"
	msgSaved          = "Saved! 👌"
	msgAlreadyExists  = "You have already have this note in your list 🤗"
)
