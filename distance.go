package geogopher

import "math"

type Distance struct {
	meters float64
}

func NewDistanceFromMillimeters(mm float64) Distance {
	return Distance{meters: mm / 1000}
}

func NewDistanceFromCentimeters(cm float64) Distance {
	return Distance{meters: cm / 100}
}

func NewDistanceFromMeters(m float64) Distance {
	return Distance{meters: m}
}

func NewDistanceFromKilometers(km float64) Distance {
	return Distance{meters: km * 1000}
}

func NewDistanceFromYards(yards float64) Distance {
	return Distance{meters: yards / 1.0936}
}

func NewDistanceFromFeet(feet float64) Distance {
	return Distance{meters: feet / 3.2808}
}

func NewDistanceFromMiles(miles float64) Distance {
	return Distance{meters: miles * 1609.34}
}

func NewDistanceFromNauticalMiles(nmi float64) Distance {
	return Distance{meters: nmi * 1852}
}

func AlmostEqual(a Distance, b Distance) bool {
	return math.Abs(a.Meters()-b.Meters()) <= 0.01
}

func (d Distance) Millimeters() float64 {
	return d.meters * MillimeterToMeterRatio
}

func (d Distance) Centimeters() float64 {
	return d.meters * CentimeterToMeterRatio
}

func (d Distance) Meters() float64 {
	return d.meters
}

func (d Distance) Kilometers() float64 {
	return d.meters / MeterToKilometerRatio
}

func (d Distance) Yards() float64 {
	return d.meters * MeterToYardRatio
}

func (d Distance) Miles() float64 {
	return d.meters / MeterToMileRatio
}

func (d Distance) Feet() float64 {
	return d.meters * MeterToFootRatio
}

func (d Distance) NauticalMiles() float64 {
	return d.meters / MeterToNauticalMileRatio
}
