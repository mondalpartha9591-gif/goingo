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

var bookings []Booking
var mu sync.Mutex

func main() {
 http.HandleFunc("/logo", func(w http.ResponseWriter, r *http.Request) {
  http.ServeFile(w, r, "goingo.jpeg")
 })
 http.HandleFunc("/manifest.json", func(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  fmt.Fprint(w, `{"name":"GoinGo","short_name":"GoinGo","start_url":"/book","display":"standalone","background_color":"#e0f2f5","theme_color":"#00b894","icons":[{"src":"/logo","sizes":"192x192","type":"image/jpeg"}]}`)
 })
 http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  http.Redirect(w, r, "/book", 302)
 })
 http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
   r.ParseForm()
   b := Booking{r.FormValue("name"), r.FormValue("phone"), r.FormValue("from"), r.FormValue("to")}
   mu.Lock()
   bookings = append(bookings, b)
   mu.Unlock()
   fmt.Fprint(w, `<h2 style="font-family:Arial;text-align:center">Booking Done! <br><a href="/book">Back</a></h2>`)
   return
  }
  fmt.Fprint(w, `<html><head><title>GoinGo</title><meta name="viewport" content="width=device-width, initial-scale=1"><link rel="manifest" href="/manifest.json"><meta name="theme-color" content="#00b894"><style>body{font-family:Arial;text-align:center;background:#e0f2f5;padding:20px}.card{background:white;max-width:380px;margin:20px auto;padding:25px;border-radius:20px;box-shadow:0 10px 20px rgba(0,0,0,.15)}</style></head><body><div class="card"><img src="/logo" style="max-width:200px"><h2>Book Your Ride</h2><form method="POST"><input name="name" placeholder="Your Name" required style="width:90%;padding:10px;margin:5px"><br><input name="phone" placeholder="Phone" required style="width:90%;padding:10px;margin:5px"><br><input name="from" placeholder="From" required style="width:90%;padding:10px;margin:5px"><br><input name="to" placeholder="To" required style="width:90%;padding:10px;margin:5px"><br><br><button style="padding:12px 30px;background:#00b894;color:white;border:none;border-radius:10px">Book Now</button></form></div><br><a href="/admin">Admin</a></body></html>`)
 })
 http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
  mu.Lock()
  defer mu.Unlock()
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprint(w, "<h2>All Bookings</h2><table border=1 style='margin:auto'><tr><th>Name</th><th>Phone</th><th>From</th><th>To</th></tr>")
  for _, b := range bookings {
   fmt.Fprintf(w, "<tr><td>%s</td><td>%s</td><td>%s</td><td>%s</td></tr>", b.Name, b.Phone, b.From, b.To)
  }
  fmt.Fprint(w, "</table>")
 })
 log.Println("Starting on :10000")
 http.ListenAndServe(":10000", nil)
}
