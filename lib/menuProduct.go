package lib

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/paimanbandi/rupiah"
)

type foods struct {
	Name  string
	Price int
}

type CartItem struct {
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

func fetchData(CacheFile string) []foods {
	resp, err := http.Get("https://raw.githubusercontent.com/abdullah-fikri/koda-b4-golang-weekly-data/main/main.json")
	if err != nil {
		fmt.Println("failed fetch data")
	}

	body, _ := io.ReadAll(resp.Body)
	var FoodsMenu []foods
	json.Unmarshal(body, &FoodsMenu)
	os.WriteFile(CacheFile, body, 0644)
	return FoodsMenu
}

func (c *CartItem) Menu(cart *[]CartItem, temps *[]temp, CacheFile string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			return
		}
	}()
	var input, qty int

	var FoodsMenu []foods

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

	info, err := os.Stat(CacheFile)
	if os.IsNotExist(err) {
		fmt.Println("Cache not found, fetching data...")
		FoodsMenu = fetchData(CacheFile)
	} else {
		age := time.Since(info.ModTime())
		if age >= 15*time.Second {
			fmt.Println("Cache expired, fetching agin...")
			FoodsMenu = fetchData(CacheFile)
		} else {
			fmt.Println("used cache...")
			data, err := os.ReadFile(CacheFile)
			if err != nil {
				fmt.Println("NOT reading cache")
			} else {
				json.Unmarshal(data, &FoodsMenu)
			}
		}
	}
	fmt.Fprintln(w, "No\tMenu\tHarga")

	for i, food := range FoodsMenu {
		hargaRupiah := rupiah.FormatInt64ToRp(int64(food.Price))
		fmt.Fprintf(w, "%d\t%s\t%s\n", i+1, food.Name, hargaRupiah)
	}

	w.Flush()

	fmt.Print("\n\n\n0. Kembali\n\nChoose a product (number): ")
	fmt.Scan(&input)

	if input == 0 {
		return
	}
	if input > len(FoodsMenu) {
		panic("Pilihan produk tidak valid")
	}

	chosen := FoodsMenu[input-1]

	fmt.Print("\n\n0. Kembali   \nquantity: ")
	fmt.Scan(&qty)
	if qty == 0 {
		return
	}
	if qty < 0 {
		panic("Quantity harus lebih dari 0")
	}

	*temps = append(*temps, temp{
		name:  chosen.Name,
		price: chosen.Price,
		qty:   qty,
		total: chosen.Price * qty,
	})

	fmt.Printf("%s x%d ditambahkan ke keranjang.\n", chosen.Name, qty)
	fmt.Print("\n\n1. Pesan lagi \nPress enter to back..")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	switch text {
	case "1":
		c.Menu(cart, temps, CacheFile)
	default:
		panic("Opsi yang anda masukkan tidak sesuai")
	}
}
