package handler

import (
	"encoding/json"
	"net/http"
	"pandemix/entity/web"
	"pandemix/model"
	"time"
)

func GetCovidData(w http.ResponseWriter, r *http.Request) {
	listCovid := web.ResponseCovid{
		Date: time.Now().Format("2006-01-02 15:04:05"),
	}

	for _, data := range model.ListPatient {
		switch data.Status {
		case "dead":
			listCovid.Dead++

		case "positive":
			listCovid.Positive++

		case "recovered":
			listCovid.Recovered++
		}

		listCovid.TotalCase++
	}

	res := web.ResponseSuccess{
		Message: "Success get covid data",
		Data:    listCovid,
	}

	response, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
