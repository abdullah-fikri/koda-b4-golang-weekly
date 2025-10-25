package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/paimanbandi/rupiah"
)

type foods struct {
	name  string
	price int
}

type cartItem struct {
	name         string
	price        int
	qty          int
	total        int
	dateCheckout time.Time
	idPesanan    int
}

type temp struct {
	name  string
	price int
	qty   int
	total int
}

func Menu(cart *[]cartItem, temps *[]temp) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			return
		}
	}()
	var input, qty int

	foodsMenu := []foods{
		{name: "Mie Gacoan Lv.0", price: 15000},
		{name: "Mie Gacoan Lv.1", price: 15000},
		{name: "Mie Gacoan Lv.2", price: 15000},
		{name: "Mie Gacoan Lv.3", price: 15000},
		{name: "Mie Hompimpa Lv.0", price: 15000},
		{name: "Mie Hompimpa Lv.1", price: 15000},
		{name: "Mie Hompimpa Lv.2", price: 15000},
		{name: "Mie Hompimpa Lv.3", price: 15000},
		{name: "Air Mineral", price: 6500},
		{name: "Ice Tea", price: 6500},
		{name: "Udang Keju", price: 12000},
		{name: "Udang Rambutan", price: 12000},
		{name: "Dimsum Ayam", price: 12000},
	}

	for i, food := range foodsMenu {
		hargaRupiah := food.price
		fmt.Printf("%d. %s - Rp %s\n", i+1, food.name, rupiah.FormatInt64ToRp(int64(hargaRupiah)))
	}

	fmt.Print("\n\n\n0. Kembali\n\nChoose a product (number): ")
	fmt.Scan(&input)

	if input == 0 {
		return
	}
	if input > len(foodsMenu) {
		panic("Pilihan produk tidak valid")
	}

	chosen := foodsMenu[input-1]

	fmt.Print("\n\n0. Kembali   \nquantity: ")
	fmt.Scan(&qty)
	if qty == 0 {
		return
	}
	if qty < 0 {
		panic("Quantity harus lebih dari 0")
	}

	*temps = append(*temps, temp{
		name:  chosen.name,
		price: chosen.price,
		qty:   qty,
		total: chosen.price * qty,
	})

	fmt.Printf("%s x%d ditambahkan ke keranjang.\n", chosen.name, qty)
	fmt.Print("\n\n1. Pesan lagi \nPress enter to back..")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	switch text {
	case "1":
		Menu(cart, temps)
	default:
		panic("Opsi yang anda masukkan tidak sesuai")
	}
}
