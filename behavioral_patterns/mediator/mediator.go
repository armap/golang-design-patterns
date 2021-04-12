package main

import "fmt"

/*
COLLEAGUE AIRPLANE INTERFACE
 */
type ColleagueAirplane interface {
	landing()
	takingOff()
}
/*
CONCRETE COLLEAGUE AIRPLANE A
*/
type colleagueAirplaneA struct {
	mediator Mediator
}
func (caa *colleagueAirplaneA) landing() {
	if !caa.mediator.canLand(caa) {
		fmt.Println("Airplane A: Landing blocked, waiting")
		return
	}
	fmt.Println("Airplane A: Landed")
}
func (caa *colleagueAirplaneA) takingOff() {
	fmt.Println("Airplane A: Taking off")
	caa.mediator.notifyTakingOff()
}
/*
CONCRETE COLLEAGUE AIRPLANE B
*/
type colleagueAirplaneB struct {
	mediator Mediator
}
func (cab *colleagueAirplaneB) landing() {
	if !cab.mediator.canLand(cab) {
		fmt.Println("Airplane B: Landing blocked, waiting")
		return
	}
	fmt.Println("Airplane B: Landed")
}
func (cab *colleagueAirplaneB) takingOff() {
	fmt.Println("Airplane B: Taking off")
	cab.mediator.notifyTakingOff()
}

/*
MEDIATOR AIRPORT
 */
type Mediator interface {
	canLand(ColleagueAirplane) bool
	notifyTakingOff()
}
type mediatorAirport struct {
	airplanesQueue []ColleagueAirplane
	isRunwayFree   bool
}
func NewMediatorAirport() *mediatorAirport {
	return &mediatorAirport{
		isRunwayFree: true,
	}
}
func (ma *mediatorAirport) canLand(colleagueAirplane ColleagueAirplane) bool {
	if ma.isRunwayFree {
		ma.isRunwayFree = false
		return true
	}
	ma.airplanesQueue = append(ma.airplanesQueue, colleagueAirplane)
	return false
}
func (ma *mediatorAirport) notifyTakingOff() {
	if !ma.isRunwayFree {
		ma.isRunwayFree = true
	}
	if len(ma.airplanesQueue) > 0 {
		firstAirplaneInQueue := ma.airplanesQueue[0]
		ma.airplanesQueue = ma.airplanesQueue[1:]
		//Permit landing
		firstAirplaneInQueue.landing()
	}
}


/*
CLIENT
 */
func main() {
	mediatorAirport := NewMediatorAirport()

	colleagueAirplaneA := &colleagueAirplaneA{
		mediator: mediatorAirport,
	}
	colleagueAirplaneB := &colleagueAirplaneB{
		mediator: mediatorAirport,
	}

	colleagueAirplaneA.landing()
	colleagueAirplaneB.landing()
	colleagueAirplaneA.takingOff()
}
