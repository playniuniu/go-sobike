package bikelib

// BikeData is base bike struct
type BikeData struct {
	Lng      float64
	Lat      float64
	CarNo    string
	CarType  string
	Distance int
}

// BikeInterface is common bike interface
type BikeInterface interface {
	GetNearbyCar() ([]BikeData, error)
}
