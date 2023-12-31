package response

import (
    "net/http"

    "github.com/gin-gonic/gin"

    "github.com/hd2yao/gin-skeleton/app/global/consts"
)

func ReturnJson(Context *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {
    Context.JSON(httpCode, gin.H{
        "code": dataCode,
        "msg":  msg,
        "data": data,
    })
}

// ReturnJsonFromString 将json字符窜以标准json格式返回（例如，从redis读取json格式的字符串，返回给浏览器json格式）
func ReturnJsonFromString(Context *gin.Context, httpCode int, jsonStr string) {
    Context.Header("Content-Type", "application/json; charset=utf-8")
    Context.String(httpCode, jsonStr)
}

// 语法糖函数封装

// Success 直接返回成功
func Success(c *gin.Context, msg string, data interface{}) {
    ReturnJson(c, http.StatusOK, consts.CurdStatusOkCode, msg, data)
}

// Fail 失败的业务逻辑
func Fail(c *gin.Context, dataCode int, msg string, data interface{}) {
    ReturnJson(c, http.StatusBadRequest, dataCode, msg, data)
    c.Abort()
}

// ErrorSystem 系统执行代码错误
func ErrorSystem(c *gin.Context, msg string, data interface{}) {
    ReturnJson(c, http.StatusInternalServerError, consts.ServerOccurredErrorCode, consts.ServerOccurredErrorMsg+msg, data)
    c.Abort()
}
