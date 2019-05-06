package main

import "net/http"

func main() {
	ConnectDB()
	_ = http.ListenAndServe(":8080", route())
}
