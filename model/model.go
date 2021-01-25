package model

const (
	LessonCol = "lesson"
)

var DBName = "lesson"

type LessonItem struct {
	Grade        int    `bson:"grade" json:"grade"`                   // 授课年级，0 表示面向全体学生
	Name         string `bson:"name" json:"name"`                     // 课程名
	Teacher      string `bson:"teacher" json:"teacher"`               // 教师名字
	ForWhom      string `bson:"for_whom" json:"for_whom"`             // 授课对象/专业，公共课为'全体学生'
	LessonNo     string `bson:"lesson_no" json:"lesson_no"`           // 学校课程号
	Kind         string `bson:"kind" json:"kind"`                     // 课程类型：选修课、专业课...
	PlaceAndTime string `bson:"place_and_time" json:"place_and_time"` // 上课地点地点，多个逗号分隔：“8314星期一第7-8节{2-18周}， 8312星期五第1-2节{3-17周(单)”
}

/*
老版本：
{
    "rank" : 2016,
    "forwho" : "计算机科学与技术,软件工程",
    "name" : "Web数据库设计与开发",
    "teacher" : "2006982685/刘巍",
    "no" : "48721003",
    "kind" : "专业课",
    "where" : "8314，8312",
    "when" : "星期一第7-8节{2-18周}|星期五第1-2节{3-17周(单)}"
}

新版：
{
    "grade" : 2016,
    "for_whom" : "计算机科学与技术,软件工程",
    "name" : "Web数据库设计与开发",
    "teacher" : "刘巍",
    "lesson_no" : "48721003",
	"kind" : "专业课",
	"place" : "8314，8312",
    "time" : "星期一第7-8节{2-18周}|星期五第1-2节{3-17周(单)}"
}
*/
