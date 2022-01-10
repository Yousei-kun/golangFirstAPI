package student

import "fmt"

type Service interface {
	FindAll() ([]Student, error)
	FindByID(ID int) (Student, error)
	Create(student StudentRequest) (Student, error)
	Update(id int, student StudentRequest) (Student, error)
	Delete(id int) (Student, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (serv *service) FindAll() ([]Student, error) {
	students, err := serv.repository.FindAll()

	return students, err
}

func (serv *service) FindByID(id int) (Student, error) {
	student, err := serv.repository.FindByID(id)

	return student, err
}

func (serv *service) Create(studentRequest StudentRequest) (Student, error) {

	price, _ := studentRequest.Score.(float64)

	student := Student{
		SID:   studentRequest.SID,
		Name:  studentRequest.Name,
		Score: price,
	}

	newStudent, err := serv.repository.Create(student)
	return newStudent, err
}

func (serv *service) Update(id int, studentRequest StudentRequest) (Student, error) {
	student, err := serv.repository.FindByID(id)

	if err != nil {
		fmt.Println("Error detected")
	}

	price, _ := studentRequest.Score.(float64)

	student.SID = studentRequest.SID
	student.Name = studentRequest.Name
	student.Score = price

	updatedStudent, errorUpdate := serv.repository.Update(student)
	return updatedStudent, errorUpdate
}

func (serv *service) Delete(id int) (Student, error) {
	student, err := serv.repository.FindByID(id)

	if err != nil {
		fmt.Println("Error detected")
	}

	deletedStudent, errorDelete := serv.repository.Delete(student)
	return deletedStudent, errorDelete
}
