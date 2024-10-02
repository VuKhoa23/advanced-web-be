package v1

import (
	"github.com/VuKhoa23/advanced-web-be/internal/service"
	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	studentService service.StudentService
}

func NewStudentHandler(studentService service.StudentService) *StudentHandler {
	return &StudentHandler{studentService: studentService}
}

func (handler *StudentHandler) GetAll(c *gin.Context) {
	students := handler.studentService.GetAllStudent(c.Request.Context())
	c.JSON(200, students)
}
