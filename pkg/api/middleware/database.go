package middleware

import (
	"booktrack/pkg/database"

	"github.com/gin-gonic/gin"
)

func SetDatabaseContext(gin *gin.Context) {
	ctx := gin.Request.Context()

	ctx = database.SetConnection(ctx, nil)

	gin.Request = gin.Request.WithContext(ctx)
}
