package model

import (
	"context"
	"log"

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
	var lesson []*LessonItem

	//	if(grade!=-1){
	cur, err := DB.Self.Database(DBName).Collection(LessonCol).Find(context.TODO(), bson.M{"name": primitive.Regex{Pattern: name}, "teacher": primitive.Regex{Pattern: teacher}, "grade": grade})
	if grade == -1 {
		cur, err = DB.Self.Database(DBName).Collection(LessonCol).Find(context.TODO(), bson.M{"name": primitive.Regex{Pattern: name}, "teacher": primitive.Regex{Pattern: teacher}})
	}
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	err = cur.All(context.Background(), &lesson)
	if err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())
	return lesson, err
}
