package main

import "fmt"

func main() {

	// Array -> fixed size, slice -> size changeable
	// Slices from Array
	bulan := [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	fmt.Println(bulan)

	slice := bulan[4:7]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	fmt.Println(days)
	holidaySlice := days[5:]
	fmt.Println(holidaySlice)
	
	holidaySlice[0] = "Malam Minggu"
	fmt.Println(days)

	daySlice2 := append(holidaySlice, "Hari Raya")
	fmt.Println(daySlice2)
	fmt.Println(cap(daySlice2))
	daySlice2[0] = "Cuti Bersama"
	fmt.Println(daySlice2)
	fmt.Println(days)

	// Create Slice
	newSlice := make([]string, 2, 5) //len, capacity

	newSlice[0] = "Audie"
	newSlice[1] = "Milson"

	fmt.Println(newSlice)

	// Copy Slice
	fmt.Println("Copy Slice")

	copySlice := make([]string, len(newSlice), cap(newSlice))
	
	copy(copySlice, newSlice) // target, source

	fmt.Println(copySlice)

	// Array vs Slice

	iniArray := [...]int {1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}