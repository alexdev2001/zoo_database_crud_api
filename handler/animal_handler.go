package handler

import (
	"alexdev2001/zoo_database_crud_api/models"
	"alexdev2001/zoo_database_crud_api/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type AnimalHanlder struct {
	service *services.AnimalService
}

func NewAnimalHanlder(service *services.AnimalService) *AnimalHanlder {
	return &AnimalHanlder{service: service}
}

// Handler function to get all students
func (h *AnimalHanlder) GetAllAnimals(w http.ResponseWriter, r *http.Request) {
	students, err := h.service.GetAnimals()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(students)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)

}

// Handler function to capture students by id
func (h *AnimalHanlder) GetAnimalById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	student, err := h.service.GetAnimalByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(student)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusAccepted)

}

// Handler function to create an animal
func (h *AnimalHanlder) CreateAnimal(w http.ResponseWriter, r *http.Request) {
	var animal models.Animal
	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	createdAnimal, err := h.service.CreateAnimal(animal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(createdAnimal)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusCreated)

}

// Hanlder function to update an animal
func (h *AnimalHanlder) UpdateAnimal(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	var animal models.Animal
	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	updatedAnimal, err := h.service.UpdateAnimal(id, animal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(updatedAnimal)
	w.WriteHeader(http.StatusOK)
}

// Handler function to delete an animal
func (h *AnimalHanlder) DeleteAnimal(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	deletedAnimal, err := h.service.DeleteAnimal(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(deletedAnimal)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
