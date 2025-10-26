# Gacoan — Aplikasi Pemesanan Makanan Berbasis Terminal

A simple **Golang CLI app** that simulates a food ordering system — inspired by "Mie Gacoan".  
Users can view menus, add items to a cart, checkout, and see past order history directly in the terminal.

---

## 🚀 Fitur Utama

✅ **Menu Makanan & Minuman**  
Tampilkan daftar menu lengkap dengan harga yang diformat otomatis ke rupiah.

✅ **Keranjang Belanja (Cart)**  
Tambah beberapa item, ubah jumlah, dan lihat total harga sebelum checkout.

✅ **Checkout & Invoice**  
Cetak invoice sederhana dengan ID transaksi dan waktu checkout.

✅ **Riwayat Transaksi (History)**  
Lihat semua transaksi sebelumnya yang tersimpan di memori selama program berjalan.

---

## 🧰 Teknologi yang Digunakan

- **Golang**
- **text/tabwriter** — untuk membuat tampilan tabel rapi di terminal
- **github.com/paimanbandi/rupiah** — untuk format harga ke Rupiah
- **bufio & os** — untuk membaca input terminal
- **time & math/rand** — untuk membuat ID pesanan dan waktu transaksi

---
