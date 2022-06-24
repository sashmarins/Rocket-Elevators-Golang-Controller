package main

import "math"

type Column struct {
	ID                 int
	_amountOfElevators int
	servedFloors       []int
	_isBasement        bool
	callButtonList     []CallButton
	elevatorsList      []Elevator
}

func NewColumn(_id, _amountOfElevators int, servedFloors []int, _isBasement bool) *Column {
	return &Column{columnID, _amountOfElevators, servedFloors, _isBasement, *createCallButtons(len(servedFloors), _isBasement), *createElevators(len(servedFloors), _amountOfElevators)}
}

var callButtonID int = 1
var elevatorID int = 1

func createCallButtons(_amountOfFloors int, _isBasement bool) *[]CallButton {
	var callButtonList []CallButton
	if _isBasement {
		var buttonFloor int = -1
		for i := 0; i < _amountOfFloors; i++ {
			callButtonList = append(callButtonList, *NewCallButton(callButtonID, buttonFloor, ""))
			buttonFloor--
			callButtonID++
		}
	} else {
		var buttonFloor int = 1
		for i := 0; i < _amountOfFloors; i++ {
			callButtonList = append(callButtonList, *NewCallButton(callButtonID, buttonFloor, ""))
			buttonFloor++
			callButtonID++
		}
	}
	return &callButtonList
}

func createElevators(_amountOfFloors int, _amountOfElevators int) *[]Elevator {
	var elevatorsList []Elevator
	for i := 0; i < _amountOfElevators; i++ {
		elevatorsList = append(elevatorsList, *NewElevator(elevatorID, _amountOfFloors))
		elevatorID++
	}
	return &elevatorsList
}

//Simulate when a user press a button on a floor to go back to the first floor
func (c *Column) requestElevator(_requestedFloor int, _direction string) *Elevator {
	var elevator = *c.findElevator(_requestedFloor, _direction)
	elevator.addNewRequest(_requestedFloor)
	elevator.addNewRequest(1)

	return &elevator
}

func (c *Column) findElevator(_requestedFloor int, _direction string) *Elevator {
	var bestElevatorInformations = *NewElevatorInformation(c.elevatorsList[0], 6, 1000000)
	if _requestedFloor == 1 {
		for _, elevator := range c.elevatorsList {
			if elevator.currentFloor == 1 && elevator.status == "stopped" {
				bestElevatorInformations = *checkIfElevatorIsBetter(1, elevator, bestElevatorInformations, _requestedFloor)
			} else if elevator.currentFloor == 1 && elevator.status == "idle" {
				bestElevatorInformations = *checkIfElevatorIsBetter(2, elevator, bestElevatorInformations, _requestedFloor)
			} else if elevator.currentFloor < 1 && elevator.direction == "up" {
				bestElevatorInformations = *checkIfElevatorIsBetter(3, elevator, bestElevatorInformations, _requestedFloor)
			} else if elevator.currentFloor > 1 && elevator.direction == "down" {
				bestElevatorInformations = *checkIfElevatorIsBetter(3, elevator, bestElevatorInformations, _requestedFloor)
			} else if elevator.status == "idle" {
				bestElevatorInformations = *checkIfElevatorIsBetter(4, elevator, bestElevatorInformations, _requestedFloor)
			} else {
				bestElevatorInformations = *checkIfElevatorIsBetter(5, elevator, bestElevatorInformations, _requestedFloor)
			}
		}
	} else {
		for _, elevator := range c.elevatorsList {
			if _requestedFloor == elevator.currentFloor && elevator.status == "stopped" && _direction == elevator.direction {
				bestElevatorInformations = *checkIfElevatorIsBetter(1, elevator, bestElevatorInformations, _requestedFloor)
			} else if _requestedFloor > elevator.currentFloor && elevator.direction == "up" && _direction == "up" {
				bestElevatorInformations = *checkIfElevatorIsBetter(2, elevator, bestElevatorInformations, _requestedFloor)
			} else if _requestedFloor < elevator.currentFloor && elevator.direction == "down" && _direction == "down" {
				bestElevatorInformations = *checkIfElevatorIsBetter(2, elevator, bestElevatorInformations, _requestedFloor)
			} else if elevator.status == "idle" {
				bestElevatorInformations = *checkIfElevatorIsBetter(4, elevator, bestElevatorInformations, _requestedFloor)
			} else {
				bestElevatorInformations = *checkIfElevatorIsBetter(5, elevator, bestElevatorInformations, _requestedFloor)
			}
		}
	}
	return &bestElevatorInformations.bestElevator
}

func checkIfElevatorIsBetter(scoreToCheck int, newElevator Elevator, bestElevatorInformations ElevatorInformation, _requestedFloor int) *ElevatorInformation {
	if scoreToCheck < bestElevatorInformations.bestScore {
		bestElevatorInformations.bestScore = scoreToCheck
		bestElevatorInformations.bestElevator = newElevator
		bestElevatorInformations.referenceGap = int(math.Abs(float64(newElevator.currentFloor) - float64(_requestedFloor)))
	} else if bestElevatorInformations.bestScore == scoreToCheck {
		var gap int = int(math.Abs(float64(newElevator.currentFloor) - float64(_requestedFloor)))
		if bestElevatorInformations.referenceGap > gap {
			bestElevatorInformations.bestElevator = newElevator
			bestElevatorInformations.referenceGap = gap
		}
	}
	return &bestElevatorInformations
}
