package gostudy

import "fmt"

// START OMIT
type OrderItem struct {
	Product  string
	Quantity int
}

func Example() {
	cart := []OrderItem{{"coke", 2}, {"pepsi", 3}, {"water", 12}}
	var cartRefs []*OrderItem

	for _, item := range cart {
		cartRefs = append(cartRefs, &item)
	}

	for _, ref := range cartRefs {
		fmt.Println(ref.Product)
	}
}

// END OMIT
