package main

import (
	"github.com/Reteger/shiki/internal/handler"
	"github.com/Reteger/shiki/internal/repository"
	"github.com/Reteger/shiki/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := repository.NewRepository()
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	r := gin.Default()
	r.GET("/api/ongoings/:days", h.GetOngoings)
	r.Run(":8080")
}
