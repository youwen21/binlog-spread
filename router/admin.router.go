package router


import (
	"binlog_spread/app/http/handler/admin"
	"binlog_spread/app/http/middleware"
	"github.com/gin-gonic/gin"
)

func adminRoute(router *gin.Engine) {
	adminOpenAPI := router.Group("/admin-api/v1")
	{
		adminOpenAPI.POST("/login", admin.AuthHandler.Login)
	}

	cacheAPI := router.Group("/admin-api/v1")
	//cacheAPI.Use(middleware.AdminToken)
	//cacheAPI.Use(middleware.BrowserCacheMiddleware)
	{
		// 菜单树
		cacheAPI.GET("/sidebar", admin.SidebarHandler.GetSideBarTree)
	}

	adminAPI := router.Group("/admin-api/v1")
	adminAPI.Use(middleware.AdminToken)
	// 框架菜单相关
	{
		// 菜单列表
		adminAPI.GET("/menu", admin.MenuHandler.Query)
		adminAPI.GET("/menu/:id", admin.MenuHandler.GetInfo)
		adminAPI.POST("/menu", admin.MenuHandler.Add)
		adminAPI.POST("/menu/:id", admin.MenuHandler.Edit)
		adminAPI.DELETE("/menu/:id", admin.MenuHandler.Delete)

		adminAPI.GET("/menu/select", admin.MenuHandler.GetSelectList)
	}

	// 数据模型相关
	{
		// 事件列表
		adminAPI.GET("/api_define", admin.ApiDefineHandler.Query)
		// 事件详情
		adminAPI.GET("/api_define/:id", admin.ApiDefineHandler.GetInfo)
		// 事件删除
		adminAPI.DELETE("/api_define/:id", admin.ApiDefineHandler.Delete)

		// 事件添加
		adminAPI.POST("/api_define", admin.ApiDefineHandler.Add)
		// 事件编辑
		adminAPI.POST("/api_define/:id", admin.ApiDefineHandler.Edit)

		// 事件对比 diff
		adminAPI.GET("/diff", admin.BinlogHandler.Diff)

		// 数据流列表
		adminAPI.GET("/binlog", admin.BinlogHandler.Query)
		// 数据流删除
		adminAPI.DELETE("/binlog/:id", admin.BinlogHandler.Delete)
		// er图 实体关系图
		adminAPI.GET("/er", admin.ERHandler.Query)
	}

	// 业务状态相关
	{
		adminAPI.GET("/state_class", admin.StateClassHandler.Query)
		adminAPI.GET("/state_class/:id", admin.StateClassHandler.GetInfo)
		adminAPI.POST("/state_class", admin.StateClassHandler.Add)
		adminAPI.POST("/state_class/:id", admin.StateClassHandler.Edit)
		adminAPI.DELETE("/state_class/:id", admin.StateClassHandler.Delete)

		adminAPI.GET("/state", admin.StateHandler.Query)
		adminAPI.GET("/state/:id", admin.StateHandler.GetInfo)
		adminAPI.POST("/state", admin.StateHandler.Add)
		adminAPI.POST("/state/:id", admin.StateHandler.Edit)
		adminAPI.DELETE("/state/:id", admin.StateHandler.Delete)

		adminAPI.GET("/state_direction", admin.StateDirectionHandler.Query)
		adminAPI.POST("/state_direction", admin.StateDirectionHandler.Add)
		adminAPI.DELETE("/state_direction/:id", admin.StateDirectionHandler.Delete)

		adminAPI.GET("/state_abnormal", admin.StateAbnormalHandler.Query)
		adminAPI.DELETE("/state_abnormal/:id", admin.StateAbnormalHandler.Delete)

		adminAPI.GET("/state_graph/:id", admin.StateGraphHandler.GetInfo)
	}

	//{
	//	adminOpenAPI.GET("/test", admin.TestHandler.TestGet)
	//}
}
