package phones

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupGenerateIMemoryDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return db, err
	}
	_, err = db.Exec(`create table customer(id int, name varchar(50), phone varchar(50));`)
	if err != nil {
		return db, err
	}
	_, err = db.Exec(`insert into customer(id,name,phone) values 
                                           (1,'ayman 1','(237) 673122155'),
                                           (2,'ayman 2','(237) 695539786'),
                                           (3,'ayman 3','(258) 848826725'),
                                           (4,'ayman 4','(251) 988200000'),
                                           (5,'ayman 5','+201148860911'),
                                           (6,'ayman 6','+201148860912'),
                                           (7,'ayman 7','+201148860913'),
                                           (8,'ayman 8','+201148860914'),
                                           (9,'ayman 9','+201148860915'),
                                           (10,'ayman 10','+201148860916');`)
	if err != nil {
		return db, err
	}

	return db, nil
}
func TestGetPhones(t *testing.T) {
	// Arrange
	db, err := setupGenerateIMemoryDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Act
	mgr := New(db)
	phons, err := mgr.GetPhones(5, 1)


	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 5, len(phons))
}
