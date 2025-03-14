package config

import (
	"log"
	"net/netip"
	"strconv"
	"strings"

	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	APIHost    string
	APIPort    string
	DBHost     string
	DBPort     string
	DBName     string
	DBUsername string
	DBPassword string
}

func New() *Config {
	env := readEnv()

	apiAddr, err := netip.ParseAddr(env.String("api.host"))
	if err != nil {
		log.Fatalf("Unable to load env variable: %s", err.Error())
	}

	apiPort, err := strconv.Atoi(env.String("api.port"))
	if err != nil {
		log.Fatalf("Unable to load env variable: %s", err.Error())
	}

	dbAddr, err := netip.ParseAddr(env.String("db.host"))
	if err != nil {
		log.Fatalf("Unable to load env variable: %s", err.Error())
	}

	dbPort, err := strconv.Atoi(env.String("db.port"))
	if err != nil {
		log.Fatalf("Unable to load env variable: %s", err.Error())
	}
	if dbPort <= 0 {
		log.Fatalf("DB port number cannot be negative.")
	}

	if env.String("db.name") == "" {
		log.Fatalf("DB name cannot be empty.")
	}

	if env.String("db.username") == "" {
		log.Fatalf("DB username cannot be empty.")
	}

	if env.String("db.password") == "" {
		log.Fatalf("DB password cannot be empty.")
	}

	return &Config{
		APIHost:    apiAddr.String(),
		APIPort:    strconv.Itoa(apiPort),
		DBHost:     dbAddr.String(),
		DBPort:     strconv.Itoa(dbPort),
		DBName:     env.String("db.name"),
		DBUsername: env.String("db.username"),
		DBPassword: env.String("db.password"),
	}
}

func readEnv() *koanf.Koanf {
	var k = koanf.New(".")

	k.Load(env.Provider("TX_TRACKER_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(strings.TrimPrefix(s, "TX_TRACKER_")), "_", ".", -1)
	}), nil)

	return k
}
