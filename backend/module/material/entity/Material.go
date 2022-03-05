package entity

type Material struct {
	id     int    `form:"id" json:"id"`
	name   string `form:"name" json:"name"`
	quantity int    `form:"quantity" json:"quantity"`
	unit 	 string `form : "unit" json : "unit"`
}

type MaterialResponse struct {
	Message string     `form:"message" json:"message"`
	Data    []Material `form:"data" json:"data"`
}

func (m *Material) GetUnit() string{
	return m.unit
}

func (m *Material) SetUnit(unit string){
	m.unit = unit
}

func (m *Material) GetName() string {
	return m.name
}

func (m *Material) SetName(name string) {
	m.name = name
}

func (m *Material) GetQuantity() int {
	return m.quantity
}

func (m *Material) SetQuantity(quantity int) {
	m.quantity = quantity
}

func (m *Material) GetId() int {
	return m.id
}

func (m *Material) SetId(id int) {
	m.id = id
}
