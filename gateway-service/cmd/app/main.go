package main

import (
	"fmt"
	handle "gateway-service/internal/server/http"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
}

func main() {
	var cfg Config

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	handler := handle.New()
	e := echo.New()
	e.GET("/cars/:id", handler.GetUserCars)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Server.Port)))
}
