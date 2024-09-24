package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code    int         `json:"code"`    // code status
	Message string      `json:"message"` // message error
	Data    interface{} `json:"data"`    // data return
}

// success response
func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, ResponseData{
		Code: code,
		Message: msg[code],
		Data: data,
	})
}

// error response
func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, ResponseData{
		Code: code,
		Message: msg[code],
		Data: nil,
	})
}