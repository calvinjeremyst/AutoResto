package service

import (
	factory "github.com/AutoResto/user/factory"
	handler "github.com/AutoResto/handler"
	
)

func RegisterOwner() error{
	db := handler.Connect()
	defer db.Close()
	
	_,errQuery := db.Exec("INSERT INTO user(id,email,password,name,phoneNumber,idRole)VALUES(?,?,?,?,?,?)",
		factory.CreateOwner().Id,
		factory.CreateOwner().Email,
		factory.CreateOwner().Password,
		factory.CreateOwner().Name,
		factory.CreateOwner().PhoneNumber,
		factory.CreateOwner().Id,
	)
	return errQuery
}


func RegisterChef()error{
	db := handler.Connect()
	defer db.Close()

	_,errQuery := db.Exec("INSERT INTO user(id,email,password,name,phoneNumber,idRole)VALUES(?,?,?,?,?,?)",
		factory.CreateChef().Id,
		factory.CreateChef().Email,
		factory.CreateChef().Password,
		factory.CreateChef().Name,
		factory.CreateChef().PhoneNumber,
		factory.CreateChef().Id,
	)
	return errQuery
}

func RegisterChef2() error{
	db := handler.Connect()
	defer db.Close()

	_,errQuery := db.Exec("INSERT INTO user(id,email,password,name,phoneNumber,idRole)VALUES(?,?,?,?,?,?)",
		factory.CreateChef2().Id,
		factory.CreateChef2().Email,
		factory.CreateChef2().Password,
		factory.CreateChef2().Name,
		factory.CreateChef2().PhoneNumber,
		factory.CreateChef2().Id,
	)
	return errQuery
}

func RegisterInventoryStaff()error{
	db := handler.Connect()
	defer db.Close()

	_,errQuery := db.Exec("INSERT INTO user(id,email,password,name,phoneNumber,idRole)VALUES(?,?,?,?,?,?)",
	factory.CreateInventoryStaff().Id,
	factory.CreateInventoryStaff().Email,
	factory.CreateInventoryStaff().Password,
	factory.CreateInventoryStaff().Name,
	factory.CreateInventoryStaff().PhoneNumber,
	factory.CreateInventoryStaff().Id,
)
return errQuery
	
}

