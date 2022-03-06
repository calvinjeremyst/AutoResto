package entity

import (
	entityRecipe "github.com/AutoResto/module/recipe/entity"
)

type Menu struct {
	Id     int                 `form:"id" json:"id"`
	Name   string              `form:"name" json:"name"`
	Price  float64             `form:"price" json:"price"`
	Recipe entityRecipe.Recipe `form:"recipe" json:"recipe"`
}

type MenuResponse struct {
	Message string `form:"message" json:"message"`
	Data    []Menu `form:"data" json:"data"`
}

// func (f *Menu) GetName() string {
// 	return f.name
// }

// func (f *Menu) SetName(name string) {
// 	f.name = name
// }

// func (f *Menu) GetPrice() float64 {
// 	return f.price
// }

// func (f *Menu) SetPrice(price float64) {
// 	f.price = price
// }

// func (f *Menu) GetId() int {
// 	return f.id
// }

// func (f *Menu) SetId(id int) {
// 	f.id = id
// }
