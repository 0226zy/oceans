package dict_query

import (
	"fmt"
	"gin-blog/pkg/e"
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueryWord(c *gin.Context) {
	word := c.Query("word")
	fmt.Println("QueryWord:" + word)
	fmt.Println(word)
	code := e.SUCCESS
	data := make(map[string]interface{})
	data["result"] = word

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
