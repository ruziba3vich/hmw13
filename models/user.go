package models

type User struct {
	Email          string
	StartDate      string
	EndDate        string
	EmailBox       []Message
	UnreadMessages []Message
}

func (u *User) MarkAllAsRead() {
	u.UnreadMessages = []Message{}
}


