package file

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"Response"
	"types"
	"io/ioutil"
)

const File = "data/data.csv"
const SaveFile = "data/res.csv"

func ReadFile(Dat *types.Datas) {

	csvfile, err := os.Open(File)
	if err != nil {
		Response.Print(fmt.Sprintf("%s\n", err))
		return
	}
	r := csv.NewReader(csvfile)

	Dat.Kilometre = make(map[int]float64)
	Dat.Price = make(map[int]float64)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			Response.Print(fmt.Sprintf("%s\n", err))
			return
		}
		if record[0] != "km" {
			Dat.Kilometre[len(Dat.Kilometre)], _ = strconv.ParseFloat(record[0], 64)
			Dat.Price[len(Dat.Price)], _ = strconv.ParseFloat(record[1], 64)
		}
	}
}

func check(e error) {
    
    if e != nil {
        Response.Print(fmt.Sprintf("%s\n", e))
    } else {
    	Response.Sucess("File created")
    }
}

func Save(Data map[int]types.HistoData) {

	str := "theta0,theta1,error\n" + fmt.Sprintf("%f,%f,%f", Data[0].Theta0, Data[0].Theta1, Data[0].Perte)
	fd := []byte(str)
    err := ioutil.WriteFile(SaveFile, fd, 0644)
    check(err)
}