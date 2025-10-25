package lib

import (
	"bufio"
	"fmt"
	"os"

	"github.com/paimanbandi/rupiah"
)

func History(history *[]cartItem) {
	var point []int
	if len(*history) == 0 {
		fmt.Print("\n\nBelum ada riwayat pembelian. \n\nPress enter to back...")
		reader := bufio.NewReader(os.Stdin)
		_, _ = reader.ReadString('\n')

		return
	}

	for _, product := range *history {
		invoice := false
		for _, p := range point {
			if p == product.idPesanan {
				invoice = true
			}
		}
		if invoice {
			continue
		}
		fmt.Println("\n\n--- INVOICE ---\n\n", product.dateCheckout, "\n\nId Pesanan 13254 - ", product.idPesanan)

		nomor := 0
		for _, item := range *history {
			if item.idPesanan == product.idPesanan {
				fmt.Println("\n\n", nomor+1, "Makanan:", item.name, " \n\nHarga:", rupiah.FormatInt64ToRp(int64(item.price)), "\n\nQuantity:", item.qty)
				nomor++
			}
		}

		point = append(point, product.idPesanan)
	}
	fmt.Print("Press enter to back...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}
