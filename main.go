package main

import (
	"flag"
	"fmt"
	"go-api/goexamples/cryptopricechecker/client"
	"log"
)

func main() {
	fiatCurrency := flag.String("fiat", "USD", "The name of the fiat currency you would like to know the price of your crypto in")
	nameofCrypto := flag.String("crypto", "BTC", "Input the name of the CryptoCurrency you would like to know the price of")
	flag.Parse()

	crypto, err := client.FetchCrypto(*fiatCurrency, *nameofCrypto)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(crypto)
}
