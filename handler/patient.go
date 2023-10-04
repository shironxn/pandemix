package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"pandemix/entity/domain"
	"pandemix/entity/web"
	"pandemix/model"
	"pandemix/util"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var listPatient = &model.ListPatient

func CreatePatient(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.LogError(w, "Failed to read request body", http.StatusBadRequest, err.Error())
		return
	}

	patient := domain.PatientEntity{
		Id:       uint(uuid.New().ID()),
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	err = json.Unmarshal(data, &patient)
	if err != nil {
		util.LogError(w, "Failed to parse request body", http.StatusBadRequest, err.Error())
		return
	}

	validate := util.Validation(w, patient)
	if len(validate) != 0 {
		util.LogError(w, "Validation error", http.StatusBadRequest, validate)
		return
	}

	*listPatient = append(*listPatient, patient)

	responseData := web.ResponsePatient{
		Name:     patient.Name,
		Age:      patient.Age,
		Gender:   patient.Gender,
		Status:   patient.Status,
		CreateAt: patient.CreateAt,
		UpdateAt: patient.UpdateAt,
	}

	res := web.ResponseSuccess{
		Message: "Success create patient data",
		Data:    responseData,
	}

	util.JsonResponse(w, res, http.StatusOK)
}

func GetPatient(w http.ResponseWriter, r *http.Request) {
	var patients []domain.PatientEntity

	for _, data := range *listPatient {
		patient := domain.PatientEntity{
			Id:       data.Id,
			Name:     data.Name,
			Age:      data.Age,
			Gender:   data.Gender,
			Status:   data.Status,
			CreateAt: data.CreateAt,
			UpdateAt: data.UpdateAt,
		}
		patients = append(patients, patient)
	}

	res := web.ResponseSuccess{
		Message: "Success get data patient",
		Data:    patients,
	}

	util.JsonResponse(w, res, http.StatusOK)
}

func UpdatePatient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	patientID, _ := strconv.Atoi(id)

	updatedPatient := domain.PatientEntity{
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	var found bool
	for i, patient := range *listPatient {
		if uint(patientID) == patient.Id {
			data, err := ioutil.ReadAll(r.Body)
			if err != nil {
				util.LogError(w, "Failed to read request body", http.StatusBadRequest, err.Error())
				return
			}

			err = json.Unmarshal(data, &updatedPatient)
			if err != nil {
				util.LogError(w, "Failed to parse request body", http.StatusBadRequest, err.Error())
				return
			}

			updatedPatient = domain.PatientEntity{
				Id:       patient.Id,
				Name:     updatedPatient.Name,
				Age:      updatedPatient.Age,
				Gender:   updatedPatient.Gender,
				Status:   updatedPatient.Status,
				CreateAt: patient.CreateAt,
				UpdateAt: updatedPatient.UpdateAt,
			}

			(*listPatient)[i] = updatedPatient

			found = true
		}
	}

	if !found {
		util.LogError(w, "Patient not found", http.StatusBadRequest, nil)
		return
	}

	res := web.ResponseSuccess{
		Message: "Success update patient data",
		Data: web.ResponsePatient{
			Name:     updatedPatient.Name,
			Age:      updatedPatient.Age,
			Gender:   updatedPatient.Gender,
			Status:   updatedPatient.Status,
			CreateAt: updatedPatient.CreateAt,
			UpdateAt: updatedPatient.UpdateAt,
		},
	}

	util.JsonResponse(w, res, http.StatusOK)
}

func DeletePatient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	patientID, _ := strconv.Atoi(id)

	var found bool
	for i, patient := range *listPatient {
		if uint(patientID) == patient.Id {
			*listPatient = append((*listPatient)[:i], (*listPatient)[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		util.LogError(w, "Patient not found", http.StatusNotFound, nil)
		return
	}

	res := web.ResponseSuccess{
		Message: "Success delete patient data",
		Data:    nil,
	}

	util.JsonResponse(w, res, http.StatusOK)
}

func GetPatientByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	patientID, _ := strconv.Atoi(id)

	var found bool
	var foundPatient domain.PatientEntity

	for _, data := range *listPatient {
		if uint(patientID) == data.Id {
			foundPatient = data
			found = true
			break
		}
	}

	if found {
		res := web.ResponseSuccess{
			Message: "Success get patient data by id",
			Data: web.ResponsePatient{
				Name:     foundPatient.Name,
				Age:      foundPatient.Age,
				Gender:   foundPatient.Gender,
				Status:   foundPatient.Status,
				CreateAt: foundPatient.CreateAt,
				UpdateAt: foundPatient.UpdateAt,
			},
		}

		util.JsonResponse(w, res, http.StatusOK)

	} else {
		util.LogError(w, "Patient not found", http.StatusNotFound, nil)
	}
}
