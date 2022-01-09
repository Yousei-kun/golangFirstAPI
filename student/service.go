package student

type Service interface {
	FindAll() ([]Student, error)
	FindByID(ID int) (Student, error)
	Create(student StudentRequest) (Student, error)
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
