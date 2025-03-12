package main

import "fmt"

//vehicle struct common
type Motor struct{
	Brand string
	Speed int
}

// Method for Vehicle
func (m Motor) Move(){
	fmt.Println(m.Brand, "is moving at a speed of", m.Speed, "mph")
}

//car struct embeds motor
type Car struct{
	Motor
	Transmission string
}
//scooter struct embedds motor
type Scooter struct{
	Motor
	Gears int
}
type Plane struct{
	Motor 
	Carrying_capacity int
}

//override  Move () for plane 
func (p Plane) Move(){
	fmt.Println(p.Brand, "is Flying at", p.Speed, "mph with a a carrying_capacity of", p.Carrying_capacity,"passengers!")
}

func main(){
	car      := Car{Motor: Motor{Brand: "Bugatti cheyron", Speed: 360}, Transmission:"Automatic"}
	scooter := Scooter{Motor: Motor{Brand:"Honda", Speed: 180}, Gears: 6}
	plane := Plane{Motor: Motor{Brand:"Emirates", Speed:540}, Carrying_capacity: 300}

	car.Move()
	scooter.Move()
	plane.Move()


}