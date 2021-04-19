package data

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"testing"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatal(err)
	}
	return db, mock
}

func NewGorm(db *sql.DB) *gorm.DB {
	dialector := postgres.New(postgres.Config{
		DriverName: "postgres",
		Conn:       db,
	})
	gormDb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return gormDb
}

var testDoctor = Doctors{
	FirstName: "Daniil",
	LastName:  "Domaskanov",
}

func TestDoctorData_ReadAllErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewDoctor(gormDb)
	mock.ExpectQuery(readAllUsersQuery).WillReturnError(errors.New("something went wrong..."))
	products, err := data.ReadAll()
	assert.Error(err)
	assert.Empty(products)
}

func TestDoctorData_CreateDoctorErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewDoctor(gormDb)
	mock.ExpectBegin()
	mock.ExpectExec(insertDoctor).WithArgs(testDoctor.DoctorId, testDoctor.FirstName, testDoctor.LastName, testDoctor.DateOfBirth, testDoctor.Salary, testDoctor.CurrentBusyState,
		testDoctor.GenderId, testDoctor.SpecialityId).
		WillReturnError(errors.New("something went wrong..."))
	mock.ExpectCommit()
	id, err := data.CreateDoctor(testDoctor)
	assert.Error(err)
	assert.Equal(id, -1)
}

func TestDoctorData_DeleteDoctorErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewDoctor(gormDb)
	mock.ExpectBegin()
	mock.ExpectExec(deleteDoctor).
		WithArgs(testDoctor.DoctorId).
		WillReturnError(errors.New("something went wrong..."))
	mock.ExpectCommit()
	err := data.DeleteDoctor(testDoctor.DoctorId)
	assert.Error(err)
}
