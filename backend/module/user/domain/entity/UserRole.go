package entity

import "fmt"

type UserRole int

const (
	Chef UserRole = iota
	InventoryStaff
	Owner
)

var userEnum = map[UserRole]string{
	Chef:           "Chef",
	InventoryStaff: "InventoryStaff",
	Owner:          "Owner",
}

func (r UserRole) String() string {
	value, isExist := userEnum[r]
	if isExist {
		return value
	} else {
		fmt.Println("Not Set Enum !! ")
		return ""
	}
}

func (u *User) SetUserRoleString(str string) {
	for key, val := range userEnum {
		if val == str {
			u.roles = key
			return
		}
	}
	u.roles = -1
}
