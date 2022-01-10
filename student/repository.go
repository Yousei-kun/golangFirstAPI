package student

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Student, error)
	FindByID(ID int) (Student, error)
	Create(student Student) (Student, error)
	Update(student Student) (Student, error)
	Delete(student Student) (Student, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo *repository) FindAll() ([]Student, error) {
	var students []Student
	errorFindAll := repo.db.Order("ID").Find(&students).Error

	return students, errorFindAll
}

func (repo *repository) FindByID(id int) (Student, error) {
	var student Student
	errorFindByID := repo.db.Find(&student, id).Error

	return student, errorFindByID
}

func (repo *repository) Create(student Student) (Student, error) {
	errorCreate := repo.db.Create(&student).Error

	return student, errorCreate
}

func (repo *repository) Update(student Student) (Student, error) {
	errorUpdate := repo.db.Save(&student).Error

	return student, errorUpdate
}

func (repo *repository) Delete(student Student) (Student, error) {
	errorDelete := repo.db.Delete(&student).Error

	return student, errorDelete
}
