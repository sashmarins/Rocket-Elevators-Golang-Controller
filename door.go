package main

type Door struct {
	_ID int
}

func NewDoor(_ID int) *Door {
	return &Door{_ID}
}
