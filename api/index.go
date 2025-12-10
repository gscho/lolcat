package handler

import (
	"net/http"
	"fmt"
)

// func LolcatPageHandler(w http.ResponseWriter, r *http.Request) {
// 	html := `
// 		<!DOCTYPE html>
// 		<html>
// 		<head>
// 			<title>Random LOLCAT</title>
// 			<style>
// 				body { background: #f2f2f2; text-align: center; font-family: sans-serif; }
// 				img { max-width: 90%%; margin-top: 40px; border-radius: 10px; }
// 				h1 { color: #333; margin-top: 20px; }
// 			</style>
// 		</head>
// 		<body>
// 			<h1>Random LOLCAT 😹</h1>
// 			<img src="https://cataas.com/cat/says/LOL" alt="LOLCAT">
// 		</body>
// 		</html>
// 	`

// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	w.Write([]byte(html))
// }

func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}