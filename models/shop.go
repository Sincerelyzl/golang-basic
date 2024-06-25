package models

type Product struct {
	Name  string
	Price float64
}

type Cart struct {
	Product []Product
}

func (p *Product) ChangeName(name string) {
	p.Name = name
}

func (p *Product) ChangePrice(price float64) {

	if price > 0 {
		p.Price = price
	}

}

func (c *Cart) AddProduct(product Product) {
	c.Product = append(c.Product, product)
}
