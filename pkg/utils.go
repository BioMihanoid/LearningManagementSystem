package pkg

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetID(c *gin.Context, param string) (int, error) {
	res := c.Param(param)
	id, err := strconv.Atoi(res)
	if err != nil {
		return 0, err
	}
	return id, nil
}
