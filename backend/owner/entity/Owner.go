package Owner

import (
	User "github.com/AutoResto/user/entity"
)

type Owner struct {
	User  User.User
	Owner User.UserRole
}
