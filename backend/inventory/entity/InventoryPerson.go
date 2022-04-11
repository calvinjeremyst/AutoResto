package entity

import (
	User "github.com/AutoResto/user/entity"
)

type InventoryPerson struct {
	User            User.User
	inventoryperson User.UserRole
}
