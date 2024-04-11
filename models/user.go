package models

import "errors"

type User struct {
	Email          string
	StartDate      string
	EndDate        string
	EmailBox       []Message
	UnreadMessages []Message
	HasRoom        bool
}

func (u *User) MarkAllAsRead() {
	u.UnreadMessages = []Message{}
}

func (u User) GetAllMessages() []Message {
	return u.EmailBox
}

func (u User) GetUnreadMessages() []Message {
	return u.UnreadMessages
}

func (u *User) RecieveMessage(m Message) {
	u.EmailBox = append(u.EmailBox, m)
	u.UnreadMessages = append(u.UnreadMessages, m)
}

func (u *User) ReadMessage(index int, fromBox bool) error {
	if fromBox {
		if len(u.EmailBox) > index {
			u.EmailBox[index].Status = false
			return nil
		} else {
			return errors.New("message not found")
		}
	} else {
		if len(u.UnreadMessages) > index {
			u.UnreadMessages[index].Status = false
			return nil
		} else {
			return errors.New("message not found")
		}
	}
}

func (u User) OrderRoom(h Hotel, roomNum uint, nights uint) (bool, error) {
	err := h.ReservateRoom(roomNum, &UserWithNights{u, nights})
	if err == nil {
		return true, nil
	}
	return false, err
}
