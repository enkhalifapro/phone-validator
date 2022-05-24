package phones

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
)

type DBConnector interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type Manager struct {
	db           DBConnector
	countriesMap map[string]Country
}

func New(db DBConnector) *Manager {
	return &Manager{
		db:           db,
		countriesMap: initCountries(),
	}
}

func initCountries() map[string]Country {
	countriesMap := make(map[string]Country)
	countriesMap["Cameroon"] = Country{Code: "+237", Name: "Cameroon", RegExp: `\(237\)\ ?[2368]\d{7,8}$`}
	countriesMap["Ethiopia"] = Country{Code: "+251", Name: "Ethiopia", RegExp: `\(251\)\ ?[1-59]\d{8}$`}
	countriesMap["Morocco"] = Country{Code: "+212", Name: "Morocco", RegExp: `\(212\)\ ?[5-9]\d{8}$`}
	countriesMap["Mozambique"] = Country{Code: "+258", Name: "Mozambique", RegExp: `\(258\)\ ?[28]\d{7,8}$`}
	countriesMap["Uganda"] = Country{Code: "+256", Name: "Uganda", RegExp: `\(256\)\ ?\d{9}$`}

	return countriesMap
}

func (m *Manager) GetPhones(limit int, skip int) ([]*Phone, error) {
	fmt.Println("deeeeee")
	fmt.Println(limit)
	fmt.Println(skip)
	rows, err := m.db.Query("SELECT id, name, phone FROM customer LIMIT ? OFFSET ?", limit, skip)
	if err != nil {
		return nil, err
	}

	phones := make([]*Phone, 0)
	for rows.Next() {
		phone := &Phone{}
		err = rows.Scan(&phone.ID, &phone.Name, &phone.Phone)
		if err != nil {
			return nil, err
		}
		m.enrichCountry(phone)
		phones = append(phones, phone)
	}
	return phones, nil
}

func (m *Manager) GetPhonesByCountry(limit int, skip int, countryName string) ([]*Phone, error) {
	country, ok := m.countriesMap[countryName]
	if !ok {
		return []*Phone{}, errors.New("unknown country")
	}
	rows, err := m.db.Query("SELECT id, name, phone FROM customer WHERE phone REGEXP ? LIMIT ? OFFSET ?", country.RegExp, limit, skip)
	if err != nil {
		return nil, err
	}

	phones := make([]*Phone, 0)
	for rows.Next() {
		phone := &Phone{}
		err = rows.Scan(&phone.ID, &phone.Name, &phone.Phone)
		if err != nil {
			return nil, err
		}
		m.enrichCountry(phone)
		phones = append(phones, phone)
	}
	return phones, nil
}

func (m *Manager) GetCountries() map[string]Country {
	return m.countriesMap
}

func (m *Manager) enrichCountry(phone *Phone) {
	for _, c := range m.countriesMap {
		match, _ := regexp.MatchString(c.RegExp, phone.Phone)
		if match {
			phone.CountryName = c.Name
			phone.CountryCode = c.Code
			phone.State = "valid"
			return
		}
	}
	phone.CountryName = "UNKNOWN"
	phone.CountryCode = "UNKNOWN"
	phone.State = "not valid"
}
