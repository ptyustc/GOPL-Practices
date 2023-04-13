package unitconversion

import "fmt"

type Inch float64
type Meter float64
type Pound float64
type KG float64

func (i Inch) String() string {
	return fmt.Sprintf("%g inch", i)
}

func (m Meter) String() string {
	return fmt.Sprintf("%g m", m)
}

func (p Pound) String() string {
	return fmt.Sprintf("%g pound", p)
}

func (kg KG) String() string {
	return fmt.Sprintf("%g kg", kg)
}
