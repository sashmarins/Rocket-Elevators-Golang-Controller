package main

//FloorRequestButton is a button on the pannel at the lobby to request any floor
type FloorRequestButton struct {
	_ID        int
	_floor     int
	_direction string
}

func NewFloorRequestButton(_ID int, _floor int, _direction string) *FloorRequestButton {
	return &FloorRequestButton{_ID, _floor, _direction}
}
