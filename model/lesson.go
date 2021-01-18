package model

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 批量新建课程文档数据
func CreateMultipleClassDocs(instances []*ClassItem) error {
	var docs []interface{}
	for _, instance := range instances {
		docs = append(docs, instance)
	}

	_, err := DB.Self.Database(DBName).Collection(ClassCol).InsertMany(context.TODO(), docs)

	return err
}

// 获取文档数据
func GetClassDoc(name, teacher string,grade int) ([]*ClassItem, error) {
	var class []*ClassItem

	cur,err := DB.Self.Database(DBName).Collection(ClassCol).Find(context.TODO(), bson.M{"name": primitive.Regex{Pattern: name}, "teacher": primitive.Regex{Pattern:teacher}, "grade": grade})
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	err = cur.All(context.Background(), &class)
	if err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())	
	return class, err
}

func GetClassDocNoGrade(name, teacher string) ([]*ClassItem, error) {
	var class []*ClassItem

	cur,err := DB.Self.Database(DBName).Collection(ClassCol).Find(context.TODO(), bson.M{"name": primitive.Regex{Pattern: name}, "teacher": primitive.Regex{Pattern:teacher}})
	if err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	err = cur.All(context.Background(), &class)
	if err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())	
	return class, err
}