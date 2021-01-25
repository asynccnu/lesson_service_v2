## lesson_service_v2

匣子 Golang 版蹭课服务

### Env

```shell
export CCNUBOX_LESSON_DB_NAME='lesson_20_21_1' # 库名，规定格式，20_21_1为2020-2021学年第一学期的课程
export CCNUBOX_LESSON_DB_URL='' # mongoDB URL
```

### Run

```shell
make
./main
```

### Script

拉取选课手册课程

```shell
go run main.go -p course.xlsx
```
