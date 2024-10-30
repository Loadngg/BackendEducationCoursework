package database

import (
	"backend/internal/backend/models"
	"gopkg.in/reform.v1"
)

type LecturesPostgres struct {
	db *reform.DB
}

func NewLecturesPostgres(db *reform.DB) *LecturesPostgres {
	return &LecturesPostgres{db: db}
}

func (a *LecturesPostgres) GetAll() ([]models.Lectures, error) {
	var lecturesList []models.Lectures

	lectures, err := a.db.SelectAllFrom(models.LecturesTable, "")
	if err != nil {
		return nil, err
	}

	for _, l := range lectures {
		item := *l.(*models.Lectures)
		lecturesList = append(lecturesList, item)
	}

	return lecturesList, nil
}

func (a *LecturesPostgres) GetById(id string) (*models.Lectures, error) {
	var lecture models.Lectures

	err := a.db.SelectOneTo(&lecture, "WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	return &lecture, nil
}

func (a *LecturesPostgres) GetAllMaterials(lectureId string) ([]models.LectureMaterials, error) {
	var materialsList []models.LectureMaterials

	materials, err := a.db.SelectAllFrom(models.LectureMaterialsTable, "WHERE lecture_id=$1", lectureId)
	if err != nil {
		return nil, err
	}

	for _, m := range materials {
		item := *m.(*models.LectureMaterials)
		materialsList = append(materialsList, item)
	}

	return materialsList, nil
}
