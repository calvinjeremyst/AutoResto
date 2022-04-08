package state

type state interface {
	addItem()
	requestItem()
	dispenseItem()
}
