package main

type ProductDetails struct {
	price int64
	brand string
}

type shirt struct {
	ProductDetails
	size  int64
	color string
}

func (s shirt) CalculatePrice() int64 {
	clothDiscount := float64(s.price) * .20
	return s.price - int64(clothDiscount)
}

type monitor struct {
	ProductDetails
	size       string
	resolution string
}

func (m monitor) CalculatePrice() int64 {
	electornicDiscount := float64(m.price) * .20
	return m.price - int64(electornicDiscount)
}

type purchasableInterface interface {
	CalculatePrice() int64
}

var cart []purchasableInterface

func addToCart(products purchasableInterface) {
	cart = append(cart, products)
}

func getCartTotal(products ...purchasableInterface) int64 {
	var total int64
	for _, product := range products {
		total += product.CalculatePrice()
	}

	return total
}
