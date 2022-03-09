package entity

type Location struct {
	Id       int    `form:"id" json:"id"`
	Loc_name string `form:"loc_name" json:"loc_name"`
}

type LocationResponse struct {
	Message string     `form:"message" json:"message"`
	Data    []Location `form:"data" json:"data"`
}
