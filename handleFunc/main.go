package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	/*
		Informasi host/port perlu dimasukkan ke dalam property .Addr milik objek server.
		Lalu dari objek tersebut panggil method .ListenAndServe() untuk start web server.

		Kelebihan menggunakan http.Server salah satunya adalah kemampuan untuk
		mengubah beberapa konfigurasi default web server Go. Contohnya bisa dilihat
		pada kode berikut, timeout untuk read request dan write request diubah
		menjadi 10 detik.
	*/
	var address = "localhost:1337"
	fmt.Println("Server started at", address)

	http.HandleFunc("/", handlerIndex)
	// http handleFunc with closure func
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Again"))
	})

	http.HandleFunc("GET /person/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		w.Write([]byte("Selamat siang " + name))
	})

	http.HandleFunc("POST /articles", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Buat artikel baru disini"))
	})

	server := new(http.Server)
	server.Addr = address

	server.ReadTimeout = time.Second * 10
	server.WriteTimeout = time.Second * 10

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Hello World"
	w.Write([]byte(message))
}
