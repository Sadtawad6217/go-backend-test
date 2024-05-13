package handlers

import (
	"gobackend/pkg/core"

	"github.com/gofiber/fiber/v2"
)

type ExamHandler struct {
	ExamService core.ExamService
}

func NewExamHandler(examService core.ExamService) *ExamHandler {
	return &ExamHandler{ExamService: examService}
}

func (h *ExamHandler) GetExamByID(c *fiber.Ctx) error {
	// Implement logic to retrieve exam data by ID
}

func (h *ExamHandler) GetAllExams(c *fiber.Ctx) error {
	// Implement logic to retrieve all exams
}
