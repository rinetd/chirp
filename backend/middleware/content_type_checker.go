package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: (if needed) support more than only JSON content type

// ContentTypeChecker check if request has set proper content-type.
func ContentTypeChecker() gin.HandlerFunc {
	return func(context *gin.Context) {
		contentType := context.Request.Header.Get("Content-Type")
		if contentType == "application/json" || contentType == "application/x-www-form-urlencoded" || contentType == "multipart/form-data" {
			context.Next()
		} else {
			context.AbortWithError(
				http.StatusUnsupportedMediaType,
				errors.New("Required content-type: application/json"),
			)
			return
		}
		// if contentType != "application/json" && contentType != "application/x-www-form-urlencoded" {
		// 	context.AbortWithError(
		// 		http.StatusUnsupportedMediaType,
		// 		errors.New("Required content-type: application/json"),
		// 	)
		// 	return
		// }
		//
		// context.Next()
	}
}
