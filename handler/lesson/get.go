package lesson

import (
	"strconv"

	"github.com/asynccnu/lesson_service_v2/handler"
	"github.com/asynccnu/lesson_service_v2/model"
	"github.com/asynccnu/lesson_service_v2/pkg/errno"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Get(c *gin.Context) {
	var err error

	name := c.DefaultQuery("name", "") // 课程名，必填
	if name == "" {
		handler.SendBadRequest(c, errno.ErrQuery, nil, "The 'name' is required.")
		return
	}

	teacher := c.DefaultQuery("teacher", "")                  // 教师姓名
	grade, err := strconv.Atoi(c.DefaultQuery("grade", "-1")) // 年级
	if err != nil {
		handler.SendBadRequest(c, errno.ErrQuery, nil, "The 'grade' is wrong.")
		return
	}

	lessones, err := model.GetClassDoc(name, teacher, grade)
	if err != nil && err != mongo.ErrNoDocuments {
		handler.SendError(c, errno.ErrGetClasses, nil, err.Error())
		return
	}

	handler.SendResponse(c, nil, lessones)
}
