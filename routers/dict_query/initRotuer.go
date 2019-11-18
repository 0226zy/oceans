package dict_query

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	query_api := r.Group("/dict")
	{
		query_api.GET("/query/:word", QueryWord)
	}
}
