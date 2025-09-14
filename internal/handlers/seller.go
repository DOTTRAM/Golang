package handlers

import (
	"Webserver/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SellerHandler struct {
	srv *services.SellerService
}

func NewSellerHandler(srv *services.SellerService) *SellerHandler {
	return &SellerHandler{srv: srv}
}

func (h *SellerHandler) Create(c *gin.Context) {
	var body struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	seller, err := h.srv.Create(c.Request.Context(), body.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": seller})
}

func (h *SellerHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	seller, err := h.srv.GetByID(c.Request.Context(), id)
	if err != nil {
		if err.Error() == "seller not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": seller})
}

func (h *SellerHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	seller, err := h.srv.Update(c.Request.Context(), id, body.Name)
	if err != nil {
		if err.Error() == "seller not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": seller})
}

func (h *SellerHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.srv.Delete(c.Request.Context(), id)
	if err != nil {
		if err.Error() == "seller not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *SellerHandler) GetAll(c *gin.Context) {
	sellers, err := h.srv.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": sellers})
}
