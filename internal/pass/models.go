package pass

type Observer struct {
	Latitude  float64
	Longitude float64
	Elevation float64
}

type Request struct {
	SatelliteID  string
	Observer     Observer
	Start        string
	End          string
	MinElevation float64
}

type Pass struct {
	AOS          string
	LOS          string
	DurationSec  int
	MaxElevation float64
	Visibility   string
}
