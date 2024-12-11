package services

import (
	"alexdev2001/zoo_database_crud_api/models"
	"errors"
	"gorm.io/gorm"
)

type AnimalService struct {
	db *gorm.DB
}

func NewAnimalService(db *gorm.DB) *AnimalService {
	return &AnimalService{db: db}
}

// function to get all animals
func (h *AnimalService) GetAnimals() ([]*models.Animal, error) {
	var animals []*models.Animal
	result := h.db.Find(&animals)
	if result.Error != nil {
		return nil, result.Error
	}
	return animals, nil
}

// function to get an animal by id
func (h *AnimalService) GetAnimalByID(id int) (*models.Animal, error) {
	var animal models.Animal
	result := h.db.First(&animal, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &animal, nil
}

// function to create animal
func (h *AnimalService) CreateAnimal(animal models.Animal) (*models.Animal, error) {
	result := h.db.Create(&animal)
	if result.Error != nil {
		return nil, result.Error
	}

	return &animal, nil
}

// function to update animal
func (h *AnimalService) UpdateAnimal(id int, animal models.Animal) (*models.Animal, error) {
	result := h.db.Model(&animal).Where("id = ?", id).Updates(animal)
	if result.Error != nil {
		return nil, result.Error
	}
	return &animal, nil
}

// function to delete animal
func (h *AnimalService) DeleteAnimal(id int) (*models.Animal, error) {
	var animal models.Animal
	result := h.db.Delete(&animal, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("animal not found")
	}
	return &animal, nil
}
