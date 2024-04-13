package models

import (
	"errors"
	"fmt"
)

type Hotel struct {
	Name          string
	NumberOfRooms int
	Rooms         map[uint]UserWithNights
}

func (h *Hotel) GetEmailContents(eContent EmailContent) string {
	text := `from %s...
	Dear %s %s,\n your room reservation for room %d for %d night(s) is confirmed. Have a nice day !`
	return fmt.Sprintf(text, h.Name,
		eContent.title,
		eContent.name,
		eContent.roomNum,
		eContent.nights)
}

func (h *Hotel) SendEmail(contents string, to *User) {
	sth := Message{From: (*h).Name, MessageContext: contents, MessageLook: "from " + (*h).Name + " " + contents[:15]}
	(*to).RecieveMessage(&sth)
}

func (h *Hotel) ReservateRoom(roomNum uint, u *UserWithNights) error {
	if (*h).Rooms[roomNum].User.HasRoom {
		return errors.New("this room is already in use")
	} else {
		(*h).Rooms[roomNum] = (*u)
	}
	return nil
}
