package service

import (
	factory "github.com/AutoResto/dp/factory"
	conn "github.com/AutoResto/dp/singleton"
)

func RegisterOwner() error {
	db := conn.Connect()
	defer db.Close()

	_, errQuery := db.Exec("INSERT INTO user(id,email,password,name,phoneNumber,idRole)VALUES(?,?,?,?,?,?)",
		factory.CreateOwner().Id,
		factory.CreateOwner().Email,
		factory.CreateOwner().Password,
		factory.CreateOwner().Name,
		factory.CreateOwner().PhoneNumber,
		factory.CreateOwner().Id,
	)
	return errQuery
}

func RegisterChef() error {
	db := conn.Connect()
	defer db.Close()

	_, errQuery := db.Exec("INSERT INTO user(id,email,password,name,phoneNumber,idRole)VALUES(?,?,?,?,?,?)",
		factory.CreateChef().Id,
		factory.CreateChef().Email,
		factory.CreateChef().Password,
		factory.CreateChef().Name,
		factory.CreateChef().PhoneNumber,
		factory.CreateChef().Id,
	)
	return errQuery
}

func RegisterChef2() error {
	db := conn.Connect()
	defer db.Close()

	_, errQuery := db.Exec("INSERT INTO user(id,email,password,name,phoneNumber,idRole)VALUES(?,?,?,?,?,?)",
		factory.CreateChef2().Id,
		factory.CreateChef2().Email,
		factory.CreateChef2().Password,
		factory.CreateChef2().Name,
		factory.CreateChef2().PhoneNumber,
		factory.CreateChef2().Id,
	)
	return errQuery
}

func RegisterInventoryStaff() error {
	db := conn.Connect()
	defer db.Close()

	_, errQuery := db.Exec("INSERT INTO user(id,email,password,name,phoneNumber,idRole)VALUES(?,?,?,?,?,?)",
		factory.CreateInventoryStaff().Id,
		factory.CreateInventoryStaff().Email,
		factory.CreateInventoryStaff().Password,
		factory.CreateInventoryStaff().Name,
		factory.CreateInventoryStaff().PhoneNumber,
		factory.CreateInventoryStaff().Id,
	)
	return errQuery

}

func RegisterInventoryStaff2() error {
	db := conn.Connect()
	defer db.Close()

	_, errQuery := db.Exec("INSERT INTO user(id,email,password,name,phoneNumber,idRole)VALUES(?,?,?,?,?,?)",
		factory.CreateInventoryStaff2().Id,
		factory.CreateInventoryStaff2().Email,
		factory.CreateInventoryStaff2().Password,
		factory.CreateInventoryStaff2().Name,
		factory.CreateInventoryStaff2().PhoneNumber,
		factory.CreateInventoryStaff2().Id,
	)
	return errQuery
}
