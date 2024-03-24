package geogopher

import (
	"testing"
)

// mockLocation is a mock implementation of the Location interface for testing.
type mockLocation struct {
	coords Coordinates
}

func (m mockLocation) GetCoordinates() Coordinates {
	return m.coords
}

// TestCalculateDistanceOnEarth tests the CalculateDistanceOnEarth function with known distances.
func TestCalculateDistanceOnEarth(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		loc1     Location
		loc2     Location
		expected Distance
	}{
		{
			name:     "New York City to Los Angeles",
			loc1:     mockLocation{coords: Coordinates{Latitude: 40.7128, Longitude: -74.0060}},
			loc2:     mockLocation{coords: Coordinates{Latitude: 34.0522, Longitude: -118.2437}},
			expected: NewDistanceFromKilometers(3935.746255),
		},
		{
			name:     "Same Location",
			loc1:     mockLocation{coords: Coordinates{Latitude: 40.7128, Longitude: -74.0060}},
			loc2:     mockLocation{coords: Coordinates{Latitude: 40.7128, Longitude: -74.0060}},
			expected: NewDistanceFromKilometers(0),
		},
		// Add more test cases as needed
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			distance := CalculateDistanceOnEarth(tc.loc1, tc.loc2)
			if !AlmostEqual(distance, tc.expected) {
				t.Errorf("Distance was incorrect, got: %f, want: %f.", distance.Kilometers(), tc.expected)
			}
		})
	}
}
