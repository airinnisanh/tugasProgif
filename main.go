package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type barang struct {
	ID_Barang int
	Nama_Barang string
	Kategori string
	Harga int
}

type barangjson struct {
	ID_Barang int `json: "ID_Barang, omitempty"`
	Nama_Barang string `json: "Nama_Barang, omitempty"`
	Kategori string `json: "Kategori, omitempty"`
	Harga int `json: "Harga, omitempty"`
}

func main() {
	port := 8080
	http.HandleFunc("/databarang/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case "GET":
				s := r.URL.Query().Get("Parameter")
				if (strings.Compare("urutharga", s) == 0) {
					GetBarangByHarga(w,r)
				} else
				if (s != "") {
					GetBarangByKategori(w,r,s)
				} else {
					GetBarang(w,r)
				}
			case "POST":
				InputBarang(w,r)	
			default:
				http.Error(w,"invalid",405)
		}
	})
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func GetBarang(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/data_barang")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	x:= barang{}

	rows, err := db.Query("SELECT * FROM data_barang")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&x.ID_Barang, &x.Nama_Barang, &x.Kategori, &x.Harga)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(&x)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func GetBarangByKategori(w http.ResponseWriter, r *http.Request, Kategori string) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/data_barang")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	x := barang{}

	rows, err := db.Query("SELECT * FROM data_barang WHERE Kategori like?", Kategori)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&x.ID_Barang, &x.Nama_Barang, &x.Kategori, &x.Harga)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(&x)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func GetBarangByHarga(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/data_barang")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	x := barang{}

	rows, err := db.Query("SELECT * FROM data_barang ORDER BY Harga DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&x.ID_Barang, &x.Nama_Barang, &x.Kategori, &x.Harga)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(&x)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}	
}


func InputBarang(w http.ResponseWriter, r *http.Request) {
	var barangnew barangjson

	dec := json.NewDecoder(r.Body)
	err:=dec.Decode(&barangnew)
	if err != nil{
		log.Fatal(err)
	}
	defer r.Body.Close()

	db,err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/data_barang")
	if err != nil{
		log.Fatal(err)
	}

	stmt, err := db.Prepare("INSERT INTO data_barang (Nama_Barang, Kategori, Harga) VALUES (?,?,?)")
	if err != nil{
		log.Fatal(err)
	}

	_, err = stmt.Exec(barangnew.Nama_Barang, barangnew.Kategori, barangnew.Harga)
}