package service

import (
	"backend/internal/backend/database"
	"backend/internal/backend/models"
)

type LecturesService struct {
	db database.Lectures
}

func NewLecturesService(db database.Lectures) *LecturesService {
	return &LecturesService{db: db}
}

func (s *LecturesService) GetAll() ([]models.Lectures, error) {
	return s.db.GetAll()
}

func (s *LecturesService) GetById(id string) (*models.Lectures, error) {
	return s.db.GetById(id)
}

func (s *LecturesService) GetAllMaterials(id string) ([]models.LectureMaterials, error) {
	return s.db.GetAllMaterials(id)
}
