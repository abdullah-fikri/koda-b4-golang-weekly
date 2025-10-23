package lib

import (
	"fmt"
)

func MainMenu() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Print(r)
			MainMenu()
		}
	}()
	var input string
	cart := []cartItem{}
	history := []cartItem{}
	for {
		fmt.Println("\n\n--- GACOAN DELIVERY ---")
		fmt.Println(`

1. Menu Makanan
2. Checkout Makanan
3. History 
4. Keluar`)

		fmt.Print("Choose a menu: ")
		fmt.Scan(&input)

		switch input {
		case "1":
			Menu(&cart)
		case "2":
			Cart(&cart, &history)
		case "3":
			History(&history)
		case "4":
			return
		default:
			panic("Opsi tidak ditemukan atau sesuai, Masukkan opsi 1 - 4  ")
		}
	}
}
