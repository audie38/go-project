package main

import "fmt"

func main(){
	var (
		name1 = "Audie"
		name2 = "Audie"
		nilaiAkhir = 90
		absensi = 80
	)

	result := name1 == name2
	fmt.Println(result)

	result = 100 > 200
	fmt.Println(result)

	result = 100 != 200
	fmt.Println(result)

	lulusNilaiAkhir := nilaiAkhir > 80
	lulusAbsensi := absensi > 80
	lulus := lulusNilaiAkhir && lulusAbsensi
	fmt.Println("Lulus: ", lulus)

}