package main

//Button on a floor or basement to go back to lobby
type CallButton struct {
	_ID        int
	_floor     int
	_direction string
}

func NewCallButton(_ID int, _floor int, _direction string) *CallButton {
	return &CallButton{_ID, _floor, _direction}
}
