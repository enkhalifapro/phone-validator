package api

import (
	"encoding/json"
	"github.com/enkhalifapro/phone-validator/api/mocks"
	"github.com/enkhalifapro/phone-validator/phones"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListPhones(t *testing.T) {
	// Arrange
	r := httprouter.New()
	mgrMock := &mocks.PhoneManager{}
	phonesRes := make([]*phones.Phone, 0)
	phonesRes = append(phonesRes, &phones.Phone{ID: 1,
		Name:        "EPHREM KINFE",
		CountryCode: "+251",
		CountryName: "Ethiopia",
		State:       "valid",
	})
	phonesRes = append(phonesRes, &phones.Phone{ID: 1,
		Name:        "Ayman Hassan",
		CountryCode: "+202",
		CountryName: "UNKNOWN",
		State:       "not valid",
	})
	mgrMock.On("GetPhones", 10, 0).Return(phonesRes, nil)
	phoneHandler := NewPhoneHandler(mgrMock)
	r.GET("/phones", phoneHandler.ListPhones)
	req, _ := http.NewRequest("GET", "/phones", nil)

	// Act
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	b, err := ioutil.ReadAll(rr.Result().Body)
	assert.Nil(t, err)
	res := make([]phones.Phone, 0)
	err = json.Unmarshal(b, &res)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(res))
	assert.Equal(t, "valid",res[0].State)
	assert.Equal(t, "not valid",res[1].State)
}
