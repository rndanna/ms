package handle

import (
	"encoding/json"
	"fmt"
	"gateway-service/internal/models"
	"gateway-service/pkg/api"
	"gateway-service/pkg/utils"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type hanlder struct {
	base api.BaseClient
}

func New() *hanlder {
	return &hanlder{
		base: api.BaseClient{
			HTTPClient: &http.Client{
				Timeout: 10 * time.Second,
			},
		},
	}
}

func (h *hanlder) GetUserCars(c echo.Context) (err error) {
	var cars []models.Car

	defer func() {
		utils.ResponseFunc(c, err, cars)
	}()

	url := "http://localhost:8081/cars/"

	id := c.Param("id")

	req, err := http.NewRequest("GET", fmt.Sprint(url, id), nil)
	if err != nil {
		err = fmt.Errorf("failed to create new request due to error: %v", err)
	}

	response, err := h.base.SendRequest(req)
	if err != nil {
		err = fmt.Errorf("failed to send request due to error: %v", err)
	}

	if err = json.Unmarshal(response, &cars); err != nil {
		err = fmt.Errorf("failed to Unmarshal due to error: %v", err)
	}

	fmt.Println(cars)
	return
}

// func (h *hanlder) GetUserEngines(c echo.Context) (err error) {
// 	var cars []models.Car
// 	var engines []models.Engine
// 	var response models.CarResponse
// 	var engineResqp models.EngineResponse

// 	defer func() {
// 		utils.ResponseFunc(c, err, engines)
// 	}()

// 	url := "http://localhost:8081/cars/"

// 	id := c.Param("id")

// 	client := &http.Client{}

// 	req, err := http.NewRequest("GET", fmt.Sprint(url, id), nil)
// 	if err != nil {
// 		err = fmt.Errorf("errAPI GetEngine : %w", err)
// 	}

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		err = fmt.Errorf("errAPI GetEngine : %w", err)
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		err = fmt.Errorf("errAPI GetEngine : %w", err)
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		err = fmt.Errorf("errAPI GetEngine : %w", err)
// 	}

// 	err = json.Unmarshal(body, &response)
// 	if err != nil {
// 		err = fmt.Errorf("errAPI GetEngine Unmarshal: %w", err)
// 	}

// 	cars = append(cars, response.Data...)

// 	for _, car := range cars {
// 		url := "http://localhost:8082/engine/"

// 		req, err = http.NewRequest("GET", fmt.Sprint(url, car.EngineID), nil)
// 		if err != nil {
// 			err = fmt.Errorf("errAPI GetEngine : %w", err)
// 		}

// 		resp, err := client.Do(req)
// 		if err != nil {
// 			err = fmt.Errorf("errAPI GetEngine : %w", err)
// 		}

// 		if resp.StatusCode != http.StatusOK {
// 			err = fmt.Errorf("errAPI GetEngine : %w", err)
// 		}

// 		body, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			err = fmt.Errorf("errAPI GetEngine : %w", err)
// 		}

// 		if err := json.Unmarshal(body, &engineResqp); err != nil {
// 			err = fmt.Errorf("errAPI GetEngine Unmarshal: %w", err)
// 		}

// 	}
// 	engines = append(engines, engineResqp.Data)

// 	return
// }

// func (h *hanlder) GetCarEngine(c echo.Context) (err error) {
// 	var car models.Car
// 	var engine models.Engine
// 	var response models.GetCarDTO
// 	var engineResqp models.EngineResponse

// 	defer func() {
// 		utils.ResponseFunc(c, err, engine)
// 	}()

// 	url := "http://localhost:8081/car/"

// 	id := c.Param("id")

// 	client := &http.Client{}

// 	req, err := http.NewRequest("GET", fmt.Sprint(url, id), nil)
// 	if err != nil {
// 		err = fmt.Errorf("errAPI GetEngine : %w", err)
// 	}

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		err = fmt.Errorf("errAPI GetEngine : %w", err)
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		err = fmt.Errorf("errAPI GetEngine : %w", err)
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		err = fmt.Errorf("errAPI GetEngine : %w", err)
// 	}

// 	err = json.Unmarshal(body, &response)
// 	if err != nil {
// 		err = fmt.Errorf("errAPI GetEngine Unmarshal: %w", err)
// 	}

// 	car = response.Data

// 	url = "http://localhost:8082/engine/"

// 	req, err = http.NewRequest("GET", fmt.Sprint(url, car.EngineID), nil)
// 	if err != nil {
// 		err = fmt.Errorf("errAPI GetEngine : %w", err)
// 	}

// 	resp, err = client.Do(req)
// 	if err != nil {
// 		err = fmt.Errorf("errAPI GetEngine : %w", err)
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		err = fmt.Errorf("errAPI GetEngine : %w", err)
// 	}

// 	body, err = io.ReadAll(resp.Body)
// 	if err != nil {
// 		err = fmt.Errorf("errAPI GetEngine : %w", err)
// 	}

// 	err = json.Unmarshal(body, &engineResqp)
// 	if err != nil {
// 		err = fmt.Errorf("errAPI GetEngine Unmarshal: %w", err)
// 	}
// 	engine = engineResqp.Data

// 	return
// }

// func (h *hanlder) GetUserEnginesByBrand(c echo.Context) (err error) {
// 	var cars []models.Car
// 	var engines []models.Engine
// 	var response models.CarResponse
// 	var engineResqp models.EngineResponse

// 	defer func() {
// 		utils.ResponseFunc(c, err, engines)
// 	}()

// 	url := "http://localhost:8081/cars/"

// 	id := c.Param("id")

// 	client := &http.Client{}

// 	req, err := http.NewRequest("GET", fmt.Sprint(url, id), nil)
// 	if err != nil {
// 		err = fmt.Errorf("errAPI GetEngine : %w", err)
// 	}

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		err = fmt.Errorf("errAPI GetEngine : %w", err)
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		err = fmt.Errorf("errAPI GetEngine : %w", err)
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		err = fmt.Errorf("errAPI GetEngine : %w", err)
// 	}

// 	err = json.Unmarshal(body, &response)
// 	if err != nil {
// 		err = fmt.Errorf("errAPI GetEngine Unmarshal: %w", err)
// 	}

// 	cars = append(cars, response.Data...)

// 	for _, car := range cars {
// 		url := "http://localhost:8082/engine/"

// 		req, err = http.NewRequest("GET", fmt.Sprint(url, car.EngineID), nil)
// 		if err != nil {
// 			err = fmt.Errorf("errAPI GetEngine : %w", err)
// 		}

// 		resp, err := client.Do(req)
// 		if err != nil {
// 			err = fmt.Errorf("errAPI GetEngine : %w", err)
// 		}

// 		if resp.StatusCode != http.StatusOK {
// 			err = fmt.Errorf("errAPI GetEngine : %w", err)
// 		}

// 		body, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			err = fmt.Errorf("errAPI GetEngine : %w", err)
// 		}

// 		if err := json.Unmarshal(body, &engineResqp); err != nil {
// 			err = fmt.Errorf("errAPI GetEngine Unmarshal: %w", err)
// 		}

// 	}
// 	engines = append(engines, engineResqp.Data)

// 	return
// }
