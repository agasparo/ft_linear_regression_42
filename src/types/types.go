package types

import (

)

type Datas struct {

	Kilometre 	map[int]float64
	Price		map[int]float64
}

type Learning struct {

	LearningRate 	float64
	MaxIterations	int
	Theta0			float64
	Theta1			float64
	LengthK			float64
	LengthP			float64
	Perte			float64
}

type Historique struct {

	Table map[int]HistoData
}

type HistoData struct {

	Theta0  float64
	Theta1  float64
	Perte	float64
}