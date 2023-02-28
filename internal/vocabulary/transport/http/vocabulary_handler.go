package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/minhquy1903/flashcard-api/internal/vocabulary/dto"
	"github.com/minhquy1903/flashcard-api/internal/vocabulary/service"
)

type VocabularyHandler struct {
	vocabSvc *service.VocabularyService
}

func NewVocabularyHandler(e *echo.Echo, vs *service.VocabularyService) {
	handler := &VocabularyHandler{
		vocabSvc: vs,
	}

	e.GET("/vocabularies", handler.GetVocabularies)
	e.POST("/vocabularies", handler.CreateVocabulary)
	// e.PUT("/vocabularies/:id", handler.UpdateVocabulary)
	// e.DELETE("/vocabularies/:id", handler.DeleteVocabulary)
}

func (h *VocabularyHandler) CreateVocabulary(c echo.Context) error {
	var req dto.CreateVocabularyInput

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.vocabSvc.Create(c.Request().Context(), req.ToModel()); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, nil)
}

func (h *VocabularyHandler) GetVocabularies(c echo.Context) error {
	var req dto.CreateVocabularyInput

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := h.vocabSvc.GetListVocab(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, nil)
}
