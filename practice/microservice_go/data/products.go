package data

import (
	"fmt"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,customSku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

var ErrProductNotFound = fmt.Errorf("Product not found.")

type GenericError struct {
	Message string `json:"message"`
}

func NewGenericError(msg string) *GenericError {
	return &GenericError{msg}
}

type Products []*Product

func NewProduct() *Product {
	return &Product{}
}

// func (p *Products) ToJSON(w io.Writer) error {
// 	enc := json.NewEncoder(w)
// 	return enc.Encode(p)
// }

// func (p *Product) FromJSON(r io.Reader) error {
// 	e := json.NewDecoder(r)
// 	return e.Decode(p)
// }

// func (p *Product) Validate() error {
// 	validate := validator.New()
// 	validate.RegisterValidation("customSku", validateSKU)
// 	return validate.Struct(p)
// }

func GetProducts() Products {
	return productList
}

func GetProductByID(id int) (*Product, error) {
	index := getIndex(id)
	if index == -1 {
		return nil, ErrProductNotFound
	}

	return productList[index], nil
}

func AddProduct(p Product) {
	p.ID = getNextId()
	productList = append(productList, &p)
}

func UpdateProduct(id int, p *Product) error {
	pos := getIndex(id)
	if pos == -1 {
		return ErrProductNotFound
	}

	p.ID = id
	productList[pos] = p
	return nil
}

func DeleteProduct(id int) error {
	pos := getIndex(id)
	if pos == -1 {
		return ErrProductNotFound
	}
	productList = append(productList[:pos], productList[(pos+1):]...)
	return nil
}

func getIndex(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

func getNextId() int {
	return productList[len(productList)-1].ID + 1
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Milky Coffee",
		Price:       2.45,
		SKU:         "latte",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Strong Coffee without Milk",
		Price:       1.99,
		SKU:         "latte",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
