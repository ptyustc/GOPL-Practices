package unitconversion

func (i Inch) IToM() Meter {
	return Meter(0.3048 * i)
}

func (m Meter) MToI() Inch {
	return Inch(m / 0.3048)
}

func (p Pound) PToKG() KG {
	return KG(0.4536 * p)
}

func (kg KG) KGToP() Pound {
	return Pound(kg / 0.4536)
}
