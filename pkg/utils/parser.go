package utils

import "github.com/gin-gonic/gin"

// ParseRequestBody parses the request body and returns the target interface
func ParseRequestBody(c *gin.Context, target interface{}) (interface{}, error) {
	if err := c.ShouldBindJSON(target); err != nil {
		return nil, err
	}
	return target, nil
}
