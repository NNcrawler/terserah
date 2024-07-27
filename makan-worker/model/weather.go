package model

import "time"

type CurrentWeather struct {
	Time        time.Time
	Text        string
	Temperature float64
	FeelsLike   float64
}

type InterpretedCondition struct {
	Time        string
	Temperature string
	Condition   string
}

// hot, cold, rain, normal
func (c CurrentWeather) DetermineCondition() InterpretedCondition {
	var timeOfDay, temperature, condition string

	// Determine time of day
	hour := c.Time.Hour()
	switch {
	case hour >= 6 && hour < 12:
		timeOfDay = "morning"
	case hour >= 12 && hour < 18:
		timeOfDay = "afternoon"
	default:
		timeOfDay = "night"
	}

	// Determine temperature
	switch {
	case c.Temperature < 10:
		temperature = "cold"
	case c.Temperature >= 10 && c.Temperature < 20:
		temperature = "cool"
	case c.Temperature >= 20 && c.Temperature < 30:
		temperature = "normal"
	case c.Temperature >= 30 && c.Temperature < 40:
		temperature = "hot"
	default:
		temperature = "super_hot"
	}

	// Determine condition
	switch c.Text {
	case "Clear", "Sunny":
		condition = "clear"
	case "Cloudy", "Partly Cloudy", "Overcast":
		condition = "cloudy"
	case "Rain", "Drizzle", "Showers":
		condition = "rain"
	default:
		condition = "unknown"
	}

	return InterpretedCondition{
		Time:        timeOfDay,
		Temperature: temperature,
		Condition:   condition,
	}
}
