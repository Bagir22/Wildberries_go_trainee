package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/


/*
	Фасад представляет интерфейс к сложной системе классов. Определяет одну точку взаимодейсвтия между клиентом и системой

	+
	1. Фасад предоставляет упрощенный интерфейс для взаимодействия с сложной подсистемой.

	-
	1. Фасад может стать обьектом, привязанным ко всем классам программы.
*/
type FlightBooking struct{}

func (fb *FlightBooking) BookFlight(destination string) {
	fmt.Printf("Booking flight to %s\n", destination)
}

type HotelBooking struct{}

func (hb *HotelBooking) BookHotel(checkIn, checkOut string) {
	fmt.Printf("Booking hotel from %s to %s\n", checkIn, checkOut)
}

type CarRental struct{}

func (cr *CarRental) RentCar(startDate, endDate string) {
	fmt.Printf("Rent car from %s to %s\n", startDate, endDate)
}

type Travel struct {
	flight *FlightBooking
	hotel *HotelBooking
	car *CarRental
}

func (tp *Travel) Plan(destination string, checkIn, checkOut, startDate, endDate string) {
	tp.flight.BookFlight(destination)
	tp.hotel.BookHotel(checkIn, checkOut)
	tp.car.RentCar(startDate, endDate)
}

func NewTrip() *Travel {
	return &Travel {
		flight:  &FlightBooking{},
		hotel: &HotelBooking{},
		car: &CarRental{},
	}
}

func main() {
	t := NewTrip()

	destination := "Kazan"
	checkIn := "2023-16-08"
	checkOut := "2023-25-08"
	startDate := "2023-17-08"
	endDate := "2023-24-08"

	t.Plan(destination, checkIn, checkOut, startDate, endDate)
}