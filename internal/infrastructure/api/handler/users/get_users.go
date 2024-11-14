package handler

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
)

type pathID struct {
	ID int `uri:"id" binding:"required"`
}

type queryEmail struct {
	Email string `form:"email" binding:"required"`
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	ctx, span := h.tracer.Start(c.Request.Context(), "GetUserByID")
	defer span.End()

	var p pathID
	if err := c.ShouldBindUri(&p); err != nil {
		h.logger.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UsersService.GetUserByID(ctx, p.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}

		h.logger.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	ctx, span := h.tracer.Start(c.Request.Context(), "GetUserByEmail")
	defer span.End()

	var p queryEmail
	if err := c.ShouldBindQuery(&p); err != nil {
		h.logger.Error(err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := h.UsersService.GetUserByEmail(ctx, p.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}

		h.logger.Error(err.Error())
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"user": user})
}
