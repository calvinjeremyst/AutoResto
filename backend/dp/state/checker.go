package state
type checker struct{

	hasItem state
	requestedItem state
	noItem state
	currentState state

	Id int `form : "id" json:"id"`
	Name string `form : "name" json : "name"`
	Quantity int `form : "quantity" json : "quantity"`
}


func newCondition(Id,Quantity int,Name string) *checker{

	ch := &checker{
		Id: Id,
		Name: Name,
		Quantity: Quantity,
	}

	hasItemState := hasItemState{
		checker: ch,
	}
	ch.hasItem = hasItemState.checker.hasItem
	return ch
}