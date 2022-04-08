package factory

import (
	usr "github.com/AutoResto/user/entity"
)

func CreateOwner() *usr.User {
	return &usr.User{
		Email:       "aji@gmail.com",
		Password:    "aji",
		Name:        "Aji Parama",
		PhoneNumber: "08154653",
		Position:    "1",
		Id:          1,
		Roles:       usr.Owner,
	}
}

func CreateChef() *usr.User {
	return &usr.User{
		Email:       "calvin@gmail.com",
		Password:    "calvin",
		Name:        "Calvin Jeremy",
		PhoneNumber: "081541534",
		Position:    "2",
		Id:          2,
		Roles:       usr.Chef,
	}
}

func CreateChef2() *usr.User {
	return &usr.User{
		Email:       "christian@gmail.com",
		Password:    "christian",
		Name:        "Christian Satya",
		PhoneNumber: "081573682750",
		Position:    "2",
		Id:          2,
		Roles:       usr.Chef,
	}
}

func CreateInventoryStaff() *usr.User {
	return &usr.User{
		Email:       "juniar@gmail.com",
		Password:    "juniar",
		Name:        "William Juniar",
		PhoneNumber: "0812535435",
		Position:    "3",
		Id:          3,
		Roles:       usr.InventoryStaff,
	}
}

func CreateInventoryStaff2() *usr.User {
	return &usr.User{
		Email:       "kolis@gmail.com",
		Password:    "kolis",
		Name:        "William Kolis",
		PhoneNumber: "081534354",
		Position:    "3",
		Id:          3,
		Roles:       usr.InventoryStaff,
	}
}
