package main

import "fmt"

type Address struct {
	City, Province, Country string
}

type Man struct{
	Name string
}

func (man *Man) Married(){
	man.Name = "Mr. " + man.Name
}

func ChangeAddressToIndonesia(address Address){
	address.Country = "Indonesia"
}

func ChangeAddrToIndonesia(address *Address){
	address.Country = "Indonesia"
}

func main() {
	addr1 := Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "Indonesia",
	}

	// pointer reference to addr1 
	addr2 := &addr1 // use pointer (& symbol) to change go default pass by value to pass by reference
	addr2.City = "Jakarta"

	// if reassign, addr1 will not change (pointer reference to struct address)
	addr2 = &Address{
		City: "Jakarta",
		Province: "DKI Jakarta",
		Country: "Indonesia",
	}

	fmt.Println("Address1 not change")
	fmt.Println("Addr1", addr1)
	fmt.Println("Addr2",*addr2)

	var addr3 *Address = &addr1

	// use the *operator to make the address 1 changes as well (pointer reference to the whole addr3 instance)
	*addr3 = Address{
		City: "Malang",
		Province: "Jawa Timur",
		Country: "Indonesia",
	}
	
	fmt.Println("Address1 change")
	fmt.Println("Addr1", addr1)
	fmt.Println("Addr3",*addr3)

	// create new pointer instance
	var addr4 *Address = new(Address)
	addr4.City = "Singapore"
	ChangeAddressToIndonesia(*addr4) // data will not changed
	fmt.Println(addr4)
	ChangeAddrToIndonesia(addr4) // data will change
	fmt.Println(addr4)

	ichigo := Man{Name: "Kurosaki"}
	ichigo.Married()
	fmt.Println(ichigo)

}