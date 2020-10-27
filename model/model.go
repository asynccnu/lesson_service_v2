package model

// LessonModel ... 课程结构
type LessonModel struct {
	Grade    int              `json:"grade"`     // 授课年级，公共课为 0
	ForWhom  string           `json:"for_whom"`  // 授课对象/专业，公共课为'全体学生'
	Name     string           `json:"name"`      // 课程名
	Teacher  string           `json:"teacher"`   // 教师，多个教室逗号隔开
	LessonNo string           `json:"lesson_no"` // 学校课程号
	Kind     string           `json:"kind"`      // 课程类型：选修课、专业课...
	Info     []*ClassInfoItem `json:"info"`      // 上课时间地点信息
	// Place    string `json:"place"`
	// Time     string `json:"time"`
}

// ClassInfoItem ... 上课信息，时间&地点
type ClassInfoItem struct {
	Place string `json:"place"`
	Time  string `json:"time"`
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
    "when" : "星期一第7-8节{2-18周}|星期五第1-2节{3-17周(单)}
}
*/
