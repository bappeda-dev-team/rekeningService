package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"rekeningService/model/web"
	"rekeningService/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RekeningControllerImpl struct {
	RekeningService service.RekeningService
}

func NewRekeningControllerImpl(rekeningService service.RekeningService) *RekeningControllerImpl {
	return &RekeningControllerImpl{
		RekeningService: rekeningService,
	}
}

// @Summary Create Rekening
// @Description Create new Rekening
// @Tags Rekening
// @Accept json
// @Produce json
// @Param data body web.RekeningCreateRequest true "Rekening Create Request"
// @Success 201 {object} web.WebResponse{data=web.RekeningResponse} "Created"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /rekening [post]
func (controller *RekeningControllerImpl) Create(c echo.Context) error {
	rekeningCreateRequest := web.RekeningCreateRequest{}
	err := c.Bind(&rekeningCreateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	rekeningResponse, err := controller.RekeningService.Create(c.Request().Context(), rekeningCreateRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
			Data:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   rekeningResponse,
	})
}

// @Summary Update Rekening
// @Description Update Rekening
// @Tags Rekening
// @Accept json
// @Produce json
// @Param id path int true "Rekening ID"
// @Param data body web.RekeningUpdateRequest true "Rekening Update Request"
// @Success 200 {object} web.WebResponse{data=web.RekeningResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /rekening/{id} [put]
func (controller *RekeningControllerImpl) Update(c echo.Context) error {
	rekeningUpdateRequest := web.RekeningUpdateRequest{}
	err := c.Bind(&rekeningUpdateRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   err.Error(),
		})
	}

	rekeningUpdateRequest.Id, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD_REQUEST",
			Data:   "Format ID tidak valid",
		})
	}

	rekeningResponse, err := controller.RekeningService.Update(c.Request().Context(), rekeningUpdateRequest)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || err.Error() == "rekening tidak ditemukan" {
			return c.JSON(http.StatusNotFound, web.WebResponse{
				Code:   http.StatusNotFound,
				Status: "NOT_FOUND",
			})
		}
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
			Data:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   rekeningResponse,
	})
}

// @Summary Delete Rekening
// @Description Delete Rekening
// @Tags Rekening
// @Accept json
// @Produce json
// @Param id path int true "Rekening ID"
// @Success 200 {object} web.WebResponse "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 404 {object} web.WebResponse "Not Found"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /rekening/{id} [delete]
func (controller *RekeningControllerImpl) Delete(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "ID tidak boleh kosong",
		})
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "Format ID tidak valid",
			Data:    err.Error(),
		})
	}

	err = controller.RekeningService.Delete(c.Request().Context(), idInt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || err.Error() == "rekening tidak ditemukan" {
			return c.JSON(http.StatusNotFound, web.WebResponse{
				Code:    http.StatusNotFound,
				Status:  "NOT_FOUND",
				Message: fmt.Sprintf("Rekening dengan ID: %s tidak ditemukan.", id),
			})
		}
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "Terjadi kesalahan saat menghapus data",
			Data:    err.Error(),
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: fmt.Sprintf("Rekening dengan ID: %s berhasil dihapus.", id),
	})
}

// @Summary Get Rekening by ID
// @Description Get Rekening detail by ID
// @Tags Rekening
// @Accept json
// @Produce json
// @Param id path int true "Rekening ID"
// @Success 200 {object} web.WebResponse{data=web.RekeningResponse} "OK"
// @Failure 400 {object} web.WebResponse "Bad Request"
// @Failure 404 {object} web.WebResponse "Not Found"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /rekening/{id} [get]
func (controller *RekeningControllerImpl) FindById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "Format ID tidak valid.",
		})
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, web.WebResponse{
			Code:    http.StatusBadRequest,
			Status:  "BAD_REQUEST",
			Message: "Format ID tidak valid.",
		})
	}

	rekeningResponse, err := controller.RekeningService.FindById(c.Request().Context(), idInt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || err.Error() == "rekening tidak ditemukan" {
			return c.JSON(http.StatusNotFound, web.WebResponse{
				Code:    http.StatusNotFound,
				Status:  "NOT_FOUND",
				Message: fmt.Sprintf("Rekening dengan ID: %s tidak ditemukan.", id),
			})
		}
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:    http.StatusInternalServerError,
			Status:  "INTERNAL_SERVER_ERROR",
			Message: "Terjadi kesalahan fatal.",
		})
	}

	if rekeningResponse.Id == 0 {
		return c.JSON(http.StatusNotFound, web.WebResponse{
			Code:    http.StatusNotFound,
			Status:  "NOT_FOUND",
			Message: fmt.Sprintf("Rekening dengan ID: %s tidak ditemukan.", id),
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   rekeningResponse,
	})
}

// @Summary Get All Rekening
// @Description Get All Rekening
// @Tags Rekening
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse{data=[]web.RekeningResponse} "OK"
// @Failure 500 {object} web.WebResponse "Internal Server Error"
// @Router /rekening [get]
func (controller *RekeningControllerImpl) FindAll(c echo.Context) error {
	rekeningResponses, err := controller.RekeningService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, web.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL_SERVER_ERROR",
			Data:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   rekeningResponses,
	})
}
