package tempconv

import "fmt"

// Celcius type
type Celcius float64

// Fahrenheit type
type Fahrenheit float64

const (
	// AbsoluteZeroC represents absolute zero in celcius
	AbsoluteZeroC Celcius = -273.15

	// FreezingC represents water's freezing temp in celcius
	FreezingC Celcius = 0.0

	// BoilingC represents water's boiling temp in celcius
	BoilingC Celcius = 100.0
)

// CToF converts celcius to fahrenheit
func CToF(c Celcius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC convert fahrenheit to celcius
func FToC(f Fahrenheit) Celcius { return Celcius((f - 32) * 5 / 9) }

func (c Celcius) String() string { return fmt.Sprintf("%g°C", c) }

func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
