package ui

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

type APIServer struct {
	engine *gin.Engine
}

func (a *APIServer) registry()  {
	APIServerInit(a.engine)
}


func (a *APIServer) init()  {

}

func APIServerInit(r *gin.Engine)  {

	// docs

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	v1 := r.Group("/v1/api")

	readersAll(v1)
	indexRegistry(v1)
	deleteReader(v1)
	createReader(v1)
	updataReader(v1)
}

func readersAll(r *gin.RouterGroup)  {
	r.GET("lookall", ReaderSelectAll)
}

func deleteReader(r *gin.RouterGroup)  {
	r.GET("delete/:url", DeleteReaderSingel)
}

func createReader(r *gin.RouterGroup)  {
	r.POST("create", CreateReader)
}

func updataReader(r *gin.RouterGroup)  {
	r.GET("updata/:url", UpdataReader)
}

func (a *APIServer) Start()  {
	a.registry()
	a.engine.Run(":8080")
}

func New() *APIServer  {
	return &APIServer{
		engine: gin.Default(),
	}
}

func indexRegistry(r *gin.RouterGroup) {
	r.GET("", HelloWorld)

}

func HelloWorld(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Hello World!"},
	)
}