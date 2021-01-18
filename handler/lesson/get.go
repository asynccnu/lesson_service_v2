package lesson

import (
	//"fmt"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/asynccnu/lesson_service_v2/handler"
	"github.com/asynccnu/lesson_service_v2/model"
	//"github.com/asynccnu/lesson_service_v2/model/model.go"
	"github.com/asynccnu/lesson_service_v2/pkg/errno"
	"go.mongodb.org/mongo-driver/mongo"
)

func Get(c *gin.Context) {
	var err error
	var class []*model.ClassItem
	name := c.DefaultQuery("name","")				//课程名字：数字逻辑
	teacher := c.DefaultQuery("teacher","")	//老师名字：陈曙
	gradet	:= c.DefaultQuery("grade","")			//面向对象：0：公共课，17：2017级

	if (gradet == ""){
		class, err = model.GetClassDocNoGrade(name,teacher)
	}else{
	grade, err := strconv.Atoi(gradet)
	if err != nil {
		handler.SendBadRequest(c, errno.ErrQuery, nil, "The 'grade' is wrong.")
		return
	}
	class, err = model.GetClassDoc(name,teacher,grade)
	}
	if mongo.ErrNoDocuments == err {
		// class = &model.ClassItem{
		// 		Grade:	grade,    
		// 		Name:	name,
		// 		Teacher:	teacher,  
		// 		//List:     make([]*model.ClassItem, 0),
		// }
		handler.SendError(c, errno.ErrGetClasses, nil, err.Error())
		return
	}
	handler.SendResponse(c, nil, class)
}
