package main

import (
	"fmt"
	"os"
	"strconv"
)

type Teman struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var dataTeman = map[int]Teman{
	1: {Nama: "Steve Rogers", Alamat: "Brooklyn", Pekerjaan: "Captain America", Alasan: "Switch career"},
	2: {Nama: "Tony Stark", Alamat: "Manhattan", Pekerjaan: "Iron Man", Alasan: "Iseng aja"},
	3: {Nama: "Peter Parker", Alamat: "Queens", Pekerjaan: "Spider-Man", Alasan: "Magang di Stark Industries"},
	4: {Nama: "Bucky Barnes", Alamat: "Brooklyn", Pekerjaan: "Winter Soldier", Alasan: "Ngikut Steve"},
	5: {Nama: "Bruce Banner", Alamat: "Mexico", Pekerjaan: "Hulk", Alasan: "Penasaran sama Golang"},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Masukkan: go run biodata.go <nomor_absen>")
		return
	}

	absen, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Input salah. Mohon input angka yang tepat.")
		return
	}
	
	teman, found := dataTeman[absen]
	if !found {
		fmt.Println("Teman dengan absen", absen, "tidak ditemukan.")
		return
	}

	displayTemanData(teman, absen)
}

func displayTemanData(teman Teman, absen int) {
	fmt.Println("Data teman dengan absen ke:", absen)
	fmt.Println("Nama:", teman.Nama)
	fmt.Println("Alamat:", teman.Alamat)
	fmt.Println("Pekerjaan:", teman.Pekerjaan)
	fmt.Println("Alasan memilih kelas Golang:", teman.Alasan)
}