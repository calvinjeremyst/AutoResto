package entity

import (
	"encoding/json"
	"log"
	"net/http"

	controller "github.com/AutoResto/controller"
	"github.com/gin-gonic/gin"
)

type User struct {
	Email       string   `form:"email" json:"email"`
	Password    string   `form:"password" json:"password"`
	Name        string   `form:"name" json:"name"`
	PhoneNumber string   `form:"phonenumber" json:"phonenumber"`
	Position    string   `form:"position" json:"position"`
	Id          int      `form:"id" json:"id"`
	Roles       UserRole `form:"role" json:"role"`
}

type UserResponse struct {
	Message string `form:"message" json:"message"`
	Data    []User `form:"data" json:"data"`
}

type LoginResponse struct {
	Message string `form:"message" json:"message"`
	Type    string `form:"userType" json:"userType"`
}

func Login(c *gin.Context) {
	db := controller.Connect()
	defer db.Close()

	var userData User
	err := json.NewDecoder(c.Request.Body).Decode(&userData)

	if err != nil {
		log.Println(err)
	}

	query := "SELECT id,name,position FROM user WHERE email = '" + userData.Email + "' AND password = '" + string(userData.Password) + "'"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var user User
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name, &user.Position); err != nil {
			log.Fatal(err.Error())
		}
	}

	var response LoginResponse
	if user.Password == userData.Password {
		controller.GenerateToken(c, user.Id, user.Name, user.Email)
		response.Message = "Login Success"
		sendLoginSuccessResponse(c, response)
	} else {
		response.Message = "Login Error"
		sendLoginErrorResponse(c, response)
	}

}

func sendLoginSuccessResponse(c *gin.Context, log LoginResponse) {
	c.JSON(http.StatusOK, log)
}

func sendLoginErrorResponse(c *gin.Context, err LoginResponse) {
	c.JSON(http.StatusBadRequest, err)

}
