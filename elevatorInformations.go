package main

type ElevatorInformation struct {
	bestElevator Elevator
	bestScore    int
	referenceGap int
}

func NewElevatorInformation(bestElevator Elevator, bestScore int, referenceGap int) *ElevatorInformation {
	return &ElevatorInformation{bestElevator, bestScore, referenceGap}
}
