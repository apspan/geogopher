package geogopher

import "math"

type Location interface {
	GetCoordinates() Coordinates
}

// Coordinates represents geographical coordinates with latitude and longitude.
type Coordinates struct {
	Latitude  float64
	Longitude float64
}

// NewCoordinates creates a new instance of Coordinates with the given latitude and longitude.
func NewCoordinates(lat, long float64) *Coordinates {
	return &Coordinates{
		Latitude:  lat,
		Longitude: long,
	}
}

// CalculateDistanceOnEarth calculates the great-circle distance between two locations on Earth
// and returns the distance as a Distance struct.
func CalculateDistanceOnEarth(loc1, loc2 Location) Distance {
	lat1 := loc1.GetCoordinates().Latitude * math.Pi / 180
	lat2 := loc2.GetCoordinates().Latitude * math.Pi / 180
	deltaLat := (loc2.GetCoordinates().Latitude - loc1.GetCoordinates().Latitude) * math.Pi / 180
	deltaLong := (loc2.GetCoordinates().Longitude - loc1.GetCoordinates().Longitude) * math.Pi / 180

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(deltaLong/2)*math.Sin(deltaLong/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return NewDistanceFromKilometers(EarthsRadiusInKm * c)
}
