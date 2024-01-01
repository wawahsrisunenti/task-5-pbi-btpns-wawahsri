package main

import (
	"net/http"

	"PBI_BTPN/config"
	"PBI_BTPN/router"
)

func main() {
	config.Load()
	r := router.Init()
	r.StaticFS("/uploads", http.Dir("uploads"))
	r.Run(":8080")
}
