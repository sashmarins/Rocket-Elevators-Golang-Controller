package main

type Battery struct {
	ID                         int
	_amountOfColumns           int
	_amountOfFloors            int
	_amountOfBasements         int
	_amountOfElevatorPerColumn int
	status                     string
	floorRequestButtonsList    []FloorRequestButton
	columnsList                []Column
}

func NewBattery(_id, _amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn int) *Battery {
	return &Battery{_id, _amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn, "running", *createFloorRequestButtons(_amountOfFloors, _amountOfBasements), *createColumns(_amountOfColumns, _amountOfFloors, _amountOfElevatorPerColumn, _amountOfBasements)}
}

var columnID int = 1
var floorRequestButtonID int = 1

func createBasementColumn(_amountOfBasements int, _amountOfElevatorPerColumn int) *Column {
	var servedFloors []int
	var floor int = -1
	for i := 0; i < _amountOfBasements; i++ {
		servedFloors = append(servedFloors, floor)
		floor--
	}
	columnID++
	basementColumn := *NewColumn(columnID, _amountOfBasements, servedFloors, true)
	return &basementColumn
}

func createColumns(_amountOfColumns int, _amountOfFloors int, _amountOfElevatorPerColumn int, _amountOfBasements int) *[]Column {
	var columnsList []Column
	// var amountOfFloorsPerColumn int = int(math.Ceil(float64(_amountOfFloors) / float64(_amountOfColumns)))
	var amountOfFloorsPerColumn int = 20
	var floor int = 1
	if _amountOfBasements > 0 {
		columnsList = append(columnsList, *createBasementColumn(_amountOfBasements, _amountOfElevatorPerColumn))
		_amountOfColumns--
	}
	for i := 0; i < _amountOfColumns; i++ {
		var servedFloors []int
		for h := 0; h < amountOfFloorsPerColumn; h++ {
			if floor <= _amountOfFloors {
				servedFloors = append(servedFloors, floor)
				floor++
			}
		}
		columnID++
		columnsList = append(columnsList, *NewColumn(columnID, _amountOfElevatorPerColumn, servedFloors, false))
	}
	// if _amountOfBasements > 0 {
	// 	basementColumn := *createBasementColumn(_amountOfBasements)
	// 	return &basementColumn
	// }
	return &columnsList
}

func createFloorRequestButtons(_amountOfFloors int, _amountofBasements int) *[]FloorRequestButton {
	var floorRequestButtonsList []FloorRequestButton
	// if _amountofBasements > 0 {
	// 	*createBasementFloorRequestButtons(_amountofBasements)
	// 	return &floorRequestButtonsList
	// }
	var buttonFloor int = 1
	for i := 0; i < _amountOfFloors; i++ {
		floorRequestButtonsList = append(floorRequestButtonsList, *NewFloorRequestButton(floorRequestButtonID, buttonFloor, ""))
		buttonFloor++
		floorRequestButtonID++
	}
	if _amountofBasements > 0 {
		BasementRequestButtons := *createBasementFloorRequestButtons(_amountofBasements)
		return &BasementRequestButtons
	}
	return &floorRequestButtonsList
}

func createBasementFloorRequestButtons(_amountOfBasements int) *[]FloorRequestButton {
	var floorRequestButtonsList []FloorRequestButton
	var buttonFloor int = -1
	for i := 0; i < _amountOfBasements; i++ {
		floorRequestButtonsList = append(floorRequestButtonsList, *NewFloorRequestButton(floorRequestButtonID, buttonFloor, ""))
		buttonFloor--
		floorRequestButtonID++
	}
	return &floorRequestButtonsList
}

func (b *Battery) findBestColumn(_requestedFloor int) *Column {
	for _, j := range b.columnsList {
		for _, x := range j.servedFloors {
			if x == _requestedFloor {
				return &j
			}
		}
	}
	return nil
}

//Simulate when a user press a button at the lobby
func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {
	var column = *b.findBestColumn(_requestedFloor)
	var elevator = *column.findElevator(1, _direction)
	elevator.addNewRequest(1)
	elevator.addNewRequest(_requestedFloor)

	return &column, &elevator
}
