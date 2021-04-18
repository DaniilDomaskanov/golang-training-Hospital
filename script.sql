CREATE TABLE speciality
(
	speciality_id SERIAL PRIMARY KEY,
	speciality_name CHARACTER VARYING(50) NOT NULL
);
CREATE TABLE disease
(
	disease_id SERIAL PRIMARY KEY,
	disease_name CHARACTER VARYING(50) NOT NULL
);
CREATE TABLE gender
(
	gender_id SERIAL PRIMARY KEY,
	gender_name CHARACTER VARYING(10) NOT NULL
);
CREATE TABLE doctor
(
	doctor_id SERIAL PRIMARY KEY,
	first_name CHARACTER VARYING(50) NOT NULL,
	last_name CHARACTER VARYING(50) NOT NULL,
	date_of_birth DATE NOT NULL,
	salary NUMERIC NOT NULL,
	current_busy_state BOOLEAN NOT NULL,
	gender_id INTEGER REFERENCES gender (gender_id) ,
    speciality_id INTEGER REFERENCES speciality (speciality_id)
);
CREATE TABLE nurse
(
	nurse_id SERIAL PRIMARY KEY,
    first_name CHARACTER VARYING(50) NOT NULL,
    last_name CHARACTER VARYING(50) NOT NULL,
    date_of_birth DATE NOT NULL,
	count_patient NUMERIC NOT NULL,
	is_busy BOOLEAN NOT NULL,
	gender_id INTEGER REFERENCES gender (gender_id),
	speciality_id INTEGER REFERENCES speciality (speciality_id)
);
CREATE TABLE ward
(
	ward_id SERIAL PRIMARY KEY,
	name CHARACTER VARYING(50) NOT NULL,
	location CHARACTER VARYING(50) NOT NULL,
	quartialization_date DATE NOT NULL,
	capacity NUMERIC NOT NULL,
	is_empty BOOLEAN NOT NULL,
	doctor_id INTEGER REFERENCES doctor (doctor_id),
	nurse_id INTEGER REFERENCES nurse (nurse_id)
);
CREATE TABLE patient
(
	patient_id SERIAL PRIMARY KEY,
	first_name CHARACTER VARYING(50) NOT NULL,
	last_name CHARACTER VARYING(50) NOT NULL,
	date_of_birth DATE NOT NULL,
	number_days NUMERIC NOT NULL,
	is_sick BOOLEAN NOT NULL,
	doctor_id INTEGER REFERENCES doctor (doctor_id),
    nurse_id INTEGER REFERENCES nurse (nurse_id),
    ward_id INTEGER REFERENCES ward (ward_id),
    gender_id INTEGER REFERENCES gender (gender_id),
    disease_id INTEGER REFERENCES disease (disease_id)
);
CREATE TABLE service
(
	service_id SERIAL PRIMARY KEY,
	name CHARACTER VARYING(50) NOT NULL,
	date_of_birth DATE NOT NULL,
	price NUMERIC NOT NULL,
	is_done BOOLEAN NOT NULL,
    doctor_id INTEGER REFERENCES doctor (doctor_id),
    patient_id INTEGER REFERENCES patient (patient_id)
);
INSERT INTO speciality (speciality_name) VALUES
('General Surgery'),
('Main Nurse'),
('Assistant nurse'),
('Neurology'),
('Infectious Diseases');
INSERT INTO disease (disease_name) VALUES
('Anthrax'),
('Botulism'),
('Brucellosis'),
('COVID-19'),
('Ebola virus');
INSERT INTO gender (gender_name) VALUES
('Female'),
('Male');
INSERT INTO doctor (first_name,last_name,date_of_birth,salary,current_busy_state
					,gender_id,speciality_id) VALUES
('Daniil','Domaskanov','2000-09-12',1250.12,FALSE,2,1),
('Dmitry','Putkov','1999-02-13',5000.54,TRUE,2,4),
('Victor','Pashkevich','2001-07-25',570.14,FALSE,2,5),
('Eva','Dushkevich','2002-10-22',790.77,TRUE,1,1),
('Nikita','Miladovski','1995-04-18',680.56,FALSE,2,5);
INSERT INTO nurse (first_name,last_name,date_of_birth,count_patient,is_busy
					,gender_id,speciality_id) VALUES
('Anya','Udodva','1988-06-9',5,TRUE,1,2),
('Natalya','Seliverstova','1976-01-29',10,TRUE,1,3),
('Evgenia','Gapanovich','1990-03-15',7,TRUE,1,3),
('Alina','Kosenkova','2000-09-12',11,TRUE,1,3),
('Veronika','Dolgolapteva','1991-02-10',12,FALSE,2,2);
INSERT INTO ward (name,location,quartialization_date,capacity,is_empty
					,doctor_id,nurse_id) VALUES
('12 ward','1-st floor','2020-01-10',6,FALSE,1,1),
('24 ','2-st floor','2020-01-12',4,FALSE,2,2),
('36','3-st floor','2020-01-13',3,FALSE,3,3),
('49','4-st floor','2020-01-14',8,TRUE,4,4),
('55','5-st floor','2020-01-16',2,FALSE,5,5);
INSERT INTO patient (first_name,last_name,date_of_birth,number_days,is_sick
					,doctor_id,nurse_id,ward_id,gender_id,disease_id) VALUES
('Amari','Fitzgerald','1999-02-12',14,TRUE,1,1,1,2,1),
('Buddy','Chandler','2000-09-12',12,TRUE,2,2,2,2,2),
('Yousaf','Schultz','2001-05-17',10,TRUE,3,3,3,2,3),
('Katherine','Mejia','2002-12-14',8,TRUE,4,4,4,1,4),
('Esme-Rose','Warner','2003-11-18',2,TRUE,5,5,5,2,5);
INSERT INTO service (name,price,is_done,date_of_birth,doctor_id,patient_id) VALUES
('Emergency room',1250.12,TRUE,'1999-02-12',1,1),
('Surgical',1400.124,FALSE,'2000-09-12',2,2),
('X ray/radiology',1567.19,TRUE,'2001-05-17',3,3),
('Laboratory',125.78,FALSE,'2002-12-14',4,4),
('Blood',55.12,TRUE,'2003-11-18',5,5);




