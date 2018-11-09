package response

import "github.com/gin-gonic/gin"

func Success(data interface{}) gin.H {
	return Response(200, "success", data)
}

func Fail(msg string, code int) gin.H {
	return Response(code, msg, nil)
}

func Response(code int, msg string, data interface{}) gin.H {
	return gin.H{"code":code, "msg":msg , "data":data}
}
