package models


type Hotel struct {
	name string
	numberOfRooms int
	rooms map[uint]User
}
