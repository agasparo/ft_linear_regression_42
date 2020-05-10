package norm

import (
	"maths"
	"types"
)

func Normalize(Data types.Datas, kilometrage float64) (float64) {

	return ((kilometrage - maths.Min(Data.Kilometre)) / (maths.Max(Data.Kilometre) - maths.Min(Data.Kilometre)))
}

func Denormalize(Data types.Datas, prix float64) (float64) {

	return ((prix * (maths.Max(Data.Price) - maths.Min(Data.Price))) + maths.Min(Data.Price))
}

func NormalizeAllData(Data *types.Datas) {

	minK := maths.Min(Data.Kilometre)
	maxK := maths.Max(Data.Kilometre)
	
	for i := 0; i < len(Data.Kilometre); i++ {
		Data.Kilometre[i] = (Data.Kilometre[i] - minK) / (maxK - minK)
	}

	minP := maths.Min(Data.Price)
	maxP := maths.Max(Data.Price)
	
	for i := 0; i < len(Data.Price); i++ {
		Data.Price[i] = (Data.Price[i] - minP) / (maxP - minP)
	}
}