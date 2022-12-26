package main

import (
	"fmt"
	"math"
)

const (
	radius = 6371 // Earth's radius in kilometers
)

type Coordinate struct {
	lat, lon float64
}

// CalculateDistance returns the distance in kilometers between two coordinates.
func CalculateDistance(a, b Coordinate) float64 {
	lat1, lon1 := a.lat, a.lon
	lat2, lon2 := b.lat, b.lon

	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180

	lat1 = lat1 * math.Pi / 180
	lat2 = lat2 * math.Pi / 180

	c := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)
	d := 2 * math.Atan2(math.Sqrt(c), math.Sqrt(1-c))

	return radius * d
}

func main() {
	coordA := Coordinate{lat: 52.2296756, lon: 21.0122287}
	coordB := Coordinate{lat: 41.89193, lon: 12.51133}

	distance := CalculateDistance(coordA, coordB)
	fmt.Printf("Distance between A and B is %.2f kilometers\n", distance)
}
