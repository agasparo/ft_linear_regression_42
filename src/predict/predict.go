package main

import(
	"fmt"
	"file"
	"types"
	"Response"
	"input"
	"maths"
)

func main() {

	Data := types.HistoData{}
	Csv := types.Datas{}
	file.ReadFile(&Csv)
	file.ReadResp(&Data)

	kilometrage := input.ReadSTDIN()
	tmp_kilometrage := Normalize(Csv, kilometrage)

	res := Data.Theta0 + (Data.Theta1 * tmp_kilometrage)
	Response.Sucess(fmt.Sprintf("For %f km : %f $\n", kilometrage, Denormalize(Csv, res)))
}

func Normalize(Data types.Datas, kilometrage float64) (float64) {

	return ((kilometrage - maths.Min(Data.Kilometre)) / (maths.Max(Data.Kilometre) - maths.Min(Data.Kilometre)))
}

func Denormalize(Data types.Datas, prix float64) (float64) {

	return ((prix * (maths.Max(Data.Price) - maths.Min(Data.Price))) + maths.Min(Data.Price))
}