package model

import (
	"context"

	"github.com/asynccnu/lesson_service_v2/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 批量新建课程文档数据
func CreateMultipleLessonDocs(instances []*LessonItem) error {
	var docs []interface{}
	for _, instance := range instances {
		docs = append(docs, instance)
	}

	_, err := DB.Self.Database(DBName).Collection(LessonCol).InsertMany(context.TODO(), docs)

	return err
}

// 获取文档数据
func GetClassDoc(name, teacher string, grade int) ([]*LessonItem, error) {
	var lesson = make([]*LessonItem, 0)

	// 匹配模式：课程名，教师，年级
	// 若年级不为 -1，则匹配年级
	pattern := bson.M{"name": primitive.Regex{Pattern: name}, "teacher": primitive.Regex{Pattern: teacher}}
	if grade != -1 {
		pattern = bson.M{"name": primitive.Regex{Pattern: name}, "teacher": primitive.Regex{Pattern: teacher}, "grade": grade}
	}

	cur, err := DB.Self.Database(DBName).Collection(LessonCol).Find(context.TODO(), pattern)
	if err != nil {
		log.Error("Find docs error." + err.Error())
		return lesson, err
	}

	defer cur.Close(context.TODO())

	if err := cur.Err(); err != nil {
		log.Error("Cursor error." + err.Error())
		return lesson, err
	}

	err = cur.All(context.TODO(), &lesson)
	if err != nil {
		log.Error("Docs decode to slice error." + err.Error())
		return lesson, err
	}

	return lesson, err
}
