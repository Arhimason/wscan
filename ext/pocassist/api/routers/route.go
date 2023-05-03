package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wscan/ext/pocassist/api/middleware/jwt"
	"wscan/ext/pocassist/api/routers/v1/auth"
	"wscan/ext/pocassist/api/routers/v1/plugin"
	"wscan/ext/pocassist/api/routers/v1/scan/result"
	scan2 "wscan/ext/pocassist/api/routers/v1/scan/scan"
	"wscan/ext/pocassist/api/routers/v1/scan/task"
	"wscan/ext/pocassist/api/routers/v1/vulnerability"
	"wscan/ext/pocassist/api/routers/v1/webapp"
	_ "wscan/ext/pocassist/docs" // docs is generated by Swag CLI, you have to import it.
	"wscan/ext/pocassist/pkg/conf"
	log "wscan/ext/pocassist/pkg/logging"
)

func Setup() {
	// gin 的【运行模式】运行时就已经确定 无法做到热加载
	gin.SetMode(conf.GlobalConfig.ServerConfig.RunMode)
}

func InitRouter(port string) {
	router := gin.Default()

	// ui
	router.StaticFS("/ui", BinaryFileSystem("web/build"))
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/ui")
	})

	// api
	v1 := router.Group("/api/v1")
	{
		v1.POST("/user/login", auth.Login)
		userRoutes := v1.Group("/user")
		userRoutes.Use(jwt.JWT())
		{
			userRoutes.POST("/self/resetpwd/", auth.Reset)
			userRoutes.GET("/info", auth.Self)
			userRoutes.GET("/logout", auth.Logout)
		}

		pluginRoutes := v1.Group("/poc")
		pluginRoutes.Use(jwt.JWT())
		{
			// all
			pluginRoutes.GET("/", plugin.Get)
			// 增
			pluginRoutes.POST("/", plugin.Add)
			// 改
			pluginRoutes.PUT("/:id/", plugin.Update)
			// 详情
			pluginRoutes.GET("/:id/", plugin.Detail)
			// 删
			pluginRoutes.DELETE("/:id/", plugin.Delete)
			// 测试单个poc
			pluginRoutes.POST("/run/", plugin.Run)
			// 上传yaml文件
			pluginRoutes.POST("/upload/", plugin.UploadYaml)
			// 下载yaml文件
			pluginRoutes.POST("/download/", plugin.DownloadYaml)
		}

		vulRoutes := v1.Group("/vul")
		vulRoutes.Use(jwt.JWT())
		{
			// basic
			vulRoutes.GET("/basic/", vulnerability.Basic)
			// all
			vulRoutes.GET("/", vulnerability.Get)
			// 增
			vulRoutes.POST("/", vulnerability.Add)
			// 改
			vulRoutes.PUT("/:id/", vulnerability.Update)
			// 详情
			vulRoutes.GET("/:id/", vulnerability.Detail)
			// 删
			vulRoutes.DELETE("/:id/", vulnerability.Delete)
		}

		appRoutes := v1.Group("/product")
		appRoutes.Use(jwt.JWT())
		{
			// all
			appRoutes.GET("/", webapp.Get)
			// 增
			appRoutes.POST("/", webapp.Add)
			// 改
			appRoutes.PUT("/:id/", webapp.Update)
			// 详情
			appRoutes.GET("/:id/", webapp.Detail)
			// 删
			appRoutes.DELETE("/:id/", webapp.Delete)
		}

		scanRoutes := v1.Group("/scan")
		scanRoutes.Use(jwt.JWT())
		{
			scanRoutes.POST("/url/", scan2.Url)
			scanRoutes.POST("/raw/", scan2.Raw)
			scanRoutes.POST("/list/", scan2.List)
		}

		taskRoutes := v1.Group("/task")
		taskRoutes.Use(jwt.JWT())
		{
			// all
			taskRoutes.GET("/", task.Get)
			// 删
			taskRoutes.DELETE("/:id/", task.Delete)
		}

		resultRoutes := v1.Group("/result")
		resultRoutes.Use(jwt.JWT())
		{
			// all
			resultRoutes.GET("/", result.Get)
			// 删
			resultRoutes.DELETE("/:id/", result.Delete)
		}

	}
	router.Run(":" + port)
	log.Info("server start at port:", port)
}
