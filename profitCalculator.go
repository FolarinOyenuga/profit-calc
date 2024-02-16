package main

import (
	"fmt"
)

func main() {
	const taxRate float64 = 19
	var quantitySold, unitPrice, grossRevenue, corporateTax, netRevenue float64

	// get values for quantity and price
	fmt.Print("Enter the number of items sold: ")
	fmt.Scan(&quantitySold)
	fmt.Print("Enter the unit price/item: ")
	fmt.Scan(&unitPrice)

	// calculate gross revenue
	grossRevenue = quantitySold * unitPrice
	fmt.Printf("Your revenue before tax is %.2f £\n", grossRevenue)

	// calculate corporate tax
	corporateTax = grossRevenue * (taxRate / 100)
	fmt.Printf("Your corporate tax stands at %.2f £\n", corporateTax)

	// calculate revenue after tax
	netRevenue = grossRevenue - corporateTax
	fmt.Printf("Your actual revenue is %.2f £\n", netRevenue)

	// calculate ratio of gross revenue to net revenue
	var ratio = grossRevenue / netRevenue
	fmt.Printf("The ratio of gross revenue to net revenue is %.2f \n", ratio)
}
