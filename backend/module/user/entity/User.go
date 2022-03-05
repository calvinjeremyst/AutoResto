package entity

type User struct {
	email       string   `form:"email" json:"email"`
	password    string   `form:"password" json:"password"`
	name        string   `form:"name" json:"name"`
	phoneNumber string   `form:"phonenumber" json:"phonenumber"`
	position    string   `form:"position" json:"position"`
	id          int      `form:"id" json:"id"`
	roles       UserRole `form:"role" json:"role"`
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

func (u *User) GetPassword() string {
	return u.password
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
