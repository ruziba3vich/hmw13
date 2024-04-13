// Sizning hamkasblaringizdan biri noyob main.go faylidan iborat dastur yaratdi
// Kodning barqarorligini yaxshilash uchun sizdan kodni qayta tahrirlashingiz so'raladi (refactor).
// Yangi kod tashkilotini taklif qiling:

// Qaysi paketlarni yaratishingiz kerak?

// Yangi katalog yaratish kerakmi?

package main

import (
	"fmt"
	m "hmw13/models"
	"math"
	"math/rand"
	"time"
)

func main() {
	var hotels []m.Hotel
	CreateDummyHotels(&hotels)

	u := m.User{}

	fmt.Println("Enter your email : ")
	fmt.Scan(u.Email)
	status := 1

	for status == 1 {
		fmt.Println("    1 -> Order a room")
		fmt.Println("    2 -> See all messages")
		fmt.Println("    3 -> See unread messages")
		fmt.Println("    4 -> Get info about my room")
		fmt.Print("  --> ")
		fmt.Scan(&status)
		if status == 1 {
			for i := 0; i < len(hotels); i++ {
				fmt.Println(i+1, hotels[i])
			}
			fmt.Print("Enter the order number of the hotel : ")
			var ordNum uint
			fmt.Scan(&ordNum)
			fmt.Print("How many nights you wanna book the room for ? : ")
			var nights uint
			fmt.Scan(&nights)
			rooms := GetAvailableRooms(hotels[ordNum], 0, []uint{})

			height := GetClosestPerfectSquare(len(rooms))

			for i := range rooms {
				fmt.Print(i, " ")
				if i%height == 0 {
					fmt.Println()
				}
			}
			var roomNumber uint
			fmt.Print("Choose one of the rooms you see above : ")
			fmt.Scan(&roomNumber)
			ok, err := u.OrderRoom(hotels[ordNum], roomNumber, nights)
			if ok {
				fmt.Println("Check your e-mail for results")
			} else {
				fmt.Println(err)
			}
		} else if status == 2 {
			messages := u.GetAllMessages()
			for i := range messages {
				fmt.Println(messages[i].MessageLook)
			}
			var choice int
			fmt.Println()
			fmt.Println("	1 -> Mark all messages as read")
			fmt.Println("	2 -> Open a message")
			fmt.Print("  -> ")
			fmt.Scan(&choice)
			if choice == 1 {
				u.MarkAllAsRead()
			} else {
				fmt.Print("   Enter the corresponding number to the messgae : ")
				var correspondingNumber int
				fmt.Scan(&correspondingNumber)
				err := u.ReadMessage(correspondingNumber, true)
				if err == nil {
					fmt.Println(u.EmailBox[correspondingNumber])
				} else {
					fmt.Println(err)
				}
			}
		} else if status == 3 {
			messages := u.GetUnreadMessages()
			for i := range messages {
				fmt.Println(i+1, messages[i].MessageLook)
			}
			var choice int
			fmt.Println()
			fmt.Println("	1 -> Mark all messages as read")
			fmt.Println("	2 -> Open a message")
			fmt.Print("  -> ")
			fmt.Scan(&choice)
			if choice == 1 {
				u.MarkAllAsRead()
			} else {
				fmt.Print("   Enter the corresponding number to the messgae : ")
				var correspondingNumber int
				fmt.Scan(&correspondingNumber)
				err := u.ReadMessage(correspondingNumber, false)
				if err == nil {
					fmt.Println(u.EmailBox[correspondingNumber])
				} else {
					fmt.Println(err)
				}
			}
		} else if status == 4 {
			if u.HasRoom {
				fmt.Println("You've got a room starting from", u.StartDate, " to", u.EndDate)
			} else {
				fmt.Println("You ain't got a room yet . . .")
			}
		} else {
			break
		}
		fmt.Print("You want to keep the process going ? 1 / Yes ; 2 / No : ")
		fmt.Scan(&status)
	}

	fmt.Println("Thank you for using the program !")
}

func CreateDummyHotels(hotels *[]m.Hotel) {
	htls := []string{

		"The Grandview",
		"The Excelsior",
		"The Fairmont",
		"The Landmark",
		"The Ambassador",

		"The Loft",
		"The Metro",
		"The Urbanite",
		"The Skyline",
		"The Element",

		"The Briar Rose",
		"The Willow Inn",
		"The Firefly Lodge",
		"The Hummingbird Retreat",
		"The Moonlight Bay",

		"The Riverview",
		"The Mountainside",
		"The Seaside",
		"The Palm Breeze",
		"The Forest Glen",
	}

	for _, hotelName := range htls {
		seed := time.Now().UnixNano()
		source := rand.NewSource(seed)
		myRand := rand.New(source)

		randomNumber := myRand.Intn(500)
		randomNumber += 400

		hotel := m.Hotel{
			Name:          hotelName,
			NumberOfRooms: randomNumber,
			Rooms:         map[uint]m.UserWithNights{},
		}

		for i := 1; i < randomNumber; i++ {
			hotel.Rooms[uint(i)] = m.UserWithNights{}
		}
		*hotels = append(*hotels, hotel)
	}
}

func GetAvailableRooms(h m.Hotel, ind uint, availableRooms []uint) (rooms []uint) {
	if int(ind) < len(h.Rooms) {
		if !h.Rooms[ind].User.HasRoom {
			availableRooms = append(availableRooms, ind)
			return GetAvailableRooms(h, ind+1, availableRooms)
		}
	}
	return availableRooms
}

func GetClosestPerfectSquare(num int) int {
	if num < 1 {
		return 0
	}
	sqrt := math.Sqrt(float64(num))
	if sqrt == math.Floor(sqrt) {
		return num
	}
	less := GetClosestPerfectSquare(num - 1)
	greater := GetClosestPerfectSquare(num + 1)
	if num-less < greater-num {
		return less
	} else {
		return greater
	}
}
