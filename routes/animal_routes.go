package routes

import (
	"alexdev2001/zoo_database_crud_api/handler"
	"alexdev2001/zoo_database_crud_api/services"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB) (r *mux.Router) {
	router := mux.NewRouter()

	// initialize handler and service instances
	animalService := services.NewAnimalService(db)
	animalHandler := handler.NewAnimalHanlder(animalService)

	router.HandleFunc("/animals/create", animalHandler.CreateAnimal).Methods("POST")
	router.HandleFunc("/animals", animalHandler.GetAllAnimals).Methods("GET")
	router.HandleFunc("/animals/{id}", animalHandler.GetAnimalById).Methods("GET")
	router.HandleFunc("/animals/update/{id}", animalHandler.UpdateAnimal).Methods("PUT")
	router.HandleFunc("/animals/delete/{id}", animalHandler.DeleteAnimal).Methods("DELETE")

	return router

}
