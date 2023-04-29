package main

import "fmt"

type Address struct {
	city, province, country string
}

func main() {
	addr1 := Address{
		city:     "Subang",
		province: "Jawa Barat",
		country:  "Indonesia",
	}

	// pointer reference to addr1 
	addr2 := &addr1 // use pointer (& symbol) to change go default pass by value to pass by reference
	addr2.city = "Jakarta"

	// if reassign, addr1 will not change (pointer reference to struct address)
	addr2 = &Address{
		city: "Jakarta",
		province: "DKI Jakarta",
		country: "Indonesia",
	}

	fmt.Println("Address1 not change")
	fmt.Println("Addr1", addr1)
	fmt.Println("Addr2",*addr2)

	var addr3 *Address = &addr1

	// use the *operator to make the address 1 changes as well (pointer reference to the whole addr3 instance)
	*addr3 = Address{
		city: "Malang",
		province: "Jawa Timur",
		country: "Indonesia",
	}
	
	fmt.Println("Address1 change")
	fmt.Println("Addr1", addr1)
	fmt.Println("Addr3",*addr3)

	// create new pointer instance
	var addr4 *Address = new(Address)
	addr4.city = "Singapore"
	fmt.Println(addr4)
}