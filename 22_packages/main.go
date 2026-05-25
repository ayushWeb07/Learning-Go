package main

import (
	"github.com/ayushWeb07/Learning-Go/auth"
	"github.com/ayushWeb07/Learning-Go/session"
)

func main() {
	auth.AuthWithGoogle()
	auth.AuthWithCredentials()

	session.GetUserSession()
}
