package api

import (
	"fmt"
	"publisher/services"

	"github.com/gin-gonic/gin"
)

func ApiPublisher(c *gin.Context) {
	channel := c.Query("channel")
	if b, err := c.GetRawData(); err == nil {
		fmt.Println(string(b))
		if num, err := services.Publish(channel, string(b)); err == nil {
			c.Writer.WriteString(fmt.Sprintf("Success Number:(%d)", num))
			return
		} else {
			fmt.Printf("publisher failed")
			c.Writer.WriteString("Failed")
			return
		}
	}
	c.Writer.WriteString("Failed")
	return
}
