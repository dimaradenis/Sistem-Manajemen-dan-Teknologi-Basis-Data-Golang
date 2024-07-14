package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Update(id int, s *model.Student) error
	Delete(id int) error
	FetchWithClass() (*[]model.StudentClass, error)
}

type studentRepoImpl struct {
	db *gorm.DB
}

func NewStudentRepo(db *gorm.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	var students []model.Student
	err := s.db.Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	if err := s.db.Create(&student).Error; err != nil {
		return err
	}
	return nil
}

func (s *studentRepoImpl) Update(id int, student *model.Student) error {
	if err := s.db.Model(&model.Student{}).Where("id = ?", id).Updates(student).Error; err != nil {
		return err
	}
	return nil
}

func (s *studentRepoImpl) Delete(id int) error {
	var student model.Student
	err := s.db.Where("id = ?", id).First(&student).Error
	if err != nil {
		return err
	}

	err = s.db.Delete(&student).Error
	return err
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	var student model.Student
	if err := s.db.First(&student, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (s *studentRepoImpl) FetchWithClass() (*[]model.StudentClass, error) {
	students := []model.StudentClass{}
	if err := s.db.Table("students").
		Select("students.*, classes.name AS class_name, classes.professor AS professor, classes.room_number AS room_number").
		Joins("JOIN classes ON students.class_id = classes.id").
		Scan(&students).Error; err != nil {
		return &students, err
	}
	return &students, nil
}
