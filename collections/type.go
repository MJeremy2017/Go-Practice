package main

import "fmt"

func main() {
	type user struct {
		name string
		email string
		ext int
		privileged bool
	}

	var bill user  // initialise with var

	fmt.Println(bill)

	lisa := user{
		name: "Lisa bailey",
		email: "lisa@email.com",
		ext: 123,
		privileged:true,
	}

	fmt.Println(lisa)

	type admin struct {
		person user
		level string
	}

	fred := admin{
		person: user{
			"fred",
			"fred@email.com",
			123,
			true,
		},
		level: "super",
	}
	fmt.Println(fred)

	type Duration int16

	dur := Duration(100)
	fmt.Println(dur)
	fmt.Println(&dur)  // the address if dur
	fmt.Println(*&dur)

	john := usr{"john", "john@email.com"}
	john.notify()
	john.changeEmail("newJohn@email.com")
	fmt.Println(john.email)

	// pointer(the address) can call it
	lily := &usr{"lily", "lily@email.com"}
	lily.notify()

	// pointer receiver
	foley := &usr{"foley", "foley@email.com"}
	foley.changeEmail("newFoley@email.com")
	fmt.Println(*foley)

}

// method
// value receiver
type usr struct {
	name string
	email string
}

func (u usr) notify() {
	fmt.Printf("Sending notification to %s <%s> \n", u.name, u.email)
}

// pointer receiver
func (u *usr) changeEmail(newEmail string)  {
	oldEmail := u.email
	u.email = newEmail
	fmt.Printf("Email changed from %s to %s", oldEmail, u.email)
}
