package main

import (
	//	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"path/filepath"

	Et "github.com/jof4002/Everything"
)

type fileInfo struct {
	Name  string    `json:"name"`
	Path  string    `json:"path"`
	Size  int64     `json:"size"`
	Fdate time.Time `json:"fdate"`
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/search", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		//keyword:=c.Param("keyword")
		keyword := c.Query("keyword")
		//fileInfoArray = doEverythingSearch(keyword)
		fileInfoArray := []fileInfo{}
		if keyword == "" {
			fileInfoArray = append(fileInfoArray, fileInfo{})
		} else {
			Et.Walk(keyword, func(path string, info Et.FileInfo, err error) error {
				name := filepath.Base(path)
				size := info.Size()
				tmod := info.ModTime()
				//fmt.Println("info: ",name,path, size, tmod.Format("2006-01-02 15:04"))
				//fileInfoArray = append(fileInfoArray,fileInfo{name,path,size,tmod})
				fileInfoArray = append(fileInfoArray, fileInfo{name, path, size, tmod})
				return nil
			})
		}

		c.JSON(200, gin.H{
			/*
				"message": fileInfoArray,
				"keyword":keyword,
			*/
			"data": fileInfoArray,
		})
	})

	r.GET("/checkstatus", func(c *gin.Context) {
		//fmt.Println("aaa",fileInfoArray)
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	//fileInfoArray = []fileInfo{}
	r.Run() // listen and serve on 0.0.0.0:8080
}
