package calcdistance

import "math"

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

func Distance(coordinate1, coordinate2 Coordinate, unit ...string) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * coordinate1.Latitude / 180)
	radlat2 := float64(PI * coordinate2.Latitude / 180)

	theta := float64(coordinate1.Longitude - coordinate2.Longitude)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}

	return dist
}
