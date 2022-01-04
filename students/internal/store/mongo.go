package store

import (
	"context"
	"fmt"

	"github.com/ArsalanKm/students/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDB struct {
	db *mongo.Database
}

const Collection = "students"
const DatabaseName = "sbu"

func NewMongoDBStore(db *mongo.Database) MongoDB {
	return MongoDB{
		db: db.Client().Database(DatabaseName),
	}
}

func (m MongoDB) Save(s model.Student) error {
	_, error := m.db.Collection(Collection).InsertOne(context.TODO(), s)
	if error != nil {
		return fmt.Errorf("mongodb insert failed %w", error)
	}
	return nil
}

func (m MongoDB) LoadById(id string) (model.Student, error) {
	return model.Student{}, nil
}
func (m MongoDB) Load() ([]model.Student, error) {
	var students []model.Student
	records, err := m.db.Collection(Collection).Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("mongo find faile %w", err)
	}
	for records.Next(context.TODO()) {
		var student model.Student

		if err := records.Decode(&student); err != nil {
			return students, fmt.Errorf("mongo db record decoding failed %w", err)
		}

		students = append(students, student)
	}
	return students, nil

}
