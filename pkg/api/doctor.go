package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"golang-training-Hospital/pkg/data"
	"log"
	"net/http"
	"strconv"
)

//DoctorApi - struct that's represent api layer between data layer and user layer
type DoctorApi struct {
	data *data.DoctorData
}

//ServeDoctorResources - create a system of routing in web application
func ServeDoctorResources(r *mux.Router, data data.DoctorData) {
	api := &DoctorApi{data: &data}
	r.HandleFunc("/doctors", api.getAllDoctors).Methods("GET")
	r.HandleFunc("/create", api.createDoctor).Methods("POST")
	r.HandleFunc("/update", api.updateDoctor).Methods("PUT")
	r.HandleFunc("/delete", api.deleteDoctor).Methods("DELETE")
}

//getAllDoctors - get all records from doctors table
//writer - perform response object from the server.The type is http.ResponseWriter
//request - perform request from the user.The type is *http.Request
func (a DoctorApi) getAllDoctors(writer http.ResponseWriter, request *http.Request) {
	doctors, err := a.data.ReadAll()
	if err != nil {
		_, err := writer.Write([]byte("got an error when tried to get doctors"))
		if err != nil {
			log.Println(err)
		}
	}
	err = json.NewEncoder(writer).Encode(doctors)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

//createDoctor - create doctor record in the doctors table
//writer - perform response object from the server.The type is http.ResponseWriter
//request - perform request from the user.The type is *http.Request
func (a DoctorApi) createDoctor(writer http.ResponseWriter, request *http.Request) {
	doctor := new(data.Doctors)
	err := json.NewDecoder(request.Body).Decode(&doctor)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if doctor == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = a.data.CreateDoctor(*doctor)
	if err != nil {
		log.Println("doctor hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

//updateDoctor - updateDoctor doctor record in the doctors table
//writer - perform response object from the server.The type is http.ResponseWriter
//request - perform request from the user.The type is *http.Request
func (a DoctorApi) updateDoctor(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil {
		log.Printf("failed format id: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	doctor := new(data.Doctors)
	err = json.NewDecoder(request.Body).Decode(&doctor)
	if doctor == nil {
		log.Printf("failed empty JSON\n")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = a.data.UpdateDoctor(id, *doctor)
	if err != nil {
		log.Println("doctor hasn't been updated")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}

//deleteDoctor - delete doctor record in the doctors table
//writer - perform response object from the server.The type is http.ResponseWriter
//request - perform request from the user.The type is *http.Request
func (a DoctorApi) deleteDoctor(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil {
		log.Printf("failed format id: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = a.data.DeleteDoctor(id)
	if err != nil {
		log.Println("doctor hasn't been deleted")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
}
