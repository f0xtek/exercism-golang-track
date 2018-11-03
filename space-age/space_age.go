package space

// Planet name
type Planet string

const earthYearsInSeconds int = 31557600

var planets = []struct {
	name          string
	orbitalPeriod float64
}{
	{
		name:          "Earth",
		orbitalPeriod: 1.0,
	},
	{
		name:          "Mercury",
		orbitalPeriod: 0.2408467,
	},
	{
		name:          "Venus",
		orbitalPeriod: 0.61519726,
	},
	{
		name:          "Mars",
		orbitalPeriod: 1.8808158,
	},
	{
		name:          "Jupiter",
		orbitalPeriod: 11.862615,
	},
	{
		name:          "Saturn",
		orbitalPeriod: 29.447498,
	},
	{
		name:          "Uranus",
		orbitalPeriod: 84.016846,
	},
	{
		name:          "Neptune",
		orbitalPeriod: 164.79132,
	},
}

var ageOnPlanet float64

// Age function, given an age in seconds, returns your age in years on the given planet.
func Age(ageInSeconds float64, planet Planet) float64 {
	for _, pl := range planets {
		if pl.name == string(planet) {
			ageInYears := ageInSeconds / float64(earthYearsInSeconds)
			ageOnPlanet = ageInYears / pl.orbitalPeriod
		}
	}
	return ageOnPlanet

}
