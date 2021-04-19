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
	DoctorId:         1,
	FirstName:        "Daniil",
	LastName:         "Domaskanov",
	DateOfBirth:      "2000-09-12",
	Salary:           "12542.12",
	CurrentBusyState: false,
	GenderId:         2,
	SpecialityId:     1,
}

var updatedValues = Doctors{
	DoctorId:  1,
	FirstName: "Vasya",
}

func TestDoctorData_ReadAll(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewDoctor(gormDb)
	rows := sqlmock.NewRows([]string{"doctor_id", "first_name", "last_name", "date_of_birth", "salary", "current_busy_state", "gender_id", "speciality_id"}).
		AddRow(testDoctor.DoctorId, testDoctor.FirstName, testDoctor.LastName, testDoctor.DateOfBirth, testDoctor.Salary,
			testDoctor.CurrentBusyState, testDoctor.GenderId, testDoctor.SpecialityId)
	mock.ExpectQuery(selectAllRows).WillReturnRows(rows)
	products, err := data.ReadAll()
	assert.NoError(err)
	assert.NotEmpty(products)
	assert.Equal(products[0], testDoctor)
	assert.Len(products, 1)
}

func TestDoctorData_DeleteDoctor(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewDoctor(gormDb)
	mock.ExpectBegin()
	mock.ExpectExec(deleteDoctor).
		WithArgs(testDoctor.DoctorId).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := data.DeleteDoctor(testDoctor.DoctorId)
	assert.NoError(err)
}

func TestDoctorData_ReadAllErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewDoctor(gormDb)
	mock.ExpectQuery(selectAllRows).WillReturnError(errors.New("something went wrong..."))
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
	mock.ExpectExec(insertDoctor).WithArgs(testDoctor.FirstName, testDoctor.LastName, testDoctor.DateOfBirth, testDoctor.Salary,
		testDoctor.CurrentBusyState, testDoctor.GenderId, testDoctor.SpecialityId, testDoctor.DoctorId).
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

func TestDoctorData_UpdateDoctorErr(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()
	gormDb := NewGorm(db)
	data := NewDoctor(gormDb)
	mock.ExpectBegin()
	mock.ExpectExec(updateDoctor).
		WithArgs(testDoctor.DoctorId).
		WillReturnError(errors.New("something went wrong..."))
	mock.ExpectCommit()
	_, err := data.UpdateDoctor(testDoctor.DoctorId, updatedValues)
	assert.Error(err)
}
