package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"os"
	docs "vitalring-seoul-go-api/docs"
	"vitalring-seoul-go-api/middleware"
	"vitalring-seoul-go-api/routes"
)

// @title VTRing API
// @version 0.1.56 (24022201)
// @description VTRing API server.
/*//@basePath /v2*/
// @schemes https http
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization
func main() {
	router := gin.Default()

	//router.Use(middleware.AllowedIp("127.0.0.1,192.168.0.1,::1,1.230.141.67")) // 허용된 ip만
	router.Use(middleware.UseCors())
	router.Use(middleware.UseMySql())
	router.Use(middleware.UseAuth())
	router.Use(middleware.UseFirestore())
	router.Use(middleware.UseMessage())
	router.Use(middleware.UseStorageBucket())

	setupSwagger(router)

	setupRoutesDefault(router)
	//setupRoutesGJSGOffice(router)

	runServer(router)
}

func setupSwagger(router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/api-docs", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
}

func setupRoutesDefault(r *gin.Engine) {
	r.GET("/", routes.RootHandler)

	r.POST("/ping", middleware.Authorization(), routes.PingPostHandler)
	r.GET("/ping", middleware.Authorization(), routes.PingGetHandler)
	r.GET("/count-ppg-all", middleware.Authorization(), routes.CountPpgAllGetHandler)
	r.GET("/count-all-days", middleware.Authorization(), routes.CountAllDaysGetHandler)

	r.GET("/data-api", middleware.Authorization(), routes.DataApiGetHandler)
	r.GET("/data-api-last", middleware.Authorization(), routes.DataApiLastGetHandler)
	r.GET("/data-api-between", middleware.Authorization(), routes.DataApiBetweenGetHandler)
	r.GET("/sleep-api", middleware.Authorization(), routes.SleepApiHandler)
	r.GET("/stress-api", middleware.Authorization(), routes.StressApiHandler)
	r.GET("/grid-api", middleware.Authorization(), routes.GridApiHandler)
	r.GET("/graph-api", routes.GraphHandler)
	r.GET("/graph-period-api", routes.GraphPeriodHandler)
	r.GET("/activity-api", middleware.Authorization(), routes.ActivityApiHandler)
	r.GET("/condition-api", middleware.Authorization(), routes.ConditionApiHandler)

	r.GET("/anomaly-abs-api", middleware.Authorization(), routes.AnomalyAbsHandler)
	r.GET("/anomaly-datas-api", middleware.Authorization(), routes.AnomalyDatasHandler)
	r.GET("/anomaly-schedule-api", routes.AnomalyScheduleApi)
	r.GET("/anomaly-datas-schedule-api", routes.AnomalyDatasUpdateScheduleApi)

	r.GET("/sum-users-days", routes.SumUsersDaysGetHandler)
	r.GET("/sum-vital-data", routes.SumVitalDataGetHandler)
	r.GET("/sum-hrv-report", routes.SumHrvReportGetHandler)
	r.GET("/count-data-users", routes.CountDataUsersGetHandler)
	r.GET("/count-auth-users", routes.CountAuthUsersGetHandler)
	r.GET("/sum-app-used", routes.SumAppUsedGetHandler)

	r.GET("/auth-users", middleware.Authorization(), middleware.OnlyAdmin(), routes.AuthUsersGetHandler)
	//r.POST("/auth-users", middleware.Authorization(), middleware.OnlyAdmin(), routes.AuthUsersPostHandler)
	r.PUT("/auth-users", middleware.Authorization(), middleware.OnlyAdmin(), routes.AuthUsersPutHandler)
	r.POST("/admin-push", middleware.Authorization(), middleware.OnlyAdmin(), routes.AdminPushPostHandler)
	r.GET("/lastupdate", routes.LastUpdateGetHandler)
	r.GET("/push-message", routes.PushMessageGetHandler)

	r.GET("/my-info", middleware.Authorization(), routes.MyInfoGetHandler)
	r.PUT("/my-info-doc", middleware.Authorization(), routes.MyInfoDocPutHandler)

	r.GET("/user-device-info", middleware.Authorization(), routes.UserDeviceInfoGetHandler)

	r.GET("/users", middleware.Authorization(), routes.GroupUsersGetHandler)
	r.DELETE("/user-sleep", middleware.Authorization(), routes.UserSleepDeleteHandler)
	r.PUT("/user-sync-finish", middleware.Authorization(), routes.UserSyncFinishPutHandler)

	r.GET("/device-firmware", middleware.Authorization(), routes.DeviceFirmwareGetHandler)

	r.GET("/group-notice", middleware.Authorization(), routes.GroupNoticeGetHandler)
	r.POST("/group-notice", middleware.Authorization(), routes.GroupNoticePostHandler)
	r.PUT("/group-notice", middleware.Authorization(), routes.GroupNoticePutHandler)
	r.DELETE("/group-notice", middleware.Authorization(), routes.GroupNoticeDeleteHandler)

	r.PUT("/group-info-doc", middleware.Authorization(), routes.GroupInfoDocPutHandler)
	r.PUT("/group-setting-info-doc", middleware.Authorization(), routes.GroupSettingInfoDocPutHandler)
	r.GET("/group-setting-noti", middleware.Authorization(), routes.GroupSettingNotiGetHandler)
	r.PUT("/group-setting-noti", middleware.Authorization(), routes.GroupSettingNotiPutHandler)

	r.GET("/group-user", middleware.Authorization(), routes.GroupUserGetHandler)
	r.PUT("/group-user", middleware.Authorization(), routes.GroupUserPutHandler)
	r.PUT("/group-user-json", middleware.Authorization(), routes.GroupUserJsonPutHandler)
	r.POST("/group-user", middleware.Authorization(), routes.GroupUserPostHandler)
	r.DELETE("/group-user", middleware.Authorization(), routes.GroupUserDeleteHandler)

	r.GET("/group-user-report-hrv", middleware.Authorization(), routes.GroupUserReportHrvGetHandler)
	r.GET("/group-users", middleware.Authorization(), routes.GroupUsersGetHandler)
	r.GET("/group-users-scores", middleware.Authorization(), routes.GroupUsersScoresGetHandler)
	r.GET("/group-users-statistics", middleware.Authorization(), routes.GroupUsersStatisticsGetHandler)
	r.GET("/group-users-anomaly", middleware.Authorization(), routes.GroupUsersAnomalyGetHandler)
	r.GET("/group-user-anomaly", middleware.Authorization(), routes.GroupUserAnomalyGetHandler)
	r.GET("/group-user-total-anomaly", middleware.Authorization(), routes.GroupUserTotalAnomalyGetHandler)
	r.GET("/group-user-anomaly-dayMillis", middleware.Authorization(), routes.GroupUserAnomalyDayMillisGetHandler)

	r.GET("/user-notifications", middleware.Authorization(), routes.UserNotificationsGetHandler)
	r.PUT("/user-notifications-read", middleware.Authorization(), routes.UserNotificationsReadPutHandler)

	r.PUT("/user-day-score", middleware.Authorization(), routes.UserDayScorePutHandler)
	r.GET("/user-email-check", routes.UserEmailCheckGetHandler)

	r.GET("/user-af-datas", middleware.Authorization(), routes.UserAfDatasGetHandler)
	r.PUT("/user-af-datas-sync", routes.UserAfDatasSyncPutHandler)
	r.GET("/user-af-raw", middleware.Authorization(), routes.UserAfRawGetHandler)
	r.POST("/user-af-raw", middleware.Authorization(), routes.UserAfRawPostHandler)
	r.POST("/user-af-raw-userid", routes.UserAfRawUserIdPostHandler)

	r.GET("/group-device", middleware.Authorization(), routes.GroupDeviceGetHandler)
	r.POST("/group-device", middleware.Authorization(), routes.GroupDevicePostHandler)
	r.PUT("/group-device", middleware.Authorization(), routes.GroupDevicePutHandler)
	r.DELETE("/group-device", middleware.Authorization(), routes.GroupDeviceDeleteHandler)

	r.POST("/inquire", middleware.Authorization(), routes.InquirePostHandler)
	r.GET("/inquire", middleware.Authorization(), routes.InquireGetHandler)

	r.GET("/user-message", middleware.Authorization(), routes.UserMessageGetHandler)
	r.GET("/user-message-new", middleware.Authorization(), routes.UserMessageNewGetHandler)
}

func setupRoutesGJSGOffice(r *gin.Engine) {
	r.GET("/", routes.RootHandler)

	r.GET("/activity-api", routes.ActivityApiHandler)
	r.GET("/condition-api", routes.ConditionApiHandler)
	r.GET("/graph-api", routes.GraphHandler)
	r.GET("/graph-period-api", routes.GraphPeriodHandler)
	r.GET("/grid-api", routes.GridApiHandler)
	r.GET("/sleep-api", routes.SleepApiHandler)
	r.GET("/stress-api", routes.StressApiHandler)

	r.GET("/data-api", routes.DataApiGetHandler)
	r.GET("/data-api-between", routes.DataApiBetweenGetHandler)
	r.GET("/data-api-last", routes.DataApiLastGetHandler)

	r.GET("/user-af-raw", routes.UserAfRawGetHandler)
	r.POST("/user-af-raw", middleware.Authorization(), routes.UserAfRawPostHandler)

	r.GET("/user-device-info", routes.UserDeviceInfoGetHandler)
}

func runServer(router *gin.Engine) {
	// PORT environment variable is provided by Cloud Run.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Print("Hello from Cloud Run! The container started successfully and is listening for HTTP requests on $PORT")

	address := ":" + port
	log.Printf("Listening on port http://localhost:%s/swagger/index.html", port)

	// Run the router and handle any potential error
	if err := router.Run(address); err != nil {
		log.Fatal(err)
	}
}
