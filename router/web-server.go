package router

import (
	"binlog_spread/conf"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)


func StartHttpServer() {
	router := SetupRouter()

	err := router.Run(conf.Config.WebListen)
	if err != nil {
		log.Fatal("setup router err:", err)
		return
	}
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	//router.Use(middleware.CORSMiddleware())

	// 静态文件
	router.StaticFile("/favicon.ico", "./static/favicon.ico")   // 单文件
	router.Static("/dist", "./static/dist")                     // 目录下的文件
	router.Static("/plugins", "./static/plugins")               // 目录下的文件
	router.Static("/AdminLTE-3.0.5", "./static/AdminLTE-3.0.5") // 目录下的文件
	router.Static("/admin", "./static/admin")                   // 目录下的文件
	//router.StaticFS("/more_static", http.Dir("my_file_system"))  // 目录正的文件，定制file.System服务

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/admin/entrance/login.html")
	})

	// admin 路由 /admin-api/
	adminRoute(router)

	// 未匹配到任何路由
	router.NoRoute(func(c *gin.Context) {
		// router.HandleContext()
		c.AbortWithStatus(http.StatusNotFound)

	})

	return router
}
