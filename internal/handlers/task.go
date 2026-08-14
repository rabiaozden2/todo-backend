package handlers

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rabiaozden/todo-backend/internal/models"
	"gorm.io/gorm"
)

type TaskHandler struct{ DB *gorm.DB }
type taskRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Category    string     `json:"category"`
	DueDate     *time.Time `json:"due_date"`
}

func userID(c *gin.Context) uuid.UUID { return c.MustGet("userID").(uuid.UUID) }
func (h *TaskHandler) List(c *gin.Context) {
	var tasks []models.Task
	if err := h.DB.Where("user_id = ?", userID(c)).Order("created_at DESC").Find(&tasks).Error; err != nil {
		c.JSON(500, gin.H{"error": "gorevler getirilemedi"})
		return
	}
	c.JSON(200, tasks)
}
func (h *TaskHandler) Create(c *gin.Context) {
	var req taskRequest
	if c.ShouldBindJSON(&req) != nil || strings.TrimSpace(req.Title) == "" {
		c.JSON(400, gin.H{"error": "baslik zorunludur"})
		return
	}
	task := models.Task{UserID: userID(c), Title: strings.TrimSpace(req.Title), Description: strings.TrimSpace(req.Description), Status: "todo", Category: req.Category, DueDate: req.DueDate}
	if err := h.DB.Create(&task).Error; err != nil {
		c.JSON(500, gin.H{"error": "gorev olusturulamadi"})
		return
	}
	c.JSON(201, task)
}
func (h *TaskHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "gecersiz gorev id"})
		return
	}
	var task models.Task
	if h.DB.Where("id = ? AND user_id = ?", id, userID(c)).First(&task).Error != nil {
		c.JSON(404, gin.H{"error": "gorev bulunamadi"})
		return
	}
	var req taskRequest
	if c.ShouldBindJSON(&req) != nil {
		c.JSON(400, gin.H{"error": "gecersiz istek"})
		return
	}
	if req.Title != "" {
		task.Title = strings.TrimSpace(req.Title)
	}
	if req.Description != "" {
		task.Description = strings.TrimSpace(req.Description)
	}
	if req.Status != "" {
		task.Status = req.Status
	}
	if req.Category != "" {
		task.Category = req.Category
	}
	if req.DueDate != nil {
		task.DueDate = req.DueDate
	}
	h.DB.Save(&task)
	c.JSON(200, task)
}
func (h *TaskHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "gecersiz gorev id"})
		return
	}
	result := h.DB.Where("id = ? AND user_id = ?", id, userID(c)).Delete(&models.Task{})
	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "gorev bulunamadi"})
		return
	}
	c.Status(http.StatusNoContent)
}
