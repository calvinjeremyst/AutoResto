package entity

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

type Claims struct{
	ID  int `json : "id"`
	Name string `json : "name"`
	Position string `json : "position"`
	jwt.StandardClaims

}

var tokenName = "token"
var jwtKey = []byte("")

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) GetPassword() string {
	return u.password
}


func (u *User) SetPassword(password string) {
	u.password = password
}


func (u *User) GetName() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) GetPhoneNumber() string {
	return u.phoneNumber
}

func (u *User) SetPhoneNumber(phoneNumber string) {
	u.phoneNumber = phoneNumber
}

func (u *User) GetPosition() string {
	return u.position
}

func (u *User) SetPosition(position string) {
	u.position = position
}

func (u *User) GetId() int {
	return u.id
}

func (u *User) SetId(id int) {
	u.id = id
}

func (u *User) GetRoles() UserRole {
	return u.roles
}

func (u *User) SetRoles(roles UserRole) {
	u.roles = roles
}

func connect() *sql.DB{
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/autoresto?parseTime=true&loc=Asia%2FJakarta")
	if err != nil {
		log.Fatal(err)
	}
	return db

}

func sendLoginSuccessResponse(c *gin.Context,log LoginResponse){
	c.JSON(http.StatusBadRequest,log)

}

func sendLoginErrorResponse(c *gin.Context, err LoginResponse){
	c.JSON(http.StatusBadRequest,err)

}

func generateToken(c *gin.Context, id int, name string, position string){
	tokenExpiryTime := time.Now().Add(60 * time.Minute)

	claims := &Claims{
		ID: id,
		Name : name,
		Position : position,
		StandardClaims : jwt.StandardClaims{
			ExpiresAt: tokenExpiryTime.Unix(),
		},

	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256,claims)
	signedToken,err := token.SignedString(jwtKey)

	if err != nil{
		return
	}

	c.SetCookie(tokenName,signedToken,1000,"/","localhost",false,true)
}

func Login(c *gin.Context){
	db := connect()
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

	if user.password == userData.phoneNumber{

		generateToken(c,user.id,user.name,user.position)
		response.Message = "Login Success"
		sendLoginSuccessResponse(c,response)

	}else{
		response.Message = "Login Error"
		sendLoginErrorResponse(c,response)
	}
	

	
}
