package aritmatika

import "math"

type Kubus struct {
	Sisi float64
}

func (K Kubus) Luas() float64 {
	return 6.0 * math.Pow(K.Sisi, 2)
}

func (K Kubus) Keliling() float64 {
	return 12 * K.Sisi
}
