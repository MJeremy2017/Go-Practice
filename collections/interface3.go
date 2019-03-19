package main

import "fmt"

type user2 struct {
	name string
	email string
}

type admin2 struct {
	user2  // embed type user
	level string
}

type notifier2 interface {
	notify()
}

func sendNotification2(n notifier2) {
	n.notify()
}

func (u *user2) notify() {
	fmt.Printf("Sending email %s <%s> \n", u.name, u.email)
}

func main()  {
	lily := admin2{
		user2{"lily", "lily@email.com"},
		"super",
	}

	// can access inner type method
	lily.notify()
	// inner method is also promoted
	lily.user2.notify()
	// outer type also implement inner type interface
	sendNotification2(&lily)
}
