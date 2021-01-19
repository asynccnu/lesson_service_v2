package script

import (
	"fmt"
	"regexp"
	// "strings"
	"github.com/asynccnu/lesson_service_v2/log"
	"github.com/asynccnu/lesson_service_v2/model"

	"github.com/tealeg/xlsx"
	"go.uber.org/zap"
)

// 从选课手册中解析获得课程信息
func GetCourseInfoFromClassFile(channel chan *model.ClassItem, file *xlsx.File) {

	gradeMap := map[string]int{"共课": 0, "017级": 17, "018级": 18, "019级": 19, "020级": 20}

	t, err := regexp.Compile("[\u4e00-\u9fa5]+") //如果匹配失败，可能是，半角全角，或者多打的括号
	if err != nil {
		log.Error("Regexp compile failed", zap.String("reason", err.Error()))
		return
	}

	// 选课手册（0-15列）中第 10，11，12 列为上课时间，第 13，14，15 列为上课地点，第 2 列为课程名字，第 9 列为老师名字
	placeAndtime := " "
	for _, sheet := range file.Sheets {
		//fmt.Println(sheet.Name)
		grade := sheet.Name
		gradeflag := grade[len(grade)-6:]

		// 遍历课程数据
		for _, row := range sheet.Rows {

			// 遍历一行课程数据中的多个时间、地点
			for j := 10; j <= 14; j += 2 {

				date := row.Cells[j].String()
				place := row.Cells[j+1].String()
				if date == "" || place == "" {
					continue
				}

				if date != "上课时间1" && date != "上课时间2" && date != "上课时间3" {
					placeAndtime += date + place + " , "
				}
			}

			forwhom := "all"
			if gradeMap[gradeflag] != 0 {
				forwhom = row.Cells[16].String()
			}
			//fmt.Println(sheet.Name)
			name := row.Cells[2].String()
			teachertmp := row.Cells[8].String()
			lessonno := row.Cells[1].String()

			// 正则匹配
			teachertmp2 := t.FindStringSubmatch(teachertmp)
			teacher := ExtractTeacher(teachertmp2)
			kind := ExtractLessonNo(lessonno)

			channel <- &model.ClassItem{
				Grade:        gradeMap[gradeflag],
				ForWhom:      forwhom,
				Name:         name,
				LessonNo:     lessonno,
				Kind:         kind,
				Teacher:      teacher,
				PlaceAndTime: placeAndtime,
			}
			placeAndtime = ""

		}
	}
	fmt.Println("Parsing class file OK")
}

//处理课程编号得出课程性质
func ExtractLessonNo(lessonno string) string {
	//性质分两种
	kind1Map := map[string]string{"0": "通识必修课", "1": "专业必修课", "2": "专业选修课", "3": "通识选修课", "4": "专业课", "5": "通识核心课"}
	kind2Map := map[string]string{"0": "公共必修课及专业课", "1": "数学与自然科学类", "2": "哲学与社会科学类", "3": "人文与艺术类", "4": "教育学与心理学类"}

	c1 := lessonno[3:4]
	c2 := lessonno[4:5]

	kind := kind1Map[c1] + "," + kind2Map[c2]
	return kind
}

func ExtractTeacher(teachertmp2 []string) string {
	length := len(teachertmp2)
	var teacher string

	for i := 0; i < length; i++ {
		teacher = teachertmp2[i] + ","
	}
	return teacher
}
