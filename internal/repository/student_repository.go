package repository

import (
	"context"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/entity"
)

type StudentRepository interface {
	GetAllStudentQuery(c context.Context) []entity.Student
}
