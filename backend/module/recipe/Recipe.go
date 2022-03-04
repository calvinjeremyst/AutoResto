package recipe

type Recipe struct {
	id          int    `form:"id" json:"id"`
	description string `form:"description" json:"description"`
}

type RecipeResponse struct {
	Message string   `form:"message" json:"message"`
	Data    []Recipe `form:"data" json:"data"`
}

func (r *Recipe) GetId() int {
	return r.id
}

func (r *Recipe) SetId(id int) {
	r.id = id
}

func (r *Recipe) GetDescription() string {
	return r.description
}

func (r *Recipe) SetDescription(description string) {
	r.description = description
}
