package data

const readAllUsersQuery = "SELECT doctor_id, first_name, last_name, date_of_birth, salary, current_busy_state, gender_id, speciality_id FROM doctors"
const innerJoin = "SELECT doctors.doctor_id, speciality.speciality_name FROM doctors INNER JOIN speciality ON doctors.speciality_id = speciality.speciality_id"
const insertDoctor = `INSERT INTO "doctors" ("first_name","last_name","date_of_birth","salary","current_busy_state","gender_id","speciality_id","doctor_id") VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING "doctor_id"`
const deleteDoctor = "DELETE FROM doctors WHERE doctor_id = 1"
const updateDoctor = `UPDATE "doctors" SET "first_name"=$1 WHERE "id" = $2`
