package tempconv

// CToF converts celcius to fahrenheit
func CToF(c Celcius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC convert fahrenheit to celcius
func FToC(f Fahrenheit) Celcius { return Celcius((f - 32) * 5 / 9) }

// CToK converts celcius to kelvin
func CToK(c Celcius) Kelvin { return Kelvin(c + 273.15) }

// KToC convert kelvin to celcius
func KToC(k Kelvin) Celcius { return Celcius(k - 273.15) }

// FToK converts fahrenheit to kelvin
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

// KToF convert kelvin to fahrenheit
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }
