package lib

import (
	"fmt"
	"os"
	"path/filepath"
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
	CacheFile := filepath.Join(os.TempDir(), "GacoanApp.json")
	for {
		fmt.Println("\x1bc--- GACOAN DELIVERY ---")
		fmt.Println(`

1. Menu Makanan
2. Checkout Makanan
3. History 
4. Option
5. Keluar`)

		fmt.Print("Choose a menu: ")
		fmt.Scan(&input)
		fmt.Print("\x1bc")

		switch input {
		case "1":
			c.Menu(&cart, &temps, CacheFile)
		case "2":
			c.Cart(&cart, &history, &temps)
		case "3":
			History(&history)
		case "4":
			Option(CacheFile)
		case "5":
			return
		default:
			panic("Opsi tidak ditemukan atau sesuai, Masukkan opsi 1 - 4  ")
		}
	}
}
