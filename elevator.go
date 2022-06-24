package main

import "sort"

type Elevator struct {
	ID                    int
	currentFloor          int
	status                string
	_amountOfFloors       int
	direction             string
	completedRequestsList []int
	floorRequestsList     []int
	//door

}

var completedRequestsList []int
var floorRequestsList []int

func NewElevator(_ID int, _amountOfFloors int) *Elevator {
	return &Elevator{_ID, 1, "idle", _amountOfFloors, "", completedRequestsList, floorRequestsList}
}

//not sure what to return
// (e *Elevator)??
func (e *Elevator) move() {
	for len(floorRequestsList) > 0 {
		var destination int = floorRequestsList[0]
		e.status = "moving"
		if e.currentFloor < destination {
			e.direction = "up"
			e.sortFloorList()
			for e.currentFloor < destination {
				e.currentFloor++
			}

		} else if e.currentFloor > destination {
			e.direction = "down"
			e.sortFloorList()
			for e.currentFloor > destination {
				e.currentFloor--
			}
		}
		e.direction = ""
		e.status = "stopped"
		completedRequestsList = append(completedRequestsList, floorRequestsList[0])
		floorRequestsList = floorRequestsList[1:]
	}

}

func (e *Elevator) sortFloorList() *[]int {
	if e.direction == "up" {
		sort.Slice(floorRequestsList, func(i, j int) bool {
			return floorRequestsList[i] < floorRequestsList[j]
		})
	} else {
		sort.Slice(floorRequestsList, func(i, j int) bool {
			return floorRequestsList[i] > floorRequestsList[j]
		})
	}
	return &floorRequestsList
}

func (e *Elevator) addNewRequest(_requestedFloor int) *[]int {
	for _, j := range e.floorRequestsList {
		if j != _requestedFloor {
			e.floorRequestsList = append(e.floorRequestsList, _requestedFloor)
		}
	}
	return &e.floorRequestsList
}
