CREATE TABLE Speciality 
(
	IdSpeciality SERIAL PRIMARY KEY,
	SpecialityName CHARACTER VARYING(50) NOT NULL
);
CREATE TABLE Desease 
(
	IdDesease SERIAL PRIMARY KEY,
	DeseaseName CHARACTER VARYING(50) NOT NULL
);
CREATE TABLE Gender 
(
	IdGender SERIAL PRIMARY KEY,
	GenderName CHARACTER VARYING(10) NOT NULL
);
CREATE TABLE Doctor
(
	IdDoctor SERIAL PRIMARY KEY,
	Fname CHARACTER VARYING(50) NOT NULL,
	Lname CHARACTER VARYING(50) NOT NULL,
	DateOfBirth DATE NOT NULL,
	Salary NUMERIC NOT NULL,
	CurrentBuzyState BOOLEAN NOT NULL,
	IdGender INTEGER REFERENCES Gender (IdGender) ON DELETE CASCADE,
	IdSpeciality INTEGER REFERENCES Speciality (IdSpeciality) ON DELETE CASCADE
);
CREATE TABLE Nurse
(
	IdNurse SERIAL PRIMARY KEY,
	Fname CHARACTER VARYING(50) NOT NULL,
	Lname CHARACTER VARYING(50) NOT NULL,
	DateOfBirth DATE NOT NULL,
	CountPatient NUMERIC NOT NULL,
	IsBuzy BOOLEAN NOT NULL,
	IdGender INTEGER REFERENCES Gender (IdGender) ON DELETE CASCADE,
	IdSpeciality INTEGER REFERENCES Speciality (IdSpeciality) ON DELETE CASCADE
);
CREATE TABLE Ward
(
	IdWard SERIAL PRIMARY KEY,
	Name CHARACTER VARYING(50) NOT NULL,
	Location CHARACTER VARYING(50) NOT NULL,
	QuartializationDate DATE NOT NULL,
	Capacity NUMERIC NOT NULL,
	IsEmpty BOOLEAN NOT NULL,
	IdDoctor INTEGER REFERENCES Doctor (IdDoctor) ON DELETE CASCADE,
	IdNurse INTEGER REFERENCES Nurse (IdNurse) ON DELETE CASCADE
);
CREATE TABLE Patient
(
	IdPatient SERIAL PRIMARY KEY,
	Fname CHARACTER VARYING(50) NOT NULL,
	Lname CHARACTER VARYING(50) NOT NULL,
	DateOfBirth DATE NOT NULL,
	NumberDays NUMERIC NOT NULL,
	IsSick BOOLEAN NOT NULL,
	IdDoctor INTEGER REFERENCES Doctor (IdDoctor) ON DELETE CASCADE,
	IdNurse INTEGER REFERENCES Nurse (IdNurse) ON DELETE CASCADE,
	IdWard INTEGER REFERENCES Ward (IdWard) ON DELETE CASCADE,
	IdGender INTEGER REFERENCES Gender (IdGender) ON DELETE CASCADE,
	IdDesease INTEGER REFERENCES Desease (IdDesease) ON DELETE CASCADE
);
CREATE TABLE Service
(
	IdService SERIAL PRIMARY KEY,
	Name CHARACTER VARYING(50) NOT NULL,
	DateOfBirth DATE NOT NULL,
	Price NUMERIC NOT NULL,
	IsDone BOOLEAN NOT NULL,
	IdDoctor INTEGER REFERENCES Doctor (IdDoctor) ON DELETE CASCADE,
	IdPatient INTEGER REFERENCES Patient (IdPatient) ON DELETE CASCADE
);
INSERT INTO Speciality (SpecialityName) VALUES 
('General Surgery'),
('Main Nurse'),
('Assistant nurse'),
('Neurology'),
('Infectious Diseases');
INSERT INTO Desease (DeseaseName) VALUES 
('Anthrax'),
('Botulism'),
('Brucellosis'),
('COVID-19'),
('Ebola virus');
INSERT INTO Gender (GenderName) VALUES 
('Female'),
('Male');
INSERT INTO Doctor (Fname,Lname,DateOfBirth,Salary,CurrentBuzyState
					,IdGender,IdSpeciality) VALUES 
('Daniil','Domaskanov','2000-09-12',1250.12,FALSE,2,1),
('Dmitry','Putkov','1999-02-13',5000.54,TRUE,2,4),
('Victor','Pashkevich','2001-07-25',570.14,FALSE,2,5),
('Eva','Dushkevich','2002-10-22',790.77,TRUE,1,1),
('Nikita','Miladovski','1995-04-18',680.56,FALSE,2,5);
INSERT INTO Nurse (Fname,Lname,DateOfBirth,CountPatient,isBuzy
					,IdGender,IdSpeciality) VALUES 
('Anya','Udodva','1988-06-9',5,TRUE,1,2),
('Natalya','Seliverstova','1976-01-29',10,TRUE,1,3),
('Evgenia','Gapanovich','1990-03-15',7,TRUE,1,3),
('Alina','Kosenkova','2000-09-12',11,TRUE,1,3),
('Veronika','Dolgolapteva','1991-02-10',12,FALSE,2,2);
INSERT INTO Ward (Name,Location,QuartializationDate,Capacity,IsEmpty
					,IdDoctor,IdNurse) VALUES 
('12 ward','1-st floor','2020-01-10',6,FALSE,1,1),
('24 ','2-st floor','2020-01-12',4,FALSE,2,2),
('36','3-st floor','2020-01-13',3,FALSE,3,3),
('49','4-st floor','2020-01-14',8,TRUE,4,4),
('55','5-st floor','2020-01-16',2,FALSE,5,5);
INSERT INTO Patient (Fname,Lname,DateOfBirth,NumberDays,IsSick
					,IdDoctor,IdNurse,IdWard,IdGender,IdDesease) VALUES 
('Amari','Fitzgerald','1999-02-12',14,TRUE,1,1,1,2,1),
('Buddy','Chandler','2000-09-12',12,TRUE,2,2,2,2,2),
('Yousaf','Schultz','2001-05-17',10,TRUE,3,3,3,2,3),
('Katherine','Mejia','2002-12-14',8,TRUE,4,4,4,1,4),
('Esme-Rose','Warner','2003-11-18',2,TRUE,5,5,5,2,5);
INSERT INTO Service (Name,Price,IsDone,DateOfService,IdDoctor,IdPatient) VALUES 
('Emergency room',1250.12,TRUE,'1999-02-12',1,1),
('Surgical',1400.124,FALSE,'2000-09-12',2,2),
('X ray/radiology',1567.19,TRUE,'2001-05-17',3,3),
('Laboratory',125.78,FALSE,'2002-12-14',4,4),
('Blood',55.12,TRUE,'2003-11-18',5,5);




