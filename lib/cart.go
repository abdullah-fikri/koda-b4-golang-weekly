package lib

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/paimanbandi/rupiah"
)

type CartInterface interface {
	Cart(cart *[]CartItem, history *[]CartItem, temps *[]temp)
	Menu(cart *[]CartItem, temps *[]temp)
	MainMenu()
}

func (c *CartItem) Cart(cart *[]CartItem, history *[]CartItem, temps *[]temp) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			return
		}
	}()
	var input string
	now := time.Now()
	randomInt := rand.Intn(1000)
	fmt.Println("\nIsi keranjang:")
	if len(*temps) == 0 {
		fmt.Println("Keranjang kosong.")
	} else {
		totalHarga := 0
		for i, item := range *temps {
			fmt.Printf("%d. %s \nHarga: %s \nquantity: %dpcs \nSubTotal: %s\n", i+1, item.name, rupiah.FormatInt64ToRp(int64(item.price)), item.qty, rupiah.FormatInt64ToRp(int64(item.total)))
			totalHarga += item.total
		}
		fmt.Printf("\nTotal harga:  %s\n", rupiah.FormatInt64ToRp(int64(totalHarga)))
	}

	fmt.Print("\nBayar Sekarang? (y/n): ")
	fmt.Scan(&input)

	switch input {
	case "y":

		if len(*temps) <= 0 {
			panic("Keranjang Kosong, tidak dapat melakukan checkout!")
		}

		fmt.Println("\n\n--- INVOICE --- \n", now, "\nId Pesanan: 13254 -", randomInt, "\n\n\nPembayaran berhasil. Terima kasih sudah memesan!", "\n\n\nPress enter to back...")
		reader := bufio.NewReader(os.Stdin)
		_, _ = reader.ReadString('\n')

		for _, t := range *temps {
			*cart = append(*cart, CartItem{
				name:         t.name,
				price:        t.price,
				qty:          t.qty,
				total:        t.total,
				dateCheckout: now,
				idPesanan:    randomInt,
			})
		}
		*temps = []temp{}

		*history = append(*history, *cart...)
		*cart = []CartItem{}
	case "n":
		fmt.Print("Kembali ke menu utama...")
	default:
		panic("Opsi yang anda masukkan tidak valid")
	}
}
