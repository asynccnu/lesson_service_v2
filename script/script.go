package script

import (
	"fmt"

	"github.com/asynccnu/lesson_service_v2/log"
	"github.com/asynccnu/lesson_service_v2/model"
	"github.com/tealeg/xlsx"

	"go.uber.org/zap"
)

// 解析并导入课程数据
func SyncImportLessonData(filePath string) {
	// 打开 Excel 文件
	file, err := xlsx.OpenFile(filePath)
	if err != nil {
		log.Fatal("Open xlsxFlie failed", zap.String("reason", err.Error()))
		return
	}

	channel := make(chan *model.LessonItem, 10)

	// 解析获取获取课程信息
	go func() {
		defer close(channel)
		GetLessonInfoFromClassFile(channel, file)
	}()

	for item := range channel {
		instances = append(instances, &model.LessonItem{
			Grade:        item.Grade,
			ForWhom:      item.ForWhom,
			Name:         item.Name,
			LessonNo:     item.LessonNo,
			Kind:         item.Kind,
			Teacher:      item.Teacher,
			PlaceAndTime: item.PlaceAndTime,
		})
	}

	// 导入空闲教室数据至数据库
	ImportDataToDB()

	fmt.Println("Import data into DB OK")
}
