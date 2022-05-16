package api

import (
	"encoding/json"
	"github.com/enkhalifapro/phone-validator/phones"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CountryHandler struct {
	phoneManager PhoneManager
}

type PhoneManager interface {
	GetCountries() ([]*phones.Country, error)
}

func NewCountryHandler(mgr PhoneManager) *CountryHandler {
	return &CountryHandler{phoneManager: mgr}
}

func (c *CountryHandler) List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	countries, err := c.phoneManager.GetCountries()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	j, err := json.Marshal(countries)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	renderJson(w, j, http.StatusOK)
}
