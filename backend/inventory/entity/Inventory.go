package entity

import (
	model "github.com/AutoResto/location/entity"
)

type Inventory struct {
	Id       int            `form:"id" json:"id"`
	Capacity int            `form:"capacity" json:"capacity"`
	Location model.Location `form:"location" json:"location"`
}
