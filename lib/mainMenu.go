package lib

import (
	"fmt"
)

func (c *CartItem) MainMenu() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Print(r)
			c.MainMenu()
		}
	}()
	var input string
	cart := []CartItem{}
	history := []CartItem{}
	temps := []temp{}
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
			c.Menu(&cart, &temps)
		case "2":
			c.Cart(&cart, &history, &temps)
		case "3":
			History(&history)
		case "4":
			return
		default:
			panic("Opsi tidak ditemukan atau sesuai, Masukkan opsi 1 - 4  ")
		}
	}
}
