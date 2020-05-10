package main

import (
	"file"
	"math"
	"types"
	"norm"
	"Response"
)

func main() {

	Data := types.Datas{}
	res := file.ReadFile(&Data)
	if res == 2 {
		Response.Print("You have a problem with the save file")
		return
	}
	if res == 1 {
		return
	}
	norm.NormalizeAllData(&Data)
	L := types.Learning{ 0.3, 100000, 0, 0, float64(len(Data.Kilometre)), float64(len(Data.Price)), 0 }
	Histo := types.Historique{}
	Train(&L, Data, &Histo)
	file.Save(Histo.Table)
}

func Train(L *types.Learning, Data types.Datas, Histo *types.Historique) {

	tmpTheta0 := L.Theta0
	tmpTheta1 := L.Theta1
	Histo.Table = make(map[int]types.HistoData)

	for i := 0; i < L.MaxIterations; i++ {

		S1, S2 := Somme(Data, L)
		tmpTheta0 -= L.LearningRate * (S1 / L.LengthK) 
		tmpTheta1 -= L.LearningRate * (S2 / L.LengthP)
		L.Perte = MoindreCarre(L, Data)
		Compare(L, Histo, tmpTheta0, tmpTheta1)
		tmpTheta0 = L.Theta0
		tmpTheta1 = L.Theta1
	}
}

func Compare(L *types.Learning, Histo *types.Historique, tmpTheta0 float64, tmpTheta1 float64) {

	if len(Histo.Table) == 1 {

		if Histo.Table[0].Perte > L.Perte {
			nData := types.HistoData{ tmpTheta0, tmpTheta1, L.Perte }
			Histo.Table[0] = nData
			L.LearningRate *= 1.05
		} else {
			L.LearningRate *= 0.5
		}
	} else {
		nData := types.HistoData{ tmpTheta0, tmpTheta1, L.Perte }
		Histo.Table[0] = nData
	}
	L.Theta0 = Histo.Table[0].Theta0
	L.Theta1 = Histo.Table[0].Theta1
	L.Perte = Histo.Table[0].Perte
}

func MoindreCarre(L *types.Learning, Data types.Datas) (float64) {

	var err float64

	for i := 0; i < len(Data.Kilometre); i++ {
		err += math.Pow((Data.Price[i] - (L.Theta1 * Data.Kilometre[i] + L.Theta0)), 2)
	}
	return (err / L.LengthK)
}

func Somme(Data types.Datas, L *types.Learning) (float64, float64) {

	var S1, S2 float64 

	for i := 0; i < len(Data.Kilometre); i++ {

		S1 += (L.Theta0 + L.Theta1 * Data.Kilometre[i]) - Data.Price[i]
		S2 += (L.Theta0 + L.Theta1 * Data.Kilometre[i] - Data.Price[i]) * Data.Kilometre[i]
	}
	return S1, S2
}