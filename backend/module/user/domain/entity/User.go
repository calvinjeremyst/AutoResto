package entity

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	email       string `form:"email" json:"email"`
	password    []byte `form:"password" json:"password"`
	name        string `form:"name" json:"name"`
	phoneNumber string `form:"phoneNumber" json:"phoneNumber"`
	position    string `form:"position" json:"position"`
	id          int    `form:"id" json:"id"`
	roles       UserRole
}

type UserResponse struct {
	Message string `form:"message" json:"message"`
	Data    []User `form:"data" json:"data"`
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

func (u *User) GetName() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) GetPhoneNumber() string {
	return u.phoneNumber
}

func (u *User) SetPhoneNumber(phoneNumber string) {
	u.phoneNumber = phoneNumber
}

func (u *User) GetPosition() string {
	return u.position
}

func (u *User) SetPosition(position string) {
	u.position = position
}

func (u *User) GetId() int {
	return u.id
}

func (u *User) SetId(id int) {
	u.id = id
}

func (u *User) GetRoles() UserRole {
	return u.roles
}

func (u *User) SetRoles(roles UserRole) {
	u.roles = roles
}
