package courbe

import (
	"os"
	"github.com/wcharczuk/go-chart"
	"runtime"
	"types"
	"os/exec"
	"fmt"
)

func Init(Data types.Datas, Histo map[int]types.HistoData) {

	All := AddPoint(Data, []chart.Series{})
	Draw(All)
}

func AddPoint(Data types.Datas, All []chart.Series) ([]chart.Series) {

	var tabx, taby, tmp []float64

	for i := 0; i < len(Data.Kilometre); i++ {
		
		tabx = append(tabx, Data.Kilometre[i], Data.Kilometre[i] + 0.005)
		taby = append(taby, Data.Price[i], Data.Price[i] + 0.005)
		
		All = append(All, chart.ContinuousSeries {
	        XValues: tabx,
	        YValues: taby,
	    })

	    tabx = tmp
	    taby = tmp
	}
	return (All)
}

/*func Normalize(Data types.Datas, kilometrage float64) (float64) {

	return ((kilometrage - maths.Min(Data.Kilometre)) / (maths.Max(Data.Kilometre) - maths.Min(Data.Kilometre)))
}

func Denormalize(Data types.Datas, prix float64) (float64) {

	return ((prix * (maths.Max(Data.Price) - maths.Min(Data.Price))) + maths.Min(Data.Price))
}*/

func Draw(All []chart.Series) {

	fmt.Println(All)

	var cmd *exec.Cmd

	graph := chart.Chart{
		Series: All,
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)

    if runtime.GOOS == "linux" {
        cmd = exec.Command("feh", "output.png")
    } else {
        cmd = exec.Command("sh", "catimg.sh", "output.png")
    }
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println(string(stdout))
}