package main

import (
	"fmt"
)

const taxRate float64 = 19

func main() {
	var quantitySold, unitPrice, grr, corporateTax, netRevenue float64

	// get values for quantity and price
	textRequest("Enter the number of items sold: ")
	fmt.Scan(&quantitySold)

	textRequest("Enter the unit price/item: ")
	fmt.Scan(&unitPrice)

	// calculate gross revenue, corportate tax, net revenue and profit ratio
	// if a function returns multiple values, like this one, you can store it in those variables
	// grossRevenue = quantitySold * unitPrice
	// corporateTax = grossRevenue * (taxRate / 100)
	// netRevenue = grossRevenue - corporateTax
	// ratio = grossRevenue / netRevenue
	revNums(quantitySold, unitPrice)
	grr, corporateTax, netRevenue, ratio := revNums(quantitySold, unitPrice)

	fmt.Printf("Your revenue before tax is %.2f £\n", grr)
	fmt.Printf("Your corporate tax stands at %.2f £\n", corporateTax)
	fmt.Printf("Your actual revenue is %.2f £\n", netRevenue)
	fmt.Printf("The ratio of gross revenue to net revenue is %.2f \n", ratio)
}

// now to write a custom function that'll replace fmt used to print text
func revNums(quantitySold, unitPrice float64) (float64, float64, float64, float64) {
	grr := quantitySold * unitPrice
	corporateTax := grr * (taxRate / 100)
	netRevenue := grr - corporateTax
	ratio := grr / netRevenue
	return grr, corporateTax, netRevenue, ratio
}

func textRequest(text string) {
	fmt.Println(text)
}
