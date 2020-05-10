package maths

func Max(data map[int]float64) (float64) {

	max := data[0]

	for i := 0; i < len(data); i++ {
		
		if data[i] > max {
			max = data[i]
		}
	}
	return (max)
}

func Min(data map[int]float64) (float64) {

	min := data[0]

	for i := 0; i < len(data); i++ {

		if data[i] < min {
			min = data[i]
		}
	}
	return (min)
}

func GetSize(nb float64) (int) {

	s := 0

	for i := nb; i > 1.1; i /= 10 {
		s++
	}
	return (s)
}

func TransSize(nb int) (int, int) {

	s := 1
	a := 0

	if nb < 0 {
		nb = nb * -1
		a = 1
	}

	for i := 0; i < nb; i++ {
		s *= 10
	}
	return s, a
}