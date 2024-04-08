// Sizning hamkasblaringizdan biri noyob main.go faylidan iborat dastur yaratdi
// Kodning barqarorligini yaxshilash uchun sizdan kodni qayta tahrirlashingiz so'raladi (refactor).
// Yangi kod tashkilotini taklif qiling:

// Qaysi paketlarni yaratishingiz kerak?

// Yangi katalog yaratish kerakmi?

package main

import "fmt"

func main() {
	// first reservation
	customerName := "Doe"
	customerEmail := "john.doe@example.com"
	var nights uint = 12
	emailContents := getEmailContents("M", customerName, nights)
	sendEmail(emailContents, customerEmail)
	createAndSaveInvoice(customerName, nights, 145.32)
}

// send an email
func sendEmail(contents string, to string) {
	// ...
	// ...
}

// prepare email template
func getEmailContents(title string, name string, nights uint) string {
	text := "Dear %s %s,\n your room reservation for %d night(s) is confirmed. Have a nice day !"
	return fmt.Sprintf(text,
		title,
		name,
		nights)
}

// create the invoice for the reservation
func createAndSaveInvoice(name string, nights uint, price float32) {
	// ...
}
