package gContext

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func HandleAPIError(ctx *gin.Context, errMsg string, err error, statusCode int) bool {
	if err != nil {
		method := ctx.Request.Method
		path := ctx.Request.URL.Path
		clientIP := ctx.ClientIP()
		userAgent := ctx.Request.UserAgent()
		queryParams := ctx.Request.URL.RawQuery

		errCode := fmt.Sprintf("status_code %d", statusCode)
		errMethod := fmt.Sprintf("method %s", method)
		errPath := fmt.Sprintf("path %s", path)
		errIP := fmt.Sprintf("client_ip %s", clientIP)
		errUserAgent := fmt.Sprintf("user_agent %s", userAgent)
		errQueryParams := fmt.Sprintf("query_params %s", queryParams)

		log.Printf(
			"An error has occurred: |%s| |%s| |%s| |%s| |%s| |%s|\n",
			errCode,
			errMethod,
			errPath,
			errIP,
			errUserAgent,
			errQueryParams,
		)
		log.Printf("Stack Trace: \n%+v\n", err)

		ctx.JSON(statusCode, gin.H{"error": errMsg})
		return true
	}
	return false
}
