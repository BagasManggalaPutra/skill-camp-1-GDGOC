package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


// 1. VARIABLE & TIPE DATA

const pi float64 = 3.14

var nama string = "Chandra"
var umur int = 18
var aktif bool = true

// 7. STRUCT

type Mahasiswa struct {
	Nama  string
	Umur  int
	Aktif bool
}

// 8. FUNCTION

func tambah(a int, b int) int {
	return a + b
}

func tampilkanMahasiswa(m Mahasiswa) {
	fmt.Println("Nama :", m.Nama)
	fmt.Println("Umur :", m.Umur)
	fmt.Println("Aktif:", m.Aktif)
}

// 9. POINTER

func ubahUmur(u *int) {
	*u = *u + 1
}

func main() {
	// 2. OPERATOR ARITMATIKA

	a := 10
	b := 3

	fmt.Println("Tambah:", a+b)
	fmt.Println("Kurang:", a-b)
	fmt.Println("Kali  :", a*b)
	fmt.Println("Bagi  :", a/b)
	fmt.Println("Mod   :", a%b)

	// 4. OPERATOR LOGIKA

	fmt.Println(a > b)
	fmt.Println(a == b)
	fmt.Println(a != b)

	// 5. PERCABANGAN
	
	nilai := 80

	if nilai >= 75 {
		fmt.Println("Lulus")
	} else {
		fmt.Println("Tidak Lulus")
	}

	hari := "Senin"
	switch hari {
	case "Senin":
		fmt.Println("Hari kerja")
	case "Minggu":
		fmt.Println("Hari libur")
	default:
		fmt.Println("Hari tidak dikenal")
	}

	// 6. PERULANGAN
	
	fmt.Println("Perulangan for:")
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	fmt.Println("Repeat until (pakai for):")
	i := 1
	for {
		fmt.Println(i)
		i++
		if i > 3 {
			break
		}
	}

	numbers := []int{1, 2, 3}
	fmt.Println("Foreach:")
	for _, v := range numbers {
		fmt.Println(v)
	}

	// STRUCT, FUNCTION, POINTER
	
	mhs := Mahasiswa{
		Nama:  "Bagas",
		Umur:  20,
		Aktif: true,
	}

	tampilkanMahasiswa(mhs)

	fmt.Println("Umur sebelum:", mhs.Umur)
	ubahUmur(&mhs.Umur)
	fmt.Println("Umur sesudah:", mhs.Umur)

	fmt.Println("Hasil tambah:", tambah(5, 7))

	// PENGENALAN GIN
	
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Gin",
			"nama":    nama,
			"umur":    umur,
			"aktif":   aktif,
			"pi":      pi,
		})
	})

	r.GET("/mahasiswa", func(c *gin.Context) {
		c.JSON(http.StatusOK, mhs)
	})

	r.Run(":8080")
}
