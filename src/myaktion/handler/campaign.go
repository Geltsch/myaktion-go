package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/geltsch/myaktion-go/src/myaktion/model"
	"github.com/geltsch/myaktion-go/src/myaktion/service"
)

/*func CreateCampaign(w http.ResponseWriter, r *http.Request) {
var campaign model.Campaign
err := json.NewDecoder(r.Body).Decode(&campaign)
if err != nil {
	log.Printf("Can't serialize request body to campaign struct: %v", err)
	http.Error(w, err.Error(), http.StatusBadRequest)
	return
}
if err := service.CreateCampaign(&campaign); err != nil {*/

/*func CreateCampaign(w http.ResponseWriter, r *http.Request) {
	var campaign *model.Campaign
	campaign, err := getCampaign(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := service.CreateCampaign(campaign); err != nil {
		log.Printf("Error calling service CreateCampaign: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(campaign); err != nil {
		log.Printf("Failure encoding value to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}*/

/*func GetCampaigns(w http.ResponseWriter, _ *http.Request) {
	campaigns, err := service.GetCampaigns()
	if err != nil {
		log.Printf("Error calling service GetCampaigns: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(campaigns); err != nil {
		log.Printf("Failure encoding value to JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}*/

func CreateCampaign(w http.ResponseWriter, r *http.Request) {
	campaign, err := getCampaign(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err = service.CreateCampaign(campaign); err != nil {
		log.Printf("Error calling service CreateCampaign: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, campaign)
}

func GetCampaigns(w http.ResponseWriter, _ *http.Request) {
	campaigns, err := service.GetCampaigns()
	if err != nil {
		log.Printf("Error calling service GetCampaigns: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	sendJson(w, campaigns)
}

func GetCampaign(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	campaign, err := service.GetCampaign(id)
	if err != nil {
		log.Errorf("Failure retrieving campaign with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if campaign == nil {
		http.Error(w, "404 campaign not found", http.StatusNotFound)
		return
	}
	sendJson(w, campaign)
}

func UpdateCampaign(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	campaign, err := getCampaign(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	campaign, err = service.UpdateCampaign(id, campaign)
	if err != nil {
		log.Errorf("Failure updating campaign with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if campaign == nil {
		http.Error(w, "404 campaign not found", http.StatusNotFound)
		return
	}
	sendJson(w, campaign)
}

func DeleteCampaign(w http.ResponseWriter, r *http.Request) {
	id, err := getId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	campaign, err := service.DeleteCampaign(id)
	if err != nil {
		log.Errorf("Failure deleting campaign with ID %v: %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if campaign == nil {
		http.Error(w, "404 campaign not found", http.StatusNotFound)
		return
	}
	sendJson(w, result{Success: "OK"})
}

func getCampaign(r *http.Request) (*model.Campaign, error) {
	var campaign model.Campaign
	err := json.NewDecoder(r.Body).Decode(&campaign)
	if err != nil {
		log.Errorf("Can't serialize request body to campaign struct: %v", err)
		return nil, err
	}
	return &campaign, nil
}
