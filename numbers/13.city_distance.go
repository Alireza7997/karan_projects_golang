package numbers

import (
	"fmt"
	"math"

	"github.com/kelvins/geocoder"
)

const earthRadius = 6371

// Distance Between Two Cities :
// Calculates the distance between two cities and allows the user to specify a unit of distance.
// This program may require finding coordinates for the cities like latitude and longitude.
func DistanceBetweenCities() {
	var (
		city1    string
		country1 string
		city2    string
		country2 string
		distance float64
	)

	fmt.Print("Enter city1 and its country's name(e.g. Berlin Germany):\n")
	fmt.Scanln(&city1, &country1)

	fmt.Print("Enter city2 and its country's name(e.g. Paris France):\n")
	fmt.Scanln(&city2, &country2)

	city1Address := geocoder.Address{
		Country: country1,
		City:    city1,
	}

	location1, err := geocoder.Geocoding(city1Address)
	if err != nil {
		fmt.Println("Could not get the location: ", err)
	}

	city2Address := geocoder.Address{
		Country: country2,
		City:    city2,
	}

	location2, err := geocoder.Geocoding(city2Address)
	if err != nil {
		fmt.Println(err)
		return
	}

	distance = DistanceCalculator(location1.Latitude, location1.Longitude, location2.Latitude, location2.Longitude)

	fmt.Printf("Distance between %s and %s is: %.2f", city1, city2, distance)

}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func DistanceCalculator(lat1, lon1, lat2, lon2 float64) float64 {
	phi1 := degreesToRadians(lat1)
	phi2 := degreesToRadians(lat2)
	deltaPhi := degreesToRadians(lat2 - lat1)
	deltaLambda := degreesToRadians(lon2 - lon1)

	a := math.Sin(deltaPhi/2)*math.Sin(deltaPhi/2) +
		math.Cos(phi1)*math.Cos(phi2)*
			math.Sin(deltaLambda/2)*math.Sin(deltaLambda/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}
