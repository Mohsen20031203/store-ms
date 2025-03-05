package main

import (
	"fmt"
)

func CalculateGoldPrice(ouncePrice float64, usdToToman float64) float64 {
	gramsPerOunce := 31.1035
	pricePerGram24 := (ouncePrice * usdToToman) / gramsPerOunce

	pricePerGram18 := pricePerGram24 * 0.75
	return pricePerGram18
}

func main() {
	ouncePrice := 2918.96
	usdToToman := 90673.0

	goldPrice := CalculateGoldPrice(ouncePrice, usdToToman)
	fmt.Printf("قیمت هر گرم طلای 18 عیار: %.0f تومان\n", goldPrice)
}
