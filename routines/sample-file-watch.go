package main

import(
	"os"
	"fmt"
	"io/ioutil"
	"encoding/csv"
	"strings"
	"strconv"
	"runtime"
)

const (
	FOLDER_PATH = "/Users/andrepinto/Documents/labs/go-samples/routines/files"
)

func main()  {

	runtime.GOMAXPROCS(2)

	for  {
		dir, _ := os.Open(FOLDER_PATH)
		files, _ := dir.Readdir(-1)

		for _, file := range files{
			path := FOLDER_PATH + "/" + file.Name()

			f, _ := os.Open(path)

			content, _ := ioutil.ReadAll(f)

			go readContent(string(content), path)
			f.Close()
			os.Remove(path)


		}
	}
}

func readContent(data string, path string)  {
	rd := csv.NewReader(strings.NewReader(data))
	rcds, _ := rd.ReadAll()

	for _, record := range rcds{
		prd := new(Product)
		prd.Name=record[0]
		prd.Price, _=strconv.ParseFloat(record[1], 64)
		prd.Brand=record[2]

		fmt.Println(prd)
	}

}

type Product struct {
	Name string
	Price float64
	Brand string
}
