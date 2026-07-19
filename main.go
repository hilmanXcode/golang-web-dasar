package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	/*
		Fungsi http.HandleFunc() digunakan untuk keperluan routing.
		Parameter pertama adalah rute dan parameter ke-2 adalah handler-nya.
	*/
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	/*
		Pada contoh di atas, ketika rute yang tidak terdaftar
		diakses, maka secara otomatis handler rute / yang dipanggil.
	*/

	var address = "localhost:1337"
	fmt.Printf("server started at %s\n", address)

	err := http.ListenAndServe(address, nil)

	if err != nil {
		log.Fatal(err)
	}
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello World"
	w.Write([]byte(message))
}

/*
	Method Write() milik parameter pertama (yang bertipe http.ResponseWriter),
	digunakan untuk meng-output-kan data ke HTTP response. Argumen method adalah data
	yang ingin dijadikan output, dituliskan dalam bentuk []byte.

	Pada contoh ini, data yang akan kita tampilkan bertipe string, maka perlu
	dilakukan casting dari string ke []byte. Praktiknya bisa dilihat seperti pada kode
	di atas, di bagian w.Write([]byte(message))
*/
