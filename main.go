package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/asynccnu/lesson_service_v2/config"
	"github.com/asynccnu/lesson_service_v2/log"
	"github.com/asynccnu/lesson_service_v2/model"
	"github.com/asynccnu/lesson_service_v2/router"
	"github.com/asynccnu/lesson_service_v2/router/middleware"
	"github.com/asynccnu/lesson_service_v2/script"

	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// 2021-2022.2.xlsx
var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")

	// 选课手册 Excel 文件路径，若不为空则启动蹭课数据导入脚本
	excelFilePath = pflag.StringP("path", "p", "", "Excel file path.")
)

func main() {
	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// logger sync
	defer log.SyncLogger()

	// init db
	model.DB.Init()
	defer model.DB.Close()

	// 选课课程数据导入
	if *excelFilePath != "" {
		script.SyncImportLessonData(*excelFilePath)
		return
	}

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// Create the Gin engine.
	g := gin.New()

	// Routes.
	router.Load(
		// Cores.
		g,

		// MiddleWares.
		middleware.Logging(),
		middleware.RequestId(),
	)

	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.",
				zap.String("reason", err.Error()))
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Info(
		fmt.Sprintf("Start to listening the incoming requests on http address: %s", viper.GetString("addr")))
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the router")
}
