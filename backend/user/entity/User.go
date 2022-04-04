package entity

import (
	"encoding/json"
	"log"
	"net/http"

	handler "github.com/AutoResto/handler"
	controller "github.com/AutoResto/user/controller"
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

type LoginResponse struct {
	Message string `form:"message" json:"message"`
	Type    string `form:"userType" json:"userType"`
}

type LogoutResponse struct {
	Message string `form:"message" json:"message"`
}
func Login(c *gin.Context) {
	db := handler.Connect()
	defer db.Close()

	var userData User
	err := json.NewDecoder(c.Request.Body).Decode(&userData)

	if err != nil {
		log.Println(err)
	}

	query := "SELECT * FROM user WHERE email = '" + userData.Email +"'"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var user User
	
	for rows.Next() {
		if err := rows.Scan(&user.Id,&user.Email,&user.Password, &user.Name,&user.PhoneNumber,&user.Position,&user.Roles); err != nil {
			log.Fatal(err.Error())
		}
	}
	var response LoginResponse
	if user.Password == userData.Password && err == nil{
		controller.GenerateToken(c, user.Id, user.Name, user.Email)
		response.Message = "Login Success"
		response.Type = user.Position
		sendLoginSuccessResponse(c, response)
	} else {
		response.Message = "Login Error"
		sendLoginErrorResponse(c, response)
	}

}

func Logout(c *gin.Context) {

	controller.ResetUserToken(c)
	var response LogoutResponse
	response.Message = "Logout Success"
	sendLogoutSuccessResponse(c, response)
}

func sendLoginSuccessResponse(c *gin.Context, log LoginResponse) {
	c.JSON(http.StatusOK, log)
}

func sendLoginErrorResponse(c *gin.Context, err LoginResponse) {
	c.JSON(http.StatusBadRequest, err)

}

func sendLogoutSuccessResponse(c *gin.Context, err LogoutResponse){
	c.JSON(http.StatusOK,err)
}

func sendLogoutErrorResponse(c *gin.Context,err LogoutResponse){
	c.JSON(http.StatusBadRequest,err)
}