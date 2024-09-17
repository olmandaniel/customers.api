package utils

import (
	"regexp"
	"time"
)

func ValidateDNI(dni string) bool {
	// Usar una expresión regular para verificar que el DNI tenga exactamente 8 dígitos
	re := regexp.MustCompile(`^\d{8}$`)
	return re.MatchString(dni)
}

func ValidatePhone(dni string) bool {
	// Usar una expresión regular para verificar que el DNI tenga exactamente 8 dígitos
	re := regexp.MustCompile(`^\d{9}$`)
	return re.MatchString(dni)
}

func DurationToYears(d time.Duration) float64 {
	const (
		secondsPerMinute     = 60
		minutesPerHour       = 60
		hoursPerDay          = 24
		daysPerYear          = 365.25
		nanosecondsPerSecond = 1e9
	)

	nanosecondsPerYear := nanosecondsPerSecond * secondsPerMinute * minutesPerHour * hoursPerDay * daysPerYear

	years := float64(d) / float64(nanosecondsPerYear)
	return years
}

func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
