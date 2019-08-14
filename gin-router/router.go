package gin_router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	BaseReq struct {
		Username string `json:"username"`
	}
	BaseResp struct {
		Success bool `json:"success"`
		Data
		Data1 interface{} `json:"data1"`
		Data2 Data        `json:"data2"`
		Data3 *Data       `json:"data3"`
	}
	Data struct {
		Message string `json:"message"`
	}
)

// @Router /ping/{name} [post]
// @Param name path string true "name"
// @Param age query int false "age"
// @Param req_param body gin_router.BaseReq true "req param"
// @Success 200 {object} gin_router.BaseResp"
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

func GinEngine() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", Ping)

	router.Static("/swagger-ui", "/home/qydev/go/src/gin-swagger/swagger-ui/")
	router.Static("/swagger-json", "/home/qydev/go/src/gin-swagger/docs/")

	return router
}
