package calcdistance

import "testing"

var testCases = []struct {
	x        Coordinate
	y        Coordinate
	distance float64
	unit     string
}{
	{
		x:        Coordinate{Latitude: 22.55, Longitude: 43.12},
		y:        Coordinate{Latitude: 13.45, Longitude: 100.28},
		distance: 6094.2511949585905,
		unit:     "K",
	},
}

func TestCalcDistance(t *testing.T) {
	for _, tc := range testCases {
		dis := Distance(tc.x, tc.y, tc.unit)
		if tc.distance != dis {
			t.Errorf("fail: want %v,%v  -> %v got %v ", tc.x, tc.y, tc.distance, dis)
		}
	}
}
