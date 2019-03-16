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
}
