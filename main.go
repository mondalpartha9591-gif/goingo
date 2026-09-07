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
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Booking struct {
	Name, Phone, From, To string
}

var (
	bookings []Booking
	mu       sync.Mutex
)

func main() {
	// Logo
	http.HandleFunc("/logo", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "goingo.jpeg")
	})

	// Home -> book e pathabe
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/book", 302)
	})

	// Booking Form
	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			b := Booking{
				Name:  r.FormValue("name"),
				Phone: r.FormValue("phone"),
				From:  r.FormValue("from"),
				To:    r.FormValue("to"),
			}
			mu.Lock()
			bookings = append(bookings, b)
			mu.Unlock()
			fmt.Fprint(w, `<h2 style="font-family:Arial;text-align:center">✅ Booking Confirmed for `+b.Name+`!<br><br><a href="/book">Go Back</a> | <a href="/admin">View All Bookings</a></h2>`)
			return
		}
		// GET - tomar ager design ta e rakhlam
		fmt.Fprint(w, `<html><head><title>GoinGo</title><meta name="viewport" content="width=device-width, initial-scale=1">
		<style>body{font-family:Arial;text-align:center;background:#e0f2f5;padding:20px}.card{background:white;max-width:380px;margin:20px auto;padding:25px;border-radius:20px;box-shadow:0 10px 20px rgba(0,0,0,.15)}</style>
		</head><body>
		<div class="card">
		<img src="/logo" style="max-width:200px"><h2>Book Your Ride</h2>
		<form method="POST">
		<input name="name" placeholder="Your Name" required style="width:90%;padding:10px;margin:5px"><br>
		<input name="phone" placeholder="Phone Number" required style="width:90%;padding:10px;margin:5px"><br>
		<input name="from" placeholder="From" required style="width:90%;padding:10px;margin:5px"><br>
		<input name="to" placeholder="To" required style="width:90%;padding:10px;margin:5px"><br><br>
		<button style="padding:12px 30px;background:#00b894;color:white;border:none;border-radius:10px">Book Now</button>
		</form>
		</div><br><a href="/admin">Admin Panel</a>
		</body></html>`)
	})

	// Admin - ekhane sob booking dekhbe
	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		defer mu.Unlock()
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, "<h2 style='font-family:Arial'>All Bookings - Total: "+fmt.Sprint(len(bookings))+"</h2>")
		if len(bookings)==0 {
			fmt.Fprint(w, "No bookings yet<br><br><a href='/book'>Go to Booking</a>")
			return
		}
		fmt.Fprint(w, "<table border=1 cellpadding=10 style='margin:auto;border-collapse:collapse;font-family:Arial'><tr><th>#</th><th>Name</th><th>Phone</th><th>From</th><th>To</th></tr>")
		for i, b := range bookings {
			fmt.Fprintf(w, "<tr><td>%d</td><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>", i+1, b.Name, b.Phone, b.From, b.To)
		}
		fmt.Fprint(w, "</table><br><a href='/book'>Back</a>")
	})

	log.Println("Starting on :10000")
	http.ListenAndServe(":10000", nil)
}
