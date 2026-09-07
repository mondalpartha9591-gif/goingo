package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/book", 302)
	})
	http.HandleFunc("/logo", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "goingo.jpeg")
	})
	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `
		<html><head><title>GoinGo</title><meta name="viewport" content="width=device-width, initial-scale=1">
		<style>body{font-family:Arial;text-align:center;background:#e0f2fe;padding:20px}
		.card{background:white;max-width:380px;margin:20px auto;padding:25px;border-radius:20px;box-shadow:0 10px 25px rgba(0,0,0,.15)}
		img{width:140px;border-radius:20px} h1{color:#0284c7} input,select,button{width:100%;padding:12px;margin:8px 0;border-radius:10px;border:1px solid #ccc} button{background:#0284c7;color:white;font-weight:bold;border:none;cursor:pointer}</style>
		</head><body>
		<div class="card"><img src="/logo"><h1>GoinGo</h1><p>We Clean, You Shine ✨</p>
		<form action="/confirm" method="POST"><input name="name" placeholder="Your Name" required><input name="phone" placeholder="Phone" required>
		<select name="service"><option>Car Wash</option><option>Bike Wash</option><option>Full Detailing</option></select>
		<button>Book Now</button></form></div></body></html>`)
	})
	http.HandleFunc("/confirm", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `<h1 style="text-align:center;margin-top:100px;color:#0284c7">✅ Booking Confirmed! GoinGo is coming! 🚗💦</h1><center><a href="/book">Go Back</a></center>`)
	})
	log.Fatal(http.ListenAndServe(":10000", nil))
}
