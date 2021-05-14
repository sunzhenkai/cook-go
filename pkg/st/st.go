package st

type Cart struct {
	Price float64
	Count int
	Promo float64
}

func (r Cart) TotalPrice() float64 {
	return r.Price * float64(r.Count) * r.Promo
}

func (r *Cart) TotalPricePtr() float64 {
	return r.Price * float64(r.Count) * r.Promo
}

func (r Cart) Promotion(p float64) float64 {
	println("Promotion, &r: ", &r)
	r.Promo *= p
	return r.TotalPrice()
}

func (r *Cart) PromotionPtr(p float64) float64 {
	println("PromotionPtr, &r: ", r)
	r.Promo *= p
	return r.TotalPricePtr()
}
