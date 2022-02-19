package entity

type Inventory struct {
	capacity int
	location string
}

type InventoryResponse struct {
	Message string      `form:"message" json:"message"`
	Data    []Inventory `form:"data" json:"data"`
}

func (i *Inventory) GetCapacity() int {
	return i.capacity
}

func (i *Inventory) SetCapacity(capacity int) {
	i.capacity = capacity
}

func (i *Inventory) GetLocation() string {
	return i.location
}

func (i *Inventory) SetLocation(location string) {
	i.location = location
}
