package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/willb0/historicrypto/binance"
)

func main() {
	token := flag.String("token", "BTC", "A ticker of a token in all caps")
	interval := flag.String("interval", "4h", "A string interval for the data you want. supported 1s, 1m, 3m, 5m, 15m, 30m, 1h, 2h, 4h, 6h, 8h, 12h, 1d")
	number_of_days := flag.Int("days", 10, "The number of days back in time you want to go")
	features := flag.String("features", "Close,Close Time", "disabled, comma separated string with the data columns you want, available are  Open time, Open, High, Low, Close, Volume, Close time, Quote asset volume, Number of trades, Taker buy base asset volume, Taker buy quote asset volume, Ignore")

	flag.Parse()

	fmt.Printf("features: %v\n", features)
	res := binance.Binance(*token, *interval, *number_of_days, []string{"Close", "Close Time"})
	println("binanced")
	now := time.Now()
	f, err := os.Create(fmt.Sprintf("%s-%s-%dd-%d_%d.csv", *token, *interval, *number_of_days, now.Month(), now.Day()))
	if err != nil {
		log.Fatal(err)
	}
	res.WriteCSV(f)
}
