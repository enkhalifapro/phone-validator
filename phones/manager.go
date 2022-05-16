package phones

import (
	"database/sql"
	"regexp"
)

type DBConnector interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type Manager struct {
	db DBConnector
}

func New(db DBConnector) *Manager {
	return &Manager{db: db}
}

func (m *Manager) GetCountries() ([]*Country, error) {
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
		if foundCountries[c.Code] == false {
			countries = append(countries, c)
			foundCountries[c.Code] = true
		}
	}

	return countries, nil
}

func getPhoneCountry(phone string) *Country {
	countries := []*Country{}
	countries = append(countries, &Country{Code: "+237", Name: "Cameroon", RegExp: `\(237\)\ ?[2368]\d{7,8}$`})
	countries = append(countries, &Country{Code: "+251", Name: "Ethiopia", RegExp: `\(251\)\ ?[1-59]\d{8}$`})
	countries = append(countries, &Country{Code: "+212", Name: "Morocco", RegExp: `\(212\)\ ?[5-9]\d{8}$`})
	countries = append(countries, &Country{Code: "+258", Name: "Mozambique", RegExp: `\(258\)\ ?[28]\d{7,8}$`})
	countries = append(countries, &Country{Code: "+256", Name: "Uganda", RegExp: `\(256\)\ ?\d{9}$`})

	for _, c := range countries {
		match, _ := regexp.MatchString(c.RegExp, phone)
		if match {
			return c
		}
	}

	return &Country{
		Code: "",
		Name: "UNKNOWN",
	}

}
