package services

import (
	"go-parent/repositories"
)

type StudentService interface {
	GetStudents() ([]repositories.student, error)
}

type studentService struct {
	repo repositories.StudentRepository
}

func NewStudentService(repo repositories.StudentRepository) StudentService {
	return &studentService{
		repo: repo,
	}
}

func (s *studentService) GetStudents() ([]repositories.Student, error) {
	students, err := s.repo.GetStudents()
	if err != nil {
		return nil, err
	}
	return students, nil
}

// 其他业务逻辑方法...
