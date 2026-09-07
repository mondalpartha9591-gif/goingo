package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `
		<body style="font-family:Poppins, sans-serif; text-align:center; padding:40px; background:#f0f9ff">
			<h1 style="color:#0284c7; font-size:40px;">Goingo</h1>
			<p style="color:#0ea5e9; letter-spacing:4px;">We Clean, You Shine</p>
			<h2>🚗 Your Car Wash API is LIVE!</h2>
			<a href="/book" style="display:inline-block; margin-top:20px; padding:15px 30px; background:#0284c7; color:white; text-decoration:none; border-radius:10px;">Go to Booking Page</a>
			<p style="margin-top:30px;">API URL: /book , /health</p>
		</body>
		`)
	})

	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `
		<body style="font-family:Poppins; text-align:center; padding:50px; background:#f0f9ff">
			<h1 style="color:#0ea5e9">Goingo - We Clean, You Shine ✨</h1>
			<h2>Book Your Car Wash Now</h2>
			<div style="background:white; padding:20px; border-radius:15px; max-width:400px; margin:20px auto; box-shadow:0 4px 10px rgba(0,0,0,0.1)">
				<p>✅ Basic Wash - ₹199</p>
				<p>✨ Premium Shine - ₹399</p>
				<p>💎 Full Detailing - ₹799</p>
				<button style="padding:15px 30px; background:#0284c7; color:white; border:none; border-radius:10px; font-size:18px; margin-top:10px">Book Now</button>
			</div>
		</body>
		`)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Running on http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}