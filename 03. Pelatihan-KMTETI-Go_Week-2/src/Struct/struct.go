package main

import "fmt"

// Struct in Go
type mahasiswa struct {
	Nama string
	Angkatan string
	NIM string
	IsAsdos bool
	IPK float64
}

func main(){
	var m1 mahasiswa
	m1.Nama = "Benaya"
	m1.Angkatan = "2022"
	m1.NIM = "22/TK"
	m1.IsAsdos = false
	m1.IPK = 3.51

	fmt.Println(m1)

	m2 := mahasiswa{
		Nama: "Imanuela",
		Angkatan: "2022",
		NIM: "22/TK",
		IsAsdos: true,
		IPK: 3.75,
	}

	m2.IsAsdos = false

	fmt.Println(m2)

	// using Slice of Struct
	var mahasiswaList []mahasiswa
	mahasiswaList = append(mahasiswaList, m1)
	mahasiswaList = append(mahasiswaList, m2)
	fmt.Println(mahasiswaList)

	
}