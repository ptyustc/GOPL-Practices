package tempconv

func (c Celsius) CToF() Fahrenheit {
	return Fahrenheit(c*9/5 - 32)
}

func (f Fahrenheit) FToC() Celsius {
	return Celsius((f + 32) * 5 / 9)
}

func (k Kelvin) KToC() Celsius {
	return Celsius(k + Kelvin(AbsoluteZero))
}

func (c Celsius) CToK() Kelvin {
	return Kelvin(c - AbsoluteZero)
}
