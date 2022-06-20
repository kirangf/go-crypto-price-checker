package model

import "fmt"

type CryptoResponse []struct {
	Name              string `json:"name"`
	Price             string `json:"price"`
	Rank              string `json:"rank"`
	High              string `json:"high"`
	CirculatingSupply string `json:"circulating_supply"`
}

//Output is exported,it formats the data to plain text.
func (c CryptoResponse) Output() string {
	p := fmt.Sprintf(
		"Name: %v\n Price: %v\n Rank: %v \n High: %v\n Circulating Supply: %v",
		c[0].Name, c[0].Price, c[0].Rank, c[0].High, c[0].CirculatingSupply)

	return p
}
