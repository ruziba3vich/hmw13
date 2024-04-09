package models

import "fmt"

type Hotel struct {
	Name          string
	NumberOfRooms int
	Rooms         map[uint]User
}

func (h Hotel) getEmailContents(title string, name string, nights uint) string {
	text := `from %s...
	Dear %s %s,\n your room reservation for %d night(s) is confirmed. Have a nice day !`
	return fmt.Sprintf(text, h.Name,
		title,
		name,
		nights)
}

func (h Hotel) sendEmail(contents string, to User) {
	to.RecieveMessage(Message{From: h.Name, MessageContext: contents,})
}
