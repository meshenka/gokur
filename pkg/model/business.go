package model

// Business a client business
type Business struct {
	ID           string
	Name         string
	Localisation LatLong
	Address      string
}

// LatLong a golocation point
type LatLong struct {
	Lat  float64
	Long float64
}
