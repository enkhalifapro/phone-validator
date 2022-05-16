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
	countriesMap["+237"] = Country{Code: "+237", Name: "Cameroon", RegExp: `\(237\)\ ?[2368]\d{7,8}$`}
	countriesMap["+251"] = Country{Code: "+251", Name: "Ethiopia", RegExp: `\(251\)\ ?[1-59]\d{8}$`}
	countriesMap["+212"] = Country{Code: "+212", Name: "Morocco", RegExp: `\(212\)\ ?[5-9]\d{8}$`}
	countriesMap["+258"] = Country{Code: "+258", Name: "Mozambique", RegExp: `\(258\)\ ?[28]\d{7,8}$`}
	countriesMap["+256"] = Country{Code: "+256", Name: "Uganda", RegExp: `\(256\)\ ?\d{9}$`}

	return countriesMap
}

func (m *Manager) GetPhones() ([]*Phone, error) {
	rows, err := m.db.Query("SELECT id, name, phone FROM customer")
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
		c := m.getPhoneCountry(phone.Phone)
		phone.CountryCode = c.Code
		phone.CountryName = c.Name

		phones = append(phones, phone)
	}

	return phones, nil
}

func (m *Manager) GetPhonesByCountry(countryCode string) ([]*Phone, error) {
	country, ok := m.countriesMap[countryCode]
	if !ok {
		return []*Phone{}, errors.New("unknown country")
	}
	x:=country.RegExp
	fmt.Println(x)
	rows, err := m.db.Query("SELECT id, name, phone FROM customer WHERE phone REGEXP ?", country.RegExp)
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
		c := m.getPhoneCountry(phone.Phone)
		phone.CountryCode = c.Code
		phone.CountryName = c.Name

		phones = append(phones, phone)
	}

	return phones, nil
}

func (m *Manager) GetCountries() map[string]Country {
	return m.countriesMap
}

/*
func (m *Manager) GetCountriesFromDB() ([]*Country, error) {
	rows, err := m.db.Query("select phone from customer")
	if err != nil {
		return nil, err
	}

	countries := make([]*Country, 0)
	foundCountries := make(map[string]bool)
	for rows.Next() {
		var phone string
		err = rows.Scan(&phone)
		if err != nil {
			return nil, err
		}

		c := getPhoneCountry(phone)
		if _, ok := foundCountries[c.Code]; !ok {
			countries = append(countries, c)
			foundCountries[c.Code] = true
		}
	}

	return countries, nil
}
*/
func (m *Manager) getPhoneCountry(phone string) *Country {
	for _, c := range m.countriesMap {
		match, _ := regexp.MatchString(c.RegExp, phone)
		if match {
			return &c
		}
	}

	return &Country{
		Code: "",
		Name: "UNKNOWN",
	}

}
