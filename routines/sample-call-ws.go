package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"runtime"
	"time"
)

const(
	WS_URL = "http://dev.markitondemand.com/MODApis/Api/v2/Quote/json?symbol="
)

func main()  {

	runtime.GOMAXPROCS(2)

	companies := []string{
		"NFLX",
		"googl",
		"aapl",
		"msft",
	}

	for _, company := range companies{
		go callWS(company)
	}

	time.Sleep(time.Second * 1)

}


func callWS(company string)  {
	resp, _ := http.Get(WS_URL+company)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	quote := new(QuoteResp)
	json.Unmarshal(body, &quote)

	fmt.Println(quote.Name, quote.LastPrice)
}


type QuoteResp struct {
	Status string
	Name string
	Symbol string
	LastPrice float32
	Change float32
	ChangePercent float32
	Timestamp string
	MSDate float32
	MarketCap int
	Volume int
	ChangeYTD float32
	ChangePercentYTD float32
	High float32
	Low float32
	Open float32
}
