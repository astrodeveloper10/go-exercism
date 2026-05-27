// Package weather provides tools to forecast weather.
package weather

var (
    // CurrentCondition represents a weather condition.
	CurrentCondition string
    // CurrentLocation represents a location.
	CurrentLocation  string
)

// Forecast returns a string value that is a represenation
// of current condition for a given location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
