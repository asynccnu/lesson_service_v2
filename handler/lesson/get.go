package lesson

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/asynccnu/lesson_service_v2/handler"
	"github.com/asynccnu/lesson_service_v2/model"
	"github.com/asynccnu/lesson_service_v2/pkg/errno"
	"go.mongodb.org/mongo-driver/mongo"
)

func Get(c *gin.Context) {
	var err error
	var class []*model.ClassItem
	name := c.DefaultQuery("name","")				//课程名字：数字逻辑
	teacher := c.DefaultQuery("teacher","")	//老师名字
	gradetemp	:= c.DefaultQuery("grade","")			//面向对象

	if (gradetemp == ""){
		class, err = model.GetClassDocNoGrade(name,teacher)
	}else{
	grade, err := strconv.Atoi(gradetemp)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrQuery, nil, "The 'grade' is wrong.")
		return
	}
	class, err = model.GetClassDoc(name,teacher,grade)
	}
	if mongo.ErrNoDocuments == err {
		handler.SendError(c, errno.ErrGetClasses, nil, err.Error())
		return
	}
	handler.SendResponse(c, nil, class)
}
