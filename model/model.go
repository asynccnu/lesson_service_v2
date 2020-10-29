package model

// LessonModel ... 课程结构
type LessonModel struct {
	Grade    int    `json:"grade"`     // 授课年级，0 表示面向全体学生
	ForWhom  string `json:"for_whom"`  // 授课对象/专业，公共课为'全体学生'
	Name     string `json:"name"`      // 课程名
	Teacher  string `json:"teacher"`   // 教师，多个教师逗号隔开
	LessonNo string `json:"lesson_no"` // 学校课程号
	Kind     string `json:"kind"`      // 课程类型：选修课、专业课...
	Place    string `json:"place"`     // 上课地点，多个地点逗号分隔
	Time     string `json:"time"`      // 上课时间，多个时间'|'分隔（逗号可能包含在给定的时间串里，如"{1-3周,5,6周}"）
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
