package handlers

import (
	"github.com/gin-gonic/gin"
	"go-parent/services"
	"net/http"
	_ "strconv"
)

type StudentHandler struct {
	service *services.StudentService
}

func NewStudentHandler(service *services.StudentService) *StudentHandler {
	return &StudentHandler{service}
}

func (h *StudentHandler) GetStudents(c *gin.Context) {
	students, err := h.service.GetStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, students)
}

// 其他处理函数...
