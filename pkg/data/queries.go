package data

const (
	insertDoctor  = `INSERT INTO "doctors" ("first_name","last_name","date_of_birth","salary","current_busy_state","gender_id","speciality_id","doctor_id") VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING "doctor_id"`
	deleteDoctor  = `DELETE FROM "doctors" WHERE "doctors"."doctor_id" = $1`
	updateDoctor  = `UPDATE "doctors" SET "first_name"=$1 WHERE "doctor_id" = $2`
	selectAllRows = `SELECT * FROM "doctors"`
)
