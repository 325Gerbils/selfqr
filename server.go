package main

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/skip2/go-qrcode"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := "localhost:8080/" + r.URL.Path
		png, err := qrcode.Encode(path, qrcode.Medium, 256)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		b64 := base64.StdEncoding.EncodeToString(png)
		fmt.Fprintf(w, "<html><head></head><body><img src=\"data:image/png;base64,"+b64+"\"></body></html>")
	})
	http.ListenAndServe(":8080", nil)
}
