package data

import "testing"

func TestCheckValidation(t *testing.T){
	p:= &Product{
		Name: "ProductName",
		Price: 1,
		SKU: "abcd-def",
	}

	err:= p.Validate()

	if err!= nil{
		t.Fatal(err)
	}
}