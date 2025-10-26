package lib

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"text/tabwriter"
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

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
		fmt.Println("============================================================")
		fmt.Fprintln(w, "No\tMenu\tHarga\tQuantity\tSub Total")

		for i, item := range *temps {
			hargaRupiah := rupiah.FormatInt64ToRp(int64(item.price))
			hargaSubTotalRupiah := rupiah.FormatInt64ToRp(int64(item.total))
			fmt.Fprintf(w, "%d\t%s\t%s\t%d pcs\t%s\n",
				i+1,
				item.name,
				hargaRupiah,
				item.qty,
				hargaSubTotalRupiah,
			)
			totalHarga += item.total
		}

		w.Flush()

		fmt.Printf("\nTotal harga: %s\n", rupiah.FormatInt64ToRp(int64(totalHarga)))
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
