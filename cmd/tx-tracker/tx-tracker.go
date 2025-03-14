package main

import (
	"github.com/waiyan13/tx-tracker/config"
	"github.com/waiyan13/tx-tracker/internal/db"
)

func main() {
	cfg := config.New()

	conn := db.Connect(cfg)
	defer conn.Close()
}
