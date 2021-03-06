package creationaldesignpattern

import "testing"

func TestDrink(t *testing.T) {
	beer := &MUSU{}
	drink := InitDrink(beer)
	drink.Drink()
	s := beer.GetResult()
	if s != "Drink MUSU" {
		t.Fatal("Error Drink")
	}
}

func TestDrink2(t *testing.T) {
	beer := &QingDao{}
	drink := InitDrink(beer)
	drink.Drink()
	s := beer.GetResult()
	if s != "Drink QingDao" {
		t.Fatal("Error Drink")
	}
}
