package utils

import "github.com/gin-gonic/gin"

func RespondWithJSON(c *gin.Context, code int, message interface{}) {
	c.JSON(code, gin.H{"message": message})
}

func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"error": message})
}
