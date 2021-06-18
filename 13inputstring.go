package main

import "fmt"

var (
	namadepan, namabelakang, s string
)

func main() {
	fmt.Printf("Masukkan nama anda : ")
	fmt.Scanln(&namadepan, &namabelakang)
	fmt.Printf("Hello  %s %s", namadepan, namabelakang)
	fmt.Printf(" Saya Golang bahasa yang menyenangkan")
}
