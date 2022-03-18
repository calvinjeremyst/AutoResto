package entity

type MaterialResponse struct {
	Message string     `form:"message" json:"message"`
	Data    []Material `form:"data" json:"data"`
}