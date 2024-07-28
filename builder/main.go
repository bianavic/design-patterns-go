package main

import "fmt"

func main() {
	myAddress := CreateAddress().
		SetStreet("123 Main St").
		SetNumber(100).
		SetCity("San Francisco").
		SetCountry("USA")

	fmt.Println("My address is: ", myAddress)
}

type Address struct {
	street  string
	number  int32
	city    string
	country string
}

func CreateAddress() *Address {
	return &Address{}
}

func (a *Address) SetStreet(streetName string) *Address {
	a.street = streetName
	return a
}

func (a *Address) SetNumber(number int32) *Address {
	a.number = number
	return a
}

func (a *Address) SetCity(cityName string) *Address {
	a.city = cityName
	return a
}

func (a *Address) SetCountry(countryName string) *Address {
	a.country = countryName
	return a
}
