package script

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/asynccnu/lesson_service_v2/log"
	"github.com/asynccnu/lesson_service_v2/model"

	"github.com/tealeg/xlsx"
	"go.uber.org/zap"
)

// 从选课手册中解析获得课程信息
func GetLessonInfoFromClassFile(channel chan *model.LessonItem, file *xlsx.File) {

	// 匹配教师，如果匹配失败，可能是，半角全角，或者多打的括号
	t, err := regexp.Compile("[\u4e00-\u9fa5]+")
	if err != nil {
		log.Error("Regexp compile failed", zap.String("reason", err.Error()))
		return
	}

	// 选课手册（0-15列）中第 10，11，12 列为上课时间，第 13，14，15 列为上课地点，第 2 列为课程名字，第 9 列为老师名字
	//	var build strings.Builder
	// var placeAndTime []model.Item
	// var itemTmp model.Item
	for _, sheet := range file.Sheets {

		// 解析年级：公共课 => 0，2020级 => 2020
		grade, err := ExtractGrade(sheet.Name)
		if err != nil {
			log.Error("extract grade error", zap.String("reason", err.Error()))
			return
		}
		// 遍历课程数据
		for rowIndex, row := range sheet.Rows {
			if rowIndex == 0 {
				continue
			}
			var placeAndTime []model.Item
			var itemTmp model.Item
			// 遍历一行课程数据中的多个时间、地点
			for j := 10; j <= 14; j += 2 {

				itemTmp.Time = row.Cells[j].String()
				itemTmp.Place = row.Cells[j+1].String()
				if itemTmp.Time == "" || itemTmp.Place == "" {
					continue
				}

				placeAndTime = append(placeAndTime, itemTmp)
				// build.WriteString(date)
				// build.WriteString(place)
				// build.WriteString(",")
			}

			//placeAndTime := build.String()
			forWhom := "all"
			if grade != 0 {
				forWhom = row.Cells[16].String()
			}

			name := row.Cells[2].String()
			teacherTmp := row.Cells[8].String()
			lessonNo := row.Cells[1].String()

			// 正则匹配
			teacherTmp2 := t.FindStringSubmatch(teacherTmp)
			teacher := ExtractTeacher(teacherTmp2)
			kind := ExtractLessonNo(lessonNo)

			channel <- &model.LessonItem{
				Grade:        grade,
				ForWhom:      forWhom,
				Name:         name,
				LessonNo:     lessonNo,
				Kind:         kind,
				Teacher:      teacher,
				PlaceAndTime: placeAndTime,
			}

		}
	}
	fmt.Println("Parsing class file OK")
}

//处理课程编号得出课程性质
func ExtractLessonNo(lessonno string) string {
	//性质分两种
	kind1Map := map[string]string{"0": "通识必修课", "1": "专业必修课", "2": "专业选修课", "3": "通识选修课", "4": "专业课", "5": "通识核心课"}
	kind2Map := map[string]string{"0": "公共必修课及专业课", "1": "数学与自然科学类", "2": "哲学与社会科学类", "3": "人文与艺术类", "4": "教育学与心理学类"}

	temp1 := lessonno[3:4]
	temp2 := lessonno[4:5]

	kind := kind1Map[temp1] + "," + kind2Map[temp2]
	return kind
}

func ExtractTeacher(teacherTmp2 []string) string {
	length := len(teacherTmp2)
	var build strings.Builder

	for i := 0; i < length; i++ {
		build.WriteString(teacherTmp2[i])
		build.WriteString(",")
	}
	teacher := build.String()
	return teacher
}

// 解析年级
func ExtractGrade(gradeStr string) (int, error) {
	if strings.Contains(gradeStr, "公共课") {
		return 0, nil
	}

	// 正则
	r := regexp.MustCompile("([0-9]{4}).*")

	matchGroups := r.FindStringSubmatch(gradeStr)
	if len(matchGroups) < 1 {
		return 0, errors.New("mathch failed")
	}

	grade, err := strconv.Atoi(matchGroups[1])
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Atoi %s failed", matchGroups[1]))
	}

	return grade, nil
}
