package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, `<!DOCTYPE html><html><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1"><link href="https://fonts.googleapis.com/css2?family=Poppins:wght@400;600;700;800&display=swap" rel="stylesheet"><style>*{font-family:'Poppins',sans-serif;margin:0;padding:0;box-sizing:border-box}body{background:#f4f6fb;color:#222} .top{background:#fff;padding:14px 20px;display:flex;justify-content:space-between;align-items:center} .logo{font-weight:800;font-size:26px;color:#0d4de3} .logo small{display:block;font-size:10px;color:#666;font-weight:400;margin-top:-4px} .hero{margin:16px;background:linear-gradient(135deg,#0a1931,#1e3a8a);border-radius:22px;padding:24px;color:#fff;display:flex;justify-content:space-between;align-items:center} .hero h1{font-size:21px;line-height:1.2} .hero p{font-size:12px;opacity:.8;margin-top:6px} .hero .btn{background:#3b82f6;border:none;color:#fff;padding:10px 20px;border-radius:24px;margin-top:14px;font-weight:700} .sec{padding:16px} .card{background:#fff;border-radius:16px;padding:16px;display:flex;justify-content:space-between;align-items:center;margin-bottom:12px;box-shadow:0 6px 16px rgba(0,0,0,.05)} .card b{font-size:14px} .card small{color:#777;font-size:11px} .price{color:#0d4de3;font-weight:800;font-size:18px} .book{background:#fff;margin:16px;border-radius:20px;padding:20px;box-shadow:0 8px 24px rgba(0,0,0,.06)} input{width:100%;padding:14px;border:1.5px solid #e2e8f0;border-radius:12px;margin-top:12px} .submit{width:100%;background:#0d4de3;color:#fff;border:none;padding:14px;border-radius:12px;margin-top:16px;font-weight:700;font-size:16px}</style></head><body><div class="top"><div class="logo">GoinGo<small>We Clean, You Shine</small></div><div style="background:#eef2ff;padding:6px 12px;border-radius:20px;font-size:12px">📍 Kolkata</div></div><div class="hero"><div><h1>Professional Car Washing at Your Doorstep</h1><p>Book in 60 seconds & relax. Our experts come to you!</p><button class="btn">Book Now</button></div><div style="font-size:68px">🚗</div></div><div class="sec"><h3>Our Services</h3><br><div class="card"><div><b>🪣 Bucket Wash</b><br><small>Exterior wash & tyre shine</small></div><div class="price">₹299</div></div><div class="card"><div><b>💦 Pressure Wash</b><br><small>High pressure + foam wash</small></div><div class="price">₹499</div></div><div class="card"><div><b>✨ 360° Deep Cleaning</b><br><small>Interior + exterior + vacuum</small></div><div class="price">₹999</div></div></div><div class="book"><h3>Book Your Wash Now</h3><p style="font-size:12px;color:#666">Get confirmed in 5 minutes</p><form method="POST" action="/book"><input name="name" placeholder="Your Name" required><input name="phone" placeholder="Phone Number" required><button class="submit" type="submit">Book Now - It's Free</button></form></div></body></html>`)
	})
	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		phone := r.FormValue("phone")
		fmt.Printf("🔥 NEW BOOKING: %s - %s\n", name, phone)
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, `<body style="font-family:Poppins,sans-serif;text-align:center;padding:60px"><h1 style="color:green">✅ Thank You %s!</h1><p>Booking for %s received.</p><p>We will call you in 5 mins.</p><br><a href="/" style="background:#0d4de3;color:white;padding:12px 24px;border-radius:10px;text-decoration:none">Go Home</a></body>`, name, phone)
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Running on http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}