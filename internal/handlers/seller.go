package handlers

import (
	"Webserver/internal/services"

	"github.com/gin-gonic/gin"
)

type SellerHandler struct {
	srv *services.SellerService
}

// NewSellerHandler Constructor
func NewSellerHandler(srv *services.SellerService) *SellerHandler {
	return &SellerHandler{srv: srv}
}

func (h *SellerHandler) Create(c *gin.Context) {
	var body struct {
		Name string `json:"name"`
	}

	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"}) // ← Добавить обработку ошибки
		return
	}

	seller, err := h.srv.Create(c, body.Name) // ← Исправить на err
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()}) // ← Обработать ошибку сервиса
		return
	}

	c.JSON(201, gin.H{"data": seller})
}

func (h *SellerHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	seller, err := h.srv.GetByID(c, id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Seller not found"})
		return
	}
	c.JSON(200, gin.H{"data": seller})
}

func (h *SellerHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Name string `json:"name"`
	}

	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	seller, err := h.srv.Update(c, id, body.Name)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": seller})
}

func (h *SellerHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	var err = h.srv.Delete(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(204, nil) // 204 No Content
}

func (h *SellerHandler) GetAll(context *gin.Context) {

}
