package data

import (
	"fmt"

	"gorm.io/gorm"
)

type Doctors struct {
	DoctorId         int    `gorm:"primaryKey; doctor_id; autoIncrement"`
	FirstName        string `gorm:"first_name"`
	LastName         string `gorm:"last_name"`
	DateOfBirth      string `gorm:"date_of_birth"`
	Salary           string
	CurrentBusyState bool `gorm:"current_busy_state"`
	GenderId         int  `gorm:"gender_id"`
	SpecialityId     int  `gorm:"speciality_id"`
}

type DoctorData struct {
	db *gorm.DB
}

func NewDoctor(db *gorm.DB) *DoctorData {
	return &DoctorData{db: db}
}

func (d DoctorData) ReadAll() ([]Doctors, error) {
	var doctors []Doctors
	result := d.db.Find(&doctors)
	if result.Error != nil {
		return nil, fmt.Errorf("can't read doctors from database, error: %w", result.Error)
	}
	return doctors, nil
}

func (d DoctorData) CreateDoctor(doctor Doctors) (int, error) {
	result := d.db.Create(&doctor)
	if result.Error != nil {
		return -1, fmt.Errorf("can't create doctor to database, error: %w", result.Error)
	}
	return doctor.DoctorId, nil
}

func (d DoctorData) UpdateDoctor(id int, updatedValues Doctors) (int, error) {
	var doctor Doctors
	resSearch := d.db.First(&doctor, id)
	if resSearch.Error != nil {
		return -1, fmt.Errorf("can't find doctor record to database using id - %d, error: %w", id, resSearch.Error)
	}
	result := d.db.Model(&doctor).Select("first_name", "last_name").Updates(updatedValues)
	if result.Error != nil {
		return -1, fmt.Errorf("can't update doctor to database using id - %d, error: %w", id, result.Error)
	}
	return doctor.DoctorId, nil
}

func (d DoctorData) DeleteDoctor(id int) error {
	result := d.db.Delete(&Doctors{}, id)
	if result.Error != nil {
		return fmt.Errorf("can't delete doctor to database using id - %d, error: %w", id, result.Error)
	}
	return nil
}

func (d DoctorData) ExecInnerJoin() error {
	rows, err := d.db.Table("doctors").Select("doctors.doctor_id, speciality.speciality_name").Joins(
		"INNER JOIN speciality ON doctors.speciality_id = speciality.speciality_id").Rows()
	if err != nil {
		return fmt.Errorf("can't execute INNER JOIN query, error: %w", err)
	}
	var (
		doctorId       int
		specialityName string
	)
	for rows.Next() {
		rows.Scan(&doctorId, &specialityName)
	}
	return nil
}
