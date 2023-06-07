package repositories

import (
	"github.com/jmoiron/sqlx"
	"go-parent/models"
)

type StudentRepository interface {
	GetStudents() ([]models.Student, error)
}

type studentRepository struct {
	db *sqlx.DB
}

func NewStudentRepository(db *sqlx.DB) StudentRepository {
	return &studentRepository{
		db: db,
	}
}

func (r *studentRepository) GetStudents() ([]models.Student, error) {
	var students []models.Student
	err := r.db.Select(&students, "SELECT * FROM students")
	if err != nil {
		return nil, err
	}
	return students, nil
}

// 其他数据库操作方法...
