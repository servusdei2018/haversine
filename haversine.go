// Package haversine provides functions for calculating distances between two
// geographic points using the Haversine formula.
//
// The Haversine formula calculates the shortest distance between two points on
// the surface of a sphere, given their latitudes and longitudes.
//
// This distance is commonly used to calculate distances between two locations on Earth.
//
// Example:
//
//	package main
//
//	import (
//		"fmt"
//
//	   	"github.com/servusdei2018/haversine"
//	)
//
//	func main() {
//		// Coordinates of New York City
//		lat1 := 40.7128
//		lon1 := -74.0060
//
//		// Coordinates of Los Angeles
//		lat2 := 34.0549
//		lon2 := -118.2426
//
//		distance, err := haversine.Haversine(lat1, lon1, lat2, lon2)
//		if err != nil {
//			fmt.Println("Error:", err)
//			return
//		}
//
//		fmt.Printf("Distance between New York City and Los Angeles: %.2f km\n", distance)
//	}
package haversine

import (
	"errors"
	"math"
)

// earthRadius represents the Earth's radius in kilometers.
const earthRadius = 6371

// degToRad converts degrees to radians.
func degToRad(deg float64) (rad float64) {
	return deg * math.Pi / 180
}

// isValidLatitude checks if the given latitude is within valid range [-90, 90].
func isValidLatitude(lat float64) bool {
	return lat >= -90 && lat <= 90
}

// isValidLongitude checks if the given longitude is within valid range [-180, 180].
func isValidLongitude(lon float64) bool {
	return lon >= -180 && lon <= 180
}

// Haversine calculates the distance in kilometers between two geographic
// coordinates using the Haversine formula.
//
// Latitude and longitude values are expected to be in degrees.
func Haversine(lat1, lon1, lat2, lon2 float64) (distance float64, err error) {
	if !isValidLatitude(lat1) || !isValidLatitude(lat2) || !isValidLongitude(lon1) || !isValidLongitude(lon2) {
		return -1, errors.New("haversine: invalid latitude or longitude values")
	}

	dLat := degToRad(lat2 - lat1)
	dLon := degToRad(lon2 - lon1)

	a := math.Pow(math.Sin(dLat/2), 2) + math.Cos(degToRad(lat1))*math.Cos(degToRad(lat2))*math.Pow(math.Sin(dLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c, nil
}
