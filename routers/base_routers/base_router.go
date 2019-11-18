package base_routers

import (
	"gin-blog/pkg/e"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type TodoHandlerT func(bs *BaseRouterService)

type BaseRouterService struct {
	Context *gin.Context
	todo    TodoHandlerT
	Resp    gin.H
	//TODO logger
}

func DefaultBaseRouterService(c *gin.Context, todo TodoHandlerT) *BaseRouterService {
	return &BaseRouterService{c, todo, gin.H{}}
}

func (this *BaseRouterService) SetSuccessStatus(code int, msg string) {
	this.Resp["code"] = code
	if len(msg) == 0 {
		this.Resp["msg"] = e.GetMsg(code)
	} else {
		this.Resp["msg"] = msg
	}
}

func (this *BaseRouterService) SetFailedStatus(code int, msg string) {
	this.Resp["code"] = code
	if len(msg) == 0 {
		this.Resp["msg"] = e.GetMsg(code)
	} else {
		this.Resp["msg"] = msg
	}
}

func (this *BaseRouterService) Execute() {
	// TODO befor todo
	this.todo(this)
	this.Context.JSON(http.StatusOK, this.Resp)
	// TODO after todo
}

func (this *BaseRouterService) SetRespData(data map[string]interface{}) {
	this.Resp["data"] = data
}

func Struct2Map(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}
