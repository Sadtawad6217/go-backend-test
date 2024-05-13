package repository

import (
	"go-backend-test/pkg/common/database"
	"go-backend-test/pkg/core"
)

type ExamRepository struct{}

func NewExamRepository() *ExamRepository {
	return &ExamRepository{}
}

func (r *ExamRepository) GetExamByID(id string) (*core.Exam, error) {
	var exam core.Exam
	if err := database.DB.First(&exam, id).Error; err != nil {
		return nil, err
	}
	return &exam, nil
}

func (r *ExamRepository) GetAllExams() ([]*core.Exam, error) {
	var exams []*core.Exam
	if err := database.DB.Find(&exams).Error; err != nil {
		return nil, err
	}
	return exams, nil
}