## Description

Application for mantaining hospital with such operations as in table below:


|             Path            | Method | Description                           | Body example                                                                                                                                                                                                                     |
|:---------------------------:|--------|---------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| /doctors                   | GET    | get all doctors                      | {"DoctorId":3,"FirstName":"Victor","LastName":"Pashkevich","DateOfBirth":"2001-07-25","Salary":"570.14","CurrentBusyState":false,"GenderId":2,"SpecialityId":5},{"DoctorId":4,"FirstName":"Eva","LastName":"Dushkevich","DateOfBirth":"2002-10-22","Salary":"790.77","CurrentBusyState":true,"GenderId":1,"SpecialityId":1},{"DoctorId":5,"FirstName":"Nikita","LastName":"Miladovski","DateOfBirth":"1995-04-18","Salary":"680.56","CurrentBusyState":false,"GenderId":2,"SpecialityId":5},{"DoctorId":2,"FirstName":"Dmitry","LastName":"Putkov","DateOfBirth":"1999-02-13","Salary":"5000.54","CurrentBusyState":true,"GenderId":2,"SpecialityId":4},{"DoctorId":1,"FirstName":"Danik","LastName":"Voinov","DateOfBirth":"2000-09-12","Salary":"1250.12","CurrentBusyState":false,"GenderId":2,"SpecialityId":1} |
| /create                   | POST   | create new doctor instance                    |                                                                                                                                                                                                                                  |
| /update/id              | PUT    | update doctor by the id                 | {
    "firstName":"Danik",
    "lastName":"Voronin"
}                                                                                                                                  |
| /delete/id              | DELETE | delete doctor by the id              |                                                                                                                                                                                                                                  |

## Usage 
1. Run server on port `8080`
	`go run cmd/main.go`
2.  Open URL
`http://localhost:8080/`

## Usage unit tests
To run unit tests type:
`go test ./...`

