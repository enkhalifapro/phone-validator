package api

import (
	"encoding/json"
	"fmt"
	"github.com/enkhalifapro/phone-validator/phones"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PhoneHandler struct {
	phoneManager PhoneManager
}

type PhoneManager interface {
	GetPhones() ([]*phones.Phone, error)
	GetPhonesByCountry(countryCode string) ([]*phones.Phone, error)
	GetCountries() map[string]phones.Country
}

func NewPhoneHandler(mgr PhoneManager) *PhoneHandler {
	return &PhoneHandler{phoneManager: mgr}
}

func (c *PhoneHandler) ListPhones(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	phons, err := c.phoneManager.GetPhonesByCountry("+237")
	if err != nil {
		fmt.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	j, err := json.Marshal(phons)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	renderJson(w, j, http.StatusOK)
}

func (c *PhoneHandler) ListCountries(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	countries := c.phoneManager.GetCountries()
	j, err := json.Marshal(countries)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	renderJson(w, j, http.StatusOK)
}
