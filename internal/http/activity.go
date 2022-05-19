package http

import (
	"net/http"
	"pil/domain"
	"pil/internal/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ActivityEchoController struct {
	Service domain.AllServices
}

func (s *ActivityEchoController) GetAllController(c echo.Context) error {
	activity := s.Service.GetAllActivity()
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":  activity,
		"total": len(activity),
	})
}

func (s *ActivityEchoController) GetByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid parameter id",
		})
	}
	activity := s.Service.GetByIdActivity(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": activity,
	})
}

func (s *ActivityEchoController) UpdateController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid parameter id",
		})
	}
	dataRaw := echo.Map{}
	if err := c.Bind(&dataRaw); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error internal server",
		})
	}

	judul := dataRaw["judul"].(string)
	deskripsi := dataRaw["deskripsi"].(string)
	link_yt := dataRaw["link_yt"].(string)
	link_drive := dataRaw["link_drive"].(string)

	data := model.Activity{
		Judul:     judul,
		Deskripsi: deskripsi,
		LinkYt:    link_yt,
		LinkDrive: link_drive,
	}

	s.Service.UpdateActivity(id, data)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data updated",
	})

}

func (s *ActivityEchoController) DeleteController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid parameter id",
		})
	}
	s.Service.DeleteActivity(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data deleted",
	})
}

func (s *ActivityEchoController) CreateController(c echo.Context) error {
	dataRaw := echo.Map{}
	if err := c.Bind(&dataRaw); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error internal server",
		})
	}

	judul := dataRaw["judul"].(string)
	deskripsi := dataRaw["deskripsi"].(string)
	link_yt := dataRaw["link_yt"].(string)
	link_drive := dataRaw["link_drive"].(string)

	data := model.Activity{
		Judul:     judul,
		Deskripsi: deskripsi,
		LinkYt:    link_yt,
		LinkDrive: link_drive,
	}

	s.Service.CreateActivity(data)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Data created",
	})

}
