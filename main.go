package main

import "fmt"

type user struct {
	id       int64
	name     string
	surname  string
	login    string
	password int64
	role     string
	remove   bool
}

func DoAdmin(myUser *user) {
	(*myUser).role = "admin"
}
func main() {
	User := user{
		id:       1,
		name:     "ikrom",
		surname:  "Rahimov",
		login:    "ikromr04@gmail.com",
		password: 12345678,
		role:     "user",
		remove:   false,
	}
	fmt.Println(User)
	DoAdmin(&User)
	fmt.Println(User)
}
