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