package entity

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

func (p Product) StatusProduct() string {
	var Status string
	if p.Stock < 3 {
		Status = "Stock Sedikit"
	} else if p.Stock < 10 {
		Status = "Stock agak banyakan"
	}
	return Status
}
