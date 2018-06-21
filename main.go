package main

import (
	"fmt"
	"time"
	"github.com/rachmathidayatdev/go_crud_mongodb/config"
	"github.com/rachmathidayatdev/go_crud_mongodb/src/modules/profile/model"
	"github.com/rachmathidayatdev/go_crud_mongodb/src/modules/profile/repository"
	"github.com/gorilla/mux"
	"encoding/json"
	"net/http"
	"log"
)

//main
func main() {
	fmt.Println("Rest API GOLANG")

	router := mux.NewRouter()
	router.HandleFunc("/get-all-profile", getAllProfiles).Methods("GET")
	router.HandleFunc("/get-profile/{id}", getProfileById).Methods("GET")
	router.HandleFunc("/save-profile", saveProfile).Methods("POST")
	router.HandleFunc("/update-profile", updateProfile).Methods("PUT")
	router.HandleFunc("/delete-profile/{id}", deleteProfile).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8300", router))
}

//saveProfile
func saveProfile(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMongoDB()

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	profileRepository := repository.NewProfileRepositoryMongo(db, "profile")

	decoder := json.NewDecoder(r.Body)

	var p model.Profile
	var paramBody model.Profile
	
	decoder.Decode(&paramBody)

	p.ID = paramBody.ID
	p.FirstName = paramBody.FirstName
	p.LastName = paramBody.LastName
	p.Email = paramBody.Email
	p.Password = paramBody.Password
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err = profileRepository.Save(&p)

	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode("Berhasil save...")
	}
}

//updateProfile
func updateProfile(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMongoDB()

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	profileRepository := repository.NewProfileRepositoryMongo(db, "profile")

	decoder := json.NewDecoder(r.Body)

	var p model.Profile
	var paramBody model.Profile
	
	decoder.Decode(&paramBody)

	p.ID = paramBody.ID
	p.FirstName = paramBody.FirstName
	p.LastName = paramBody.LastName
	p.Email = paramBody.Email
	p.Password = paramBody.Password
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err = profileRepository.Update(paramBody.ID, &p)

	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode("Berhasil update...")
	}
}

//deleteProfile
func deleteProfile(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMongoDB()

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	profileRepository := repository.NewProfileRepositoryMongo(db, "profile")

	params := mux.Vars(r)
	
	err = profileRepository.Delete(params["id"])

	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode("Berhasil delete...")
	}
}

//getProfileById
func getProfileById(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMongoDB()

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	profileRepository := repository.NewProfileRepositoryMongo(db, "profile")

	params := mux.Vars(r)

	id := params["id"]

	profile, err := profileRepository.FindByID(id)

	if err != nil {
		json.NewEncoder(w).Encode(err)
	} 
	
	json.NewEncoder(w).Encode(profile)
}

//getAllProfiles
func getAllProfiles(w http.ResponseWriter, r *http.Request) {
	db, err := config.GetMongoDB()

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	profileRepository := repository.NewProfileRepositoryMongo(db, "profile")

	profiles, err := profileRepository.FindAll()

	if err != nil {
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(profiles)
}