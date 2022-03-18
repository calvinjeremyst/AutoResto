package Chef

import (
	User "github.com/AutoResto/user/entity"
)

type Owner struct{
	User User.User
	Chef User.UserRole
}
