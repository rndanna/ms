package handle

import (
	"engine-service/internal/models"
	"engine-service/internal/storage/postgresql"
	"engine-service/pkg/utils"
	"fmt"

	"github.com/labstack/echo/v4"
)

type handle struct {
	db *postgresql.DB
}

func New(db *postgresql.DB) *handle {
	return &handle{
		db: db,
	}
}

func (h *handle) GetEngine(c echo.Context) (err error) {
	var engine *models.Engine

	defer func() {
		utils.ResponseFunc(c, err, engine)
	}()

	id := c.Param("id")

	engine, err = h.db.GetEngine(id)
	if err != nil {
		err = fmt.Errorf("err GetEngine : %w", err)
	}

	return
}
