package service

import (
	"context"
	"github.com/VuKhoa23/advanced-web-be/internal/domain/model"
)

type StudentService interface {
	GetAllStudent(ctx context.Context) []model.Student
}
