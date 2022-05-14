package main

import (
	"github.com/gin-gonic/gin"
	"github.com/linengcs/go-project/controller"
	"github.com/linengcs/go-project/repository"
	"os"
)

func main() {
	if err := Init("./data/"); err != nil {
		os.Exit(1)
	}
	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := controller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})
	err := r.Run()
	if err != nil {
		return
	}
}

func Init(filePath string) error {
	if err := repository.Init(filePath); err != nil {
		return err
	}
	return nil
}
