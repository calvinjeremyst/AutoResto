package service

import (
	"fmt"

	"github.com/AutoResto/backend/module/user/domain/persistence"
)

var pm persistence.UserPersistenceMgr

func SetPersistenceMgr(o persistence.UserPersistenceMgr) {
	pm = o
}

func Login(email string, password string) bool {
	registeredUser := pm.GetUserByAttribute("email", email)
	if len(registeredUser) != 1 {
		return false
	}
	defer registeredUser[0].SetRawPassword(nil)
	if registeredUser[0].ComparePassword(password) && registeredUser[0].GetId() > 0 {
		fmt.Println("Nama User Logged :" + registeredUser[0].GetName())
		return true
	}
	return false
}
