package store

import "github.com/ArsalanKm/students/internal/model"

// Memory Students is an implementation of Student store which uses memory as storage.
type MemoryStudent struct {
	students map[string]model.Student
}

func NewMomoryStudent() *MemoryStudent{
	return &MemoryStudent{
		students: make(map[string]model.Student),
	}
}
func (ms *MemoryStudent) Save(student model.Student) error{
	if _ , ok := ms.students[student.ID]; ok {
		return ErrStudentDuplicate
	}
	ms.students[student.ID]=student;
	return nil
}

func (ms *MemoryStudent) Load(Id string) (model.Student, error){
	if _,ok:=ms.students[Id];ok{
		return ms.students[Id],nil
	}
	return model.Student{},ErrStudentNotFund
}
