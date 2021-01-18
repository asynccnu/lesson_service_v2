package script

import (
	"github.com/asynccnu/lesson_service_v2/log"
	"github.com/asynccnu/lesson_service_v2/model"

	"go.uber.org/zap"
)

var instances []*model.ClassItem

// type CourseItem struct{
// 	Grade    int    `bson:"grade" json:"grade"`     // 授课年级，0 表示面向全体学生
// 	Name     string `bson:"name" json:"name"`      // 课程名
// 	Teacher  string `bson:"teacher" json:"teacher"`   // 教师名字
// 	ForWhom  string `bson:"for_whom" json:"for_whom"`  // 授课对象/专业，公共课为'全体学生'
// 	LessonNo string `bson:"lesson_no" json:"lesson_no"` // 学校课程号
// 	Kind     string `bson:"kind" json:"kind"`      // 课程类型：选修课、专业课...
// 	PlaceAndTime    string `bson:"placeandtime" json:"placeandtime"`     // 上课地点地点，多个逗号分隔：“8314星期一第7-8节{2-18周} ， 8312星期五第1-2节{3-17周(单)”
// }



// 导入数据至数据库
func ImportDataToDB() {
	// 批量插入数据
	if err := model.CreateMultipleClassDocs(instances); err != nil {
		log.Error("Inserting multiple data failed", zap.String("reason", err.Error()))
	}
}
