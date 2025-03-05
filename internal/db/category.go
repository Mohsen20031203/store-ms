package db

import "hello/internal/db/models"

func (s *Storege) CreateCategory(category *models.Category) error {
	err := s.DB.Create(category).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Storege) GetCategory(id int64) (models.Category, error) {
	var category models.Category

	err := s.DB.First(&category, id).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (s *Storege) ListCategory() ([]models.Category, error) {
	var category []models.Category

	err := s.DB.Order("id desc").Find(&category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (s *Storege) UpdateCategory(Name string) error {
	var category models.Category

	err := s.DB.First(&category, Name).Error
	if err != nil {
		return err
	}
	return nil
}
