package entity

type Material struct {
	id     int    `form:"id" json:"id"`
	name   string `form:"name" json:"name"`
	weight int    `form:"weight" json:"weight"`
}

type MaterialResponse struct {
	Message string     `form:"message" json:"message"`
	Data    []Material `form:"data" json:"data"`
}

func (m *Material) GetName() string {
	return m.name
}

func (m *Material) SetName(name string) {
	m.name = name
}

func (m *Material) GetWeight() int {
	return m.weight
}

func (m *Material) SetWeigh(weight int) {
	m.weight = weight
}

func (m *Material) GetId() int {
	return m.id
}

func (m *Material) SetId(id int) {
	m.id = id
}
