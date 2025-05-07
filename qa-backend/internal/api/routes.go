package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	ConfigQuestionEndpoints(r)
	// Add more like: ConfigUserEndpoints(r), ConfigGroupEndpoints(r)
}
