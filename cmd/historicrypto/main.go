package main

import (
	"github.com/willb0/historicrypto/binance"
)

func main() {
	
	binance.Binance("ETH","1h",20,[]string{"Close","Close Time"})
	println("binanced")

	



}
