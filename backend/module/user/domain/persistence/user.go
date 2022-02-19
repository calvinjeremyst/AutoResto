package persistence

import "github.com/AutoResto/backend/module/user/domain/entity"

type UserPersistenceMgr interface {
	Login(email string, password string) bool
	GetUserByAttribute(attribute string, value interface{}) []entity.User
}
