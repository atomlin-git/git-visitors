package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

var visitors_count = 0

func main() {
	visitors_count_, err := os.ReadFile("visitors")
	if err != nil {
		panic("cant read visitors")
	}

	visitors_count, err := strconv.Atoi(string(visitors_count_))
	if err != nil {
		visitors_count = 0
	}

	http.HandleFunc("/git", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")

		visitors_count++
		w.Write([]byte(fmt.Sprintf("<svg width=\"300\" height=\"20\" xmlns=\"http://www.w3.org/2000/svg\"><foreignObject width=\"300\" height=\"20\"><div xmlns=\"http://www.w3.org/1999/xhtml\"><li style=\"color: #9198a1; list-style-type: none; font-family: var(--fontStack-monospace, ui-monospace, SFMono-Regular, SF Mono, Menlo, Consolas, Liberation Mono, monospace);\">visitors count - %d</li></div></foreignObject></svg>", visitors_count)))
		os.WriteFile("visitors", []byte(strconv.FormatInt(int64(visitors_count), 10)), 0644)
	})
	http.ListenAndServe(":80", nil)
}
