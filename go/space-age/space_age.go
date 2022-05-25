package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	var secondsInDay float64 = 60 * 60 * 24
	givenTime := seconds / secondsInDay
	earthDays := 365.25

	switch planet {
	case "Mercury":
		return givenTime / (earthDays * 0.2408467)
	case "Venus":
		return givenTime / (earthDays * 0.61519726)
	case "Earth":
		return givenTime / earthDays
	case "Mars":
		return givenTime / (earthDays * 1.8808158)
	case "Jupiter":
		return givenTime / (earthDays * 11.862615)
	case "Saturn":
		return givenTime / (earthDays * 29.447498)
	case "Uranus":
		return givenTime / (earthDays * 84.016846)
	case "Neptune":
		return givenTime / (earthDays * 164.79132)
	default:
		return 0
	}
}
