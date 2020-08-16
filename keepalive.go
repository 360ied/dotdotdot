package dotdotdot

import (
	"fmt"
	"net/http"
)

func keepalive() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "I'm alive")

		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8080", nil)
}
