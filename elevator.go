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

// var completedRequestsList []int
// var floorRequestsList []int

func NewElevator(ID int, _amountOfFloors int) *Elevator {
	return &Elevator{ID, 1, "idle", _amountOfFloors, "", []int{}, []int{}}
}

//not sure what to return
// (e *Elevator)??
func (e *Elevator) move() {
	for len(e.floorRequestsList) != 0 {
		e.sortFloorList()
		var destination int = e.floorRequestsList[0]
		e.status = "moving"
		// e.sortFloorList()
		if e.currentFloor < destination {
			e.direction = "up"
			// e.sortFloorList()
			for e.currentFloor < destination {
				e.currentFloor++
			}

		} else if e.currentFloor > destination {
			e.direction = "down"
			// e.sortFloorList()
			for e.currentFloor > destination {
				e.currentFloor--
			}
		} else {
			// e.currentFloor = destination
			e.direction = ""
			e.completedRequestsList = append(e.completedRequestsList, destination)
			e.floorRequestsList = append(e.floorRequestsList[:0], e.floorRequestsList[1:]...)
			e.status = "stopped"
		}
	}
	e.status = "idle"
	e.direction = ""

}

func (e *Elevator) sortFloorList() *[]int {
	if e.direction == "up" {
		// sort.Slice(floorRequestsList, func(i, j int) bool {
		// 	return floorRequestsList[i] < floorRequestsList[j]
		sort.Ints(e.floorRequestsList)
	}
	if e.direction == "down" {
		// sort.Slice(floorRequestsList, func(i, j int) bool {
		// 	// return floorRequestsList[i] > floorRequestsList[j]
		sort.Sort(sort.Reverse(sort.IntSlice(e.floorRequestsList)))
	}
	return &e.floorRequestsList
}

func (e *Elevator) addNewRequest(_requestedFloor int) *[]int {
	// for _, j := range e.floorRequestsList {
	// 	if j != _requestedFloor {
	e.floorRequestsList = append(e.floorRequestsList, _requestedFloor)
	// 	}
	// }
	e.move()
	return &e.floorRequestsList
}
