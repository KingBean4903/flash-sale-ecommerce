package handler

import (
	"fmt"
	"net/http"
)

func StartHTTPServer() {

	http.HandleFunc("/place-order", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,  "Order Received ")
	})

	http.ListenAndServe(":8700", nil)
}
