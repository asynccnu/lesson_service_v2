package model

var DBName = "class"

const (
	ClassCol = "class"
)

type ClassItem struct{
	Grade    int    `bson:"grade" json:"grade"`     // 授课年级，0 表示面向全体学生
	Name     string `bson:"name" json:"name"`      // 课程名
	Teacher  string `bson:"teacher" json:"teacher"`   // 教师名字
	ForWhom  string `bson:"for_whom" json:"for_whom"`  // 授课对象/专业，公共课为'全体学生'
	LessonNo string `bson:"lesson_no" json:"lesson_no"` // 学校课程号
	Kind     string `bson:"kind" json:"kind"`      // 课程类型：选修课、专业课...
	PlaceAndTime    string `bson:"placeandtime" json:"placeandtime"`     // 上课地点地点，多个逗号分隔：“8314星期一第7-8节{2-18周} ， 8312星期五第1-2节{3-17周(单)”
}

// ClassModel ... 课程结构
// type ClassModel struct {
// 	Grade    int    `bson:"grade" json:"grade"`     // 授课年级，0 表示面向全体学生
// 	Name     string `bson:"name" json:"name"`      // 课程名
//     Teacher  string `bson:"teacher" json:"teacher"`   // 教师，多个教师逗号隔开
//     List     []*ClassItem   `bson:"list" json:"list"`   
// 	//Time     string `json:"time"`      // 上课时间，多个时间'|'分隔（逗号可能包含在给定	"github.com/asynccnu/lesson_service_v2/model"的时间串里，如"{1-3周,5,6周}"）
// }

// type ClassItem struct{
//     ForWhom  string `bson:"for_whom" json:"for_whom"`  // 授课对象/专业，公共课为'全体学生'
// 	LessonNo string `bson:"lesson_no" json:"lesson_no"` // 学校课程号
// 	Kind     string `bson:"kind" json:"kind"`      // 课程类型：选修课、专业课...
// 	PlaceAndTime    string `bson:"placeandtime" json:"placeandtime"`     // 上课地点地点，多个逗号分隔：“8314星期一第7-8节{2-18周} ， 8312星期五第1-2节{3-17周(单)”
// }
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
