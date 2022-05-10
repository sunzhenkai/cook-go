package main

import (
	"cooking-go/pkg/st"
	"fmt"
)

func main() {
	ra := st.Cart{Price: 4, Count: 5, Promo: 1}
	println("main, &ra: ", &ra)
	fmt.Printf("%v\n", ra.TotalPrice())
	fmt.Printf("%v\n", ra.TotalPricePtr())
	fmt.Printf("%v\n", (&ra).TotalPricePtr())

	println("-----")
	fmt.Printf("%v\n", ra.Promotion(0.9))
	fmt.Printf("%v\n", ra.TotalPrice())
	fmt.Printf("%v\n", ra.TotalPricePtr())

	println("-----")
	rb := st.Cart{Price: 4, Count: 5, Promo: 1}
	fmt.Printf("%v\n", rb.PromotionPtr(0.9))
	fmt.Printf("%v\n", rb.TotalPrice())
	fmt.Printf("%v\n", rb.TotalPricePtr())
}
