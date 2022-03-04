package entity

type Food struct {
	id    int     `form:"id" json:"id"`
	name  string  `form:"name" json:"name"`
	price float64 `form:"price" json:"price"`
}

type FoodResponse struct {
	Message string `form:"message" json:"message"`
	Data    []Food `form:"data" json:"data"`
}

func (f *Food) GetName() string {
	return f.name
}

func (f *Food) SetName(name string) {
	f.name = name
}

func (f *Food) GetPrice() float64 {
	return f.price
}

func (f *Food) SetPrice(price float64) {
	f.price = price
}

func (f *Food) GetId() int {
	return f.id
}

func (f *Food) SetId(id int) {
	f.id = id
}
