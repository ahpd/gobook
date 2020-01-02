// Package tempconv performs Celcius and Fahrenheit computations
package tempconv

import "fmt"

// Celcius type
type Celcius float64

// Fahrenheit type
type Fahrenheit float64

// Kelvin type
type Kelvin float64

const (
	// AbsoluteZeroC represents absolute zero in celcius
	AbsoluteZeroC Celcius = -273.15

	// FreezingC represents water's freezing temp in celcius
	FreezingC Celcius = 0.0

	// BoilingC represents water's boiling temp in celcius
	BoilingC Celcius = 100.0

	// AbsoluteZeroK represents absolute zero in kelvin
	AbsoluteZeroK = 0.0

	// FreezingK represents water's freezing temp in kelvin
	FreezingK = 273.15

	// BoilingK represents water's boiling temp in kelvin
	BoilingK = 373.15
)

func (c Celcius) String() string { return fmt.Sprintf("%g°C", c) }

func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func (k Kelvin) String() string { return fmt.Sprintf("%g°K", k) }
