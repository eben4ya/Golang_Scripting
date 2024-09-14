package main

import (
	"fmt"
)

func main() {
    var celsius float64
    var choice int

    fmt.Print("Masukkan suhu dalam Celsius: ")
    fmt.Scanln(&celsius)

    fmt.Println("Pilih konversi suhu: ")
    fmt.Println("1. Fahrenheit")
    fmt.Println("2. Kelvin")
    fmt.Println("3. Reamur")
    fmt.Print("Masukkan pilihan (1-3): ")
    fmt.Scanln(&choice)

    switch choice {
    case 1:
        fahrenheit := (celsius * 9 / 5) + 32
        fmt.Printf("Suhu dalam Fahrenheit: %.2f°F\n", fahrenheit)
    case 2:
        kelvin := celsius + 273.15
        fmt.Printf("Suhu dalam Kelvin: %.2fK\n", kelvin)
    case 3:
        reamur := celsius * 4 / 5
        fmt.Printf("Suhu dalam Reamur: %.2f°Re\n", reamur)
    default:
        fmt.Println("Pilihan tidak valid. Silakan masukkan angka 1-3.")
    }
}
