package api

import "github.com/gin-gonic/gin"

func NewRouter(passHandler *PassHandler) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/passes", passHandler.GetPasses)
	}

	return r
}
