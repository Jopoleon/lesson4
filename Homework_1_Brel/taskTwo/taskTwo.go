package taskTwo

type Product struct {
	Name     string
	Category string
	Price    float64
}

func (p Product) convertToProductItem() ProductItem {
	return ProductItem{p.Name, p.Price}
}

type ProductItem struct {
	Name  string
	Price float64
}

func Execute(slice []Product) map[string][]ProductItem {
	result := make(map[string][]ProductItem)
	for _, product := range slice {
		result[product.Category] = append(result[product.Category], product.convertToProductItem())
	}
	return result
}
