package lib

import (
	"fmt"
	"os"
)

func Option(CacheFile string) {
	var input int
	fmt.Print("1. Clear Cache \n0. Kembali \n\nChoose: ")
	fmt.Scan(&input)

	if input == 1 {
		err := os.Remove(CacheFile)
		if err != nil {
			fmt.Println("gagal menghapus")
		}
	} else {
		return
	}

}
