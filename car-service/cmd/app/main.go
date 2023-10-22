package main

import (
	handle "car-service/internal/server/http"
	"car-service/internal/storage/postgresql"
	"database/sql"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	Databse struct {
		Postgresql struct {
			Host     string `mapstructure:"host"`
			Port     string `mapstructure:"port"`
			Username string `mapstructure:"username"`
			Password string `mapstructure:"password"`
			Name     string `mapstructure:"name"`
		} `mapstructure:"postgresql"`
	} `mapstructure:"database"`
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

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Databse.Postgresql.Username,
		cfg.Databse.Postgresql.Password,
		cfg.Databse.Postgresql.Host,
		cfg.Databse.Postgresql.Port,
		cfg.Databse.Postgresql.Name,
	)

	pool, err := sql.Open("postgres",
		dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	db := postgresql.New(pool)
	handle := handle.New(db)

	e := echo.New()
	e.GET("/cars/:id", handle.GetUserCars)
	e.GET("/car/:id", handle.GetCar)
	e.GET("/cars/user/:id", handle.GetIDEnginesByUser)
	e.GET("/cars/brand", handle.GetIDEnginesByBrand)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Server.Port)))
}
