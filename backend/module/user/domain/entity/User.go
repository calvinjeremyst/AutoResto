package entity

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	email       string
	password    []byte
	name        string
	phoneNumber string
	position    string
	id          int
	roles       UserRole
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) GetHashPassword() []byte {
	return u.password
}

func (u *User) SetHashPassword(rawPassword string) {
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	u.password = hashedPassword
}

func (u *User) SetRawPassword(password []byte) {
	u.password = password
}

func (u *User) ComparePassword(rawPassword string) bool {
	// Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword(u.password, []byte(rawPassword))
	if err != nil { // nil means it is a match
		return false
	}
	return true
}
