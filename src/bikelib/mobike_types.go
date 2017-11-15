package bikelib

import (
	"fmt"
)

// MobikeJSON struct is api json response
type MobikeJSON struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Biketype int         `json:"biketype"`
	Object   []MobikeCar `json:"object"`
}

// MobikeCar struct is used for mobike car
type MobikeCar struct {
	DistID   string      `json:"distId"`
	DistX    float64     `json:"distX"`
	DistY    float64     `json:"distY"`
	DistNum  int         `json:"distNum"`
	Distance string      `json:"distance"`
	BikeIds  string      `json:"bikeIds"`
	Biketype int         `json:"biketype"`
	Type     int         `json:"type"`
	Boundary interface{} `json:"boundary"`
}

func (el MobikeCar) String() string {
	return fmt.Sprintf("Mobike Num: %v Pos: ( %v, %v )", el.BikeIds, el.DistX, el.DistY)
}
