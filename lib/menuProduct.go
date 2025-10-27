package lib

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"text/tabwriter"
	"time"

	"github.com/joho/godotenv"
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

func defaultTime(key string, defaultValue string) string {
	godotenv.Load()
	val, exist := os.LookupEnv(key)
	if !exist {
		return defaultValue
	}
	return val
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

func search(food foods, channel chan foods, wg *sync.WaitGroup, input string) {
	defer wg.Done()
	if strings.Contains(strings.ToLower(food.Name), strings.ToLower(input)) {
		channel <- food
	}
}

func (c *CartItem) Menu(cart *[]CartItem, temps *[]temp, CacheFile string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
			return
		}
	}()
	var input, qty int
	var wg sync.WaitGroup
	var FoodsMenu []foods

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)

	info, err := os.Stat(CacheFile)
	if os.IsNotExist(err) {
		FoodsMenu = fetchData(CacheFile)
	} else {
		timeDefault, _ := strconv.Atoi(defaultTime("TIMEDEFAULT", "15"))
		age := time.Since(info.ModTime())
		if age >= time.Duration(timeDefault)*time.Minute {
			FoodsMenu = fetchData(CacheFile)

		} else {
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
	fmt.Println("\n\n99. Cari")

	fmt.Print("\n\n\n0. Kembali\n\nChoose a product (number): ")
	fmt.Scan(&input)

	switch input {
	case 0:
		return
	case 99:
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nMasukkan nama makanan: ")
		searchInput, _ := reader.ReadString('\n')
		searchInput = strings.TrimSpace(searchInput)

		channelFoods := make(chan foods, len(FoodsMenu))

		for _, food := range FoodsMenu {
			wg.Add(1)
			go search(food, channelFoods, &wg, searchInput)
		}

		wg.Wait()
		close(channelFoods)

		wResult := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
		fmt.Fprintln(wResult, "No\tMenu\tHarga")

		i := 1
		for result := range channelFoods {
			harga := rupiah.FormatInt64ToRp(int64(result.Price))
			fmt.Fprintf(wResult, "%d\t%s\t%s\n", i, result.Name, harga)
			i++
		}

		if i == 1 {
			fmt.Println(" Tidak ada makanan yang cocok.")
		} else {
			wResult.Flush()
		}

		fmt.Println("Choose product (number): ")
		fmt.Scan(&input)
	}
	if input == 0 {
		return
	}
	if input > len(FoodsMenu) {
		panic("Pilihan produk tidak valid")
	}

	fmt.Print("\n\n0. Kembali   \nquantity: ")
	fmt.Scan(&qty)
	if qty == 0 {
		return
	}
	if qty < 0 {
		panic("Quantity harus lebih dari 0")
	}
	chosen := FoodsMenu[input-1]

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
