***Notes***

1. 查询时，可以输入年级，老师，课程任意一个，两个，或者全部
2. 课程，以及老师允许少输，不允许错输

***获取课程***

| Method | Query              | URL           |
| ------ | ------------------ | ------------- |
| GET    | Name,teacher,grade |ccnubox.muxixyz.com/api/lesson/v2 |

***URL Params***

```
name:数据结构
teacher:董石
grade: 2019
```

**Response Data**				

```
{
    "code": 0,
    "message": "OK",
    "data": [
        {
            "grade": 2019,
            "name": "数据结构",
            "teacher": "董石,",
            "for_whom": "数字媒体技术",
            "lesson_no": "48900004",
            "kind": "通识必修课,公共必修课及专业课",
            "place_and_time": [
                {
                    "place": "N311",
                    "time": "星期一第1-2节{4-20周}"
                },
                {
                    "place": "信技实验室",
                    "time": "星期三第5-6节{4-12周}"
                }
            ]
        },
        {
            "grade": 2019,
            "name": "数据结构",
            "teacher": "董石,",
            "for_whom": "数字媒体技术",
            "lesson_no": "48900004",
            "kind": "通识必修课,公共必修课及专业课",
            "place_and_time": [
                {
                    "place": "N311",
                    "time": "星期二第7-8节{4-20周}"
                },
                {
                    "place": "信技实验室",
                    "time": "星期四第1-2节{4-12周}"
                }
            ]
        },
        {
            "grade": 2019,
            "name": "数据结构实验",
            "teacher": "董石,",
            "for_whom": "数字媒体技术",
            "lesson_no": "48900008",
            "kind": "通识必修课,公共必修课及专业课",
            "place_and_time": [
                {
                    "place": "信技实验室",
                    "time": "星期三第5-6节{13-20周}"
                }
            ]
        },
        {
            "grade": 2019,
            "name": "数据结构实验",
            "teacher": "董石,",
            "for_whom": "数字媒体技术",
            "lesson_no": "48900008",
            "kind": "通识必修课,公共必修课及专业课",
            "place_and_time": [
                {
                    "place": "信技实验室",
                    "time": "星期四第1-2节{13-20周}"
                }
            ]
        }
    ]
}
```

