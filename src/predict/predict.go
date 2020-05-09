package main

import(
	"fmt"
	"file"
	"types"
	"Response"
	"input"
)

func main() {

	Data := types.HistoData{}
	file.ReadResp(&Data)
	kilometrage := input.ReadSTDIN()
	res := Data.Theta0 + (Data.Theta1 * kilometrage)
	Response.Sucess(fmt.Sprintf("For %f km : %f $\n", kilometrage, res))
}