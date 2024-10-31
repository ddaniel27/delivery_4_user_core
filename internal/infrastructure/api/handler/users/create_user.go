package handler

import (
	"ddaniel27/usercore/internal/core/dto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) CreateUser(c *gin.Context) {
	ctx, span := h.tracer.Start(c.Request.Context(), "CreateUser")
	defer span.End()

	var body dto.CreateUserDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		h.logger.Error(fmt.Sprintf("Failed to bind JSON: %s", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.UsersService.CreateUser(ctx, &body); err != nil {
		h.logger.Error(fmt.Sprintf("Failed to create user: %s", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"user_created": false, "error": err.Error()})
		return
	}

	h.userCounter.Add(ctx, 1)
	c.JSON(http.StatusCreated, gin.H{"user_created": true})
}
