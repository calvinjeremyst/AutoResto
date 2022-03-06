package entity

import (
	
	"encoding/json"
	"log"
	"net/http"
	controller "github.com/AutoResto/controller"
	"github.com/gin-gonic/gin"	
)

type User struct {
	email       string   `form:"email" json:"email"`
	password    string   `form:"password" json:"password"`
	name        string   `form:"name" json:"name"`
	phoneNumber string   `form:"phonenumber" json:"phonenumber"`
	position    string   `form:"position" json:"position"`
	id          int      `form:"id" json:"id"`
	roles       UserRole `form:"role" json:"role"`
}

type UserResponse struct {
	Message string `form:"message" json:"message"`
	Data    []User `form:"data" json:"data"`
}

type LoginResponse struct{
	Message string `form : "message" json : "message"`
	Roles string `form : "roles" json : "roles"`
}

func sendLoginSuccessResponse(c *gin.Context,log LoginResponse){
	c.JSON(http.StatusOK,log)

}

func sendLoginErrorResponse(c *gin.Context, err LoginResponse){
	c.JSON(http.StatusBadRequest,err)

}
func Login(c *gin.Context){
	db := controller.Connect()
	defer db.Close()

	var userData User
	err := json.NewDecoder(c.Request.Body).Decode(&userData)

	if err != nil{
		log.Println(err)
	}

	query := "SELECT id,name,position FROM user WHERE email = '"+userData.email+"' AND password = '"+string(userData.password)+"'"

	rows,err := db.Query(query)
	if err != nil{
		log.Println(err)
	}

	var user User
	for rows.Next(){
		if err := rows.Scan(&user.id,&user.name,&user.position); err!= nil {
			log.Fatal(err.Error())
		}
		
	}
	
	var response LoginResponse

	if user.password == userData.password{
		controller.GenerateToken(c,user.id,user.name,user.email)
		response.Message = "Login Success"
		sendLoginSuccessResponse(c,response)

	}else{
		response.Message = "Login Error"
		sendLoginErrorResponse(c,response)
	}
	
}
