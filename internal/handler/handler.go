package handler

import (
	"net/http"
	"strconv"

	"github.com/Reteger/shiki/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc service.Service
}

func NewHandler(svc service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) GetOngoings(c *gin.Context) {
	daysStr := c.Param("days")
	days, _ := strconv.Atoi(daysStr)
	resp, err := h.svc.GetOngoings(days)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}
