package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	addr := ":8106"
	fmt.Printf("🧘 🧘 ZenFlow Yoga - Meditation & Wellness Studio Server đang chạy tại http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
