package routes

import (
	"pil/config"
	"pil/internal/http"
	"pil/repository"
	"pil/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func AllRoutesAPI(e *echo.Echo, db *gorm.DB, conf config.Config) {
	repo := repository.NewRepository(db)
	service := service.NewService(repo, conf)
	c := http.ActivityEchoController{
		Service: service,
	}
	e.POST("/activity", c.CreateController)
	e.GET("/activity", c.GetAllController)
	e.GET("/activity/:id", c.GetByIdController)
	e.PUT("/activity/:id", c.UpdateController)
	e.DELETE("/activity/:id", c.DeleteController)
}
