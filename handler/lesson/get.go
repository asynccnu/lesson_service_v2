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
	var grade int
	var class []*model.LessonItem
	name := c.DefaultQuery("name", "")       //课程名字：数字逻辑
	teacher := c.DefaultQuery("teacher", "") //老师名字
	gradeTemp := c.DefaultQuery("grade", "") //面向对象

	if gradeTemp != "" {
		grade, err = strconv.Atoi(gradeTemp)
		if err != nil {
			handler.SendBadRequest(c, errno.ErrQuery, nil, "The 'grade' is wrong.")
			return
		}
	} else {
		grade = -1
	}

	class, err = model.GetClassDoc(name, teacher, grade)

	if mongo.ErrNoDocuments == err {
		handler.SendError(c, errno.ErrGetClasses, nil, err.Error())
		return
	}
	handler.SendResponse(c, nil, class)
}
