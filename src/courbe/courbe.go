package courbe

import (
	"os"
	"github.com/wcharczuk/go-chart"
	"runtime"
	"types"
	"os/exec"
	"fmt"
	"maths"
	"norm"
)

func Init(Data types.Datas, Histo types.HistoData) {

	All := AddPoint(Data, []chart.Series{}, Histo)
	All = AddCourbe(Data, All, Histo)
	Draw(All)
}

func AddPoint(Data types.Datas, All []chart.Series, Histo types.HistoData) ([]chart.Series) {

	var tabx, taby, tmp []float64

	s := maths.GetSize(maths.Max(Data.Kilometre))
	Add := float64(1 * maths.TransSize(s - 4))

	for i := 0; i < len(Data.Kilometre); i++ {

		tabx = append(tabx, Data.Kilometre[i], Data.Kilometre[i] + Add)
		taby = append(taby, Data.Price[i], Data.Price[i] + Add)
		
		All = append(All, chart.ContinuousSeries {
	        XValues: tabx,
	        YValues: taby,
	    })

	    tabx = tmp
	    taby = tmp
	}
	return (All)
}

func AddCourbe(Data types.Datas, All []chart.Series, Histo types.HistoData) ([]chart.Series) {

	var tabx, taby []float64

	for i := 0; i < len(Data.Kilometre); i++ {

		x := Histo.Theta0 + (Histo.Theta1 * norm.Normalize(Data, Data.Kilometre[i]))
		y := norm.Denormalize(Data, x)

		tabx = append(tabx, Data.Kilometre[i])
		taby = append(taby, y)
	}
	All = append(All, chart.ContinuousSeries {
		Name:    "f(x)",
	    XValues: tabx,
	    YValues: taby,
	})
	return (All)	
}

func Draw(All []chart.Series) {

	var cmd *exec.Cmd

	graph := chart.Chart{
		Series: All,
	}

	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
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