package api

import (
	"encoding/json"
	"fmt"
	"github.com/enkhalifapro/phone-validator/phones"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type PhoneHandler struct {
	phoneManager PhoneManager
}

type PhoneManager interface {
	GetPhones(limit int, skip int) ([]*phones.Phone, error)
	GetPhonesByCountry(limit int, skip int, countryName string) ([]*phones.Phone, error)
	GetCountries() map[string]phones.Country
}

func NewPhoneHandler(mgr PhoneManager) *PhoneHandler {
	return &PhoneHandler{phoneManager: mgr}
}

func (c *PhoneHandler) ListPhones(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pageNumber, err := strconv.Atoi(r.URL.Query().Get("pageNumber"))
	if err != nil {
		pageNumber = 1
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil {
		pageSize = 10
	}

	phons, err := c.phoneManager.GetPhones(pageSize, (pageNumber-1)*pageSize)
	if err != nil {
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

func (c *PhoneHandler) GetPhonesByCountry(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	pageNumber, err := strconv.Atoi(r.URL.Query().Get("pageNumber"))
	if err != nil {
		pageNumber = 1
	}
	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil {
		pageSize = 10
	}
	countryName := params.ByName("countryName")
	phons, err := c.phoneManager.GetPhonesByCountry(pageSize, (pageNumber-1)*pageSize, countryName)
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
