package script

import (
	"github.com/asynccnu/lesson_service_v2/log"
	"github.com/asynccnu/lesson_service_v2/model"

	"go.uber.org/zap"
)

var instances []*model.ClassItem


// 导入数据至数据库
func ImportDataToDB() {
	// 批量插入数据
	if err := model.CreateMultipleClassDocs(instances); err != nil {
		log.Error("Inserting multiple data failed", zap.String("reason", err.Error()))
	}
}
