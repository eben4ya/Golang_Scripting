package main

import "fmt"

type Mahasiswa struct {
	Nama     string
	Angkatan string	
	NIM      string
}

func updateNIMbyRef(m *Mahasiswa) {
	// update NIM
	m.NIM = "22/TK/2022"
}

func main(){
	mhs1 := Mahasiswa{
		Nama: "Benaya",
		Angkatan: "2022",
		NIM: "22/TK",
	}

	fmt.Println(mhs1)

	// update NIM
	updateNIMbyRef(&mhs1)
	fmt.Println(mhs1)
}