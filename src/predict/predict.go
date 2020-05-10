package main

import(
	"fmt"
	"file"
	"types"
	"Response"
	"input"
	"maths"
	"courbe"
)

func main() {

	Data := types.HistoData{}
	Csv := types.Datas{}
	file.ReadFile(&Csv)
	file.ReadResp(&Data)

	kilometrage, _ := input.ReadSTDIN("Choose a km", 0)
	tmp_kilometrage := Normalize(Csv, kilometrage)

	res := Data.Theta0 + (Data.Theta1 * tmp_kilometrage)
	Response.Sucess(fmt.Sprintf("For %f km : %f $\n", kilometrage, Denormalize(Csv, res)))

	_, Graph := input.ReadSTDIN("See graph : [y] yes / [n] no", 1)
	if Graph == "y" {
		courbe.Init(Csv, Data)
	}
}

func Normalize(Data types.Datas, kilometrage float64) (float64) {

	return ((kilometrage - maths.Min(Data.Kilometre)) / (maths.Max(Data.Kilometre) - maths.Min(Data.Kilometre)))
}

func Denormalize(Data types.Datas, prix float64) (float64) {

	return ((prix * (maths.Max(Data.Price) - maths.Min(Data.Price))) + maths.Min(Data.Price))
}