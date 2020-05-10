package main

import(
	"fmt"
	"file"
	"types"
	"Response"
	"input"
	"norm"
	"courbe"
)

func main() {

	Data := types.HistoData{}
	Csv := types.Datas{}
	file.ReadFile(&Csv)
	res := file.ReadResp(&Data)

	if res == 0 && (Data.Theta0 != 0 || Data.Theta1 != 0) {
		kilometrage, _ := input.ReadSTDIN("Choose a km", 0)
		tmp_kilometrage := norm.Normalize(Csv, kilometrage)

		res := Data.Theta0 + (Data.Theta1 * tmp_kilometrage)
		Response.Sucess(fmt.Sprintf("For %f km : %f $\n", kilometrage, norm.Denormalize(Csv, res)))

		_, Graph := input.ReadSTDIN("See graph : [y] yes / [n] no", 1)
		if Graph == "y" {
			courbe.Init(Csv, Data)
		}
		return
	}
	Response.Sucess("0\n")
}