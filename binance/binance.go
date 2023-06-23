package binance

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-gota/gota/dataframe"
)

type BinanceTimestamp struct {
}

func Binance(ticker, interval string, number_of_days int, cols []string) dataframe.DataFrame {
	base_url := "https://data.binance.vision"
	params := fmt.Sprintf("/data/spot/daily/klines/%sUSDT/%s/%sUSDT-%s", ticker, interval, ticker, interval)

	complete_url := base_url + params
	start := time.Now().AddDate(0, 0, -1)
	end := start.AddDate(0, 0, -number_of_days)
	df := dataframe.DataFrame{}
	for d := start; !d.Before(end); d = d.AddDate(0, 0, -1) {
		year, month, day := d.Date()
		new_url := complete_url + fmt.Sprintf("-%02d-%02d-%02d.zip", year, month, day)
		new_df := GetCsvFromUrl(new_url)
		df = df.Concat(new_df)
	}
	return df
}

func GetCsvFromUrl(url string,) dataframe.DataFrame {
	println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	archive, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		log.Fatal(err)
	}
	for _, zipFile := range archive.File {
		fmt.Println("Reading file:", zipFile.Name)
		contents, err := zipFile.Open()
		if err != nil {
			log.Fatal(err)
		}
		zipFile.Open()

		names := dataframe.Names("Open time",
			"Open",
			"High",
			"Low",
			"Close",
			"Volume",
			"Close time",
			"Quote asset volume",
			"Number of trades",
			"Taker buy base asset volume",
			"Taker buy quote asset volume",
			"Ignore")
		noHeader := dataframe.HasHeader(false)
		df := dataframe.ReadCSV(contents,names,noHeader)
		println("read")
		return df
	}
	return dataframe.DataFrame{}

}
