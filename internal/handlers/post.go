package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Recrusion/blog-api/internal/domain"
	"github.com/labstack/echo/v4"
)

func (h *Handlers) CreatePost(c echo.Context) error {
	var post domain.Post
	if err := c.Bind(&post); err != nil {
		c.Logger().Error(fmt.Errorf("Invalid request: %v", err))
		return c.JSON(400, map[string]string{"error": "Invalid request"})
	}

	err := h.handler.CreatePost(&post)
	if err != nil {
		c.Logger().Error(fmt.Errorf("Failed to create post: %v", err))
		return c.JSON(400, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, post)
}

func (h *Handlers) GetPost(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.Logger().Error(fmt.Errorf("Invalid request: %v", err))
		return c.JSON(400, map[string]string{"error": "Invalid request"})
	}

	post, err := h.handler.GetPost(id)
	if err != nil {
		c.Logger().Error(fmt.Errorf("Failed get post: %v", err))
		return c.JSON(400, map[string]string{"error": "Invalid request"})
	}

	return c.JSON(http.StatusOK, post)
}
