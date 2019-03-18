package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name string
	email string
}

// user implement notifier interface
func (u *user) notify() {
	fmt.Printf("user: %s| email: %s\n", u.name, u.email)
}

func sendNotification(n notifier) {  // n is type that implement notifier interface
	n.notify()
}


func main()  {
	bill := user{"bill", "bill@email"}
	sendNotification(&bill)
}