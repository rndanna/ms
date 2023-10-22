package handle

import (
	"car-service/internal/models"
	"car-service/internal/storage/postgresql"
	"car-service/pkg/utils"
	"fmt"

	"github.com/labstack/echo/v4"
)

type handle struct {
	db *postgresql.DB
}

var url = "http://localhost:8081/engine/"

func New(db *postgresql.DB) *handle {
	return &handle{
		db: db,
	}
}

func (h *handle) GetCar(c echo.Context) (err error) {
	var car *models.Car

	defer func() {
		utils.ResponseFunc(c, err, car)
	}()

	id := c.Param("id")

	car, err = h.db.GetCar(id)
	if err != nil {
		err = fmt.Errorf("err GetUserEngines : %w", err)

	}

	return
}

func (h *handle) GetUserCars(c echo.Context) (err error) {
	var cars []models.Car

	defer func() {
		utils.ResponseFunc(c, err, cars)
	}()

	id := c.Param("id")

	cars, err = h.db.GetUserCars(id)
	if err != nil {
		err = fmt.Errorf("err GetUserEngines : %w", err)

	}

	return
}

func (h *handle) GetIDEnginesByUser(c echo.Context) (err error) {
	var (
		cars []models.Car
		ids  []int
	)

	defer func() {
		utils.ResponseFunc(c, err, ids)
	}()

	id := c.Param("id")

	cars, err = h.db.GetUserCars(id)
	if err != nil {
		err = fmt.Errorf("err GetUserEngines : %w", err)

	}

	for _, car := range cars {
		ids = append(ids, car.EngineID)
	}

	return
}

func (h *handle) GetIDEnginesByBrand(c echo.Context) (err error) {
	var (
		cars []models.Car
		ids  []int
	)

	defer func() {
		utils.ResponseFunc(c, err, ids)
	}()

	brand := c.QueryParam("brand")
	if brand == "" {
		return fmt.Errorf("err GetCarByBrand : %w", utils.ErrEmptyQuery)
	}

	cars, err = h.db.GetCarByBrand(brand)
	if err != nil {
		err = fmt.Errorf("err GetCarByBrand : %w", err)

	}

	for _, car := range cars {
		ids = append(ids, car.EngineID)
	}

	return
}
