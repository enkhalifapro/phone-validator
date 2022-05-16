package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Provider struct {
	db DBConnector
}

type Customer struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type DBConnector interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

func New(db DBConnector) *Provider {
	return &Provider{db: db}
}

/*func (p *Provider) Query(sql string, dest ...interface{}) error {
	rows, err := p.db.Query(sql)
	if err != nil {
		return err
	}

	customers := make([]*Customer, 0)
	for rows.Next() {
		cust := &Customer{}
		err = rows.Scan(&cust.ID, &cust.Name, &cust.Phone)
		if err != nil {
			return err
		}

		rows.

		customers = append(customers, cust)
	}

	return nil
}*/

/*func (p *Provider) Test() {
	db, err := sql.Open("sqlite3", "./sample.db")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select * from customer")
	if err != nil {
		panic(err)
	}

	customers := make([]*Customer, 0)
	for rows.Next() {
		cust := &Customer{}
		err = rows.Scan(&cust.ID, &cust.Name, &cust.Phone)
		if err != nil {
			panic(err)
		}

		customers = append(customers, cust)
	}

	for _, c := range customers {
		fmt.Printf("ID: %v Name: %s Phone: %s\n", c.ID, c.Name, c.Phone)
	}
}
*/