# Belajar Go ( Golang )

Selamat datang di repositori belajar Go! Di sini Anda akan menemukan panduan dan sumber daya untuk memulai perjalanan Anda dalam mempelajari bahasa pemrograman Go (Golang).

## Apa itu Go?

Go, atau Golang, adalah bahasa pemrograman yang dikembangkan oleh Google. Go dirancang untuk menjadi sederhana, efisien, dan mudah dipelajari, sambil memberikan kemampuan untuk membangun aplikasi berskala besar dengan kinerja tinggi. Go terkenal dengan sintaksis yang jelas dan fitur-fitur seperti pengelolaan goroutine untuk concurrency.

## Mengapa Memilih Go?

- **Kinerja Tinggi**: Go mengkompilasi kode menjadi kode mesin, yang memungkinkan aplikasi Go berjalan lebih cepat daripada bahasa pemrograman lain yang bersifat interpretatif.
- **Concurrency**: Go memiliki model concurrency yang efisien dengan goroutines, yang memungkinkan penulisan aplikasi yang dapat menangani ribuan proses secara bersamaan.
- **Sederhana dan Mudah Dipelajari**: Go memiliki sintaksis yang minimalis dan mudah dipahami, bahkan untuk pemula.
- **Pengembangan Aplikasi Skala Besar**: Go cocok digunakan untuk membangun aplikasi backend yang membutuhkan skalabilitas tinggi, seperti server web, API, dan sistem distribusi.

## Langkah-Langkah untuk Mulai Belajar Go

### 1. Instalasi Go

Untuk mulai menggunakan Go, Anda perlu menginstalnya di komputer Anda. Ikuti langkah-langkah berikut:

- **Windows**: Unduh Go dari [situs resmi Go](https://golang.org/dl/) dan ikuti petunjuk instalasi untuk Windows.
- **macOS**: Anda bisa menggunakan Homebrew dengan perintah `brew install go`.
- **Linux**: Gunakan perintah paket seperti `sudo apt install golang` (Ubuntu) atau ikuti petunjuk di [situs resmi Go](https://golang.org/doc/install).

### 2. Code Go Pertama Anda

Setelah Go terinstal, buat file baru dengan nama `hello.go` dan tulis kode berikut:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
